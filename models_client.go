package modzy

import (
	"context"
	"fmt"
	"strings"

	"github.com/modzy/go-sdk/model"
)

type ModelsClient interface {
	ListModels(ctx context.Context, input *ListModelsInput) (*ListModelsOutput, error)
	GetMinimumEngines(ctx context.Context) (*GetMinimumEnginesOutput, error)
	UpdateModelProcessingEngines(ctx context.Context, input *UpdateModelProcessingEnginesInput) (*UpdateModelProcessingEnginesOutput, error)
	GetModelDetails(ctx context.Context, input *GetModelDetailsInput) (*GetModelDetailsOutput, error)
	GetModelDetailsByName(ctx context.Context, input *GetModelDetailsByNameInput) (*GetModelDetailsOutput, error)
	ListModelVersions(ctx context.Context, input *ListModelVersionsInput) (*ListModelVersionsOutput, error)
	GetRelatedModels(ctx context.Context, input *GetRelatedModelsInput) (*GetRelatedModelsOutput, error)
	GetModelVersionDetails(ctx context.Context, input *GetModelVersionDetailsInput) (*GetModelVersionDetailsOutput, error)

	// GET:/models/{model_id}/versions/{version}/sample-input
	// GetModelVersionSampleInput(ctx context.Context, input *GetModelVersionSampleInputInput) (*GetModelVersionSampleInputOutput, error)

	// GET:/models/{model_id}/versions/{version}/sample-output
	// GetModelVersionSampleOutput(ctx context.Context, input *GetModelVersionSampleOutputInput) (*GetModelVersionSampleOutputOutput, error)

	GetTags(ctx context.Context) (*GetTagsOutput, error)
	GetTagModels(ctx context.Context, input *GetTagModelsInput) (*GetTagModelsOutput, error)
}

type standardModelsClient struct {
	baseClient *standardClient
}

var _ ModelsClient = &standardModelsClient{}

func (c *standardModelsClient) GetModelVersionDetails(ctx context.Context, input *GetModelVersionDetailsInput) (*GetModelVersionDetailsOutput, error) {
	var out model.ModelVersionDetails
	url := fmt.Sprintf("/api/models/%s/versions/%s", input.ModelID, input.Version)
	_, err := c.baseClient.requestor.get(ctx, url, &out)
	if err != nil {
		return nil, err
	}

	return &GetModelVersionDetailsOutput{
		Details: out,
	}, nil
}

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

func (c *standardModelsClient) GetModelDetails(ctx context.Context, input *GetModelDetailsInput) (*GetModelDetailsOutput, error) {
	var out model.ModelDetails
	url := fmt.Sprintf("/api/models/%s", input.ModelID)
	_, err := c.baseClient.requestor.get(ctx, url, &out)
	if err != nil {
		return nil, err
	}

	return &GetModelDetailsOutput{
		Details: out,
	}, nil
}

func (c *standardModelsClient) GetRelatedModels(ctx context.Context, input *GetRelatedModelsInput) (*GetRelatedModelsOutput, error) {
	var out []model.RelatedModel
	url := fmt.Sprintf("/api/models/%s/related-models", input.ModelID)
	_, err := c.baseClient.requestor.get(ctx, url, &out)
	if err != nil {
		return nil, err
	}

	return &GetRelatedModelsOutput{
		RelatedModels: out,
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

func (c *standardModelsClient) GetTags(ctx context.Context) (*GetTagsOutput, error) {
	var items []model.ModelTag
	url := "/api/models/tags"
	_, err := c.baseClient.requestor.get(ctx, url, &items)
	if err != nil {
		return nil, err
	}

	return &GetTagsOutput{
		Tags: items,
	}, nil
}

func (c *standardModelsClient) GetTagModels(ctx context.Context, input *GetTagModelsInput) (*GetTagModelsOutput, error) {
	var out GetTagModelsOutput
	url := fmt.Sprintf("/api/models/tags/%s", strings.Join(input.TagIDs, ","))
	_, err := c.baseClient.requestor.get(ctx, url, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *standardModelsClient) GetModelDetailsByName(ctx context.Context, input *GetModelDetailsByNameInput) (*GetModelDetailsOutput, error) {
	models, err := c.ListModels(ctx, (&ListModelsInput{}).
		WithPaging(1, 1).
		WithFilterAnd(ListModelsFilterFieldName, input.Name),
	)
	if err != nil {
		return nil, err
	}
	if len(models.Models) != 1 {
		return nil, ErrNotFound
	}
	return c.GetModelDetails(ctx, &GetModelDetailsInput{ModelID: models.Models[0].ID})
}

func (c *standardModelsClient) ListModelVersions(ctx context.Context, input *ListModelVersionsInput) (*ListModelVersionsOutput, error) {
	input.Paging = input.Paging.withDefaults()

	var items []model.ModelVersion
	url := fmt.Sprintf("/api/models/%s/versions", input.ModelID)
	_, links, err := c.baseClient.requestor.list(ctx, url, input.Paging, &items)

	var out model.ModelVersionDetails
	_, err = c.baseClient.requestor.patch(ctx, url, input, &out)
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

	return &ListModelVersionsOutput{
		Versions: items,
		NextPage: nextPage,
	}, nil
}

func (c *standardModelsClient) UpdateModelProcessingEngines(ctx context.Context, input *UpdateModelProcessingEnginesInput) (*UpdateModelProcessingEnginesOutput, error) {
	isAdmin, err := c.baseClient.Accounting().HasEntitlement(ctx, "CAN_PATCH_PROCESSING_MODEL_VERSION")
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("/api/models/%s/versions/%s", input.ModelID, input.Version)
	if isAdmin {
		url = url + "/processing"
	}

	var out model.ModelVersionDetails
	_, err = c.baseClient.requestor.patch(ctx, url, input, &out)
	if err != nil {
		return nil, err
	}

	return &UpdateModelProcessingEnginesOutput{
		Details: out,
	}, nil
}
