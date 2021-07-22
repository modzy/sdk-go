package modzy

import (
	"github.com/modzy/go-sdk/model"
)

// GetModelVersionDetailsInput -
type GetModelVersionDetailsInput struct {
	Identifier string
	Version    string
}

// GetModelVersionDetailsOutput -
type GetModelVersionDetailsOutput struct {
	Details model.ModelDetails `json:"details"`
}
