package modzy_test

import (
	"testing"

	modzy "github.com/modzy/go-sdk"
)

func TestS3Key(t *testing.T) {

	r, err := modzy.S3Key("bucket", "key")()

	if r.Bucket != "bucket" {
		t.Errorf("did not pass through bucket")
	}
	if r.Key != "key" {
		t.Errorf("did not pass through bucket")
	}
	if err != nil {
		t.Errorf("did not expect err: %v", err)
	}
}