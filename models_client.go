package modzy

import (
	"context"
	"fmt"

	"github.com/modzy/go-sdk/model"
)

type ModelsClient interface {
	ListModels(ctx context.Context, input *ListModelsInput) (*ListModelsOutput, error)

	// GET:/models/processing-engines
	// ListMinimumEngines(ctx context.Context, input *ListMinimumEnginesInput) (*ListMinimumEnginesOutput, error)

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

func (c *standardModelsClient) ListModels(ctx context.Context, input *ListModelsInput) (*ListModelsOutput, error) {
	input.Paging = input.Paging.withDefaults()

	var items []model.ModelVersionSummary
	url := "/api/models"
	_, links, err := c.baseClient.requestor.list(ctx, url, input.Paging, &items)
	if err != nil {
		return nil, err
	}

	// decide if we have a next page (the next link is not always accurate?)
	var nextPage *ListModelsInput
	if _, hasNextLink := links["next"]; len(items) == input.Paging.PerPage && hasNextLink {
		nextPage = &ListModelsInput{
			Paging: input.Paging.Next(),
		}
	}

	return &ListModelsOutput{
		Models:   items,
		NextPage: nextPage,
	}, nil
}
