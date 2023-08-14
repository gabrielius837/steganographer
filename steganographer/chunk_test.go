package steganographer

import (
	"testing"
)

func TestIsCrcValid_MustBeTrue(t *testing.T) {
	chunk := Chunk{
		header: [INT_SIZE]byte{
			0x49, 0x48, 0x44, 0x52,
		},
		data: []byte{
			0x00, 0x00, 0x02, 0x80, 0x00,
			0x00, 0x01, 0xe0, 0x08, 0x06,
			0x00, 0x00, 0x00,
		},
		crc: 0x35d1dce4,
	}

	output := chunk.CalculateCrc()
	if output != chunk.crc {
		t.Fatalf("IsCrcValid failed, expected be valid, got invalid")
	}
}

func TestCalculateCrc_MustBeFalse(t *testing.T) {
	chunk := Chunk{
		header: [INT_SIZE]byte{
			0x49, 0x48, 0x44, 0x52,
		},
		data: []byte{
			0x00, 0x00, 0x02, 0x80, 0x00,
			0x00, 0x01, 0xe0, 0x08, 0x05, //changed 0x06 to 0x05
			0x00, 0x00, 0x00,
		},
		crc: 0x35d1dce4,
	}

	output := chunk.CalculateCrc()
	if output == chunk.crc {
		t.Fatalf("IsCrcValid failed, expected be invalid, got valid")
	}
}

type Wrapper struct {
	input    [4]byte
	expected uint32
	result   uint32
}

func TestGetHeaderIdentifier_VerifyResult(t *testing.T) {
	arr := []Wrapper{
		{input: [4]byte{0x49, 0x48, 0x44, 0x52}, expected: 0x49484452, result: 0},
	}
	errors := make([]Wrapper, 0, len(arr))

	for _, item := range arr {
		chunk := Chunk{header: item.input}
		item.result = chunk.GetHeaderIdentifier()
		if item.result != item.expected {
			errors = append(errors, item)
		}
	}

	if len(errors) > 0 {
		for _, error := range errors {
			t.Fatalf("getHeaderCheckSum failed for %s, expected 0x%x, got 0x%x", error.input, error.expected, error.result)
		}
	}

}
