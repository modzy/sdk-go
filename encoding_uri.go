package modzy

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"
)

type URIEncodable func() (io.Reader, error)

func URIEncodedReader(alreadyEncoded io.Reader) URIEncodable {
	return func() (io.Reader, error) {
		return alreadyEncoded, nil
	}
}

func URIEncodedString(alreadyEncoded string) URIEncodable {
	return func() (io.Reader, error) {
		return strings.NewReader(alreadyEncoded), nil
	}
}

func URIEncodedFile(alreadyEncodedFilename string) URIEncodable {
	return func() (io.Reader, error) {
		file, err := os.Open(alreadyEncodedFilename)
		if err != nil {
			return nil, errors.WithMessagef(err, "Failed to open file: %s", alreadyEncodedFilename)
		}
		return URIEncodedReader(file)()
	}
}

func URIEncodeReader(notEncodedReader io.Reader, mimeType string) URIEncodable {
	return func() (io.Reader, error) {
		if mimeType == "" {
			mimeType = "application/octet-stream"
		}
		sourceBytes, err := io.ReadAll(notEncodedReader)
		if err != nil {
			return nil, errors.WithMessage(err, "failed to read source data")
		}
		sourceBase64 := base64.StdEncoding.EncodeToString(sourceBytes)
		return strings.NewReader(fmt.Sprintf(`data:%s;base64,%s`, mimeType, sourceBase64)), nil
	}
}

func URIEncodeString(notEncodedString string, mimeType string) URIEncodable {
	return URIEncodeReader(strings.NewReader(notEncodedString), mimeType)
}

func URIEncodeFile(filename string, mimeType string) URIEncodable {
	return func() (io.Reader, error) {
		if mimeType == "" {
			mimeType = detectMimeType(filename)
		}
		file, err := os.Open(filename)
		if err != nil {
			return nil, errors.WithMessagef(err, "Failed to open file: %s", filename)
		}
		return URIEncodeReader(file, mimeType)()
	}
}

func detectMimeType(filename string) string {
	split := strings.Split(filename, ".")
	extension := split[len(split)-1]

	// TODO what are common extensions to detect for this type of data?
	switch extension {
	case "jpg":
		fallthrough
	case "jpeg":
		return "image/jpeg"
	case "png":
		return "image/png"
	}
	return ""
}
