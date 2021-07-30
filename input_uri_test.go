package modzy_test

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	modzy "github.com/modzy/go-sdk"
	"github.com/spf13/afero"
)

func TestURIEncodedReader(t *testing.T) {
	expectedReader := strings.NewReader("a")

	r, err := modzy.URIEncodedReader(expectedReader)()

	if r != expectedReader {
		t.Errorf("did not pass through reader")
	}
	if err != nil {
		t.Errorf("expected nil error: %v", err)
	}
}

func TestURIEncodedString(t *testing.T) {
	r, err := modzy.URIEncodedString("a")()

	b, _ := ioutil.ReadAll(r)
	if string(b) != "a" {
		t.Fatalf("did not read stream: %s", string(b))
	}
	if err != nil {
		t.Errorf("expected nil error: %v", err)
	}
}

func TestURIEncodedFile(t *testing.T) {
	modzy.AppFs = afero.NewMemMapFs()
	_ = afero.WriteFile(modzy.AppFs, "src/a/b", []byte("abc"), 0644)

	r, err := modzy.URIEncodedFile("src/a/b")()

	b, _ := ioutil.ReadAll(r)
	if string(b) != "abc" {
		t.Fatalf("did not read stream: %s", string(b))
	}
	if err != nil {
		t.Errorf("expected nil error: %v", err)
	}
}

func TestURIEncodedFileErr(t *testing.T) {
	modzy.AppFs = afero.NewMemMapFs()

	_, err := modzy.URIEncodedFile("src/a/b")()
	if err == nil {
		t.Errorf("expected an error")
	}
}

func TestURIEncodeReader(t *testing.T) {
	r, err := modzy.URIEncodeReader(strings.NewReader("a"), "")()

	if err != nil {
		t.Errorf("expected nil error: %v", err)
	}
	b, _ := ioutil.ReadAll(r)
	if string(b) != "data:application/octet-stream;base64,YQ==" {
		t.Fatalf("did not read stream: %s", string(b))
	}
}

type badReader struct{}

func (r *badReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("nope")
}

func TestURIEncodeBadReader(t *testing.T) {
	_, err := modzy.URIEncodeReader(&badReader{}, "")()

	if err == nil {
		t.Errorf("expected error")
	}
}

func TestURIEncodeString(t *testing.T) {
	r, err := modzy.URIEncodeString("a", "")()

	b, _ := ioutil.ReadAll(r)
	if string(b) != "data:application/octet-stream;base64,YQ==" {
		t.Fatalf("did not read stream: %s", string(b))
	}
	if err != nil {
		t.Errorf("expected nil error: %v", err)
	}
}

func TestURIEncodeFile(t *testing.T) {
	modzy.AppFs = afero.NewMemMapFs()
	_ = afero.WriteFile(modzy.AppFs, "src/a/b", []byte("a"), 0644)

	r, err := modzy.URIEncodeFile("src/a/b", "")()

	b, _ := ioutil.ReadAll(r)
	if string(b) != "data:application/octet-stream;base64,YQ==" {
		t.Fatalf("did not read stream: %s", string(b))
	}
	if err != nil {
		t.Errorf("expected nil error: %v", err)
	}
}

func TestURIEncodeFileErr(t *testing.T) {
	modzy.AppFs = afero.NewMemMapFs()

	_, err := modzy.URIEncodeFile("src/a/b", "")()
	if err == nil {
		t.Errorf("expected an error")
	}
}

var mimeTypes = map[string]string{
	"jpg":  "image/jpeg",
	"jpeg": "image/jpeg",
	"png":  "image/png",
}

func TestURIEncodeMimeTypes(t *testing.T) {
	for ext, mime := range mimeTypes {
		modzy.AppFs = afero.NewMemMapFs()
		_ = afero.WriteFile(modzy.AppFs, "file."+ext, []byte("a"), 0644)

		r, err := modzy.URIEncodeFile("file."+ext, "")()

		b, _ := ioutil.ReadAll(r)
		if string(b) != fmt.Sprintf("data:%s;base64,YQ==", mime) {
			t.Fatalf("did not encode mime type properly: %s/%s", mime, string(b))
		}
		if err != nil {
			t.Errorf("expected nil error: %v", err)
		}
	}
}
