package steganographer

import (
	"testing"
)

func Test_IsValidPngHeader_MustBeTrue(t *testing.T) {
	header := validPngHeader[:]
	result := isValidPngHeader(header)
	if result == false {
		t.Fatalf("expected for input %v to be valid png header", header)
	}
}

func Test_IsValidPngHeader_MustBeFalse_WrongLength(t *testing.T) {
	header := validPngHeader[:][:7]
	result := isValidPngHeader(header)
	if result == true {
		t.Fatalf("expected for input %v to be invalid png header", header)
	}
}

func Test_IsValidPngHeader_MustBeFalse_SpoofedByte(t *testing.T) {
	header := make([]byte, 8)
	copy(header, validPngHeader[:])
	header[0] += 1
	result := isValidPngHeader(header)
	if result == true {
		t.Fatalf("expected for input %v to be invalid png header", header)
	}
}
