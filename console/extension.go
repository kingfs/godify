package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// APIBasedExtension represents an API-based extension.
type APIBasedExtension struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	APIEndpoint string `json:"api_endpoint"`
	APIKey      string `json:"api_key"` // This will be a masked value
	CreatedAt   int64  `json:"created_at"`
}

// CreateAPIBasedExtensionRequest is the request to create an API-based extension.
type CreateAPIBasedExtensionRequest struct {
	Name        string `json:"name"`
	APIEndpoint string `json:"api_endpoint"`
	APIKey      string `json:"api_key"`
}

// UpdateAPIBasedExtensionRequest is the request to update an API-based extension.
type UpdateAPIBasedExtensionRequest struct {
	Name        string `json:"name"`
	APIEndpoint string `json:"api_endpoint"`
	APIKey      string `json:"api_key"` // Use HIDDEN_VALUE to keep the existing key
}

// GetCodeBasedExtension retrieves data for a code-based extension.
func (c *client.Client) GetCodeBasedExtension(ctx context.Context, module string) (map[string]interface{}, error) {
	var result map[string]interface{}
	path := fmt.Sprintf("/console/api/code-based-extension?module=%s", module)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAPIBasedExtensions retrieves all API-based extensions.
func (c *client.Client) GetAPIBasedExtensions(ctx context.Context) ([]APIBasedExtension, error) {
	var result []APIBasedExtension
	err := c.sendRequest(ctx, "GET", "/console/api/api-based-extension", nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateAPIBasedExtension creates a new API-based extension.
func (c *client.Client) CreateAPIBasedExtension(ctx context.Context, req *CreateAPIBasedExtensionRequest) (*APIBasedExtension, error) {
	var result APIBasedExtension
	err := c.sendRequest(ctx, "POST", "/console/api/api-based-extension", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAPIBasedExtension retrieves a single API-based extension by ID.
func (c *client.Client) GetAPIBasedExtension(ctx context.Context, id string) (*APIBasedExtension, error) {
	var result APIBasedExtension
	path := fmt.Sprintf("/console/api/api-based-extension/%s", id)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateAPIBasedExtension updates an API-based extension.
func (c *client.Client) UpdateAPIBasedExtension(ctx context.Context, id string, req *UpdateAPIBasedExtensionRequest) (*APIBasedExtension, error) {
	var result APIBasedExtension
	path := fmt.Sprintf("/console/api/api-based-extension/%s", id)
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteAPIBasedExtension deletes an API-based extension.
func (c *client.Client) DeleteAPIBasedExtension(ctx context.Context, id string) error {
	path := fmt.Sprintf("/console/api/api-based-extension/%s", id)
	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}
