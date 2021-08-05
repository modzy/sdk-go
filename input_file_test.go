package modzy_test

import (
	"io/ioutil"
	"strings"
	"testing"

	modzy "github.com/modzy/sdk-go"
	"github.com/spf13/afero"
)

func TestFileInputReader(t *testing.T) {
	expectedReader := strings.NewReader("a")
	r, err := modzy.FileInputReader(expectedReader)()
	if r != expectedReader {
		t.Errorf("did not pass through reader")
	}
	if err != nil {
		t.Errorf("expected nil error: %v", err)
	}
}

func TestFileInputFile(t *testing.T) {
	modzy.AppFs = afero.NewMemMapFs()
	_ = afero.WriteFile(modzy.AppFs, "src/a/b", []byte("file b"), 0644)

	r, err := modzy.FileInputFile("src/a/b")()

	b, _ := ioutil.ReadAll(r)
	if string(b) != "file b" {
		t.Fatalf("did not read file")
	}
	if err != nil {
		t.Errorf("expected nil error: %v", err)
	}
}

func TestFileInputFileError(t *testing.T) {
	modzy.AppFs = afero.NewMemMapFs()

	_, err := modzy.FileInputFile("not/a/file")()
	if err == nil {
		t.Errorf("expected an error")
	}
}
