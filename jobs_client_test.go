package modzy

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestGetJobDetailsHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()
	client := NewClient(serv.URL)
	_, err := client.Jobs().GetJobDetails(context.TODO(), &GetJobDetailsInput{JobIdentifier: "inputID"})
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestGetJobDetails(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected method to be GET, got %s", r.Method)
		}
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

func TestListJobsHistoryHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()
	client := NewClient(serv.URL)
	_, err := client.Jobs().ListJobsHistory(context.TODO(), (&ListJobsHistoryInput{}).WithPaging(2, 3))
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestListJobsHistory(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected method to be GET, got %s", r.Method)
		}
		if r.RequestURI != "/api/jobs/history?page=7&per-page=2" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Header().Set("Link", `<https://example>; rel="next"`)
		w.Write([]byte(`[{"jobIdentifier": "jsonID"},{"jobIdentifier": "jsonID2"}]`))
	}))
	defer serv.Close()
	client := NewClient(serv.URL)
	out, err := client.Jobs().ListJobsHistory(context.TODO(), (&ListJobsHistoryInput{}).WithPaging(2, 7))
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if out.Jobs[0].JobIdentifier != "jsonID" {
		t.Errorf("response not parsed")
	}
	if out.NextPage == nil {
		t.Errorf("expected NextPage to have a value")
	}
	if out.NextPage.Paging.Page != 8 {
		t.Errorf("expected NextPage to be next")
	}
}

func TestSubmitJobTextHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()
	client := NewClient(serv.URL)
	_, err := client.Jobs().SubmitJobText(context.TODO(), &SubmitJobTextInput{})
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestSubmitJobText(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected method to be POST, got %s", r.Method)
		}
		if r.RequestURI != "/api/jobs" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Write([]byte(`{"accountIdentifier": "jsonAccountID"}`))
	}))
	defer serv.Close()
	client := NewClient(serv.URL)
	out, err := client.Jobs().SubmitJobText(context.TODO(), &SubmitJobTextInput{
		ModelIdentifier: "modelID",
		ModelVersion:    "modelVersion",
		Explain:         true,
		Timeout:         time.Second * 9,
		Inputs: map[string]TextInputItem{
			"input-1": {
				"input-1.1": "input-1.1-value",
				"input-1.2": "input-1.2-value",
			},
			"input-2": {
				"input-2.1": "input-2.1-value",
			},
		},
	})
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if out.Response.AccountIdentifier != "jsonAccountID" {
		t.Errorf("response not parsed")
	}
}

func TestWaitForJobCompletion(t *testing.T) {
	checked := 0
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		checked++
		if checked <= 2 {
			w.Write([]byte(`{"status": "NOT_DONE"}`))
		} else {
			w.Write([]byte(`{"status": "TIMEDOUT"}`))
		}
	}))
	client := NewClient(serv.URL)
	out, err := client.Jobs().WaitForJobCompletion(context.TODO(), &GetJobDetailsInput{}, time.Millisecond)
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if out.Details.Status != "TIMEDOUT" {
		t.Errorf("response not parsed")
	}
	if checked != 3 {
		t.Errorf("checked was %d", checked)
	}
}

func TestWaitForJobCompletionCancelContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	cancel()
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status": "NOT_DONE"}`))
	}))
	client := NewClient(serv.URL)
	_, err := client.Jobs().WaitForJobCompletion(ctx, &GetJobDetailsInput{}, time.Hour)
	if err == nil {
		t.Errorf("Expected error")
	}
	if !strings.Contains(err.Error(), "Wait for job completion was canceled") {
		t.Errorf("Error was different than expected: %v", err)
	}
}

func TestWaitForJobCompletionHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	client := NewClient(serv.URL)
	_, err := client.Jobs().WaitForJobCompletion(context.TODO(), &GetJobDetailsInput{}, time.Millisecond)
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestCancelHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()
	client := NewClient(serv.URL)
	_, err := client.Jobs().CancelJob(context.TODO(), &CancelJobInput{JobIdentifier: "inputID"})
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestCancelJob(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("expected method to be DELETE, got %s", r.Method)
		}
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

func TestGetJobResultsHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()
	client := NewClient(serv.URL)
	_, err := client.Jobs().GetJobResults(context.TODO(), &GetJobResultsInput{JobIdentifier: "inputID"})
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestGetJobResults(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected method to be GET, got %s", r.Method)
		}
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

func TestGetJobFeaturesHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()
	client := NewClient(serv.URL)
	_, err := client.Jobs().GetJobFeatures(context.TODO())
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestGetJobFeatures(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected method to be GET, got %s", r.Method)
		}
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
