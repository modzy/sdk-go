package modzy

import (
	"context"
	"time"
)

// JobActions are a collection shortcut methods when dealing with a single job.  This is returned from a SubmitJob function call.
type JobActions interface {
	// GetDetails will get the details of the job
	GetDetails(ctx context.Context) (*GetJobDetailsOutput, error)
	// GetModelDetails will get the details of the related version of its model
	GetModelDetails(ctx context.Context) (*GetModelVersionDetailsOutput, error)
	// WaitForCompletion will block until the job is no longer processing
	WaitForCompletion(ctx context.Context, pollInterval time.Duration) (*GetJobDetailsOutput, error)
	// Cancel will cancel the a still-processing job
	Cancel(ctx context.Context) (*CancelJobOutput, error)
	// GetResults will get the results of a complete job
	GetResults(ctx context.Context) (*GetJobResultsOutput, error)
}

type standardJobActions struct {
	client        Client
	jobIdentifier string
}

func NewJobActions(client Client, jobIdentifier string) JobActions {
	return &standardJobActions{
		client:        client,
		jobIdentifier: jobIdentifier,
	}
}

func (j *standardJobActions) GetDetails(ctx context.Context) (*GetJobDetailsOutput, error) {
	return j.client.Jobs().GetJobDetails(ctx, &GetJobDetailsInput{
		JobIdentifier: j.jobIdentifier,
	})
}

func (j *standardJobActions) WaitForCompletion(ctx context.Context, pollInterval time.Duration) (*GetJobDetailsOutput, error) {
	return j.client.Jobs().WaitForJobCompletion(ctx, &WaitForJobCompletionInput{
		JobIdentifier: j.jobIdentifier,
	}, pollInterval)
}

func (j *standardJobActions) Cancel(ctx context.Context) (*CancelJobOutput, error) {
	return j.client.Jobs().CancelJob(ctx, &CancelJobInput{
		JobIdentifier: j.jobIdentifier,
	})
}

func (j *standardJobActions) GetResults(ctx context.Context) (*GetJobResultsOutput, error) {
	return j.client.Jobs().GetJobResults(ctx, &GetJobResultsInput{
		JobIdentifier: j.jobIdentifier,
	})
}

func (j *standardJobActions) GetModelDetails(ctx context.Context) (*GetModelVersionDetailsOutput, error) {
	jobDetails, err := j.GetDetails(ctx)
	if err != nil {
		return nil, err
	}
	return j.client.Models().GetModelVersionDetails(ctx, &GetModelVersionDetailsInput{
		ModelID: jobDetails.Details.Model.Identifier,
		Version: jobDetails.Details.Model.Version,
	})
}
