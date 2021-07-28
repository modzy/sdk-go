package modzy

import (
	"bytes"
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"
)

type JobInputable func() (*jobInputableData, error)

type jobInputableType string

const (
	jobInputableTypeString   jobInputableType = "string"
	jobInputableTypeEmbedded jobInputableType = "embedded"
	jobInputableTypeByte     jobInputableType = "byte"
)

type jobInputableData struct {
	Type jobInputableType
	Data io.Reader
}

func JobInputTextReader(textReader io.Reader) JobInputable {
	return func() (*jobInputableData, error) {
		return &jobInputableData{
			Type: jobInputableTypeString,
			Data: textReader,
		}, nil
	}
}

func JobInputText(text string) JobInputable {
	return JobInputTextReader(strings.NewReader(text))
}

func JobInputTextFile(filename string) JobInputable {
	return func() (*jobInputableData, error) {
		file, err := os.Open(filename)
		if err != nil {
			return nil, errors.WithMessagef(err, "Failed to open file: %s", filename)
		}
		return JobInputTextReader(file)()
	}
}

func JobInputURIEncodedReader(uriEncodedStream io.Reader) JobInputable {
	return func() (*jobInputableData, error) {
		return &jobInputableData{
			Type: jobInputableTypeEmbedded,
			Data: uriEncodedStream,
		}, nil
	}
}

func JobInputURIEncodedString(uriEncoded string) JobInputable {
	return JobInputURIEncodedReader(strings.NewReader(uriEncoded))
}

func JobInputURIEncodedFile(filename string) JobInputable {
	return func() (*jobInputableData, error) {
		file, err := os.Open(filename)
		if err != nil {
			return nil, errors.WithMessagef(err, "Failed to open file: %s", filename)
		}
		return JobInputURIEncodedReader(file)()
	}
}

func JobInputByteReader(dataReader io.Reader) JobInputable {
	return func() (*jobInputableData, error) {
		return &jobInputableData{
			Type: jobInputableTypeByte,
			Data: dataReader,
		}, nil
	}
}

func JobInputBytes(data []byte) JobInputable {
	return JobInputByteReader(bytes.NewReader(data))
}

func JobInputFile(filename string) JobInputable {
	return func() (*jobInputableData, error) {
		file, err := os.Open(filename)
		if err != nil {
			return nil, errors.WithMessagef(err, "Failed to open file: %s", filename)
		}
		return JobInputByteReader(file)()
	}
}
