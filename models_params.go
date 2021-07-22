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

type GetRelatedModelsInput struct {
	ModelID string
}

type GetRelatedModelsOutput struct {
	RelatedModels []model.RelatedModel `json:"relatedModels"`
}
