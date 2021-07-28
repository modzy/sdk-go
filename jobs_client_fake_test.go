package modzy

import (
	"context"
	"testing"
	"time"
)

func TestJobsClientFake(t *testing.T) {
	expectedCtx := context.WithValue(context.TODO(), "a", "b")

	calls := 0
	fake := &JobsClientFake{
		GetJobDetailsFunc: func(ctx context.Context, input *GetJobDetailsInput) (*GetJobDetailsOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			if input == nil {
				t.Errorf("input was not passed through")
			}
			return nil, nil
		},
		ListJobsHistoryFunc: func(ctx context.Context, input *ListJobsHistoryInput) (*ListJobsHistoryOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			if input == nil {
				t.Errorf("input was not passed through")
			}
			return nil, nil
		},
		SubmitJobTextFunc: func(ctx context.Context, input *SubmitJobTextInput) (*SubmitJobTextOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			if input == nil {
				t.Errorf("input was not passed through")
			}
			return nil, nil
		},
		SubmitJobEmbeddedFunc: func(ctx context.Context, input *SubmitJobEmbeddedInput) (*SubmitJobEmbeddedOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			if input == nil {
				t.Errorf("input was not passed through")
			}
			return nil, nil
		},
		WaitForJobCompletionFunc: func(ctx context.Context, input *WaitForJobCompletionInput, pollInterval time.Duration) (*GetJobDetailsOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			if input == nil {
				t.Errorf("input was not passed through")
			}
			if pollInterval != time.Second*12 {
				t.Errorf("not expected pollInterval")
			}
			return nil, nil
		},
		CancelJobFunc: func(ctx context.Context, input *CancelJobInput) (*CancelJobOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			if input == nil {
				t.Errorf("input was not passed through")
			}
			return nil, nil
		},
		GetJobResultsFunc: func(ctx context.Context, input *GetJobResultsInput) (*GetJobResultsOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			if input == nil {
				t.Errorf("input was not passed through")
			}
			return nil, nil
		},
		GetJobFeaturesFunc: func(ctx context.Context) (*GetJobFeaturesOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			return nil, nil
		},
	}

	fake.GetJobDetails(expectedCtx, &GetJobDetailsInput{})
	fake.ListJobsHistory(expectedCtx, &ListJobsHistoryInput{})
	fake.SubmitJobText(expectedCtx, &SubmitJobTextInput{})
	fake.SubmitJobEmbedded(expectedCtx, &SubmitJobEmbeddedInput{})
	fake.WaitForJobCompletion(expectedCtx, &WaitForJobCompletionInput{}, time.Second*12)
	fake.CancelJob(expectedCtx, &CancelJobInput{})
	fake.GetJobResults(expectedCtx, &GetJobResultsInput{})
	fake.GetJobFeatures(expectedCtx)

	if calls != 8 {
		t.Errorf("Did not call all of the funcs: %d", calls)
	}
}
