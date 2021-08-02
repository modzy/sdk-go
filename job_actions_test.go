package modzy

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/modzy/go-sdk/internal/model"
)

type testContextKey string

func TestJobActionsGetDetails(t *testing.T) {
	expectedCtx := context.WithValue(context.TODO(), testContextKey(testContextKey("ck")), "v")
	jobID := "jobid"

	client := &ClientFake{
		JobsFunc: func() JobsClient {
			return &JobsClientFake{
				GetJobDetailsFunc: func(ctx context.Context, input *GetJobDetailsInput) (*GetJobDetailsOutput, error) {
					if expectedCtx != ctx {
						t.Errorf("ctx not passed through")
					}
					if input.JobIdentifier != jobID {
						t.Errorf("JobIdentifier not passed through")
					}
					return nil, fmt.Errorf("made it")
				},
			}
		},
	}
	_, err := NewJobActions(client, jobID).GetDetails(expectedCtx)
	if err.Error() != "made it" {
		t.Errorf("Did not hit test code")
	}
}

func TestJobActionsWaitForCompletion(t *testing.T) {
	expectedCtx := context.WithValue(context.TODO(), testContextKey("ck"), "v")
	jobID := "jobid"

	client := &ClientFake{
		JobsFunc: func() JobsClient {
			return &JobsClientFake{
				WaitForJobCompletionFunc: func(ctx context.Context, input *WaitForJobCompletionInput, pollInterval time.Duration) (*GetJobDetailsOutput, error) {
					if expectedCtx != ctx {
						t.Errorf("ctx not passed through")
					}
					if input.JobIdentifier != jobID {
						t.Errorf("JobIdentifier not passed through")
					}
					if pollInterval != time.Second*10 {
						t.Errorf("pollInterval not passed through")
					}
					return nil, fmt.Errorf("made it")
				},
			}
		},
	}
	_, err := NewJobActions(client, jobID).WaitForCompletion(expectedCtx, time.Second*10)
	if err.Error() != "made it" {
		t.Errorf("Did not hit test code")
	}
}

func TestJobActionsCancel(t *testing.T) {
	expectedCtx := context.WithValue(context.TODO(), testContextKey("ck"), "v")
	jobID := "jobid"

	client := &ClientFake{
		JobsFunc: func() JobsClient {
			return &JobsClientFake{
				CancelJobFunc: func(ctx context.Context, input *CancelJobInput) (*CancelJobOutput, error) {
					if expectedCtx != ctx {
						t.Errorf("ctx not passed through")
					}
					if input.JobIdentifier != jobID {
						t.Errorf("JobIdentifier not passed through")
					}
					return nil, fmt.Errorf("made it")
				},
			}
		},
	}
	_, err := NewJobActions(client, jobID).Cancel(expectedCtx)
	if err.Error() != "made it" {
		t.Errorf("Did not hit test code")
	}
}

func TestJobActionsGetResults(t *testing.T) {
	expectedCtx := context.WithValue(context.TODO(), testContextKey("ck"), "v")
	jobID := "jobid"

	client := &ClientFake{
		JobsFunc: func() JobsClient {
			return &JobsClientFake{
				GetJobResultsFunc: func(ctx context.Context, input *GetJobResultsInput) (*GetJobResultsOutput, error) {
					if expectedCtx != ctx {
						t.Errorf("ctx not passed through")
					}
					if input.JobIdentifier != jobID {
						t.Errorf("JobIdentifier not passed through")
					}
					return nil, fmt.Errorf("made it")
				},
			}
		},
	}
	_, err := NewJobActions(client, jobID).GetResults(expectedCtx)
	if err.Error() != "made it" {
		t.Errorf("Did not hit test code")
	}
}

func TestJobActionsGetModelDetailsError(t *testing.T) {
	expectedCtx := context.WithValue(context.TODO(), testContextKey("ck"), "v")

	client := &ClientFake{
		JobsFunc: func() JobsClient {
			return &JobsClientFake{
				GetJobDetailsFunc: func(ctx context.Context, input *GetJobDetailsInput) (*GetJobDetailsOutput, error) {
					return nil, fmt.Errorf("nope")
				},
			}
		},
	}
	_, err := NewJobActions(client, "do not care").GetModelDetails(expectedCtx)
	if err.Error() != "nope" {
		t.Errorf("Did not hit error condition")
	}
}

func TestJobActionsGetModelDetails(t *testing.T) {
	expectedCtx := context.WithValue(context.TODO(), testContextKey("ck"), "v")
	jobID := "jobid"

	client := &ClientFake{
		JobsFunc: func() JobsClient {
			return &JobsClientFake{
				GetJobDetailsFunc: func(ctx context.Context, input *GetJobDetailsInput) (*GetJobDetailsOutput, error) {
					if expectedCtx != ctx {
						t.Errorf("ctx not passed through")
					}
					if input.JobIdentifier != jobID {
						t.Errorf("JobIdentifier not passed through")
					}

					return &GetJobDetailsOutput{
						Details: model.JobDetails{
							Model: model.ModelNamedIdentifier{
								Identifier: "modelID",
								Version:    "modelVersion",
							},
						},
					}, nil
				},
			}
		},
		ModelsFunc: func() ModelsClient {
			return &ModelsClientFake{
				GetModelVersionDetailsFunc: func(ctx context.Context, input *GetModelVersionDetailsInput) (*GetModelVersionDetailsOutput, error) {
					if expectedCtx != ctx {
						t.Errorf("ctx not passed through")
					}
					if input.ModelID != "modelID" {
						t.Errorf("model ID not passed through")
					}
					if input.Version != "modelVersion" {
						t.Errorf("model version not passed through")
					}
					return nil, fmt.Errorf("made it")
				},
			}
		},
	}
	_, err := NewJobActions(client, jobID).GetModelDetails(expectedCtx)
	if err.Error() != "made it" {
		t.Errorf("Did not hit test code")
	}
}
