package modzy

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TODO: are these type of unit tests worth writing?

func TestCancleJob(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/api/jobs/cancelID" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Write([]byte(`{"jobIdentifier": "jsonID"}`))
	}))
	defer serv.Close()
	client := NewClient(serv.URL)
	out, err := client.Jobs().CancelJob(context.TODO(), &CancelJobInput{JobIdentifier: "cancelID"})
	if err != nil {
		t.Errorf("err not nil")
	}
	if out.Details.JobIdentifier != "jsonID" {
		t.Errorf("response not parsed")
	}
}
