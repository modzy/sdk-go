package model_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/modzy/sdk-go/model"
)

type testDate struct {
	time time.Time
	str  string
}

var testDates = []testDate{
	{time.Date(1000, 5, 15, 0, 0, 0, 0, time.UTC), "1000-05-15"},
}

func TestModzyDateUnmarshal(t *testing.T) {
	for _, testDate := range testDates {
		var mt model.ModzyDate
		err := json.Unmarshal([]byte(fmt.Sprintf(`"%s"`, testDate.str)), &mt)
		if err != nil {
			t.Fatalf("error was not nil: %v", err)
		}

		if mt.UnixNano() != testDate.time.UnixNano() {
			t.Errorf("want nano %d, got %d", mt.UnixNano(), testDate.time.UnixNano())
		}
	}
}

func TestModzyDateUnmarshalNull(t *testing.T) {
	var mt model.ModzyDate
	err := json.Unmarshal([]byte(`null`), &mt)
	if err != nil {
		t.Fatalf("error was not nil: %v", err)
	}

	if !mt.IsZero() {
		t.Errorf("wanted zero time, got %s", mt)
	}
}

func TestModzyDateUnmarshalError(t *testing.T) {
	var mt model.ModzyDate
	// call unmarshal straight so it does not do the pre-check in the json code
	err := (&mt).UnmarshalJSON([]byte(`junk`))
	if err == nil {
		t.Fatalf("error was nil")
	}
	if !strings.Contains(err.Error(), "failed to parse modzy time: junk") {
		t.Fatalf("error was not expected kind: %v", err)
	}
}

func TestModzyDateMarshal(t *testing.T) {
	// when we marshal to the api, we just output the main format
	mt := model.ModzyDate{Time: testDates[0].time}
	jsb, err := json.Marshal(mt)
	if err != nil {
		t.Errorf("error was not nil: %v", err)
	}

	want := fmt.Sprintf(`"%s"`, testDates[0].str)
	got := string(jsb)

	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestModzyDateMarshalZero(t *testing.T) {
	mt := model.ModzyDate{}
	jsb, err := json.Marshal(mt)
	if err != nil {
		t.Errorf("error was not nil: %v", err)
	}

	js := string(jsb)
	if js != `null` {
		t.Errorf("wanted null, got %s", js)
	}
}
