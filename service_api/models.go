package service_api

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// I18nObject represents an internationalized string.
type I18nObject struct {
	EnUS  string `json:"en_US"`
	ZhHans string `json:"zh_Hans"`
}

// ProviderModel represents a single model offered by a provider.
type ProviderModel struct {
	Model              string                 `json:"model"`
	Label              I18nObject             `json:"label"`
	ModelType          string                 `json:"model_type"`
	Features           []string               `json:"features"`
	FetchFrom          string                 `json:"fetch_from"`
	ModelProperties    map[string]interface{} `json:"model_properties"`
	Status             string                 `json:"status"`
	LoadBalancingEnabled bool                `json:"load_balancing_enabled"`
}

// ProviderWithModels represents a provider and the list of models they offer.
type ProviderWithModels struct {
	Provider   string          `json:"provider"`
	Label      I18nObject      `json:"label"`
	IconSmall  I18nObject      `json:"icon_small"`
	IconLarge  I18nObject      `json:"icon_large"`
	Status     string          `json:"status"`
	Models     []ProviderModel `json:"models"`
}

// AvailableModelsResponse is the response for getting available models.
type AvailableModelsResponse struct {
	Data []ProviderWithModels `json:"data"`
}

// GetAvailableModels retrieves the available models for a given model type from the workspace.
func (c *client.Client) GetAvailableModels(ctx context.Context, modelType string) (*AvailableModelsResponse, error) {
	var result AvailableModelsResponse
	path := fmt.Sprintf("/v1/workspaces/current/models/model-types/%s", modelType)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
