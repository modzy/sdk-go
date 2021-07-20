// model package contains types as defined by the http api.
package model

import (
	"time"
)

type ModelIdentifier struct {
	Identifier string `json:"identifier"`
	Version    string `json:"version"`
}

type ModelNamedIdentifier struct {
	Identifier string `json:"identifier"`
	Version    string `json:"version"`
	Name       string `json:"name"`
}

type JobInputIdentifier struct {
	Identifier string `json:"identifier"`
}

type User struct {
	Identifier         string `json:"identifier"`
	FirstName          string `json:"firstName"`
	LastName           string `json:"lastName"`
	Email              string `json:"email"`
	ExternalIdentifier string `json:"externalIdentifier"`
	PictureURL         string `json:"pictureURL"`
	Status             string `json:"status"`
	Title              string `json:"title"`
	// TODO: type this
	AccessKeys []interface{} `json:"accessKeys"`
}

type Team struct {
	Identifier string `json:"identifier"`
}

type SubmitJobResponseInputs struct {
	Identifier []string `json:"identifier"`
}

type SubmitJobResponse struct {
	Model                    ModelIdentifier         `json:"model"`
	Explain                  bool                    `json:"explain"`
	Timeout                  int                     `json:"timeout"`
	AccountIdentifier        string                  `json:"accountIdentifier"`
	TotalInputs              int                     `json:"totalInputs"`
	JobIdentifier            string                  `json:"jobIdentifier"`
	JobType                  string                  `json:"jobType"`
	AccessKey                string                  `json:"accessKey"`
	JobInputs                SubmitJobResponseInputs `json:"jobInputs"`
	InputByteAmount          int                     `json:"inputByteAmount"`
	SubmittedAt              time.Time               `json:"submittedAt"`
	ImageClassificationModel bool                    `json:"imageClassificationModel"`
}

type JobSummary struct {
	JobIdentifier string               `json:"jobIdentifier"`
	Status        string               `json:"status"`
	Model         ModelNamedIdentifier `json:"model"`
}

type JobDetails struct {
	JobIdentifier            string               `json:"jobIdentifier"`
	SubmittedBy              string               `json:"submittedBy"`
	AccountIdentifier        string               `json:"accountIdentifier"`
	Model                    ModelNamedIdentifier `json:"model"`
	Status                   string               `json:"status"`
	CreatedAt                time.Time            `json:"createdAt"`
	UpdatedAt                time.Time            `json:"updatedAt"`
	SubmittedAt              time.Time            `json:"submittedAt"`
	Total                    int                  `json:"total"`
	Pending                  int                  `json:"pending"`
	Completed                int                  `json:"completed"`
	Failed                   int                  `json:"failed"`
	ElapsedTime              int                  `json:"elapsedTime"`
	QueueTime                int                  `json:"queueTime"`
	User                     User                 `json:"user"`
	JobInputs                []JobInputIdentifier `json:"jobInputs"`
	Explain                  bool                 `json:"explain"`
	HoursDeleteInput         int                  `json:"hoursDeleteInput"`
	Team                     Team                 `json:"team"`
	ImageClassificationModel bool                 `json:"imageClassificationModel"`
}

// type JobResults struct {
// 	JobIdentifier string `json:"jobIdentifier"`
// 	Total         int    `json:"total"`
// 	Completed     int    `json:"completed"`
// 	Failed        int    `json:"failed"`
// 	Finished      bool   `json:"finished"`
// 	SubmittedBy   string `json:"submittedByKey"`
// }

// type JobResult struct {
// 	Status      string          `json:"status"`
// 	Engine      string          `json:"engine"`
// 	StartTime   time.Time       `json:"startTime"`
// 	UpdateTime  time.Time       `json:"updateTime"`
// 	EndTime     time.Time       `json:"endTime"`
// 	ElapsedTime int             `json:"elapsedTime"`
// 	ResultsJSON json.RawMessage `json:"results.json"`
// }
