package model

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// 2021-07-20T01:40:11.187+0000
const DateFormat = "2006-01-02T15:04:05.999-0700"

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

	if t, err := time.Parse(DateFormat, clean); err != nil {
		return errors.WithMessage(err, "failed to parse using custom format")
	} else {
		*mt = ModzyTime{Time: t}
	}
	return nil
}

func (mt ModzyTime) MarshalJSON() ([]byte, error) {
	if mt.IsZero() {
		return []byte("null"), nil
	}
	formatted := mt.Time.Format(DateFormat)
	return []byte(fmt.Sprintf(`"%s"`, formatted)), nil
}
