package modzy

import (
	"io"
	"os"

	"github.com/pkg/errors"
)

type ChunkEncodable func() (io.Reader, error)

func ChunkEncodeFile(file *os.File) ChunkEncodable {
	return func() (io.Reader, error) {
		return file, nil
	}
}

// ChunkEncodeFilename will attempt to detect the mimeType if not provided based
// on the filename extension.
func ChunkEncodeFilename(filename string) ChunkEncodable {
	return func() (io.Reader, error) {
		file, err := os.Open(filename)
		if err != nil {
			return nil, errors.WithMessagef(err, "Failed to open file: %s", filename)
		}
		return ChunkEncodeFile(file)()
	}
}
