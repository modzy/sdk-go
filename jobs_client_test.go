package modzy

import (
	"context"
	"fmt"
	"io"
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

func TestSubmitJobEmbeddedHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()
	client := NewClient(serv.URL)
	_, err := client.Jobs().SubmitJobEmbedded(context.TODO(), &SubmitJobEmbeddedInput{})
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestSubmitJobEmbeddedNoDataReaderError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Jobs().SubmitJobEmbedded(context.TODO(), &SubmitJobEmbeddedInput{
		Inputs: map[string]EmbeddedInputItem{
			"input-1": {
				"input-1.1": func() (io.Reader, error) {
					return nil, fmt.Errorf("nope")
				},
			},
		},
	})
	if err == nil {
		t.Errorf("Expected error")
	}
}

type badReader struct{}

func (r *badReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("nope")
}

func TestSubmitJobEmbeddedFailedDataReaderError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Jobs().SubmitJobEmbedded(context.TODO(), &SubmitJobEmbeddedInput{
		Inputs: map[string]EmbeddedInputItem{
			"input-1": {
				"input-1.1": func() (io.Reader, error) {
					return &badReader{}, nil
				},
			},
		},
	})
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestSubmitJobEmbedded(t *testing.T) {
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
	out, err := client.Jobs().SubmitJobEmbedded(context.TODO(), &SubmitJobEmbeddedInput{
		ModelIdentifier: "modelID",
		ModelVersion:    "modelVersion",
		Explain:         true,
		Timeout:         time.Second * 9,
		Inputs: map[string]EmbeddedInputItem{
			"input-1": {
				"input-1.1": URIEncodeString("input-1.1-value", ""),
				"input-1.2": URIEncodeString("input-1.2-value", ""),
			},
			"input-2": {
				"input-2.1": URIEncodeString("input-2.1-value", ""),
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

func TestSubmitJobFileMaxChunkError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Jobs().SubmitJobFile(context.TODO(), &SubmitJobFileInput{})
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestSubmitJobFilePostError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.String() {
		case "/api/jobs/features":
			w.Write([]byte(`{"inputChunkMaximumSize":"1M"}`))
		case "/api/jobs":
			w.WriteHeader(500)
		default:
			t.Fatalf("An unexpected url was requested: %s", r.URL.String())
		}
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Jobs().SubmitJobFile(context.TODO(), &SubmitJobFileInput{})
	if err == nil {
		t.Errorf("Expected error")
	}
	if !strings.Contains(err.Error(), "failed to post open job") {
		t.Errorf("Error was different than expected: %v", err)
	}
}

func TestSubmitJobFileNoReaderFailure(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.String() {
		case "/api/jobs/features":
			w.Write([]byte(`{"inputChunkMaximumSize":"1M"}`))
		case "/api/jobs":
			w.Write([]byte(`{"jobIdentifier":"openJobID"}`))
		case "/api/jobs/openJobID/input-1/input-1.1":
			w.WriteHeader(500)
		case "/api/jobs/openJobID":
			// the job being closed
		default:
			t.Fatalf("An unexpected url was requested: %s", r.URL.String())
		}
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Jobs().SubmitJobFile(context.TODO(), &SubmitJobFileInput{
		Inputs: map[string]FileInputItem{
			"input-1": {
				"input-1.1": func() (io.Reader, error) {
					return nil, fmt.Errorf("nope")
				},
			},
		},
	})
	if err == nil {
		t.Fatalf("Expected error")
	}
	if !strings.Contains(err.Error(), "failed to get data reader for item") {
		t.Errorf("Error was different than expected: %v", err)
	}
}

func TestSubmitJobFileBadReaderFailure(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.String() {
		case "/api/jobs/features":
			w.Write([]byte(`{"inputChunkMaximumSize":"1M"}`))
		case "/api/jobs":
			w.Write([]byte(`{"jobIdentifier":"openJobID"}`))
		case "/api/jobs/openJobID/input-1/input-1.1":
			w.WriteHeader(500)
		case "/api/jobs/openJobID":
			// the job being closed
		default:
			t.Fatalf("An unexpected url was requested: %s", r.URL.String())
		}
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Jobs().SubmitJobFile(context.TODO(), &SubmitJobFileInput{
		Inputs: map[string]FileInputItem{
			"input-1": {
				"input-1.1": func() (io.Reader, error) {
					return &badReader{}, nil
				},
			},
		},
	})
	if err == nil {
		t.Fatalf("Expected error")
	}
	if !strings.Contains(err.Error(), "failed reading a chunk of data") {
		t.Errorf("Error was different than expected: %v", err)
	}
}

func TestSubmitJobFileChunkPostFailure(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.String() {
		case "/api/jobs/features":
			w.Write([]byte(`{"inputChunkMaximumSize":"1M"}`))
		case "/api/jobs":
			w.Write([]byte(`{"jobIdentifier":"openJobID"}`))
		case "/api/jobs/openJobID/input-1/input-1.1":
			w.WriteHeader(500)
		case "/api/jobs/openJobID":
			// the job being closed
		default:
			t.Fatalf("An unexpected url was requested: %s", r.URL.String())
		}
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Jobs().SubmitJobFile(context.TODO(), &SubmitJobFileInput{
		Inputs: map[string]FileInputItem{
			"input-1": {
				"input-1.1": ChunkReader(strings.NewReader("abc")),
			},
		},
	})
	if err == nil {
		t.Fatalf("Expected error")
	}
	if !strings.Contains(err.Error(), "failed to post a chunk of data") {
		t.Errorf("Error was different than expected: %v", err)
	}
}

func TestSubmitJobFileCloseAfterChunkPostChunkPostFailure(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.String() {
		case "/api/jobs/features":
			w.Write([]byte(`{"inputChunkMaximumSize":"1M"}`))
		case "/api/jobs":
			w.Write([]byte(`{"jobIdentifier":"openJobID"}`))
		case "/api/jobs/openJobID/input-1/input-1.1":
			// post was fine
		case "/api/jobs/openJobID/close":
			w.WriteHeader(500)
		default:
			t.Fatalf("An unexpected url was requested: %s", r.URL.String())
		}
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Jobs().SubmitJobFile(context.TODO(), &SubmitJobFileInput{
		Inputs: map[string]FileInputItem{
			"input-1": {
				"input-1.1": ChunkReader(strings.NewReader("abc")),
			},
		},
	})
	if err == nil {
		t.Fatalf("Expected error")
	}
	if !strings.Contains(err.Error(), "failed to close open job after successfully") {
		t.Errorf("Error was different than expected: %v", err)
	}
}

func TestSubmitJobFile(t *testing.T) {
	chunks := 0
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.String() {
		case "/api/jobs/features":
			w.Write([]byte(`{"inputChunkMaximumSize":"1M"}`))
		case "/api/jobs":
			w.Write([]byte(`{"jobIdentifier":"openJobID"}`))
		case "/api/jobs/openJobID/input-1/input-1.1":
			chunks++
			// post a chunk is good
		case "/api/jobs/openJobID/close":
			// final close is fine
		default:
			t.Fatalf("An unexpected url was requested: %s", r.URL.String())
		}
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	out, err := client.Jobs().SubmitJobFile(context.TODO(), &SubmitJobFileInput{
		ChunkSize: 1,
		Inputs: map[string]FileInputItem{
			"input-1": {
				"input-1.1": ChunkReader(strings.NewReader("abc")),
			},
		},
	})
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if out.Response.JobIdentifier != "openJobID" {
		t.Errorf("response not parsed")
	}
	if chunks != 3 {
		t.Errorf("Expected 3 chunks, got %d", chunks)
	}
}

func TestBadChunkSizeError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.String() {
		case "/api/jobs/features":
			w.Write([]byte(`{"inputChunkMaximumSize":"junk"}`))
		default:
			t.Fatalf("An unexpected url was requested: %s", r.URL.String())
		}
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Jobs().SubmitJobFile(context.TODO(), &SubmitJobFileInput{})
	if err == nil {
		t.Fatalf("Expected error")
	}
	if !strings.Contains(err.Error(), "failed to parse InputChunk") {
		t.Errorf("Error was different than expected: %v", err)
	}
}

func TestDefaultChunkCoverage(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.String() {
		case "/api/jobs/features":
			w.Write([]byte(`{"inputChunkMaximumSize":"0"}`))
		}
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Jobs().SubmitJobFile(context.TODO(), &SubmitJobFileInput{})
	if err == nil {
		t.Fatalf("Expected error")
	}
	// this is purely coverage -- not testing that the resulting chunk is actually 1MB
}

func TestSubmitJobS3HTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Jobs().SubmitJobS3(context.TODO(), &SubmitJobS3Input{})
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestSubmitJobS3NoKeyDefinitionReaderError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Jobs().SubmitJobS3(context.TODO(), &SubmitJobS3Input{
		Inputs: map[string]S3InputItem{
			"input-1": {
				"input-1.1": func() (*S3KeyDefinition, error) {
					return nil, fmt.Errorf("nope")
				},
			},
		},
	})
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestSubmitJobS3(t *testing.T) {
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
	out, err := client.Jobs().SubmitJobS3(context.TODO(), &SubmitJobS3Input{
		ModelIdentifier: "modelID",
		ModelVersion:    "modelVersion",
		Explain:         true,
		Timeout:         time.Second * 9,
		Inputs: map[string]S3InputItem{
			"input-1": {
				"input-1.1": S3Key("bucket", "key"),
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

func TestSubmitJobJDBCHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Jobs().SubmitJobJDBC(context.TODO(), &SubmitJobJDBCInput{})
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestSubmitJobJDBC(t *testing.T) {
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
	out, err := client.Jobs().SubmitJobJDBC(context.TODO(), &SubmitJobJDBCInput{
		ModelIdentifier:   "modelID",
		ModelVersion:      "modelVersion",
		Explain:           true,
		Timeout:           time.Second * 9,
		JDBCConnectionURL: "jdbcURL",
		DatabaseUsername:  "username",
		DatabasePassword:  "password",
		Query:             "query",
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
		if r.Method != "GET" {
			t.Errorf("expected method to be GET, got %s", r.Method)
		}
		if r.RequestURI != "/api/jobs/waitingOn" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}

		checked++
		if checked <= 2 {
			w.Write([]byte(`{"status": "NOT_DONE"}`))
		} else {
			w.Write([]byte(`{"status": "TIMEDOUT"}`))
		}
	}))
	client := NewClient(serv.URL)
	out, err := client.Jobs().WaitForJobCompletion(context.TODO(), &WaitForJobCompletionInput{
		JobIdentifier: "waitingOn",
	}, time.Millisecond)
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
	_, err := client.Jobs().WaitForJobCompletion(ctx, &WaitForJobCompletionInput{}, time.Hour)
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
	_, err := client.Jobs().WaitForJobCompletion(context.TODO(), &WaitForJobCompletionInput{}, time.Millisecond)
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestWaitForJobCompletionJobIsOpenError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status": "OPEN"}`))
	}))
	client := NewClient(serv.URL)
	_, err := client.Jobs().WaitForJobCompletion(context.TODO(), &WaitForJobCompletionInput{}, time.Millisecond)
	if err == nil {
		t.Errorf("Expected error")
	}
	if !strings.Contains(err.Error(), "Job is currently OPEN") {
		t.Errorf("Error was different than expected: %v", err)
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
