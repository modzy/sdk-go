package modzy

import (
	"context"
	"fmt"

	"github.com/modzy/go-sdk/model"
)

type ModelsClient interface {
	GetModelDetails(ctx context.Context, input *GetModelDetailsInput) (*GetModelDetailsOutput, error)
}

type standardModelsClient struct {
	baseClient *standardClient
}

var _ ModelsClient = &standardModelsClient{}

func (c *standardModelsClient) GetModelDetails(ctx context.Context, input *GetModelDetailsInput) (*GetModelDetailsOutput, error) {
	var out model.ModelDetails
	url := fmt.Sprintf("/api/models/%s/versions/%s", input.Identifier, input.Version)
	_, err := c.baseClient.requestor.get(ctx, url, &out)
	if err != nil {
		return nil, err
	}

	return &GetModelDetailsOutput{
		Details: out,
	}, nil

}
