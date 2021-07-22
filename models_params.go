package modzy

import (
	"github.com/modzy/go-sdk/model"
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

type ListModelsFilterField string

const (
	ListModelsFilterFieldModelID            ListModelsFilterField = "modelId"
	ListModelsFilterFieldAuthor             ListModelsFilterField = "author"
	ListModelsFilterFieldCreatedByEmail     ListModelsFilterField = "createdByEmail"
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

func (i *ListModelsInput) WithFilterAnd(field ListModelsFilterField, values ...string) *ListModelsInput {
	i.Paging = i.Paging.WithFilter(And(string(field), values...))
	return i
}

func (i *ListModelsInput) WithFilterOr(field ListModelsFilterField, values ...string) *ListModelsInput {
	i.Paging = i.Paging.WithFilter(Or(string(field), values...))
	return i
}

type ListModelsOutput struct {
	Models   []model.ModelVersionSummary `json:"models"`
	NextPage *ListModelsInput            `json:"nextPage"`
}
