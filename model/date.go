package model

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const (
	DateFormat = "2006-01-02"
)

type ModzyDate struct {
	time.Time
}

var _ json.Unmarshaler = &ModzyDate{}
var _ json.Marshaler = &ModzyDate{}

func (mt *ModzyDate) UnmarshalJSON(b []byte) error {
	clean := strings.Trim(string(b), `"`)
	if clean == "null" {
		return nil
	}

	for _, df := range []string{DateFormat} {
		if t, err := time.Parse(df, clean); err == nil {
			*mt = ModzyDate{Time: t}
			return nil
		}
	}

	return fmt.Errorf("failed to parse modzy time: %s", clean)
}

func (mt ModzyDate) MarshalJSON() ([]byte, error) {
	if mt.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, mt)), nil
}

func (mt ModzyDate) String() string {
	return mt.Time.Format(DateFormat)
}
