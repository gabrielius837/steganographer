package steganographer

import (
	"fmt"
	"time"
)

const (
	TIME_INDEX        = 0x74494d45
	TIME_CHUNK_LENGTH = 7
)

type TIME struct {
	timestamp time.Time
}

func (chunk *Chunk) NewTIME() (*TIME, error) {
	if chunk.GetHeaderIdentifier() != TIME_INDEX {
		return nil, NewChunkError(
			chunk.offset,
			fmt.Sprintf(
				"unexpected chunk header, expecting: tIME, got: %s",
				chunk.header,
			),
		)
	}

	if len(chunk.data) != TIME_CHUNK_LENGTH {
		return nil, NewChunkError(
			chunk.offset,
			fmt.Sprintf(
				"unexpected tIME chunk length, expecting %d, got %d",
				TIME_CHUNK_LENGTH, len(chunk.data),
			),
		)
	}

	year := int(chunk.data[1]) | int(chunk.data[0])<<8
	month := time.Month(chunk.data[2])
	day := int(chunk.data[3])
	hour := int(chunk.data[4])
	minute := int(chunk.data[5])
	second := int(chunk.data[6])
	timestamp := time.Date(
		year,
		month,
		day,
		hour,
		minute,
		second,
		0,
		time.UTC,
	)
	return &TIME{
		timestamp: timestamp,
	}, nil
}

func (time *TIME) String() string {
	return time.timestamp.String()
}
