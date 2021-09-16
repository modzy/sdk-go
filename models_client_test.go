// nolint:errcheck
package modzy

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetModelVersionDetailsHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Models().GetModelVersionDetails(context.TODO(), &GetModelVersionDetailsInput{
		ModelID: "modelID",
		Version: "modelVersion",
	})
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestGetModelVersionDetails(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected method to be GET, got %s", r.Method)
		}
		if r.RequestURI != "/api/models/modelID/versions/modelVersion" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Write([]byte(`{"sampleOutput": "some-output"}`))
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	out, err := client.Models().GetModelVersionDetails(context.TODO(), &GetModelVersionDetailsInput{
		ModelID: "modelID",
		Version: "modelVersion",
	})
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if out.Details.SampleOutput != "some-output" {
		t.Errorf("response not parsed")
	}
}

func TestGetLatestModelsHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Models().GetLatestModels(context.TODO())
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestGetLatestModels(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected method to be GET, got %s", r.Method)
		}
		if r.RequestURI != "/api/models/latest" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Write([]byte(`[{"name": "some-name"}]`))
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	out, err := client.Models().GetLatestModels(context.TODO())
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if out.Models[0].Name != "some-name" {
		t.Errorf("response not parsed")
	}
}

func TestGetMinimumEnginesHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Models().GetMinimumEngines(context.TODO())
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestGetMinimumEngines(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected method to be GET, got %s", r.Method)
		}
		if r.RequestURI != "/api/models/processing-engines" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Write([]byte(`{"minimumProcessingEnginesSum": 123}`))
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	out, err := client.Models().GetMinimumEngines(context.TODO())
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if out.Details.MinimumProcessingEnginesSum != 123 {
		t.Errorf("response not parsed")
	}
}

func TestGetModelDetailsHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Models().GetModelDetails(context.TODO(), &GetModelDetailsInput{
		ModelID: "modelID",
	})
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestGetModelDetails(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected method to be GET, got %s", r.Method)
		}
		if r.RequestURI != "/api/models/modelID" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Write([]byte(`{"description": "some-description"}`))
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	out, err := client.Models().GetModelDetails(context.TODO(), &GetModelDetailsInput{
		ModelID: "modelID",
	})
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if out.Details.Description != "some-description" {
		t.Errorf("response not parsed")
	}
}

func TestGetRelatedModelsHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Models().GetRelatedModels(context.TODO(), &GetRelatedModelsInput{
		ModelID: "modelID",
	})
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestGetRelatedModels(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected method to be GET, got %s", r.Method)
		}
		if r.RequestURI != "/api/models/modelID/related-models" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Write([]byte(`[{"name": "some-name"}]`))
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	out, err := client.Models().GetRelatedModels(context.TODO(), &GetRelatedModelsInput{
		ModelID: "modelID",
	})
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if out.RelatedModels[0].Name != "some-name" {
		t.Errorf("response not parsed")
	}
}

func TestListModelsHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()
	client := NewClient(serv.URL)
	_, err := client.Models().ListModels(context.TODO(), (&ListModelsInput{}).WithPaging(2, 3))
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestListModels(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected method to be GET, got %s", r.Method)
		}
		if r.RequestURI != "/api/models?page=7&per-page=2" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Header().Set("Link", `<https://example>; rel="next"`)
		w.Write([]byte(`[{"modelID": "jsonID"},{"modelID": "jsonID2"}]`))
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	out, err := client.Models().ListModels(context.TODO(), (&ListModelsInput{}).WithPaging(2, 7))
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if out.Models[0].ID != "jsonID" {
		t.Errorf("response not parsed")
	}
	if out.NextPage == nil {
		t.Errorf("expected NextPage to have a value")
	}
	if out.NextPage.Paging.Page != 8 {
		t.Errorf("expected NextPage to be next")
	}
}

func TestGetTagsHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Models().GetTags(context.TODO())
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestGetTags(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected method to be GET, got %s", r.Method)
		}
		if r.RequestURI != "/api/models/tags" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Write([]byte(`[{"name": "tag-name"}]`))
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	out, err := client.Models().GetTags(context.TODO())
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if out.Tags[0].Name != "tag-name" {
		t.Errorf("response not parsed")
	}
}

func TestGetTagModelsHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Models().GetTagModels(context.TODO(), &GetTagModelsInput{
		TagIDs: []string{"t1", "t2"},
	})
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestGetTagModels(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected method to be GET, got %s", r.Method)
		}
		if r.RequestURI != "/api/models/tags/t1,t2" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Write([]byte(`{"tags":[{"name": "some-name"}]}`))
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	out, err := client.Models().GetTagModels(context.TODO(), &GetTagModelsInput{
		TagIDs: []string{"t1", "t2"},
	})
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if out.Tags[0].Name != "some-name" {
		t.Errorf("response not parsed")
	}
}

func TestGetModelDetailsByNameHTTPErrorList(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Models().GetModelDetailsByName(context.TODO(), &GetModelDetailsByNameInput{
		Name: "modelName",
	})
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestGetModelDetailsByNameHTTPErrorDetails(t *testing.T) {
	calls := 0
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		switch calls {
		case 1:
			w.Write([]byte(`[{"modelId": "some-model-id"}]`))
		case 2:
			w.WriteHeader(500)
		}
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Models().GetModelDetailsByName(context.TODO(), &GetModelDetailsByNameInput{
		Name: "modelName",
	})
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestGetModelDetailsByNameErrorNoMatch(t *testing.T) {
	calls := 0
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		switch calls {
		case 1:
			w.Write([]byte(`[]`))
		}
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Models().GetModelDetailsByName(context.TODO(), &GetModelDetailsByNameInput{
		Name: "modelName",
	})
	if err != ErrNotFound {
		t.Errorf("Expected NotFound error, got %v", err)
	}
}

func TestGetModelDetailsByName(t *testing.T) {
	calls := 0
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		switch calls {
		case 1:
			if r.Method != "GET" {
				t.Errorf("expected method to be GET, got %s", r.Method)
			}
			if r.RequestURI != "/api/models?name=modelName&page=1&per-page=1" {
				t.Errorf("get url not expected: %s", r.RequestURI)
			}
			w.Write([]byte(`[{"modelId": "some-model-id"}]`))
		case 2:
			if r.Method != "GET" {
				t.Errorf("expected method to be GET, got %s", r.Method)
			}
			if r.RequestURI != "/api/models/some-model-id" {
				t.Errorf("get url not expected: %s", r.RequestURI)
			}
			w.Write([]byte(`{"description": "some-description"}`))
		}
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	out, err := client.Models().GetModelDetailsByName(context.TODO(), &GetModelDetailsByNameInput{
		Name: "modelName",
	})
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if out.Details.Description != "some-description" {
		t.Errorf("response not parsed")
	}
}

func TestListModelVersionsHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Models().ListModelVersions(context.TODO(), (&ListModelVersionsInput{}).WithPaging(2, 3))
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestListModelVersions(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected method to be GET, got %s", r.Method)
		}
		if r.RequestURI != "/api/models/modelID/versions?page=7&per-page=2" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Header().Set("Link", `<https://example>; rel="next"`)
		w.Write([]byte(`[{"version": "some-version"},{"version": "some-version-2"}]`))
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	out, err := client.Models().ListModelVersions(context.TODO(), (&ListModelVersionsInput{
		ModelID: "modelID",
	}).WithPaging(2, 7))
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if out.Versions[0].Version != "some-version" {
		t.Errorf("response not parsed")
	}
	if out.NextPage == nil {
		t.Errorf("expected NextPage to have a value")
	}
	if out.NextPage.Paging.Page != 8 {
		t.Errorf("expected NextPage to be next")
	}
}

func TestUpdateModelProcessingEnginesEntitlementError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Models().UpdateModelProcessingEngines(context.TODO(), &UpdateModelProcessingEnginesInput{})
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestUpdateModelProcessingEnginesPatchError(t *testing.T) {
	calls := 0
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		switch calls {
		case 1:
			w.Write([]byte(`[{"identifier": "no"},{"identifier": "CAN_PATCH_PROCESSING_MODEL_VERSION"}]`))
		case 2:
			w.WriteHeader(500)
		}
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Models().UpdateModelProcessingEngines(context.TODO(), &UpdateModelProcessingEnginesInput{})
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestUpdateModelProcessingEngines(t *testing.T) {
	calls := 0
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		switch calls {
		case 1:
			if r.Method != "GET" {
				t.Errorf("expected method to be GET, got %s", r.Method)
			}
			if r.RequestURI != "/api/accounting/entitlements" {
				t.Errorf("get url not expected: %s", r.RequestURI)
			}
			w.Write([]byte(`[{"identifier": "no"}]`))
		case 2:
			if r.Method != "PATCH" {
				t.Errorf("expected method to be PATCH, got %s", r.Method)
			}
			if r.RequestURI != "/api/models/modelID/versions/version" {
				t.Errorf("get url not expected: %s", r.RequestURI)
			}
			w.Write([]byte(`{"sampleOutput": "some-sample-out"}`))
		}
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	out, err := client.Models().UpdateModelProcessingEngines(context.TODO(), &UpdateModelProcessingEnginesInput{
		ModelID:                 "modelID",
		Version:                 "version",
		MinimumParallelCapacity: 15,
		MaximumParallelCapacity: 86,
	})
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if out.Details.SampleOutput != "some-sample-out" {
		t.Errorf("response not parsed")
	}
}

func TestUpdateModelProcessingEnginesAdmin(t *testing.T) {
	calls := 0
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		switch calls {
		case 1:
			w.Write([]byte(`[{"identifier": "no"},{"identifier": "CAN_PATCH_PROCESSING_MODEL_VERSION"}]`))
		case 2:
			w.Write([]byte(`{"sampleOutput": "some-sample-out"}`))
		}
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Models().UpdateModelProcessingEngines(context.TODO(), &UpdateModelProcessingEnginesInput{
		ModelID:                 "modelID",
		Version:                 "version",
		MinimumParallelCapacity: 15,
		MaximumParallelCapacity: 86,
	})
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
}

func TestGetModelVersionSampleInputHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Models().GetModelVersionSampleInput(context.TODO(), &GetModelVersionSampleInputInput{
		ModelID: "modelID",
		Version: "modelVersion",
	})
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestGetModelVersionSampleInput(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected method to be GET, got %s", r.Method)
		}
		if r.RequestURI != "/api/models/modelID/versions/modelVersion/sample-input" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Write([]byte(`{"giveme": "inputs", "here": true}`))
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	out, err := client.Models().GetModelVersionSampleInput(context.TODO(), &GetModelVersionSampleInputInput{
		ModelID: "modelID",
		Version: "modelVersion",
	})
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if out.Sample != `{"giveme":"inputs","here":true}` {
		t.Errorf("response not parsed: %s", out.Sample)
	}
}

func TestGetModelVersionSampleOutputHTTPError(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	_, err := client.Models().GetModelVersionSampleOutput(context.TODO(), &GetModelVersionSampleOutputInput{
		ModelID: "modelID",
		Version: "modelVersion",
	})
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestGetModelVersionSampleOutput(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected method to be GET, got %s", r.Method)
		}
		if r.RequestURI != "/api/models/modelID/versions/modelVersion/sample-output" {
			t.Errorf("get url not expected: %s", r.RequestURI)
		}
		w.Write([]byte(`{"heressome": "outputs", "right": true}`))
	}))
	defer serv.Close()

	client := NewClient(serv.URL)
	out, err := client.Models().GetModelVersionSampleOutput(context.TODO(), &GetModelVersionSampleOutputInput{
		ModelID: "modelID",
		Version: "modelVersion",
	})
	if err != nil {
		t.Errorf("err not nil: %v", err)
	}
	if out.Sample != `{"heressome":"outputs","right":true}` {
		t.Errorf("response not parsed: %s", out.Sample)
	}
}
