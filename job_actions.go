package modzy

import (
	"context"
	"time"
)

type JobActions interface {
	GetDetails(ctx context.Context) (*GetJobDetailsOutput, error)
	WaitForCompletion(ctx context.Context, pollInterval time.Duration) (*GetJobDetailsOutput, error)
	Cancel(ctx context.Context) (*CancelJobOutput, error)
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
	return j.client.Jobs().WaitForJobCompletion(ctx, &GetJobDetailsInput{
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
