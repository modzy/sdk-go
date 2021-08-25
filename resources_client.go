package modzy

import (
	"context"
	"sync"

	"github.com/modzy/sdk-go/model"
)

type ResourcesClient interface {
	// GetProcessingModels will get information about the model resources being used
	GetProcessingModels(ctx context.Context) (*GetProcessingModelsOutput, error)
}

type standardResourcesClient struct {
	sync.Mutex
	baseClient *standardClient
}

var _ ResourcesClient = &standardResourcesClient{}

func (c *standardResourcesClient) GetProcessingModels(ctx context.Context) (*GetProcessingModelsOutput, error) {
	var out []model.ResourcesProcessingModel
	url := "/api/resources/processing/models"
	_, err := c.baseClient.requestor.Get(ctx, url, &out)
	if err != nil {
		return nil, err
	}

	return &GetProcessingModelsOutput{
		Models: out,
	}, nil
}
