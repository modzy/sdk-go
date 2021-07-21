package modzy

import (
	"time"

	"github.com/modzy/go-sdk/model"
)

// GetJobDetailsInput -
type GetJobDetailsInput struct {
	JobIdentifier string
}

// GetJobDetailsOutput -
type GetJobDetailsOutput struct {
	Details model.JobDetails `json:"details"`
}

// // ListJobsInput -
// type ListJobsInput struct {
// 	Paging PagingInput
// }

// type ListJobsFilterField string

// const (
// 	ListJobsFilterFieldStartDate ListJobsFilterField = "startDate"
// 	ListJobsFilterFieldEndDate   ListJobsFilterField = "endDate"
// 	ListJobsFilterFieldStatus    ListJobsFilterField = "status"
// 	ListJobsFilterFieldModel     ListJobsFilterField = "model"
// 	ListJobsFilterFieldUser      ListJobsFilterField = "user"
// 	ListJobsFilterFieldAccessKey ListJobsFilterField = "accessKey" // I see "prefix" in the docs -- what does that mean?
// )

// func (i *ListJobsInput) WithPaging(perPage int, page int) *ListJobsInput {
// 	i.Paging = NewPaging(perPage, page)
// 	return i
// }

// func (i *ListJobsInput) WithFilterAnd(field ListJobsFilterField, values ...string) *ListJobsInput {
// 	i.Paging = i.Paging.WithFilter(And(string(field), values...))
// 	return i
// }

// func (i *ListJobsInput) WithFilterOr(field ListJobsFilterField, values ...string) *ListJobsInput {
// 	i.Paging = i.Paging.WithFilter(Or(string(field), values...))
// 	return i
// }

// // ListJobsOutput -
// type ListJobsOutput struct {
// 	Jobs     []model.JobSummary `json:"jobs"`
// 	NextPage *ListJobsInput     `json:"nextPage"`
// }

type ListJobsHistoryInput struct {
	Paging PagingInput
}

type ListJobsHistoryFilterField string

const (
	ListJobsHistoryFilterFieldStartDate ListJobsHistoryFilterField = "startDate"
	ListJobsHistoryFilterFieldEndDate   ListJobsHistoryFilterField = "endDate"
	ListJobsHistoryFilterFieldStatus    ListJobsHistoryFilterField = "status"
	ListJobsHistoryFilterFieldModel     ListJobsHistoryFilterField = "model"
	ListJobsHistoryFilterFieldUser      ListJobsHistoryFilterField = "user"
	ListJobsHistoryFilterFieldAccessKey ListJobsHistoryFilterField = "accessKey" // I see "prefix" in the docs -- what does that mean?
)

type ListJobsHistorySortField string

const (
	ListJobsHistorySortIdentifier    ListJobsHistorySortField = "identifier"
	ListJobsHistorySortSubmittedBy   ListJobsHistorySortField = "submittedBy"
	ListJobsHistorySortSubmittedJobs ListJobsHistorySortField = "submittedJobs"
	ListJobsHistorySortStatus        ListJobsHistorySortField = "status"
	ListJobsHistorySortCreatedAt     ListJobsHistorySortField = "createdAt"
	ListJobsHistorySortUpdatedAt     ListJobsHistorySortField = "updatedAt"
	ListJobsHistorySortSubmittedAt   ListJobsHistorySortField = "submittedAt"
	ListJobsHistorySortTotal         ListJobsHistorySortField = "total"
	ListJobsHistorySortCompleted     ListJobsHistorySortField = "completed"
	ListJobsHistorySortFail          ListJobsHistorySortField = "fail"
	ListJobsHistorySortModel         ListJobsHistorySortField = "model"
)

func (i *ListJobsHistoryInput) WithPaging(perPage int, page int) *ListJobsHistoryInput {
	i.Paging = NewPaging(perPage, page)
	return i
}

func (i *ListJobsHistoryInput) WithFilterAnd(field ListJobsHistoryFilterField, values ...string) *ListJobsHistoryInput {
	i.Paging = i.Paging.WithFilter(And(string(field), values...))
	return i
}

func (i *ListJobsHistoryInput) WithFilterOr(field ListJobsHistoryFilterField, values ...string) *ListJobsHistoryInput {
	i.Paging = i.Paging.WithFilter(Or(string(field), values...))
	return i
}

func (i ListJobsHistoryInput) WithSort(sortDirection SortDirection, sortBy ...ListJobsHistorySortField) *ListJobsHistoryInput {
	sorts := []string{}
	for _, s := range sortBy {
		sorts = append(sorts, string(s))
	}
	i.Paging.SortDirection = sortDirection
	i.Paging.SortBy = sorts
	return &i
}

type ListJobsHistoryOutput struct {
	Jobs     []model.JobSummary    `json:"jobs"`
	NextPage *ListJobsHistoryInput `json:"nextPage"`
}

type TextInputItem map[string]string

type SubmitJobTextInput struct {
	ModelIdentifier string
	ModelVersion    string
	Explain         bool
	Timeout         time.Duration
	Inputs          map[string]TextInputItem
}

type SubmitJobTextOutput struct {
	Response model.SubmitJobResponse
	JobActions
}

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
