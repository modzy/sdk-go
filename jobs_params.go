package modzy

import (
	"time"

	"github.com/modzy/sdk-go/internal/model"
)

// GetJobDetailsInput -
type GetJobDetailsInput struct {
	JobIdentifier string
}

type WaitForJobCompletionInput GetJobDetailsInput

// GetJobDetailsOutput -
type GetJobDetailsOutput struct {
	Details model.JobDetails `json:"details"`
}

type ListJobsHistoryInput struct {
	Paging PagingInput
}

// ListJobsHistoryFilterField are known field names that can be used when filtering the jobs history
type ListJobsHistoryFilterField string

const (
	ListJobsHistoryFilterFieldStartDate ListJobsHistoryFilterField = "startDate"
	ListJobsHistoryFilterFieldEndDate   ListJobsHistoryFilterField = "endDate"
	ListJobsHistoryFilterFieldStatus    ListJobsHistoryFilterField = "status"
	ListJobsHistoryFilterFieldModel     ListJobsHistoryFilterField = "model"
	ListJobsHistoryFilterFieldUser      ListJobsHistoryFilterField = "user"
	ListJobsHistoryFilterFieldAccessKey ListJobsHistoryFilterField = "accessKey" // I see "prefix" in the docs -- what does that mean?
)

// ListJobsHistoryFilterField are known field names that can be used when sorting the jobs history
type ListJobsHistorySortField string

const (
	ListJobsHistorySortFieldIdentifier    ListJobsHistorySortField = "identifier"
	ListJobsHistorySortFieldSubmittedBy   ListJobsHistorySortField = "submittedBy"
	ListJobsHistorySortFieldSubmittedJobs ListJobsHistorySortField = "submittedJobs"
	ListJobsHistorySortFieldStatus        ListJobsHistorySortField = "status"
	ListJobsHistorySortFieldCreatedAt     ListJobsHistorySortField = "createdAt"
	ListJobsHistorySortFieldUpdatedAt     ListJobsHistorySortField = "updatedAt"
	ListJobsHistorySortFieldSubmittedAt   ListJobsHistorySortField = "submittedAt"
	ListJobsHistorySortFieldTotal         ListJobsHistorySortField = "total"
	ListJobsHistorySortFieldCompleted     ListJobsHistorySortField = "completed"
	ListJobsHistorySortFieldFail          ListJobsHistorySortField = "fail"
	ListJobsHistorySortFieldModel         ListJobsHistorySortField = "model"
)

func (i *ListJobsHistoryInput) WithPaging(perPage int, page int) *ListJobsHistoryInput {
	i.Paging = NewPaging(perPage, page)
	return i
}

func (i *ListJobsHistoryInput) WithFilter(field ListJobsHistoryFilterField, value string) *ListJobsHistoryInput {
	i.Paging = i.Paging.WithFilterAnd(string(field), value)
	return i
}

func (i *ListJobsHistoryInput) WithFilterAnd(field ListJobsHistoryFilterField, values ...string) *ListJobsHistoryInput {
	i.Paging = i.Paging.WithFilterAnd(string(field), values...)
	return i
}

func (i *ListJobsHistoryInput) WithFilterOr(field ListJobsHistoryFilterField, values ...string) *ListJobsHistoryInput {
	i.Paging = i.Paging.WithFilterOr(string(field), values...)
	return i
}

func (i *ListJobsHistoryInput) WithSort(sortDirection SortDirection, sortBy ...ListJobsHistorySortField) *ListJobsHistoryInput {
	sorts := []string{}
	for _, s := range sortBy {
		sorts = append(sorts, string(s))
	}
	i.Paging.SortDirection = sortDirection
	i.Paging.SortBy = sorts
	return i
}

type ListJobsHistoryOutput struct {
	Jobs     []model.JobSummary    `json:"jobs"`
	NextPage *ListJobsHistoryInput `json:"nextPage"`
}

type SubmitJobOutput struct {
	Response model.SubmitJobResponse
	JobActions
}

type TextInputItem map[string]string

type SubmitJobTextInput struct {
	ModelIdentifier string
	ModelVersion    string
	Explain         bool
	Timeout         time.Duration
	Inputs          map[string]TextInputItem
}

type SubmitJobTextOutput = SubmitJobOutput

type EmbeddedInputItem map[string]URIEncodable

type SubmitJobEmbeddedInput struct {
	ModelIdentifier string
	ModelVersion    string
	Explain         bool
	Timeout         time.Duration
	Inputs          map[string]EmbeddedInputItem
}

type SubmitJobEmbeddedOutput = SubmitJobOutput

type FileInputItem map[string]FileInputEncodable

type SubmitJobFileInput struct {
	ModelIdentifier string
	ModelVersion    string
	Explain         bool
	Timeout         time.Duration
	// ChunkSize (in bytes) is optional -- if not provided it will use the configured MaximumChunkSize.
	// If provided it will be limited to the configured maximum;
	ChunkSize int
	Inputs    map[string]FileInputItem
}

type SubmitJobFileOutput = SubmitJobOutput

type S3InputItem map[string]S3Inputable

type SubmitJobS3Input struct {
	ModelIdentifier    string
	ModelVersion       string
	Explain            bool
	Timeout            time.Duration
	AWSAccessKeyID     string
	AWSSecretAccessKey string
	AWSRegion          string
	Inputs             map[string]S3InputItem
}

type SubmitJobS3Output = SubmitJobOutput

type SubmitJobJDBCInput struct {
	ModelIdentifier   string
	ModelVersion      string
	Explain           bool
	Timeout           time.Duration
	JDBCConnectionURL string
	DatabaseUsername  string
	DatabasePassword  string
	Query             string
}

type SubmitJobJDBCOutput = SubmitJobOutput

type CancelJobInput struct {
	JobIdentifier string `json:"jobIdentifier"`
}

type CancelJobOutput struct {
	Details model.JobDetails
}

type GetJobResultsInput struct {
	JobIdentifier string
}

type GetJobResultsOutput struct {
	Results model.JobResults
}

type GetJobFeaturesInput struct {
}

type GetJobFeaturesOutput struct {
	Features model.JobFeatures `json:"features"`
}
