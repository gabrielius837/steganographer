package steganographer

import (
	"encoding/binary"
	"fmt"
)

const (
	IHDR_INDEX  = uint32(0x49484452)
	IHDR_LENGTH = 13
)

type pixelation int

const (
	// fallback
	pixelation_unknown pixelation = -1
	// grayscale
	pixelation_0_1                  = 1
	pixelation_0_2                  = 2
	pixelation_0_4                  = 4
	pixelation_0_8                  = 8
	pixelation_0_16                 = 16
	pixelation_grayscale pixelation = 0
	// triple
	pixelation_2_8               = 2<<8 | 8
	pixelation_2_16              = 2<<8 | 16
	pixelation_triple pixelation = 1
	// palette, must have PLTE chunk
	pixelation_3_1                = 3<<8 | 1
	pixelation_3_2                = 3<<8 | 2
	pixelation_3_4                = 3<<8 | 4
	pixelation_3_8                = 3<<8 | 8
	pixelation_palette pixelation = 2
	// grayscale with alpha
	pixelation_4_8                        = 4<<8 | 8
	pixelation_4_16                       = 4<<8 | 16
	pixelation_grayscale_alpha pixelation = 3
	// triple with alpha
	pixelation_6_8                     = 6<<8 | 8
	pixelation_6_16                    = 6<<8 | 16
	pixelation_triple_alpha pixelation = 4
)

func (ihdr *IHDR) getPixelation() pixelation {
	ihdrPixelation := int(ihdr.color<<8) | int(ihdr.depth)
	switch ihdrPixelation {
	case pixelation_0_1, pixelation_0_2, pixelation_0_4, pixelation_0_8, pixelation_0_16:
		return pixelation_grayscale
	case pixelation_2_8, pixelation_2_16:
		return pixelation_triple
	case pixelation_3_1, pixelation_3_2, pixelation_3_4, pixelation_3_8:
		return pixelation_palette
	case pixelation_4_8, pixelation_4_16:
		return pixelation_grayscale_alpha
	case pixelation_6_8, pixelation_6_16:
		return pixelation_triple_alpha
	default:
		return pixelation_unknown
	}
}

type IHDR struct {
	width       uint32
	height      uint32
	depth       byte
	color       byte
	compression byte
	filter      byte
	interlace   byte
}

func (chunk *Chunk) NewIHDR() (*IHDR, error) {
	_, err := chunk.IsChunkValid([uint32_size]byte{0x49, 0x48, 0x44, 0x52})
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
		color:       chunk.data[9],
		compression: chunk.data[10],
		filter:      chunk.data[11],
		interlace:   chunk.data[12],
	}, nil
}

func (ihdr *IHDR) String() string {
	return fmt.Sprintf(
		"width: %d height: %d depth: %d color: %d compression: %d filter: %d interlace: %d",
		ihdr.width, ihdr.height, ihdr.depth, ihdr.color, ihdr.compression, ihdr.filter, ihdr.interlace,
	)
}
