package modzy

import (
	"io"

	"github.com/pkg/errors"
)

type ChunkEncodable func() (io.Reader, error)

func ChunkReader(data io.Reader) ChunkEncodable {
	return func() (io.Reader, error) {
		return data, nil
	}
}

func ChunkFile(filename string) ChunkEncodable {
	return func() (io.Reader, error) {
		file, err := AppFs.Open(filename)
		if err != nil {
			return nil, errors.WithMessagef(err, "Failed to open file: %s", filename)
		}
		return ChunkReader(file)()
	}
}
