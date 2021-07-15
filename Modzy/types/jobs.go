package types

type AllJobSummary struct {
	Jobidentifier string          `json:"jobIdentifier,omitempty"`
	Status        string          `json:"status,omitempty"`
	Model         JobModelSummary `json:"model,omitempty"`
}
type JobModelSummary struct {
	Identifier string `json:"identifier,omitempty"`
	Version    string `json:"version,omitempty"`
	Name       string `json:"name,omitempty"`
}

type AllJobsParams struct {
	PerPage int    `url:"per-page,omitempty"`
	Status  string `url:"status,omitempty"`
}

type JobDetails struct {
	Jobidentifier     string          `json:"jobIdentifier,omitempty"`
	Submittedby       string          `json:"submittedBy,omitempty"`
	Accountidentifier string          `json:"accountIdentifier,omitempty"`
	Model             JobModelSummary `json:"model,omitempty"`
	Status            string          `json:"status,omitempty"`
	Createdat         string          `json:"createdAt,omitempty"`
	Updatedat         string          `json:"updatedAt,omitempty"`
	Submittedat       string          `json:"submittedAt,omitempty"`
	Total             int             `json:"total,omitempty"`
	Pending           int             `json:"pending,omitempty"`
	Completed         int             `json:"completed,omitempty"`
	Failed            int             `json:"failed,omitempty"`
	Elapsedtime       int             `json:"elapsedTime,omitempty"`
	Queuetime         int             `json:"queueTime,omitempty"`
	User              User            `json:"user,omitempty"`
	Jobinputs         []Jobinputs     `json:"jobInputs,omitempty"`
	Explain           bool            `json:"explain,omitempty"`
	Team              Team            `json:"team,omitempty"`
}

type Accesskeys struct {
	Prefix    string `json:"prefix,omitempty"`
	Isdefault bool   `json:"isDefault,omitempty"`
}

type Jobinputs struct {
	Identifier string `json:"identifier,omitempty"`
}
