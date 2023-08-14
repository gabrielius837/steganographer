package steganographer

import (
	"fmt"
)

const (
	TEXT_INDEX = uint32(0x74455874)
)

type TEXT struct {
	text string
}

func (chunk *Chunk) NewTEXT() (*TEXT, error) {
	identifier := chunk.GetHeaderIdentifier()
	if identifier != TEXT_INDEX {
		return nil, NewChunkError(
			chunk.offset,
			fmt.Sprintf(
				"unexpected chunk header, expecting: tEXt, got: %s",
				chunk.header,
			),
		)
	}

	for i := 0; i < len(chunk.data); i++ {
		if chunk.data[i] == 0x00 {
			chunk.data[i] = 0x20
		}
	}

	return &TEXT{
		text: string(chunk.data),
	}, nil
}

func (text *TEXT) String() string {
	return text.text
}
