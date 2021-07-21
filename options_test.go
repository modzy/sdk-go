package modzy

import (
	"net/http"
	"testing"
)

func TestClientOptionWithHTTPClient(t *testing.T) {
	client := &standardClient{
		requestor: &requestor{},
	}
	httpClient := &http.Client{}
	WithHTTPClient(httpClient)(client)
	if client.requestor.httpClient != httpClient {
		t.Errorf("Option did not set httpClient")
	}
}

func TestClientOptionWithHTTPDebugging(t *testing.T) {
	client := &standardClient{
		requestor: &requestor{},
	}
	WithHTTPDebugging(true, false)(client)
	if !client.requestor.requestDebugging {
		t.Errorf("Expected requestDebugging to be true")
	}
	if client.requestor.responseDebugging {
		t.Errorf("Expected responseDebugging to be false")
	}

	client = &standardClient{
		requestor: &requestor{},
	}
	WithHTTPDebugging(false, true)(client)
	if client.requestor.requestDebugging {
		t.Errorf("Expected requestDebugging to be false")
	}
	if !client.requestor.responseDebugging {
		t.Errorf("Expected requestDebugging to be true")
	}
}
