package main

import (
	"fmt"
	"io/fs"
	"os"

	"project/steganographer"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Expecting single argument, path to png file\n")
		os.Exit(1)
	}

	path := os.Args[1]
	file, err := os.OpenFile(path, os.O_RDONLY, fs.ModeExclusive)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	defer file.Close()

	if !steganographer.HasPngHeader(file) {
		fmt.Fprintf(os.Stderr, "%s does not have valid png header", file.Name())
		os.Exit(1)
	}

	lastChunk := false
	for !lastChunk {
		chunk, err := steganographer.ReadChunk(file)
		if chunk.GetHeaderIdentifier() == steganographer.IEND_INDEX {
			lastChunk = true
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		fmt.Printf("%s\n", chunk)
		chunk.PrintChunkInfo()
	}

	if lastChunk {
		fmt.Println("success, IEND chunk encountered")
	}
}
