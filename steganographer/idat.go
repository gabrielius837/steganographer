package steganographer

import (
	"encoding/binary"
	"fmt"
)

const (
	IDAT_INDEX = uint32(0x49444154)
)

type IDAT struct {
	cm      byte
	cinfo   byte
	fcheck  byte
	fdict   byte
	flevel  byte
	data    []byte
	adler32 uint32
}

func (chunk *Chunk) NewIDAT() (*IDAT, error) {
	_, err := chunk.IsValid()
	if err != nil {
		return nil, err
	}

	if chunk.GetHeaderIdentifier() != IDAT_INDEX {
		return nil, NewChunkError(chunk.offset, fmt.Sprintf("unexpected chunk header, got: %s, expected: IDAT", chunk.header))
	}

	cm := chunk.data[0] & 0b00001111
	cinfo := chunk.data[0] & 0b11110000 >> 4

	fcheck := chunk.data[1] & 0b00001111
	fdict := chunk.data[1] & 0b00010000 >> 4
	flevel := chunk.data[1] & 0b11100000 >> 5

	length := len(chunk.data)
	data := chunk.data[2 : length-5]
	adler32 := binary.BigEndian.Uint32(chunk.data[length-5 : length])

	return &IDAT{
		cm:      cm,
		cinfo:   cinfo,
		fcheck:  fcheck,
		fdict:   fdict,
		flevel:  flevel,
		data:    data,
		adler32: adler32,
	}, nil
}

func (idat *IDAT) String() string {
	return fmt.Sprintf(
		"cm: %d, cinfo: %d, fcheck: %d, fdict: %d, flevel: %d, data length: %d, adler32: 0x%x",
		idat.cm, idat.cinfo, idat.fcheck, idat.fdict, idat.flevel, len(idat.data), idat.adler32,
	)
}

func Adler32(data []byte) uint32 {
	adler := uint32(1)

	const (
		op   = uint32(0xFFFF)
		base = uint32(65521)
	)

	var (
		s1 = adler & op
		s2 = (adler >> 16) & op
	)

	for _, value := range data {
		s1 = (s1 + uint32(value)) % base
		s2 = (s2 + s1) % base
	}

	return (s2 << 16) + s1
}
