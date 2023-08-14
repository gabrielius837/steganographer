package steganographer

import (
	"encoding/binary"
	"fmt"
)

const (
	IHDR_INDEX  = uint32(0x49484452)
	IHDR_LENGTH = 13
)

type IHDR struct {
	width       uint32
	height      uint32
	depth       byte
	colour      byte
	compression byte
	filter      byte
	interlace   byte
}

func (chunk *Chunk) NewIHDR() (*IHDR, error) {
	_, err := chunk.IsValid()
	if err != nil {
		return nil, err
	}

	if chunk.GetHeaderIdentifier() != IHDR_INDEX {
		return nil, NewChunkError(chunk.offset, fmt.Sprintf("unexpected chunk header, expecting: IHDR, got: %s", chunk.header))
	}

	width := binary.BigEndian.Uint32(chunk.data)
	height := binary.BigEndian.Uint32(chunk.data[4:8])

	return &IHDR{
		width:       width,
		height:      height,
		depth:       chunk.data[8],
		colour:      chunk.data[9],
		compression: chunk.data[10],
		filter:      chunk.data[11],
		interlace:   chunk.data[12],
	}, nil
}

func (ihdr *IHDR) String() string {
	return fmt.Sprintf(
		"width: %d height: %d depth: %d colour: %d compression: %d filter: %d interlace: %d",
		ihdr.width, ihdr.height, ihdr.depth, ihdr.colour, ihdr.compression, ihdr.filter, ihdr.interlace,
	)
}
