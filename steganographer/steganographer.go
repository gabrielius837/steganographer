package steganographer

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

func getNewFileName(filename string) string {
	extension := filepath.Ext(filename)

	timestamp := time.Now().UnixNano()
	if len(extension) == 0 {
		return fmt.Sprintf("%s_%d.png", filename, timestamp)
	}

	index := strings.LastIndexByte(filename, '.')

	if index == -1 {
		return fmt.Sprintf("%d.png", timestamp)
	}

	return fmt.Sprintf("%s_%d.png", filename[:index], timestamp)
}
