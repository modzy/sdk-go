package modzy

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TODO: are these type of unit tests worth writing?

func TestGetJobDetails(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/api/jobs/inputID" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Write([]byte(`{"jobIdentifier": "jsonID"}`))
	}))
	defer serv.Close()
	client := NewClient(serv.URL)
	out, err := client.Jobs().GetJobDetails(context.TODO(), &GetJobDetailsInput{JobIdentifier: "inputID"})
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if out.Details.JobIdentifier != "jsonID" {
		t.Errorf("response not parsed")
	}
}

func TestCancelJob(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/api/jobs/inputID" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Write([]byte(`{"jobIdentifier": "jsonID"}`))
	}))
	defer serv.Close()
	client := NewClient(serv.URL)
	out, err := client.Jobs().CancelJob(context.TODO(), &CancelJobInput{JobIdentifier: "inputID"})
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if out.Details.JobIdentifier != "jsonID" {
		t.Errorf("response not parsed")
	}
}

func TestGetJobResults(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/api/results/inputID" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Write([]byte(`{"jobIdentifier": "jsonID"}`))
	}))
	defer serv.Close()
	client := NewClient(serv.URL)
	out, err := client.Jobs().GetJobResults(context.TODO(), &GetJobResultsInput{JobIdentifier: "inputID"})
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if out.Results.JobIdentifier != "jsonID" {
		t.Errorf("response not parsed")
	}
}

func TestGetJobFeatures(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/api/jobs/features" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Write([]byte(`{"inputChunkMaximumSize": "someSize"}`))
	}))
	defer serv.Close()
	client := NewClient(serv.URL)
	out, err := client.Jobs().GetJobFeatures(context.TODO())
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if out.Features.InputChunkMaximumSize != "someSize" {
		t.Errorf("response not parsed")
	}
}
