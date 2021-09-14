// nolint:errcheck
package modzy

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/modzy/sdk-go/model"
)

func TestGetAlertsHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Dashboard().GetAlerts(context.TODO(), &GetAlertsInput{})
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestGetAlerts(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected method to be GET, got %s", r.Method)
		}
		if r.RequestURI != "/api/notifications/alerts" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Write([]byte(`[{"type":"some-type"}]`))
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	out, err := client.Dashboard().GetAlerts(context.TODO(), &GetAlertsInput{})
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if out.Alerts[0].Type != "some-type" {
		t.Errorf("response not parsed")
	}
}

func TestGetAlertDetailsHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Dashboard().GetAlertDetails(context.TODO(), &GetAlertDetailsInput{
		Type: "some-type",
	})
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestGetAlertDetails(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected method to be GET, got %s", r.Method)
		}
		if r.RequestURI != "/api/notifications/alerts/some-type" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Write([]byte(`["a","b"]`))
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	out, err := client.Dashboard().GetAlertDetails(context.TODO(), &GetAlertDetailsInput{
		Type: "some-type",
	})
	if err != nil {
		t.Fatalf("err not nil: %v", err)
	}
	if out.Type != "some-type" {
		t.Errorf("response not parsed")
	}
	if out.Entities[1] != "b" {
		t.Errorf("response not parsed")
	}
}

func TestGetDataProcessedHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Dashboard().GetDataProcessed(context.TODO(), &GetDataProcessedInput{})
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestGetDataProcessed(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected method to be GET, got %s", r.Method)
		}
		if r.RequestURI != "/api/metrics/data-processed?user-identifier=uid" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Write([]byte(`{"dataProcessed":{"percentage":0.123}}`))
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	out, err := client.Dashboard().GetDataProcessed(context.TODO(), &GetDataProcessedInput{
		UserIdentifier: "uid",
	})
	if err != nil {
		t.Fatalf("err not nil: %v", err)
	}
	if out.Summary.Percentage != 0.123 {
		t.Errorf("response not parsed")
	}
}

func TestGetPredictionsMadeHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Dashboard().GetPredictionsMade(context.TODO(), &GetPredictionsMadeInput{})
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestGetPredictionsMade(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected method to be GET, got %s", r.Method)
		}
		if r.RequestURI != "/api/metrics/predictions-made?begin-date=2001-02-03" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Write([]byte(`{"predictionsMade":{"percentage":0.123}}`))
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	out, err := client.Dashboard().GetPredictionsMade(context.TODO(), &GetPredictionsMadeInput{
		BeginDate: model.ModzyDate{Time: time.Date(2001, 02, 03, 0, 0, 0, 0, time.UTC)},
	})
	if err != nil {
		t.Fatalf("err not nil: %v", err)
	}
	if out.Summary.Percentage != 0.123 {
		t.Errorf("response not parsed: %f", out.Summary.Percentage)
	}
}

func TestGetActiveUsersHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Dashboard().GetActiveUsers(context.TODO(), &GetActiveUsersInput{})
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestGetActiveUsers(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected method to be GET, got %s", r.Method)
		}
		if r.RequestURI != "/api/metrics/active-users?end-date=2001-02-03" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Write([]byte(`[{"firstName":"fName"}]`))
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	out, err := client.Dashboard().GetActiveUsers(context.TODO(), &GetActiveUsersInput{
		EndDate: model.ModzyDate{Time: time.Date(2001, 02, 03, 0, 0, 0, 0, time.UTC)},
	})
	if err != nil {
		t.Fatalf("err not nil: %v", err)
	}
	if out.Users[0].FirstName != "fName" {
		t.Errorf("response not parsed: %s", out.Users[0].FirstName)
	}
}

func TestGetActiveModelsHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Dashboard().GetActiveModels(context.TODO(), &GetActiveModelsInput{})
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestGetActiveModels(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected method to be GET, got %s", r.Method)
		}
		if r.RequestURI != "/api/metrics/active-models?access-key=access&model-identifier=model&team-identifier=team" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Write([]byte(`[{"modelVersion":"v123"}]`))
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	out, err := client.Dashboard().GetActiveModels(context.TODO(), &GetActiveModelsInput{
		AccessKeyPrefix: "access",
		ModelIdentifier: "model",
		TeamIdentifier:  "team",
	})
	if err != nil {
		t.Fatalf("err not nil: %v", err)
	}
	if out.Models[0].Version != "v123" {
		t.Errorf("response not parsed")
	}
}

func TestGetPrometheusMetricHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Dashboard().GetPrometheusMetric(context.TODO(), &GetPrometheusMetricInput{})
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestGetPrometheusMetric(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected method to be GET, got %s", r.Method)
		}
		if r.RequestURI != "/api/metrics/prometheus/cpu-available?begin-date=2001-02-03" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Write([]byte(`{
			"data":{
				"result":[{
					"values":[
						["junk"],
						[1618290060, "0.23115264514500705"],
						["notatime", "0.23115264514500705"]
					]
				}]
			}
		}`))
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	out, err := client.Dashboard().GetPrometheusMetric(context.TODO(), &GetPrometheusMetricInput{
		Metric:    PrometheusMetricType(PrometheusMetricTypeCPUAvailable),
		BeginDate: model.ModzyDate{Time: time.Date(2001, 02, 03, 0, 0, 0, 0, time.UTC)},
	})
	if err != nil {
		t.Fatalf("err not nil: %v", err)
	}
	gotTime := out.Values[0].Time.UTC()
	if gotTime != time.Date(2021, 04, 13, 05, 01, 00, 0, time.UTC) {
		t.Errorf("response not parsed: %+v", gotTime)
	}
	if out.Values[0].Value != "0.23115264514500705" {
		t.Errorf("response not parsed: %v", out.Values[0].Value)
	}
}
