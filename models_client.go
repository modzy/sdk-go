package modzy

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/modzy/sdk-go/model"
)

type ModelsClient interface {
	// ListModels lists all models.  This supports paging and filtering.
	ListModels(ctx context.Context, input *ListModelsInput) (*ListModelsOutput, error)
	// GetLatestModels returns all of the recent models for your team
	GetLatestModels(ctx context.Context) (*GetLatestModelsOutput, error)
	// GetMinimumEngines reads the engine configuration values
	GetMinimumEngines(ctx context.Context) (*GetMinimumEnginesOutput, error)
	// UpdateModelProcessingEngines updates the engine configuration values
	UpdateModelProcessingEngines(ctx context.Context, input *UpdateModelProcessingEnginesInput) (*UpdateModelProcessingEnginesOutput, error)
	// GetModelDetails reads the details of a model
	GetModelDetails(ctx context.Context, input *GetModelDetailsInput) (*GetModelDetailsOutput, error)
	// GetModelDetailsByName reads details of a model that matches the provided name
	GetModelDetailsByName(ctx context.Context, input *GetModelDetailsByNameInput) (*GetModelDetailsOutput, error)
	// ListModelVersions lists all versions of a model.  This supports paging and filtering.
	ListModelVersions(ctx context.Context, input *ListModelVersionsInput) (*ListModelVersionsOutput, error)
	// GetRelatedModels will return all models that are related to the given model.
	GetRelatedModels(ctx context.Context, input *GetRelatedModelsInput) (*GetRelatedModelsOutput, error)
	// GetModelVersionDetails reads the details of a specific version of a model.
	GetModelVersionDetails(ctx context.Context, input *GetModelVersionDetailsInput) (*GetModelVersionDetailsOutput, error)
	// GetModelVersionSampleInput gets a simple input for the provided model version.
	GetModelVersionSampleInput(ctx context.Context, input *GetModelVersionSampleInputInput) (*GetModelVersionSampleInputOutput, error)
	// GetModelVersionSampleOutput gets a simple output for the provided model version.
	GetModelVersionSampleOutput(ctx context.Context, input *GetModelVersionSampleOutputInput) (*GetModelVersionSampleOutputOutput, error)
	// GetTags returns all tags
	GetTags(ctx context.Context) (*GetTagsOutput, error)
	// GetTagModels returns models that match the provided tags
	GetTagModels(ctx context.Context, input *GetTagModelsInput) (*GetTagModelsOutput, error)
}

type standardModelsClient struct {
	baseClient *standardClient
}

var _ ModelsClient = &standardModelsClient{}

func (c *standardModelsClient) GetModelVersionDetails(ctx context.Context, input *GetModelVersionDetailsInput) (*GetModelVersionDetailsOutput, error) {
	var out model.ModelVersionDetails
	url := fmt.Sprintf("/api/models/%s/versions/%s", input.ModelID, input.Version)
	_, err := c.baseClient.requestor.Get(ctx, url, &out)
	if err != nil {
		return nil, err
	}

	return &GetModelVersionDetailsOutput{
		Details: out,
	}, nil
}

func (c *standardModelsClient) GetLatestModels(ctx context.Context) (*GetLatestModelsOutput, error) {
	var out []model.ModelDetails
	url := "/api/models/latest"
	_, err := c.baseClient.requestor.Get(ctx, url, &out)
	if err != nil {
		return nil, err
	}

	return &GetLatestModelsOutput{
		Models: out,
	}, nil
}

func (c *standardModelsClient) GetMinimumEngines(ctx context.Context) (*GetMinimumEnginesOutput, error) {
	var out model.MinimumEngines
	url := "/api/models/processing-engines"
	_, err := c.baseClient.requestor.Get(ctx, url, &out)
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
	_, err := c.baseClient.requestor.Get(ctx, url, &out)
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
	_, err := c.baseClient.requestor.Get(ctx, url, &out)
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
	_, links, err := c.baseClient.requestor.List(ctx, url, input.Paging, &items)
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
	_, err := c.baseClient.requestor.Get(ctx, url, &items)
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
	_, err := c.baseClient.requestor.Get(ctx, url, &out)
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
	_, links, err := c.baseClient.requestor.List(ctx, url, input.Paging, &items)
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
	_, err = c.baseClient.requestor.Patch(ctx, url, input, &out)
	if err != nil {
		return nil, err
	}
	return &UpdateModelProcessingEnginesOutput{
		Details: out,
	}, nil
}

func (c *standardModelsClient) GetModelVersionSampleInput(ctx context.Context, input *GetModelVersionSampleInputInput) (*GetModelVersionSampleInputOutput, error) {
	var out interface{}
	url := fmt.Sprintf("/api/models/%s/versions/%s/sample-input", input.ModelID, input.Version)
	_, err := c.baseClient.requestor.Get(ctx, url, &out)
	if err != nil {
		return nil, err
	}

	jsonB, err := json.Marshal(out)
	return &GetModelVersionSampleInputOutput{
		Sample: string(jsonB),
	}, err
}

func (c *standardModelsClient) GetModelVersionSampleOutput(ctx context.Context, input *GetModelVersionSampleOutputInput) (*GetModelVersionSampleOutputOutput, error) {
	var out interface{}
	url := fmt.Sprintf("/api/models/%s/versions/%s/sample-output", input.ModelID, input.Version)
	_, err := c.baseClient.requestor.Get(ctx, url, &out)
	if err != nil {
		return nil, err
	}

	jsonB, err := json.Marshal(out)
	return &GetModelVersionSampleOutputOutput{
		Sample: string(jsonB),
	}, err
}
