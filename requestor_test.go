// nolint:errcheck
package modzy

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var notParsablePath = string([]byte{0x7f})

func TestGet(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("method not GET: %s", r.Method)
		}
		if r.RequestURI != "/the/path" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Write([]byte(`"some-response"`))
	}))
	defer serv.Close()

	requestor := &requestor{
		baseURL:    serv.URL,
		httpClient: defaultHTTPClient,
	}

	var into string
	_, err := requestor.Get(context.TODO(), "/the/path", &into)
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if into != "some-response" {
		t.Errorf("response not parsed into: %s", into)
	}
}

func TestListBadPath(t *testing.T) {
	requestor := &requestor{}
	_, _, err := requestor.List(context.TODO(), notParsablePath, PagingInput{}, nil)
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestList(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("method not GET: %s", r.Method)
		}
		if r.RequestURI != "/the/list?andfield=a1%3Ba2&direction=DESC&orfield=f1%2Cf2&page=4&per-page=3&sort-by=s1%2Cs2" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Header().Set("Link", `<https://example>; rel="next"`)
		w.Write([]byte(`"some-response"`))
	}))
	defer serv.Close()

	requestor := &requestor{
		baseURL:    serv.URL,
		httpClient: defaultHTTPClient,
	}

	paging := NewPaging(3, 4).
		WithFilterAnd("andfield", "a1", "a2").
		WithFilterOr("orfield", "f1", "f2").
		WithSort(SortDirectionDescending, "s1", "s2")

	var into string
	_, links, err := requestor.List(context.TODO(), "/the/list", paging, &into)
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if into != "some-response" {
		t.Errorf("response not parsed into: %s", into)
	}
	if next, ok := links["next"]; !ok || next.URI != "https://example" {
		t.Errorf("links did not parse correctly: %+v", links)
	}
}

func TestPost(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method not POST: %s", r.Method)
		}
		if r.RequestURI != "/the/post/path" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		var received string
		_ = json.NewDecoder(r.Body).Decode(&received)
		if received != "post-data" {
			t.Errorf("received payload not correct")
		}
		w.Write([]byte(`"some-response"`))
	}))
	defer serv.Close()

	requestor := &requestor{
		baseURL:    serv.URL,
		httpClient: defaultHTTPClient,
	}

	toPost := "post-data"
	var into string
	_, err := requestor.Post(context.TODO(), "/the/post/path", toPost, &into)
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if into != "some-response" {
		t.Errorf("response not parsed into: %s", into)
	}
}

func TestDelete(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("method not DELETE: %s", r.Method)
		}
		if r.RequestURI != "/the/del/path" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Write([]byte(`"some-response"`))
	}))
	defer serv.Close()

	requestor := &requestor{
		baseURL:    serv.URL,
		httpClient: defaultHTTPClient,
	}

	var into string
	_, err := requestor.Delete(context.TODO(), "/the/del/path", &into)
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if into != "some-response" {
		t.Errorf("response not parsed into: %s", into)
	}
}

func TestPostMultipart(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method not POST: %s", r.Method)
		}
		if r.RequestURI != "/the/post/path" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		if !strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data; boundary=") {
			t.Errorf("content-type header not correct: %s", r.Header.Get("Content-Type"))
		}

		received, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("test read did not work: %v", err)
		}

		if !strings.Contains(string(received), "f1data") {
			t.Errorf("received payload does not contain f1 data")
		}
		if !strings.Contains(string(received), "f2data") {
			t.Errorf("received payload does not contain f2 data")
		}

		w.Write([]byte(`"some-response"`))
	}))
	defer serv.Close()

	requestor := &requestor{
		baseURL:          serv.URL,
		httpClient:       defaultHTTPClient,
		requestDebugging: true, // for coverage
	}

	var into string
	_, err := requestor.PostMultipart(context.TODO(), "/the/post/path", map[string]io.Reader{
		"f1": strings.NewReader("f1data"),
		"f2": strings.NewReader("f2data"),
	}, &into)
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if into != "some-response" {
		t.Errorf("response not parsed into: %s", into)
	}
}

func TestPostMultipartBadReader(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`"some-response"`))
	}))
	defer serv.Close()

	requestor := &requestor{
		baseURL:    serv.URL,
		httpClient: defaultHTTPClient,
	}

	var into string
	_, err := requestor.PostMultipart(context.TODO(), "/the/post/path", map[string]io.Reader{
		"f1": &badReader{},
	}, &into)
	if err == nil {
		t.Fatalf("expected an error")
	}
	if !strings.Contains(err.Error(), "bad-reader-nope") {
		t.Errorf("error was not the type expected: %v", err)
	}
}

func TestExecuteBodyStructCannotMarshal(t *testing.T) {
	requestor := &requestor{}
	_, err := requestor.execute(context.TODO(), "", "", cannotMarshal{}, nil, nil)
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestExecuteCannotBuildRequest(t *testing.T) {
	requestor := &requestor{}
	_, err := requestor.execute(context.TODO(), notParsablePath, "GET", nil, nil, nil)
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestExecuteFailureToExecute(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("stop now")
	}))
	defer serv.Close()

	requestor := &requestor{
		httpClient: defaultHTTPClient,
	}
	_, err := requestor.execute(context.TODO(), "http://nope", "GET", nil, nil, nil)
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestExecuteWithModzyError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
		w.Write([]byte(`{"statusCode": 123}`))
	}))
	defer serv.Close()

	requestor := &requestor{
		httpClient: defaultHTTPClient,
	}
	_, err := requestor.execute(context.TODO(), serv.URL, "GET", nil, nil, nil)
	if modzyErr, is := err.(*ModzyHTTPError); !is {
		t.Errorf("expected modzy error: %v", err)
	} else {
		if modzyErr.StatusCode != 123 {
			t.Errorf("expected parsed modzy code of 123")
		}
	}
}

func TestExecuteBad200Content(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`bad-json`))
	}))
	defer serv.Close()

	requestor := &requestor{
		httpClient: defaultHTTPClient,
	}
	var into map[string]interface{}
	_, err := requestor.execute(context.TODO(), serv.URL, "GET", nil, &into, nil)
	if err == nil {
		t.Fatalf("expected error")
	}
	if !strings.Contains(err.Error(), "invalid character 'b'") {
		t.Errorf("Did not get the expected error: %v", err)
	}
}

func TestExecute204(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(204)
	}))
	defer serv.Close()

	requestor := &requestor{
		httpClient: defaultHTTPClient,
	}
	_, err := requestor.execute(context.TODO(), serv.URL, "GET", nil, nil, nil)
	if err != nil {
		t.Errorf("did not expect error: %v", err)
	}
}

func TestAuthorizationDecorator(t *testing.T) {
	// just make sure it runs (coverage), I am not asserting logs
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Test") != "Decoration" {
			t.Errorf("authorization decorator not ran")
		}
		w.WriteHeader(204)
	}))
	defer serv.Close()

	requestor := &requestor{
		httpClient: defaultHTTPClient,
		authorizationDecorator: func(toDecorate *http.Request) *http.Request {
			toDecorate.Header.Set("Test", "Decoration")
			return toDecorate
		},
	}
	resp, err := requestor.execute(context.TODO(), serv.URL, "GET", nil, nil, nil)
	if err != nil {
		t.Errorf("did not expect error: %v", err)
	}
	if resp.StatusCode != 204 {
		t.Errorf("Did not hit test server")
	}
}

func TestExecuteWithDebugging(t *testing.T) {
	// just make sure it runs (coverage), I am not asserting logs
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(204)
	}))
	defer serv.Close()

	requestor := &requestor{
		httpClient:        defaultHTTPClient,
		requestDebugging:  true,
		responseDebugging: true,
	}
	_, err := requestor.execute(context.TODO(), serv.URL, "GET", nil, nil, nil)
	if err != nil {
		t.Errorf("did not expect error: %v", err)
	}
}

type cannotMarshal struct{}

func (cannotMarshal) MarshalJSON() ([]byte, error) {
	return nil, fmt.Errorf("will-not-marshal")
}
