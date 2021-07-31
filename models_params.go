package modzy

import (
	"github.com/modzy/go-sdk/internal/model"
)

// GetModelVersionDetailsInput -
type GetModelVersionDetailsInput struct {
	ModelID string
	Version string
}

// GetModelVersionDetailsOutput -
type GetModelVersionDetailsOutput struct {
	Details model.ModelVersionDetails `json:"details"`
}

type GetModelDetailsInput struct {
	ModelID string
}

type GetModelDetailsOutput struct {
	Details model.ModelDetails `json:"details"`
}

type GetMinimumEnginesOutput struct {
	Details model.MinimumEngines `json:"details"`
}

type GetRelatedModelsInput struct {
	ModelID string
}

type GetRelatedModelsOutput struct {
	RelatedModels []model.RelatedModel `json:"relatedModels"`
}

type ListModelsInput struct {
	Paging PagingInput
}

// ListModelsFilterField are known field names that can be used when filtering the models list
type ListModelsFilterField string

const (
	ListModelsFilterFieldModelID            ListModelsFilterField = "modelId"
	ListModelsFilterFieldAuthor             ListModelsFilterField = "author"
	ListModelsFilterFieldCreatedByEmail     ListModelsFilterField = "createdByEmail"
	ListModelsFilterFieldName               ListModelsFilterField = "name"
	ListModelsFilterFieldDescription        ListModelsFilterField = "description"
	ListModelsFilterFieldIsActive           ListModelsFilterField = "isActive"
	ListModelsFilterFieldIsExpired          ListModelsFilterField = "isExpired"
	ListModelsFilterFieldIsCommercial       ListModelsFilterField = "isCommercial"
	ListModelsFilterFieldIsRecommended      ListModelsFilterField = "isRecommended"
	ListModelsFilterFieldLastActiveDateTime ListModelsFilterField = "lastActiveDateTime"
)

func (i *ListModelsInput) WithPaging(perPage int, page int) *ListModelsInput {
	i.Paging = NewPaging(perPage, page)
	return i
}

func (i *ListModelsInput) WithFilter(field ListModelsFilterField, value string) *ListModelsInput {
	i.Paging = i.Paging.WithFilterAnd(string(field), value)
	return i
}

func (i *ListModelsInput) WithFilterAnd(field ListModelsFilterField, values ...string) *ListModelsInput {
	i.Paging = i.Paging.WithFilterAnd(string(field), values...)
	return i
}

func (i *ListModelsInput) WithFilterOr(field ListModelsFilterField, values ...string) *ListModelsInput {
	i.Paging = i.Paging.WithFilterOr(string(field), values...)
	return i
}

type ListModelsOutput struct {
	Models   []model.ModelVersionSummary `json:"models"`
	NextPage *ListModelsInput            `json:"nextPage"`
}

type GetTagsOutput struct {
	Tags []model.ModelTag `json:"tags"`
}

type GetTagModelsInput struct {
	TagIDs []string
}

type GetTagModelsOutput struct {
	Tags   []model.ModelTag      `json:"tags"`
	Models []model.ModelWithTags `json:"models"`
}

type GetModelDetailsByNameInput struct {
	Name string
}

type ListModelVersionsInput struct {
	ModelID string
	Paging  PagingInput
}

// ListModelVersionsFilterField are known field names that can be used when filtering the model versions list
type ListModelVersionsFilterField string

const (
	ListModelVersionsFilterFieldVersion        ListModelVersionsFilterField = "version"
	ListModelVersionsFilterFieldCreatedAt      ListModelVersionsFilterField = "createdAt"
	ListModelVersionsFilterFieldCreatedBy      ListModelVersionsFilterField = "createdBy"
	ListModelVersionsFilterFieldStatus         ListModelVersionsFilterField = "status"
	ListModelVersionsFilterFieldIsAvailable    ListModelVersionsFilterField = "isAvailable"
	ListModelVersionsFilterFieldIsUpdateAt     ListModelVersionsFilterField = "updatedAt"
	ListModelVersionsFilterFieldIsActive       ListModelVersionsFilterField = "isActive"
	ListModelVersionsFilterFieldIsExperimental ListModelVersionsFilterField = "isExperimental"
)

func (i *ListModelVersionsInput) WithPaging(perPage int, page int) *ListModelVersionsInput {
	i.Paging = NewPaging(perPage, page)
	return i
}

func (i *ListModelVersionsInput) WithFilter(field ListModelVersionsFilterField, value string) *ListModelVersionsInput {
	i.Paging = i.Paging.WithFilterAnd(string(field), value)
	return i
}

func (i *ListModelVersionsInput) WithFilterAnd(field ListModelVersionsFilterField, values ...string) *ListModelVersionsInput {
	i.Paging = i.Paging.WithFilterAnd(string(field), values...)
	return i
}

func (i *ListModelVersionsInput) WithFilterOr(field ListModelVersionsFilterField, values ...string) *ListModelVersionsInput {
	i.Paging = i.Paging.WithFilterOr(string(field), values...)
	return i
}

type ListModelVersionsOutput struct {
	Versions []model.ModelVersion `json:"versions"`
	NextPage *ListModelsInput     `json:"nextPage"`
}

type UpdateModelProcessingEnginesInput struct {
	ModelID                 string
	Version                 string
	MinimumParallelCapacity int `json:"minimumParallelCapacity"`
	MaximumParallelCapacity int `json:"maximumParallelCapacity"`
}

type UpdateModelProcessingEnginesOutput struct {
	Details model.ModelVersionDetails `json:"details"`
}

type GetModelVersionSampleInputInput struct {
	ModelID string
	Version string
}

type GetModelVersionSampleInputOutput struct {
	Sample string `json:"sample"`
}

type GetModelVersionSampleOutputInput struct {
	ModelID string
	Version string
}

type GetModelVersionSampleOutputOutput struct {
	Sample string `json:"sample"`
}
