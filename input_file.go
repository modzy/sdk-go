package modzy

import (
	"io"

	"github.com/pkg/errors"
)

// FileInputEncodable is a data input that can be provided when using Jobs().SubmitJobFile(...).
//
// Provided implementations of this function are:
//	FileInputReader
//	FileInputFile
type FileInputEncodable func() (io.Reader, error)

func FileInputReader(data io.Reader) FileInputEncodable {
	return func() (io.Reader, error) {
		return data, nil
	}
}

func FileInputFile(filename string) FileInputEncodable {
	return func() (io.Reader, error) {
		file, err := AppFs.Open(filename)
		if err != nil {
			return nil, errors.WithMessagef(err, "Failed to open file: %s", filename)
		}
		return FileInputReader(file)()
	}
}
