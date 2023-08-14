package steganographer

import (
	"fmt"
)

type ChunkError struct {
	offset  uint32
	message string
}

func NewChunkError(offset uint32, message string) error {
	return &ChunkError{offset: offset, message: message}
}

func NewChunkWrappedError(offset uint32, error error) error {
	return &ChunkError{offset: offset, message: error.Error()}
}

func (error *ChunkError) Error() string {
	return fmt.Sprintf("offset: %d, %s", error.offset, error.message)
}

type IsValid interface {
	IsValid() error
}
