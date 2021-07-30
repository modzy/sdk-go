package modzy_test

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	modzy "github.com/modzy/go-sdk"
	"github.com/spf13/afero"
)

func TestJobInputTextReader(t *testing.T) {
	expectedReader := strings.NewReader("a")

	r, err := modzy.JobInputTextReader(expectedReader)()

	if r.Data != expectedReader {
		t.Errorf("did not pass through reader")
	}
	if r.Type != "string" {
		t.Errorf("expected type to be string, got :%s", r.Type)
	}
	if err != nil {
		t.Errorf("expected nil error: %v", err)
	}
}

func TestJobInputText(t *testing.T) {
	r, err := modzy.JobInputText("abc")()

	b, _ := ioutil.ReadAll(r.Data)
	if string(b) != "abc" {
		t.Fatalf("did not read stream: %s", string(b))
	}
	if r.Type != "string" {
		t.Errorf("expected type to be string, got :%s", r.Type)
	}
	if err != nil {
		t.Errorf("expected nil error: %v", err)
	}
}

func TestJobInputTextFile(t *testing.T) {
	modzy.AppFs = afero.NewMemMapFs()
	_ = afero.WriteFile(modzy.AppFs, "src/a/b", []byte("abc"), 0644)

	r, err := modzy.JobInputTextFile("src/a/b")()

	b, _ := ioutil.ReadAll(r.Data)
	if string(b) != "abc" {
		t.Fatalf("did not read stream: %s", string(b))
	}
	if r.Type != "string" {
		t.Errorf("expected type to be string, got :%s", r.Type)
	}
	if err != nil {
		t.Errorf("expected nil error: %v", err)
	}
}

func TestJobInputTextFileErr(t *testing.T) {
	modzy.AppFs = afero.NewMemMapFs()

	_, err := modzy.JobInputTextFile("src/a/b")()

	if err == nil {
		t.Errorf("expected an error")
	}
}

func TestJobInputURIEncoded(t *testing.T) {
	r, err := modzy.JobInputURIEncodedString("abc")()

	b, _ := ioutil.ReadAll(r.Data)
	if string(b) != "abc" {
		t.Fatalf("did not read stream: %s", string(b))
	}
	if r.Type != "embedded" {
		t.Errorf("expected type to be string, got :%s", r.Type)
	}
	if err != nil {
		t.Errorf("expected nil error: %v", err)
	}
}

func TestJobInputURIEncodedFile(t *testing.T) {
	modzy.AppFs = afero.NewMemMapFs()
	_ = afero.WriteFile(modzy.AppFs, "src/a/b", []byte("abc"), 0644)

	r, err := modzy.JobInputURIEncodedFile("src/a/b")()

	b, _ := ioutil.ReadAll(r.Data)
	if string(b) != "abc" {
		t.Fatalf("did not read stream: %s", string(b))
	}
	if r.Type != "embedded" {
		t.Errorf("expected type to be string, got :%s", r.Type)
	}
	if err != nil {
		t.Errorf("expected nil error: %v", err)
	}
}

func TestJobInputURIEncodedFileErr(t *testing.T) {
	modzy.AppFs = afero.NewMemMapFs()

	_, err := modzy.JobInputURIEncodedFile("src/a/b")()

	if err == nil {
		t.Errorf("expected an error")
	}
}

func TestJobInputByteReader(t *testing.T) {
	expectedReader := bytes.NewReader([]byte("abc"))

	r, err := modzy.JobInputByteReader(expectedReader)()

	b, _ := ioutil.ReadAll(r.Data)
	if string(b) != "abc" {
		t.Fatalf("did not read stream: %s", string(b))
	}
	if r.Type != "byte" {
		t.Errorf("expected type to be string, got :%s", r.Type)
	}
	if err != nil {
		t.Errorf("expected nil error: %v", err)
	}
}

func TestJobInputBytes(t *testing.T) {
	r, err := modzy.JobInputBytes([]byte("abc"))()

	b, _ := ioutil.ReadAll(r.Data)
	if string(b) != "abc" {
		t.Fatalf("did not read stream: %s", string(b))
	}
	if r.Type != "byte" {
		t.Errorf("expected type to be string, got :%s", r.Type)
	}
	if err != nil {
		t.Errorf("expected nil error: %v", err)
	}
}

func TestJobInputByteFile(t *testing.T) {
	modzy.AppFs = afero.NewMemMapFs()
	_ = afero.WriteFile(modzy.AppFs, "src/a/b", []byte("abc"), 0644)

	r, err := modzy.JobInputFile("src/a/b")()

	b, _ := ioutil.ReadAll(r.Data)
	if string(b) != "abc" {
		t.Fatalf("did not read stream: %s", string(b))
	}
	if r.Type != "byte" {
		t.Errorf("expected type to be string, got :%s", r.Type)
	}
	if err != nil {
		t.Errorf("expected nil error: %v", err)
	}
}

func TestJobInputByteFileErr(t *testing.T) {
	modzy.AppFs = afero.NewMemMapFs()

	_, err := modzy.JobInputFile("src/a/b")()

	if err == nil {
		t.Errorf("expected an error")
	}
}
