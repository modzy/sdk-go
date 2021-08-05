package model_test

import (
	"encoding/json"
	"testing"

	"github.com/modzy/sdk-go/internal/model"
)

func TestSortedVersions(t *testing.T) {
	var sorted model.SortedVersions
	err := json.Unmarshal([]byte(`["1.0.0","1.0.1"]`), &sorted)
	if err != nil {
		t.Fatalf("Expected no error: %v", err)
	}

	if len(sorted) != 2 {
		t.Errorf("sort not correct length")
	}
	if sorted[0] != "1.0.1" {
		t.Errorf("sort not correct")
	}
	if sorted[1] != "1.0.0" {
		t.Errorf("sort not correct")
	}
}

func TestSortedVersionsErr(t *testing.T) {
	sorted := &model.SortedVersions{}
	err := sorted.UnmarshalJSON([]byte(`junk`))
	if err == nil {
		t.Fatalf("Expected an error: %v", sorted)
	}
}
