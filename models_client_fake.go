package modzy

import (
	"context"
)

type ModelsClientFake struct {
	ListModelsFunc             func(ctx context.Context, input *ListModelsInput) (*ListModelsOutput, error)
	GetMinimumEnginesFunc      func(ctx context.Context) (*GetMinimumEnginesOutput, error)
	GetModelDetailsFunc        func(ctx context.Context, input *GetModelDetailsInput) (*GetModelDetailsOutput, error)
	GetRelatedModelsFunc       func(ctx context.Context, input *GetRelatedModelsInput) (*GetRelatedModelsOutput, error)
	GetModelVersionDetailsFunc func(ctx context.Context, input *GetModelVersionDetailsInput) (*GetModelVersionDetailsOutput, error)
	GetTagsFunc                func(ctx context.Context) (*GetTagsOutput, error)
	GetTagModelsFunc           func(ctx context.Context, input *GetTagModelsInput) (*GetTagModelsOutput, error)
}

var _ ModelsClient = &ModelsClientFake{}

func (c *ModelsClientFake) ListModels(ctx context.Context, input *ListModelsInput) (*ListModelsOutput, error) {
	return c.ListModelsFunc(ctx, input)
}

func (c *ModelsClientFake) GetMinimumEngines(ctx context.Context) (*GetMinimumEnginesOutput, error) {
	return c.GetMinimumEnginesFunc(ctx)
}
func (c *ModelsClientFake) GetModelDetails(ctx context.Context, input *GetModelDetailsInput) (*GetModelDetailsOutput, error) {
	return c.GetModelDetailsFunc(ctx, input)
}
func (c *ModelsClientFake) GetRelatedModels(ctx context.Context, input *GetRelatedModelsInput) (*GetRelatedModelsOutput, error) {
	return c.GetRelatedModelsFunc(ctx, input)
}
func (c *ModelsClientFake) GetModelVersionDetails(ctx context.Context, input *GetModelVersionDetailsInput) (*GetModelVersionDetailsOutput, error) {
	return c.GetModelVersionDetailsFunc(ctx, input)
}
func (c *ModelsClientFake) GetTags(ctx context.Context) (*GetTagsOutput, error) {
	return c.GetTagsFunc(ctx)
}
func (c *ModelsClientFake) GetTagModels(ctx context.Context, input *GetTagModelsInput) (*GetTagModelsOutput, error) {
	return c.GetTagModels(ctx, input)
}
