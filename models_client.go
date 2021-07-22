package modzy

import (
	"context"
	"fmt"

	"github.com/modzy/go-sdk/model"
)

type ModelsClient interface {
	// GET:/models
	// ListModels(ctx context.Context, input *ListModelsInput) (*ListModelsOutput, error)

	GetMinimumEngines(ctx context.Context) (*GetMinimumEnginesOutput, error)

	// PATCH:/models/{model_id}/versions/{version}
	// PATCH:/models/{model_id}/versions/{version}/processing (for admin?)
	// UpdateModelProcessingEngines(ctx context.Context, input *UpdateModelProcessingEnginesInput) (*UpdateModelProcessingEnginesOutput, error)

	// GET:/models/{model_id}
	// GetModelDetails(ctx context.Context, input *GetModelDetailsInput) (*GetModelDetailsOutput, error)

	// ListModels{name:} -> GetModelDetails(result[0])
	// GetModelDetailsByName(ctx context.Context, input *GetModelDetailsByNameInput) (*GetModelDetailsOutput, error)

	// GET:/models/{model_id}/related-models
	// ListRelatedModels(ctx context.Context, input *ListRelatedModelsInput) (*ListRelatedModelsOutput, error)

	// GET:/models/{model_id}/versions
	// ListModelVersions(ctx context.Context, input *ListModelVersionsInput) (*ListModelVersionsOutput, error)

	GetModelVersionDetails(ctx context.Context, input *GetModelVersionDetailsInput) (*GetModelVersionDetailsOutput, error)

	// GET:/models/{model_id}/versions/{version}/sample-input
	// GetModelVersionSampleInput(ctx context.Context, input *GetModelVersionSampleInputInput) (*GetModelVersionSampleInputOutput, error)

	// GET:/models/{model_id}/versions/{version}/sample-output
	// GetModelVersionSampleOutput(ctx context.Context, input *GetModelVersionSampleOutputInput) (*GetModelVersionSampleOutputOutput, error)

	// GET:/models/tags
	// ListTags(ctx context.Context, input *ListTagsInput) (*ListTagsOutput, error)

	// GET:/models/tags/{tagId}[,{tagId},...]
	// ListTagModels(ctx context.Context, input *ListTagModelsInput) (*ListTagModelsModelsOutput, error)
}

type standardModelsClient struct {
	baseClient *standardClient
}

var _ ModelsClient = &standardModelsClient{}

func (c *standardModelsClient) GetModelVersionDetails(ctx context.Context, input *GetModelVersionDetailsInput) (*GetModelVersionDetailsOutput, error) {
	var out model.ModelDetails
	url := fmt.Sprintf("/api/models/%s/versions/%s", input.Identifier, input.Version)
	_, err := c.baseClient.requestor.get(ctx, url, &out)
	if err != nil {
		return nil, err
	}

	return &GetModelVersionDetailsOutput{
		Details: out,
	}, nil
}

// GET:/models/processing-engines
func (c *standardModelsClient) GetMinimumEngines(ctx context.Context) (*GetMinimumEnginesOutput, error) {
	var out model.MinimumEngines
	url := fmt.Sprintf("/api/models/processing-engines")
	_, err := c.baseClient.requestor.get(ctx, url, &out)
	if err != nil {
		return nil, err
	}

	return &GetMinimumEnginesOutput{
		Details: out,
	}, nil
}
