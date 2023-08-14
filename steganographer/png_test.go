package steganographer

import (
	"io/fs"
	"os"
	"testing"
)

const (
	PNG_FILE  = "png.png"
	TEST_FILE = "test"
)

func TestHasPngHeader_MustBeTrue(t *testing.T) {
	file, err := os.OpenFile(PNG_FILE, os.O_RDONLY, fs.ModeExclusive)
	if err != nil {
		t.Fatalf("Could not read png file: %s, reason: %s", PNG_FILE, err.Error())
	}
	defer file.Close()

	result := HasPngHeader(file)

	if !result {
		t.Fatalf("Expecting %s to have a valid png header", PNG_FILE)
	}
}

func TestHasPngHeader_MustBeFalse(t *testing.T) {
	file, err := os.OpenFile(TEST_FILE, os.O_RDONLY, fs.ModeExclusive)
	if err != nil {
		t.Fatalf("Could not read test file: %s, reason: %s", TEST_FILE, err.Error())
	}
	defer file.Close()

	result := HasPngHeader(file)

	if result {
		t.Fatalf("Expecting %s to don't have a valid png header", TEST_FILE)
	}
}
