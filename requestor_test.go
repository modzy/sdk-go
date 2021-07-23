package modzy

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

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

func TestList(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("method not GET: %s", r.Method)
		}
		if r.RequestURI != "/the/list?direction=DESC&fand=fand-1%3Bfand-2&for=for-1%2Cfor-2&page=4&per-page=3&sort-by=s1%2Cs2" {
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
		WithFilter(And("fand", "fand-1", "fand-2")).
		WithFilter(Or("for", "for-1", "for-2")).
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
