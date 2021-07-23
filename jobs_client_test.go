package modzy

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestGetJobDetailsError(t *testing.T) {
	ctx := context.TODO()
	req := &fakeRequestor{
		GetFunc: func(ctx context.Context, path string, into interface{}) (*http.Response, error) {
			return nil, fmt.Errorf("nope")
		},
	}
	client := NewClient("base", withRequestor(req))
	_, err := client.Jobs().GetJobDetails(ctx, &GetJobDetailsInput{
		JobIdentifier: "jobID",
	})

	if err.Error() != "nope" {
		t.Errorf("did not hit fake")
	}
}

func TestGetJobDetails(t *testing.T) {
	ctx := context.TODO()
	req := &fakeRequestor{
		GetFunc: func(ctx context.Context, path string, into interface{}) (*http.Response, error) {
			if path != "/api/jobs/jobID" {
				t.Errorf("get url not expected: %s", path)
			}
			_ = json.Unmarshal([]byte(`{"jobIdentifier": "jsonID"}`), into)
			return nil, nil
		},
	}
	client := NewClient("base", withRequestor(req))
	out, _ := client.Jobs().GetJobDetails(ctx, &GetJobDetailsInput{
		JobIdentifier: "jobID",
	})
	if out.Details.JobIdentifier != "jsonID" {
		t.Errorf("payload did not parse")
	}
}

// TODO: are these type of unit tests worth writing?
