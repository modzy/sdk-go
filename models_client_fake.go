package modzy

import (
	"context"
)

type ModelsClientFake struct {
	ListModelsFunc                   func(ctx context.Context, input *ListModelsInput) (*ListModelsOutput, error)
	GetMinimumEnginesFunc            func(ctx context.Context) (*GetMinimumEnginesOutput, error)
	UpdateModelProcessingEnginesFunc func(ctx context.Context, input *UpdateModelProcessingEnginesInput) (*UpdateModelProcessingEnginesOutput, error)
	GetModelDetailsFunc              func(ctx context.Context, input *GetModelDetailsInput) (*GetModelDetailsOutput, error)
	GetModelDetailsByNameFunc        func(ctx context.Context, input *GetModelDetailsByNameInput) (*GetModelDetailsOutput, error)
	ListModelVersionsFunc            func(ctx context.Context, input *ListModelVersionsInput) (*ListModelVersionsOutput, error)
	GetRelatedModelsFunc             func(ctx context.Context, input *GetRelatedModelsInput) (*GetRelatedModelsOutput, error)
	GetModelVersionDetailsFunc       func(ctx context.Context, input *GetModelVersionDetailsInput) (*GetModelVersionDetailsOutput, error)
	GetModelVersionSampleInputFunc   func(ctx context.Context, input *GetModelVersionSampleInputInput) (*GetModelVersionSampleInputOutput, error)
	GetModelVersionSampleOutputFunc  func(ctx context.Context, input *GetModelVersionSampleOutputInput) (*GetModelVersionSampleOutputOutput, error)
	GetTagsFunc                      func(ctx context.Context) (*GetTagsOutput, error)
	GetTagModelsFunc                 func(ctx context.Context, input *GetTagModelsInput) (*GetTagModelsOutput, error)
}

var _ ModelsClient = &ModelsClientFake{}

func (c *ModelsClientFake) ListModels(ctx context.Context, input *ListModelsInput) (*ListModelsOutput, error) {
	return c.ListModelsFunc(ctx, input)
}

func (c *ModelsClientFake) GetMinimumEngines(ctx context.Context) (*GetMinimumEnginesOutput, error) {
	return c.GetMinimumEnginesFunc(ctx)
}

func (c *ModelsClientFake) UpdateModelProcessingEngines(ctx context.Context, input *UpdateModelProcessingEnginesInput) (*UpdateModelProcessingEnginesOutput, error) {
	return c.UpdateModelProcessingEnginesFunc(ctx, input)
}

func (c *ModelsClientFake) GetModelDetails(ctx context.Context, input *GetModelDetailsInput) (*GetModelDetailsOutput, error) {
	return c.GetModelDetailsFunc(ctx, input)
}

func (c *ModelsClientFake) GetModelDetailsByName(ctx context.Context, input *GetModelDetailsByNameInput) (*GetModelDetailsOutput, error) {
	return c.GetModelDetailsByNameFunc(ctx, input)
}

func (c *ModelsClientFake) ListModelVersions(ctx context.Context, input *ListModelVersionsInput) (*ListModelVersionsOutput, error) {
	return c.ListModelVersionsFunc(ctx, input)
}

func (c *ModelsClientFake) GetRelatedModels(ctx context.Context, input *GetRelatedModelsInput) (*GetRelatedModelsOutput, error) {
	return c.GetRelatedModelsFunc(ctx, input)
}

func (c *ModelsClientFake) GetModelVersionDetails(ctx context.Context, input *GetModelVersionDetailsInput) (*GetModelVersionDetailsOutput, error) {
	return c.GetModelVersionDetailsFunc(ctx, input)
}

func (c *ModelsClientFake) GetModelVersionSampleInput(ctx context.Context, input *GetModelVersionSampleInputInput) (*GetModelVersionSampleInputOutput, error) {
	return c.GetModelVersionSampleInputFunc(ctx, input)
}

func (c *ModelsClientFake) GetModelVersionSampleOutput(ctx context.Context, input *GetModelVersionSampleOutputInput) (*GetModelVersionSampleOutputOutput, error) {
	return c.GetModelVersionSampleOutputFunc(ctx, input)
}

func (c *ModelsClientFake) GetTags(ctx context.Context) (*GetTagsOutput, error) {
	return c.GetTagsFunc(ctx)
}

func (c *ModelsClientFake) GetTagModels(ctx context.Context, input *GetTagModelsInput) (*GetTagModelsOutput, error) {
	return c.GetTagModelsFunc(ctx, input)
}
