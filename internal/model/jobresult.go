package model

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type JobResults struct {
	JobIdentifier string               `json:"jobIdentifier"`
	Total         int                  `json:"total"`
	Completed     int                  `json:"completed"`
	Failed        int                  `json:"failed"`
	Finished      bool                 `json:"finished"`
	SubmittedBy   string               `json:"submittedByKey"`
	Results       map[string]JobResult `json:"results"`
}

type JobResult struct {
	Status      string                 `json:"status"`
	Engine      string                 `json:"engine"`
	StartTime   ModzyTime              `json:"startTime"`
	UpdateTime  ModzyTime              `json:"updateTime"`
	EndTime     ModzyTime              `json:"endTime"`
	ElapsedTime int                    `json:"elapsedTime"`
	Data        map[string]interface{} `json:"data"`
}

var _ json.Unmarshaler = &JobResult{}

// Custom unmarshal in order to fill in ResultData with extra fields
func (j *JobResult) UnmarshalJSON(b []byte) error {
	// make a junk type so that we don't recurse into this unmarshal
	type innerJobResult JobResult

	// get the known fields
	var inner innerJobResult
	if err := json.Unmarshal(b, &inner); err != nil {
		return errors.WithMessage(err, "failed to unmarshal base portion")
	}

	// get the extra fields
	extra := make(map[string]interface{})
	// I cannot think of a way this could error after the previous marshal passes...
	json.Unmarshal(b, &extra)

	// do not repeat the known fields within the extras
	for _, f := range []string{
		"status",
		"engine",
		"startTime",
		"updateTime",
		"endTime",
		"elapsedTime",
	} {
		delete(extra, f)
	}
	inner.Data = extra

	*j = JobResult(inner)
	return nil
}
