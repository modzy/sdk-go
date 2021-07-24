package modzy

import (
	"context"
	"fmt"
	"time"

	"github.com/modzy/go-sdk/model"
)

type JobsClient interface {
	GetJobDetails(ctx context.Context, input *GetJobDetailsInput) (*GetJobDetailsOutput, error)
	ListJobsHistory(ctx context.Context, input *ListJobsHistoryInput) (*ListJobsHistoryOutput, error)
	SubmitJobText(ctx context.Context, input *SubmitJobTextInput) (*SubmitJobTextOutput, error)
	SubmitJobEmbedded(ctx context.Context, input *SubmitJobEmbeddedInput) (*SubmitJobEmbeddedOutput, error)
	// SubmitJobFile(ctx context.Context, input *SubmitJobFileInput) (*SubmitJobFileOutput, error)
	// SubmitJobS3(ctx context.Context, input *SubmitJobS3Input) (*SubmitJobS3Output, error)
	// SubmitJobJDBC(ctx context.Context, input *SubmitJobJDBCInput) (*SubmitJobJDBCOutput, error)
	WaitForJobCompletion(ctx context.Context, input *WaitForJobCompletionInput, pollInterval time.Duration) (*GetJobDetailsOutput, error)
	CancelJob(ctx context.Context, input *CancelJobInput) (*CancelJobOutput, error)
	GetJobResults(ctx context.Context, input *GetJobResultsInput) (*GetJobResultsOutput, error)
	GetJobFeatures(ctx context.Context) (*GetJobFeaturesOutput, error)
}

type standardJobsClient struct {
	baseClient *standardClient
}

var _ JobsClient = &standardJobsClient{}

func (c *standardJobsClient) GetJobDetails(ctx context.Context, input *GetJobDetailsInput) (*GetJobDetailsOutput, error) {
	var out model.JobDetails
	url := fmt.Sprintf("/api/jobs/%s", input.JobIdentifier)
	_, err := c.baseClient.requestor.Get(ctx, url, &out)
	if err != nil {
		return nil, err
	}

	return &GetJobDetailsOutput{
		Details: out,
	}, nil
}

func (c *standardJobsClient) ListJobsHistory(ctx context.Context, input *ListJobsHistoryInput) (*ListJobsHistoryOutput, error) {
	input.Paging = input.Paging.withDefaults()

	var items []model.JobSummary
	url := "/api/jobs/history"
	_, links, err := c.baseClient.requestor.List(ctx, url, input.Paging, &items)
	if err != nil {
		return nil, err
	}

	// decide if we have a next page (the next link is not always accurate?)
	var nextPage *ListJobsHistoryInput
	if _, hasNextLink := links["next"]; len(items) == input.Paging.PerPage && hasNextLink {
		nextPage = &ListJobsHistoryInput{
			Paging: input.Paging.Next(),
		}
	}

	return &ListJobsHistoryOutput{
		Jobs:     items,
		NextPage: nextPage,
	}, nil
}

func (c *standardJobsClient) SubmitJobText(ctx context.Context, input *SubmitJobTextInput) (*SubmitJobTextOutput, error) {

	toPostSources := map[string]model.TextInputItem{}
	for k, v := range input.Inputs {
		input := map[string]string{}
		for innerK, innerV := range v {
			input[innerK] = innerV
		}
		toPostSources[k] = input
	}

	toPost := model.SubmitTextJob{
		Model: model.SubmitJobModelInfo{
			Identifier: input.ModelIdentifier,
			Version:    input.ModelVersion,
		},
		Explain: input.Explain,
		Timeout: int(input.Timeout / time.Millisecond),
		Input: model.TextInput{
			Type:    "text",
			Sources: toPostSources,
		},
	}

	var response model.SubmitJobResponse

	url := "/api/jobs"
	_, err := c.baseClient.requestor.Post(ctx, url, toPost, &response)
	if err != nil {
		return nil, err
	}

	return &SubmitJobTextOutput{
		Response:   response,
		JobActions: NewJobActions(c.baseClient, response.JobIdentifier),
	}, nil
}

func (c *standardJobsClient) SubmitJobEmbedded(ctx context.Context, input *SubmitJobEmbeddedInput) (*SubmitJobEmbeddedOutput, error) {

	toPostSources := map[string]model.TextInputItem{}
	for k, v := range input.Inputs {
		input := map[string]string{}
		for innerK, innerV := range v {
			input[innerK] = innerV
		}
		toPostSources[k] = input
	}

	toPost := model.SubmitTextJob{
		Model: model.SubmitJobModelInfo{
			Identifier: input.ModelIdentifier,
			Version:    input.ModelVersion,
		},
		Explain: input.Explain,
		Timeout: int(input.Timeout / time.Millisecond),
		Input: model.TextInput{
			Type:    "embedded",
			Sources: toPostSources,
		},
	}

	var response model.SubmitJobResponse

	url := "/api/jobs"
	_, err := c.baseClient.requestor.Post(ctx, url, toPost, &response)
	if err != nil {
		return nil, err
	}

	return &SubmitJobEmbeddedOutput{
		Response:   response,
		JobActions: NewJobActions(c.baseClient, response.JobIdentifier),
	}, nil
}

// WaitForJobCompletion will wait until the provided job is done processing.
// The minimum pollInterval is 5 seconds.
// If the provided context is canceled, this wait will error.
func (c *standardJobsClient) WaitForJobCompletion(ctx context.Context, input *WaitForJobCompletionInput, pollInterval time.Duration) (*GetJobDetailsOutput, error) {
	timer := time.NewTimer(pollInterval)

	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("Wait for job completion was canceled due to provided context being canceled")
		case <-timer.C:
			job, err := c.GetJobDetails(ctx, &GetJobDetailsInput{input.JobIdentifier})
			if err != nil {
				return nil, err
			}
			// check
			if job.Details.Status == JobStatusCanceled ||
				job.Details.Status == JobStatusCompleted ||
				job.Details.Status == JobStatusTimedOut {
				// job is done
				return job, nil
			}
			// not done -- wait and try again
			timer.Reset(pollInterval)
		}
	}
}

func (c *standardJobsClient) CancelJob(ctx context.Context, input *CancelJobInput) (*CancelJobOutput, error) {
	var response model.JobDetails

	url := fmt.Sprintf("/api/jobs/%s", input.JobIdentifier)
	_, err := c.baseClient.requestor.Delete(ctx, url, &response)
	if err != nil {
		return nil, err
	}

	return &CancelJobOutput{
		Details: response,
	}, nil
}

func (c *standardJobsClient) GetJobResults(ctx context.Context, input *GetJobResultsInput) (*GetJobResultsOutput, error) {
	var response model.JobResults

	url := fmt.Sprintf("/api/results/%s", input.JobIdentifier)
	_, err := c.baseClient.requestor.Get(ctx, url, &response)
	if err != nil {
		return nil, err
	}

	return &GetJobResultsOutput{
		Results: response,
	}, nil
}

func (c *standardJobsClient) GetJobFeatures(ctx context.Context) (*GetJobFeaturesOutput, error) {
	var response model.JobFeatures

	url := fmt.Sprintf("/api/jobs/features")
	_, err := c.baseClient.requestor.Get(ctx, url, &response)
	if err != nil {
		return nil, err
	}

	return &GetJobFeaturesOutput{
		Features: response,
	}, nil
}
