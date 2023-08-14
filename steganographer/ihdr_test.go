package steganographer

import (
	"testing"
)

const (
	width       = 0x00000280
	height      = 0x000001e0
	depth       = 8
	colour      = 6
	compression = 0
	filter      = 0
	interlace   = 0
)

func TestNewIHDR_MustBeValid(t *testing.T) {
	chunk := &Chunk{
		offset: 8,
		length: 13,
		header: [4]byte{0x49, 0x48, 0x44, 0x52},
		data:   []byte{0x00, 0x00, 0x02, 0x80, 0x00, 0x00, 0x01, 0xe0, 0x08, 0x06, 0x00, 0x00, 0x00},
		crc:    0x35d1dce4,
	}

	ihdr, err := chunk.NewIHDR()

	if err != nil {
		t.Fatalf("NewIHDR failed, got error:\n%s", err.Error())
	}

	if width != ihdr.width ||
		height != ihdr.height ||
		depth != ihdr.depth ||
		colour != ihdr.colour ||
		compression != ihdr.compression ||
		filter != ihdr.filter ||
		interlace != ihdr.interlace {
		t.Fatalf(
			"NewIHDR failed, got unexpected values:\nwidth %d, got %d\nheight %d, got %d\ndepth %d, got %d\ncolour %d, got %d\ncompression %d, got %d\nfilter %d, got %d\ninterlace %d, got %d\n",
			ihdr.width, width, ihdr.height, height, ihdr.depth, depth, ihdr.colour, colour, ihdr.compression, compression, ihdr.filter, filter, ihdr.interlace, interlace,
		)
	}

}
