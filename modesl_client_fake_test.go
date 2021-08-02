package modzy

import (
	"context"
	"testing"
)

func TestModelsClientFake(t *testing.T) {
	expectedCtx := context.WithValue(context.TODO(), testContextKey("a"), "b")

	calls := 0
	fake := &ModelsClientFake{
		ListModelsFunc: func(ctx context.Context, input *ListModelsInput) (*ListModelsOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			if input == nil {
				t.Errorf("input was not passed through")
			}
			return nil, nil
		},
		GetMinimumEnginesFunc: func(ctx context.Context) (*GetMinimumEnginesOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			return nil, nil
		},
		UpdateModelProcessingEnginesFunc: func(ctx context.Context, input *UpdateModelProcessingEnginesInput) (*UpdateModelProcessingEnginesOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			if input == nil {
				t.Errorf("input was not passed through")
			}
			return nil, nil
		},
		GetModelDetailsFunc: func(ctx context.Context, input *GetModelDetailsInput) (*GetModelDetailsOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			if input == nil {
				t.Errorf("input was not passed through")
			}
			return nil, nil
		},
		GetModelDetailsByNameFunc: func(ctx context.Context, input *GetModelDetailsByNameInput) (*GetModelDetailsOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			if input == nil {
				t.Errorf("input was not passed through")
			}
			return nil, nil
		},
		ListModelVersionsFunc: func(ctx context.Context, input *ListModelVersionsInput) (*ListModelVersionsOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			if input == nil {
				t.Errorf("input was not passed through")
			}
			return nil, nil
		},
		GetRelatedModelsFunc: func(ctx context.Context, input *GetRelatedModelsInput) (*GetRelatedModelsOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			if input == nil {
				t.Errorf("input was not passed through")
			}
			return nil, nil
		},
		GetModelVersionDetailsFunc: func(ctx context.Context, input *GetModelVersionDetailsInput) (*GetModelVersionDetailsOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			if input == nil {
				t.Errorf("input was not passed through")
			}
			return nil, nil
		},
		GetModelVersionSampleInputFunc: func(ctx context.Context, input *GetModelVersionSampleInputInput) (*GetModelVersionSampleInputOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			if input == nil {
				t.Errorf("input was not passed through")
			}
			return nil, nil
		},
		GetModelVersionSampleOutputFunc: func(ctx context.Context, input *GetModelVersionSampleOutputInput) (*GetModelVersionSampleOutputOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			if input == nil {
				t.Errorf("input was not passed through")
			}
			return nil, nil
		},
		GetTagsFunc: func(ctx context.Context) (*GetTagsOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			return nil, nil
		},
		GetTagModelsFunc: func(ctx context.Context, input *GetTagModelsInput) (*GetTagModelsOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			if input == nil {
				t.Errorf("input was not passed through")
			}
			return nil, nil
		},
	}

	fake.ListModels(expectedCtx, &ListModelsInput{})
	fake.GetMinimumEngines(expectedCtx)
	fake.UpdateModelProcessingEngines(expectedCtx, &UpdateModelProcessingEnginesInput{})
	fake.GetModelDetails(expectedCtx, &GetModelDetailsInput{})
	fake.GetModelDetailsByName(expectedCtx, &GetModelDetailsByNameInput{})
	fake.ListModelVersions(expectedCtx, &ListModelVersionsInput{})
	fake.GetRelatedModels(expectedCtx, &GetRelatedModelsInput{})
	fake.GetModelVersionDetails(expectedCtx, &GetModelVersionDetailsInput{})
	fake.GetModelVersionSampleInput(expectedCtx, &GetModelVersionSampleInputInput{})
	fake.GetModelVersionSampleOutput(expectedCtx, &GetModelVersionSampleOutputInput{})
	fake.GetTags(expectedCtx)
	fake.GetTagModels(expectedCtx, &GetTagModelsInput{})

	if calls != 12 {
		t.Errorf("Did not call all of the funcs: %d", calls)
	}
}
