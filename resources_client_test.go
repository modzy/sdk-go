package modzy

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProcessingModelsHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Resources().GetProcessingModels(context.TODO())
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestGetProcessingModels(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected method to be GET, got %s", r.Method)
		}
		if r.RequestURI != "/api/resources/processing/models" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Write([]byte(`[{"version": "ver"}]`))
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	out, err := client.Resources().GetProcessingModels(context.TODO())
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if out.Models[0].Version != "ver" {
		t.Errorf("Expected entitlement one, got %s", out.Models[0].Version)
	}
}
