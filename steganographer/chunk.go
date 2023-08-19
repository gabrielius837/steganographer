package steganographer

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

const (
	INT_SIZE = 4
)

type Chunk struct {
	offset uint32
	length uint32
	header [4]byte
	data   []byte
	crc    uint32
}

func ReadChunk(file *os.File) (*Chunk, error) {
	offset, err := file.Seek(0, io.SeekCurrent)
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, INT_SIZE)

	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}
	length := binary.BigEndian.Uint32(buffer)

	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}
	var header [INT_SIZE]byte
	for i := 0; i < INT_SIZE; i++ {
		header[i] = buffer[i]
	}

	data := make([]byte, length)
	_, err = file.Read(data)
	if err != nil {
		return nil, err
	}

	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}
	crc := binary.BigEndian.Uint32(buffer)

	return &Chunk{
		offset: uint32(offset),
		length: length,
		header: header,
		data:   data,
		crc:    crc,
	}, nil
}

func (chunk *Chunk) String() string {
	return fmt.Sprintf(
		"offset: %d length: %d header: %s crc: 0x%x",
		chunk.offset,
		chunk.length,
		chunk.header,
		chunk.crc,
	)
}

func (chunk *Chunk) GetHeaderIdentifier() uint32 {
	return uint32(chunk.header[3]) |
		uint32(chunk.header[2])<<8 |
		uint32(chunk.header[1])<<16 |
		uint32(chunk.header[0])<<24
}

func (chunk *Chunk) PrintChunkInfo() {
	parsedChunk, err := chunk.GetChunkInfo()

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	} else {
		fmt.Println(parsedChunk)
	}
}

func (chunk *Chunk) Decompress() ([]byte, error) {
	if chunk.GetHeaderIdentifier() != IDAT_INDEX {
		return nil, fmt.Errorf("can only decompress IDAT chunks")
	}

	r, err := zlib.NewReader(bytes.NewReader(chunk.data))
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	return buf.Bytes(), nil
}

func (chunk *Chunk) GetChunkInfo() (fmt.Stringer, error) {
	index := chunk.GetHeaderIdentifier()
	var parsedChunk fmt.Stringer
	var err error
	switch index {
	case IHDR_INDEX:
		parsedChunk, err = chunk.NewIHDR()
	case IDAT_INDEX:
		parsedChunk, err = chunk.NewIDAT()
		bytes, err := chunk.Decompress()
		if err == nil {
			fmt.Printf("Decompressed bytes: 0x%x\n", bytes)
		} else {
			fmt.Fprintf(os.Stderr, "Failed to decompress: %e\n", err)
		}
		// print decompressed bytes
	case TEXT_INDEX:
		parsedChunk, err = chunk.NewTEXT()
	case TIME_INDEX:
		parsedChunk, err = chunk.NewTIME()
	default:
		err = NewChunkError(
			chunk.offset,
			fmt.Sprintf("cannot read info from %s", chunk.header),
		)
	}

	return parsedChunk, err
}

func (chunk *Chunk) isChunkCritical() bool {
	switch chunk.GetHeaderIdentifier() {
	case IHDR_INDEX, PLTE_INDEX, IDAT_INDEX, IEND_INDEX:
		return true
	default:
		return false
	}
}
