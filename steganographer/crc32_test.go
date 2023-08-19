package steganographer

import (
	"testing"
)

func Test_CalculateCrc32ForChunk_MustMatch(t *testing.T) {
	chunk := &Chunk{
		offset: 12,
		length: 13,
		header: [4]byte{0x49, 0x48, 0x44, 0x52},
		data: []byte{
			0x00, 0x00, 0x00, 0x01,
			0x00, 0x00, 0x00, 0x01,
			0x08,
			0x06,
			0x00,
			0x00,
			0x00,
		},
		crc: uint32(0x1f15c489),
	}
	crc := chunk.CalculateCrc()
	if crc != chunk.crc {
		t.Fatalf("expected to get matching crc %d != %d", chunk.crc, crc)
	}
}

func Test_CalculateCrc32ForChunk_MustNotMatch(t *testing.T) {
	chunk := &Chunk{
		offset: 12,
		length: 13,
		header: [4]byte{0x49, 0x48, 0x44, 0x52},
		data: []byte{
			// first byte is spoofed should be zero
			0x01, 0x00, 0x00, 0x01,
			0x00, 0x00, 0x00, 0x01,
			0x08,
			0x06,
			0x00,
			0x00,
			0x00,
		},
		crc: uint32(0x1f15c489),
	}
	crc := chunk.CalculateCrc()
	if crc == chunk.crc {
		t.Fatalf("expected to get non matching crc %d == %d", chunk.crc, crc)
	}
}
