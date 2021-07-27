package modzy

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"
)

type URIEncodable func() (io.Reader, error)

func URIEncodedString(alreadyEncoded string) URIEncodable {
	return func() (io.Reader, error) {
		return strings.NewReader(alreadyEncoded), nil
	}
}

func URIEncodeBytes(notEncodedBytes []byte, mimeType string) URIEncodable {
	return func() (io.Reader, error) {
		return URIEncodeReader(bytes.NewReader(notEncodedBytes), mimeType)
	}
}

func URIEncodeString(notEncodedString string, mimeType string) URIEncodable {
	return func() (io.Reader, error) {
		return URIEncodeReader(strings.NewReader(notEncodedString), mimeType)
	}
}

func URIEncodeFile(file *os.File, mimeType string) URIEncodable {
	return func() (io.Reader, error) {
		return URIEncodeReader(file, mimeType)
	}
}

// URIEncodeFilename will attempt to detect the mimeType if not provided based
// on the filename extension.
func URIEncodeFilename(filename string, mimeType string) URIEncodable {
	return func() (io.Reader, error) {
		file, err := os.Open(filename)
		if err != nil {
			return nil, errors.WithMessagef(err, "Failed to open file: %s", filename)
		}
		if mimeType == "" {
			mimeType = DetectMimeType(filename)
		}
		return URIEncodeFile(file, mimeType)()
	}
}

func URIEncodeReader(source io.Reader, mimeType string) (io.Reader, error) {
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}
	sourceBytes, err := io.ReadAll(source)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to read source data")
	}
	sourceBase64 := base64.StdEncoding.EncodeToString(sourceBytes)
	return strings.NewReader(fmt.Sprintf(`data:%s;base64,%s`, mimeType, sourceBase64)), nil
}

func DetectMimeType(filename string) string {
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
