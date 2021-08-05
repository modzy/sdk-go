// model package contains types as defined by the http api.
package model

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
	SubmittedAt              ModzyTime               `json:"submittedAt"`
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
	CreatedAt                ModzyTime            `json:"createdAt"`
	UpdatedAt                ModzyTime            `json:"updatedAt"`
	SubmittedAt              ModzyTime            `json:"submittedAt"`
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

type JobFeatures struct {
	InputChunkMaximumSize string `json:"inputChunkMaximumSize"`
	MaximumInputChunks    int    `json:"maximumInputChunks"`
	MaximumInputsPerJob   int    `json:"maximumInputsPerJob"`
}

type SubmitJobModelInfo struct {
	Identifier string `json:"identifier"`
	Version    string `json:"version"`
}

type TextInputItem map[string]string

type TextInput struct {
	Type    string                   `json:"type"`
	Sources map[string]TextInputItem `json:"sources"`
}

type EmbeddedInputItem map[string]string

type EmbeddedInput struct {
	Type    string                       `json:"type"`
	Sources map[string]EmbeddedInputItem `json:"sources"`
}

type S3InputItemKey struct {
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}

type S3InputItem map[string]S3InputItemKey

type S3Input struct {
	Type            string                 `json:"type"`
	AccessKeyID     string                 `json:"accessKeyID,omitempty"`
	SecretAccessKey string                 `json:"secretAccessKey,omitempty"`
	Region          string                 `json:"region,omitempty"`
	Sources         map[string]S3InputItem `json:"sources"`
}

type JDBCInput struct {
	Type     string `json:"type"`
	URL      string `json:"url,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Driver   string `json:"driver,omitempty"`
	Query    string `json:"query,omitempty"`
}

type SubmitTextJob struct {
	Model   SubmitJobModelInfo `json:"model"`
	Explain bool               `json:"explain,omitempty"`
	Timeout int                `json:"timeout,omitempty"`
	Input   TextInput          `json:"input,omitempty"`
}

type SubmitEmbeddedJob struct {
	Model   SubmitJobModelInfo `json:"model"`
	Explain bool               `json:"explain,omitempty"`
	Timeout int                `json:"timeout,omitempty"`
	Input   EmbeddedInput      `json:"input,omitempty"`
}

type SubmitChunkedJob struct {
	Model   SubmitJobModelInfo `json:"model"`
	Explain bool               `json:"explain,omitempty"`
	Timeout int                `json:"timeout,omitempty"`
}

type SubmitS3Job struct {
	Model   SubmitJobModelInfo `json:"model"`
	Explain bool               `json:"explain,omitempty"`
	Timeout int                `json:"timeout,omitempty"`
	Input   S3Input            `json:"input,omitempty"`
}

type SubmitJDBCJob struct {
	Model   SubmitJobModelInfo `json:"model"`
	Explain bool               `json:"explain,omitempty"`
	Timeout int                `json:"timeout,omitempty"`
	Input   JDBCInput          `json:"input,omitempty"`
}
