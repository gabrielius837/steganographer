package steganographer

import (
	"testing"
)

func TestAdler32_MustMatch(t *testing.T) {
	input := []byte{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}
	expected := uint32(0x0aff020e)

	result := Adler32(input)

	if result != expected {
		t.Fatalf("Adler32 failed, got 0x%X expected 0x%X", result, expected)
	}
}
