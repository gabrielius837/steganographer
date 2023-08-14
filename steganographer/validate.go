package steganographer

const headerSize = 8

var (
	validPngHeader = [headerSize]byte{137, 80, 78, 71, 13, 10, 26, 10}
)

func isValidPngHeader(header []byte) bool {
	if len(header) != headerSize {
		return false
	}

	valid := true
	for i := 0; i < headerSize; i++ {
		match := header[i] == validPngHeader[i]
		valid = valid && match
	}

	return valid
}
