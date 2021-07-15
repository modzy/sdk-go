package types

import "time"

type Model struct {
	ModelID             string     `json:"modelId,omitempty"`
	LatestVersion       string     `json:"latestVersion,omitempty"`
	LatestActiveVersion string     `json:"latestActiveVersion,omitempty"`
	Versions            []string   `json:"versions,omitempty"`
	Author              string     `json:"author,omitempty"`
	Name                string     `json:"name,omitempty"`
	Description         string     `json:"description,omitempty"`
	Permalink           string     `json:"permalink,omitempty"`
	Features            []Features `json:"features,omitempty"`
	IsActive            bool       `json:"isActive,omitempty"`
	IsExpired           bool       `json:"isExpired,omitempty"`
	IsRecommended       bool       `json:"isRecommended,omitempty"`
	IsCommercial        bool       `json:"isCommercial,omitempty"`
	Tags                []Tags     `json:"tags,omitempty"`
	Images              []Images   `json:"images,omitempty"`
	LastActiveDateTime  time.Time  `json:"lastActiveDateTime,omitempty"`
	CreatedByEmail      string     `json:"createdByEmail,omitempty"`
	CreatedByFullName   string     `json:"createdByFullName,omitempty"`
}
type Features struct {
	Identifier  string `json:"identifier,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
type Tags struct {
	Identifier    string `json:"identifier,omitempty"`
	Name          string `json:"name,omitempty"`
	DataType      string `json:"dataType,omitempty"`
	IsCategorical bool   `json:"isCategorical,omitempty"`
}
type Images struct {
	URL          string `json:"url,omitempty"`
	Caption      string `json:"caption,omitempty"`
	RelationType string `json:"relationType,omitempty"`
}

type ModelDetail struct {
	Version               string         `json:"version,omitempty"`
	CreatedAt             time.Time      `json:"createdAt,omitempty"`
	UpdatedAt             time.Time      `json:"updatedAt,omitempty"`
	InputValidationSchema string         `json:"inputValidationSchema,omitempty"`
	Timeout               Timeout        `json:"timeout,omitempty"`
	Requirement           Requirement    `json:"requirement,omitempty"`
	ContainerImage        ContainerImage `json:"containerImage,omitempty"`
	LoadStatus            LoadStatus     `json:"loadStatus,omitempty"`
	RunStatus             RunStatus      `json:"runStatus,omitempty"`
	Inputs                []Inputs       `json:"inputs,omitempty"`
	Outputs               []Outputs      `json:"outputs,omitempty"`
	Statistics            []interface{}  `json:"statistics,omitempty"`
	IsActive              bool           `json:"isActive,omitempty"`
	IsAvailable           bool           `json:"isAvailable,omitempty"`
	Status                string         `json:"status,omitempty"`
	Model                 Model          `json:"model,omitempty"`
	Processing            Processing     `json:"processing,omitempty"`
}

type Timeout struct {
	Status int `json:"status,omitempty"`
	Run    int `json:"run,omitempty"`
}
type Requirement struct {
	GpuUnits     int    `json:"gpuUnits,omitempty"`
	CPUAmount    string `json:"cpuAmount,omitempty"`
	MemoryAmount string `json:"memoryAmount,omitempty"`
}
type ContainerImage struct {
	UploadStatus        string `json:"uploadStatus,omitempty"`
	LoadStatus          string `json:"loadStatus,omitempty"`
	UploadPercentage    int    `json:"uploadPercentage,omitempty"`
	LoadPercentage      int    `json:"loadPercentage,omitempty"`
	ContainerImageSize  int    `json:"containerImageSize,omitempty"`
	RegistryHost        string `json:"registryHost,omitempty"`
	RepositoryNamespace string `json:"repositoryNamespace,omitempty"`
	RepositoryName      string `json:"repositoryName,omitempty"`
}
type LoadStatus struct {
	Step       int    `json:"step,omitempty"`
	StepName   string `json:"stepName,omitempty"`
	Percentage int    `json:"percentage,omitempty"`
}
type Result struct {
	Model          interface{} `json:"model,omitempty"`
	JobID          interface{} `json:"jobId,omitempty"`
	DatasourceName interface{} `json:"datasourceName,omitempty"`
	Status         string      `json:"status,omitempty"`
	Engine         string      `json:"engine,omitempty"`
	Error          interface{} `json:"error,omitempty"`
	StartTime      string      `json:"startTime,omitempty"`
	EndTime        string      `json:"endTime,omitempty"`
	UpdateTime     string      `json:"updateTime,omitempty"`
	ElapsedTime    int         `json:"elapsedTime,omitempty"`
	ResultsJSON    interface{} `json:"results.json,omitempty"`
}
type RunStatus struct {
	Step       int    `json:"step,omitempty"`
	StepName   string `json:"stepName,omitempty"`
	Percentage int    `json:"percentage,omitempty"`
	Result     Result `json:"result,omitempty"`
}
type Inputs struct {
	Name               string `json:"name,omitempty"`
	AcceptedMediaTypes string `json:"acceptedMediaTypes,omitempty"`
	MaximumSize        int    `json:"maximumSize,omitempty"`
	Description        string `json:"description,omitempty"`
}
type Outputs struct {
	Name        string `json:"name,omitempty"`
	MediaType   string `json:"mediaType,omitempty"`
	MaximumSize int    `json:"maximumSize,omitempty"`
	Description string `json:"description,omitempty"`
}

type Processing struct {
	MinimumParallelCapacity int `json:"minimumParallelCapacity,omitempty"`
	MaximumParallelCapacity int `json:"maximumParallelCapacity,omitempty"`
}

type ModelRequestOpt struct {
	// These are the available query parameters as documented in https://models.modzy.com/docs/marketplace/models/retrieve-models
	Author             string    `url:"author,omitempty"`
	CreatedByEmail     string    `url:"createdByEmail,omitempty"`
	ModelId            string    `url:"modelId,omitempty"`
	Name               string    `url:"name,omitempty"`
	Description        string    `url:"description,omitempty"`
	IsActive           bool      `url:"isActive,omitempty"`
	IsExpired          bool      `url:"isExpired,omitempty"`
	IsCommercial       bool      `url:"isCommercial,omitempty"`
	IsRecommended      bool      `url:"isRecommended,omitempty"`
	LastActiveDateTime time.Time `url:"lastActiveDateTime,omitempty"`
}

type AllModelRequestOpt struct {
	// These are the available query parameters as documented in https://models.modzy.com/docs/marketplace/models/retrieve-all-models-versions
	Search                  string    `url:"search,omitempty"`
	Version                 string    `url:"version,omitempty"`
	CreatedAt               time.Time `url:"createdAt,omitempty"`
	CreatedBy               string    `url:"createdBy,omitempty"`
	Status                  string    `url:"status,omitempty"`
	IsAvailable             bool      `url:"isAvailable,omitempty"`
	UpdatedAt               time.Time `url:"updatedAt,omitempty"`
	IsActive                string    `url:"isActive,omitempty"`
	IsExperimental          bool      `url:"isExperimental,omitempty"`
	ModelId                 string    `url:"model.modelId,omitempty"`
	ModelAuthor             string    `url:"mode.author,omitempty"`
	ModelCreatedByEmail     string    `url:"model.createdByEmail,omitempty"`
	ModelCreatedByFullName  string    `url:"model.createdByFullName,omitempty"`
	ModelName               string    `url:"model.name,omitempty"`
	ModelDescription        string    `url:"model.description,omitempty"`
	ModelIsActive           bool      `url:"model.isActive,omitempty"`
	ModelIsExpired          bool      `url:"model.isExpired,omitempty"`
	ModelIsCommercial       bool      `url:"model.isCommercial,omitempty"`
	ModelIsReccommended     bool      `url:"model.isRecommended,omitempty"`
	ModelLastActiveDateTime time.Time `url:"model.lastActiveDateTime,omitempty"`
	PerPage                 int       `url:"per-page,omitempty"`
	SortBy                  string    `url:"sort-by,omitempty"`
	Direction               string    `url:"direction,omitempty"`
}

type RequestError struct {
	StatusCode           int    `json:"statusCode"`
	Status               string `json:"status"`
	Message              string `json:"message"`
	TechnicalInformation string `json:"technicalInformation"`
	ReportErrorURL       string `json:"reportErrorUrl"`
}

type ModelError struct {
	Timestamp string `json:"timestamp"`
	Path      string `json:"path"`
	Status    int    `json:"status"`
	Error     string `json:"error"`
	Message   string `json:"message"`
}

type ModelVersion struct {
	Version string `json:"version"`
}
