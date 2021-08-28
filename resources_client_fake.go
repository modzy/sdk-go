package modzy

import (
	"context"
)

// ResourcesClientFake is meant to help in mocking the ResourcesClient interface easily for unit testing.
type ResourcesClientFake struct {
	GetProcessingModelsFunc func(ctx context.Context) (*GetProcessingModelsOutput, error)
}

var _ ResourcesClient = &ResourcesClientFake{}

func (c *ResourcesClientFake) GetProcessingModels(ctx context.Context) (*GetProcessingModelsOutput, error) {
	return c.GetProcessingModelsFunc(ctx)
}
