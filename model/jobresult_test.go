package model

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestJobResultUnmarshal(t *testing.T) {
	var jr JobResult
	err := json.Unmarshal([]byte(`{"status":"s","rando":123}`), &jr)
	if err != nil {
		t.Fatalf("error was not nil: %v", err)
	}
	if jr.Status != "s" {
		t.Errorf("Expected jr.Status to be s, got %s", jr.Status)
	}
	got, has := jr.Data["rando"]
	// json parses as float first (not int)
	if !has || got != float64(123) {
		t.Errorf("Expcted jr.Data[rando] to be 123, got %t:%V", has, got)
	}
}

func TestJobResultUnmarshalError(t *testing.T) {
	var jr JobResult
	// call unmarshal straight so it does not do the pre-check in the json code
	err := (&jr).UnmarshalJSON([]byte(`junk`))
	if err == nil {
		t.Fatalf("error was nil")
	}
	if !strings.Contains(err.Error(), "failed to unmarshal base portion") {
		t.Fatalf("error was not expected kind: %v", err)
	}
}
