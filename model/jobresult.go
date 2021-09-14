package model

import (
	"encoding/json"

	"github.com/modzy/sdk-go/internal/impossible"
	"github.com/pkg/errors"
)

type JobResults struct {
	JobIdentifier string    `json:"jobIdentifier"`
	Total         int       `json:"total"`
	Completed     int       `json:"completed"`
	Failed        int       `json:"failed"`
	Finished      bool      `json:"finished"`
	SubmittedBy   string    `json:"submittedByKey"`
	Explained     bool      `json:"explained"`
	SubmittedAt   ModzyTime `json:"submittedAt"`

	JobQueueTime     int `json:"jobQueueTime"`
	JobProcessedTime int `json:"jobProcessedTime"`
	JobElapsedTime   int `json:"jobElapsedTime"`

	// next api version:
	// InitialQueueTime          int       `json:"initialQueueTime"`
	// TotalQueueTime            int       `json:"totalQueueTime"`
	// AverageModelLatency       float64   `json:"averageModelLatency"`
	// TotalModelLatency         float64   `json:"totalModelLatency"`
	// ElapsedTime               float64   `json:"elapsedTime"`
	// StartingResultSummarizing ModzyTime `json:"startingResultSummarizing"`
	// ResultSummarizing         int       `json:"resultSummarizing"`

	Results  map[string]JobResult `json:"results"`
	Failures map[string]JobResult `json:"failures"`
}

type JobResult struct {
	Status      string                 `json:"status"`
	Engine      string                 `json:"engine"`
	Error       string                 `json:"error"`
	StartTime   ModzyTime              `json:"startTime"`
	UpdateTime  ModzyTime              `json:"updateTime"`
	EndTime     ModzyTime              `json:"endTime"`
	ElapsedTime int                    `json:"elapsedTime"`
	Data        map[string]interface{} `json:"data"`
}

var _ json.Unmarshaler = &JobResult{}

// UnmarshalJSON custom unmarshal in order to fill in ResultData with extra fields
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
	impossible.HandleError(json.Unmarshal(b, &extra))

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
