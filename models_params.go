package modzy

import (
	"github.com/modzy/go-sdk/model"
)

// GetModelDetailsInput -
type GetModelDetailsInput struct {
	Identifier string
	Version    string
}

// GetModelDetailsOutput -
type GetModelDetailsOutput struct {
	Details model.ModelDetails `json:"details"`
}
