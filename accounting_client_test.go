// nolint:errcheck
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

func TestGetLicenseHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Accounting().GetLicense(context.TODO())
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestGetLicense(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected method to be GET, got %s", r.Method)
		}
		if r.RequestURI != "/api/license" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Write([]byte(`{"companyName": "cn"}`))
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	out, err := client.Accounting().GetLicense(context.TODO())
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if out.License.CompanyName != "cn" {
		t.Errorf("Expected entitlement one, got %s", out.License.CompanyName)
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

func TestListAccountingUsersHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()
	client := NewClient(serv.URL)
	_, err := client.Accounting().ListAccountingUsers(context.TODO(), (&ListAccountingUsersInput{}).WithPaging(2, 3))
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestListAccountingUsers(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected method to be GET, got %s", r.Method)
		}
		if r.RequestURI != "/api/accounting/users?page=7&per-page=2" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Header().Set("Link", `<https://example>; rel="next"`)
		w.Write([]byte(`[{"email": "jsonID"},{"email": "jsonID2"}]`))
	}))
	defer serv.Close()
	client := NewClient(serv.URL)
	out, err := client.Accounting().ListAccountingUsers(context.TODO(), (&ListAccountingUsersInput{}).WithPaging(2, 7))
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if out.Users[0].Email != "jsonID" {
		t.Errorf("response not parsed")
	}
	if out.NextPage == nil {
		t.Errorf("expected NextPage to have a value")
	}
	if out.NextPage.Paging.Page != 8 {
		t.Errorf("expected NextPage to be next")
	}
}

func TestListProjectsHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()
	client := NewClient(serv.URL)
	_, err := client.Accounting().ListProjects(context.TODO(), (&ListProjectsInput{}).WithPaging(2, 3))
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestListProjects(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected method to be GET, got %s", r.Method)
		}
		if r.RequestURI != "/api/accounting/projects?page=7&per-page=2" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Header().Set("Link", `<https://example>; rel="next"`)
		w.Write([]byte(`[{"name": "jsonID"},{"name": "jsonID2"}]`))
	}))
	defer serv.Close()
	client := NewClient(serv.URL)
	out, err := client.Accounting().ListProjects(context.TODO(), (&ListProjectsInput{}).WithPaging(2, 7))
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if out.Projects[0].Name != "jsonID" {
		t.Errorf("response not parsed")
	}
	if out.NextPage == nil {
		t.Errorf("expected NextPage to have a value")
	}
	if out.NextPage.Paging.Page != 8 {
		t.Errorf("expected NextPage to be next")
	}
}

func TestGetProjectDetailsHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()
	client := NewClient(serv.URL)
	_, err := client.Accounting().GetProjectDetails(context.TODO(), &GetProjectDetailsInput{
		ProjectID: "projid",
	})
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestGetProjectDetails(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected method to be GET, got %s", r.Method)
		}
		if r.RequestURI != "/api/accounting/projects/projid" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Write([]byte(`{"name": "jsonID"}`))
	}))
	defer serv.Close()
	client := NewClient(serv.URL)
	out, err := client.Accounting().GetProjectDetails(context.TODO(), &GetProjectDetailsInput{
		ProjectID: "projid",
	})
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if out.Project.Name != "jsonID" {
		t.Errorf("response not parsed")
	}
}
