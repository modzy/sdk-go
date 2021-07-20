package modzy

import (
	"net/http"
	"testing"
)

func TestClientOptionWithHTTPClient(t *testing.T) {
	client := &standardClient{}
	httpClient := &http.Client{}
	WithHTTPClient(httpClient)(client)
	if client.httpClient != httpClient {
		t.Errorf("Option did not set httpClient")
	}
}

func TestClientOptionWithHTTPDebugging(t *testing.T) {
	client := &standardClient{}
	WithHTTPDebugging(true, false)(client)
	if !client.requestDebugging {
		t.Errorf("Expected requestDebugging to be true")
	}
	if client.responseDebugging {
		t.Errorf("Expected responseDebugging to be false")
	}

	client = &standardClient{}
	WithHTTPDebugging(false, true)(client)
	if client.requestDebugging {
		t.Errorf("Expected requestDebugging to be false")
	}
	if !client.responseDebugging {
		t.Errorf("Expected requestDebugging to be true")
	}
}
