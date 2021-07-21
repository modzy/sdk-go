package model

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"
)

const testTimeString = "1000-05-15T22:23:24.567+0000"

var jsonTestTimeString = fmt.Sprintf(`"%s"`, testTimeString)

// .567 seconds == 5.67e+8 nanoseconds
var testTime = time.Date(1000, 5, 15, 22, 23, 24, 5.67e+8, time.UTC)

func TestModzyTimeUnmarshal(t *testing.T) {
	var mt ModzyTime
	err := json.Unmarshal([]byte(jsonTestTimeString), &mt)
	if err != nil {
		t.Fatalf("error was not nil: %v", err)
	}

	if mt.UnixNano() != testTime.UnixNano() {
		t.Errorf("want nano %d, got %d", mt.UnixNano(), testTime.UnixNano())
	}
}

func TestModzyTimeUnmarshalNull(t *testing.T) {
	var mt ModzyTime
	err := json.Unmarshal([]byte(`null`), &mt)
	if err != nil {
		t.Fatalf("error was not nil: %v", err)
	}

	if !mt.IsZero() {
		t.Errorf("wanted zero time, got %s", mt)
	}
}

func TestModzyTimeUnmarshalError(t *testing.T) {
	var mt ModzyTime
	// call unmarshal straight so it does not do the pre-check in the json code
	err := (&mt).UnmarshalJSON([]byte(`junk`))
	if err == nil {
		t.Fatalf("error was nil")
	}
	if !strings.Contains(err.Error(), "parse using custom format") {
		t.Fatalf("error was not expected kind: %v", err)
	}
}

func TestModzyTimeMarshal(t *testing.T) {
	mt := ModzyTime{Time: testTime}
	jsb, err := json.Marshal(mt)
	if err != nil {
		t.Errorf("error was not nil: %v", err)
	}

	js := string(jsb)
	if js != jsonTestTimeString {
		t.Errorf("want %s, got %s", jsonTestTimeString, js)
	}
}

func TestModzyTimeMarshalZero(t *testing.T) {
	mt := ModzyTime{}
	jsb, err := json.Marshal(mt)
	if err != nil {
		t.Errorf("error was not nil: %v", err)
	}

	js := string(jsb)
	if js != `null` {
		t.Errorf("wanted null, got %s", js)
	}
}
