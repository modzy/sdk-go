package modzy

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEntitlementsHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Accounting().GetEntitlements(context.TODO())
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestGetEntitlements(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected method to be GET, got %s", r.Method)
		}
		if r.RequestURI != "/api/accounting/entitlements" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Write([]byte(`[{"identifier": "one"}, {"identifier": "two"}]`))
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	out, err := client.Accounting().GetEntitlements(context.TODO())
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if out.Entitlements[0].Identifier != "one" {
		t.Errorf("Expected entitlement one, got %s", out.Entitlements[0].Identifier)
	}
	if out.Entitlements[1].Identifier != "two" {
		t.Errorf("Expected entitlement two, got %s", out.Entitlements[1].Identifier)
	}
}

func TestHasEntitlementHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Accounting().HasEntitlement(context.TODO(), "ENTITLEMENT")
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestHasEntitlement(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected method to be GET, got %s", r.Method)
		}
		if r.RequestURI != "/api/accounting/entitlements" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Write([]byte(`[{"identifier": "one"}, {"identifier": "two"}]`))
	}))
	defer serv.Close()

	client := NewClient(serv.URL)

	out, err := client.Accounting().HasEntitlement(context.TODO(), "do-not-have-this")
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if out != false {
		t.Errorf("Expected to not have entitlment")
	}

	out, err = client.Accounting().HasEntitlement(context.TODO(), "one")
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if out != true {
		t.Errorf("Expected to have entitlment")
	}
}
