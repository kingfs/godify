package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// APIKey represents an API key for a resource.
type APIKey struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Token      string `json:"token"`
	LastUsedAt int64  `json:"last_used_at"`
	CreatedAt  int64  `json:"created_at"`
}

// APIKeyListResponse is the response for listing API keys.
type APIKeyListResponse struct {
	Data []APIKey `json:"data"`
}

// APIKeyActionResponse is a generic response for API key actions.
type APIKeyActionResponse struct {
	Result string `json:"result"`
}

// GetAPIKeys retrieves the list of API keys for a given resource.
// resourceType can be "apps" or "datasets".
func (c *client.Client) GetAPIKeys(ctx context.Context, resourceType, resourceID string) (*APIKeyListResponse, error) {
	var result APIKeyListResponse
	path := fmt.Sprintf("/console/api/%s/%s/api-keys", resourceType, resourceID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateAPIKey creates a new API key for a given resource.
// resourceType can be "apps" or "datasets".
func (c *client.Client) CreateAPIKey(ctx context.Context, resourceType, resourceID string) (*APIKey, error) {
	var result APIKey
	path := fmt.Sprintf("/console/api/%s/%s/api-keys", resourceType, resourceID)
	err := c.sendRequest(ctx, "POST", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteAPIKey deletes an API key.
// resourceType can be "apps" or "datasets".
func (c *client.Client) DeleteAPIKey(ctx context.Context, resourceType, resourceID, apiKeyID string) (*APIKeyActionResponse, error) {
	var result APIKeyActionResponse
	path := fmt.Sprintf("/console/api/%s/%s/api-keys/%s", resourceType, resourceID, apiKeyID)
	err := c.sendRequest(ctx, "DELETE", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
