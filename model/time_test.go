package model_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/modzy/go-sdk/model"
)

type testTime struct {
	time time.Time
	str  string
}

// .567 seconds == 5.67e+8 nanoseconds
var testTimes = []testTime{
	{time.Date(1000, 5, 15, 22, 23, 24, 5.67e+8, time.UTC), "1000-05-15T22:23:24.567+0000"},
	{time.Date(1000, 5, 15, 22, 23, 24, 5.67e+8, time.UTC), "1000-05-15T22:23:24.567+00:00"},
}

func TestModzyTimeUnmarshal(t *testing.T) {
	for _, testTime := range testTimes {
		var mt model.ModzyTime
		err := json.Unmarshal([]byte(fmt.Sprintf(`"%s"`, testTime.str)), &mt)
		if err != nil {
			t.Fatalf("error was not nil: %v", err)
		}

		if mt.UnixNano() != testTime.time.UnixNano() {
			t.Errorf("want nano %d, got %d", mt.UnixNano(), testTime.time.UnixNano())
		}
	}
}

func TestModzyTimeUnmarshalNull(t *testing.T) {
	var mt model.ModzyTime
	err := json.Unmarshal([]byte(`null`), &mt)
	if err != nil {
		t.Fatalf("error was not nil: %v", err)
	}

	if !mt.IsZero() {
		t.Errorf("wanted zero time, got %s", mt)
	}
}

func TestModzyTimeUnmarshalError(t *testing.T) {
	var mt model.ModzyTime
	// call unmarshal straight so it does not do the pre-check in the json code
	err := (&mt).UnmarshalJSON([]byte(`junk`))
	if err == nil {
		t.Fatalf("error was nil")
	}
	if !strings.Contains(err.Error(), "failed to parse modzy time: junk") {
		t.Fatalf("error was not expected kind: %v", err)
	}
}

func TestModzyTimeMarshal(t *testing.T) {
	// when we marshal to the api, we just output the main format
	mt := model.ModzyTime{Time: testTimes[0].time}
	jsb, err := json.Marshal(mt)
	if err != nil {
		t.Errorf("error was not nil: %v", err)
	}

	want := fmt.Sprintf(`"%s"`, testTimes[0].str)
	got := string(jsb)

	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestModzyTimeMarshalZero(t *testing.T) {
	mt := model.ModzyTime{}
	jsb, err := json.Marshal(mt)
	if err != nil {
		t.Errorf("error was not nil: %v", err)
	}

	js := string(jsb)
	if js != `null` {
		t.Errorf("wanted null, got %s", js)
	}
}
