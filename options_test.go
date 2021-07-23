package modzy

import (
	"net/http"
	"testing"
)

func TestClientOptionWithHTTPClient(t *testing.T) {
	client := &standardClient{
		requestor: &stdRequestor{},
	}
	httpClient := &http.Client{}
	WithHTTPClient(httpClient)(client)
	if client.requestor.(*stdRequestor).httpClient != httpClient {
		t.Errorf("Option did not set httpClient")
	}
}

func TestClientOptionWithHTTPDebugging(t *testing.T) {
	client := &standardClient{
		requestor: &stdRequestor{},
	}
	WithHTTPDebugging(true, false)(client)
	if !client.requestor.(*stdRequestor).requestDebugging {
		t.Errorf("Expected requestDebugging to be true")
	}
	if client.requestor.(*stdRequestor).responseDebugging {
		t.Errorf("Expected responseDebugging to be false")
	}

	client = &standardClient{
		requestor: &stdRequestor{},
	}
	WithHTTPDebugging(false, true)(client)
	if client.requestor.(*stdRequestor).requestDebugging {
		t.Errorf("Expected requestDebugging to be false")
	}
	if !client.requestor.(*stdRequestor).responseDebugging {
		t.Errorf("Expected requestDebugging to be true")
	}
}
