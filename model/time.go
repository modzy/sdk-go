package model

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// There are multiple time formats coming back from the API
const (
	// 2021-07-20T01:40:11.187+0000
	DateFormat = "2006-01-02T15:04:05.999-0700"
	// 2021-07-20T01:40:09.560+00:00
	DateFormat2 = "2006-01-02T15:04:05.999-07:00"
)

type ModzyTime struct {
	time.Time
}

var _ json.Unmarshaler = &ModzyTime{}
var _ json.Marshaler = &ModzyTime{}

func (mt *ModzyTime) UnmarshalJSON(b []byte) error {
	clean := strings.Trim(string(b), `"`)
	if clean == "null" {
		return nil
	}

	for _, df := range []string{DateFormat, DateFormat2} {
		if t, err := time.Parse(df, clean); err == nil {
			*mt = ModzyTime{Time: t}
			return nil
		}
	}

	return fmt.Errorf("failed to parse modzy time: %s", clean)
}

func (mt ModzyTime) MarshalJSON() ([]byte, error) {
	if mt.IsZero() {
		return []byte("null"), nil
	}
	formatted := mt.Time.Format(DateFormat)
	return []byte(fmt.Sprintf(`"%s"`, formatted)), nil
}
