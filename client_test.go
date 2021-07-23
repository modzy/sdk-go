package modzy

import (
	"net/http"
	"testing"
)

func TestClientWithAPIKey(t *testing.T) {
	c := (&standardClient{
		requestor: &requestor{},
	})
	c.WithAPIKey("k")
	req := &http.Request{
		Header: http.Header{},
	}
	decorated := c.requestor.authorizationDecorator(req)

	got := decorated.Header.Get("Authorization")
	if got != "ApiKey k" {
		t.Errorf("Expected ApiKey k, got %s", got)
	}
}

func TestClientWithWithTeamKey(t *testing.T) {
	c := (&standardClient{
		requestor: &requestor{},
	})
	c.WithTeamKey("team", "teamKey")
	req := &http.Request{
		Header: http.Header{},
	}
	decorated := c.requestor.authorizationDecorator(req)

	if decorated.Header.Get("Authorization") != "Bearer teamKey" {
		t.Errorf("Expected Bearer teamKey, got %s", decorated.Header.Get("Authorization"))
	}

	if decorated.Header.Get("Modzy-Team-Id") != "team" {
		t.Errorf("Expected team, got %s", decorated.Header.Get("Modzy-Team-Id"))
	}
}

func TestClientWithOptions(t *testing.T) {
	c := (&standardClient{
		requestor: &requestor{},
	})
	c.WithOptions(WithHTTPDebugging(true, false))
	if c.requestor.requestDebugging != true {
		t.Errorf("Expected requestDebugging to be true, was not")
	}
}

func TestNewClient(t *testing.T) {
	c := NewClient("baseURL", WithHTTPDebugging(true, false))
	standardC := c.(*standardClient)
	if standardC.requestor.baseURL != "baseURL" {
		t.Errorf("baseURL not set")
	}
	if standardC.requestor.httpClient != defaultHTTPClient {
		t.Errorf("default httpClient not set")
	}
	if standardC.jobsClient == nil {
		t.Errorf("no jobs client")
	}
	if standardC.modelsClient == nil {
		t.Errorf("no models client")
	}
}

func TestClientJobs(t *testing.T) {
	jobsClient := &standardJobsClient{}
	c := (&standardClient{
		jobsClient: jobsClient,
	})
	if c.Jobs() != jobsClient {
		t.Errorf("jobsClient did not return")
	}
}

func TestClientModels(t *testing.T) {
	modelsClient := &standardModelsClient{}
	c := (&standardClient{
		modelsClient: modelsClient,
	})
	if c.Models() != modelsClient {
		t.Errorf("modelsClient did not return")
	}
}
