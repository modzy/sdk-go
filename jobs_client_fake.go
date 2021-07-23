package modzy

import (
	"context"
	"time"
)

type JobsClientFake struct {
	GetJobDetailsFunc        func(ctx context.Context, input *GetJobDetailsInput) (*GetJobDetailsOutput, error)
	ListJobsHistoryFunc      func(ctx context.Context, input *ListJobsHistoryInput) (*ListJobsHistoryOutput, error)
	SubmitJobTextFunc        func(ctx context.Context, input *SubmitJobTextInput) (*SubmitJobTextOutput, error)
	WaitForJobCompletionFunc func(ctx context.Context, input *GetJobDetailsInput, pollInterval time.Duration) (*GetJobDetailsOutput, error)
	// SubmitJobEmbeddedFunc func(ctx context.Context, input *SubmitJobEmbeddedInput) (*SubmitJobEmbeddedOutput, error)
	// SubmitJobFileFunc func(ctx context.Context, input *SubmitJobFileInput) (*SubmitJobFileOutput, error)
	// SubmitJobS3Func func(ctx context.Context, input *SubmitJobS3Input) (*SubmitJobS3Output, error)
	// SubmitJobJDBCFunc func(ctx context.Context, input *SubmitJobJDBCInput) (*SubmitJobJDBCOutput, error)
	CancelJobFunc      func(ctx context.Context, input *CancelJobInput) (*CancelJobOutput, error)
	GetJobResultsFunc  func(ctx context.Context, input *GetJobResultsInput) (*GetJobResultsOutput, error)
	GetJobFeaturesFunc func(ctx context.Context, input *GetJobFeaturesInput) (*GetJobFeaturesOutput, error)
}

var _ JobsClient = &JobsClientFake{}

func (c *JobsClientFake) GetJobDetails(ctx context.Context, input *GetJobDetailsInput) (*GetJobDetailsOutput, error) {
	return c.GetJobDetailsFunc(ctx, input)
}

func (c *JobsClientFake) ListJobsHistory(ctx context.Context, input *ListJobsHistoryInput) (*ListJobsHistoryOutput, error) {
	return c.ListJobsHistoryFunc(ctx, input)
}

func (c *JobsClientFake) SubmitJobText(ctx context.Context, input *SubmitJobTextInput) (*SubmitJobTextOutput, error) {
	return c.SubmitJobTextFunc(ctx, input)
}

func (c *JobsClientFake) WaitForJobCompletion(ctx context.Context, input *GetJobDetailsInput, pollInterval time.Duration) (*GetJobDetailsOutput, error) {
	return c.WaitForJobCompletionFunc(ctx, input, pollInterval)
}

// SubmitJobEmbedded(ctx context.Context, input *SubmitJobEmbeddedInput) (*SubmitJobEmbeddedOutput, error)
// SubmitJobFile(ctx context.Context, input *SubmitJobFileInput) (*SubmitJobFileOutput, error)
// SubmitJobS3(ctx context.Context, input *SubmitJobS3Input) (*SubmitJobS3Output, error)
// SubmitJobJDBC(ctx context.Context, input *SubmitJobJDBCInput) (*SubmitJobJDBCOutput, error)

func (c *JobsClientFake) CancelJob(ctx context.Context, input *CancelJobInput) (*CancelJobOutput, error) {
	return c.CancelJobFunc(ctx, input)
}

func (c *JobsClientFake) GetJobResults(ctx context.Context, input *GetJobResultsInput) (*GetJobResultsOutput, error) {
	return c.GetJobResultsFunc(ctx, input)
}

func (c *JobsClientFake) GetJobFeatures(ctx context.Context, input *GetJobFeaturesInput) (*GetJobFeaturesOutput, error) {
	return c.GetJobFeaturesFunc(ctx, input)
}
