package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/types"
)

// DataSourceAPIKeyBinding represents an API key binding for a data source.
type DataSourceAPIKeyBinding struct {
	ID        string `json:"id"`
	Category  string `json:"category"`
	Provider  string `json:"provider"`
	Disabled  bool   `json:"disabled"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// DataSourceAPIKeyBindingListResponse is the response for listing bindings.
type DataSourceAPIKeyBindingListResponse struct {
	Sources []DataSourceAPIKeyBinding `json:"sources"`
}

// CreateDataSourceBindingRequest is the request to create a new binding.
type CreateDataSourceBindingRequest struct {
	Category    string                 `json:"category"`
	Provider    string                 `json:"provider"`
	Credentials map[string]interface{} `json:"credentials"`
}

// GetDataSourceAPIKeyBindings retrieves all API key bindings for data sources.
func (c *client.Client) GetDataSourceAPIKeyBindings(ctx context.Context) (*DataSourceAPIKeyBindingListResponse, error) {
	var result DataSourceAPIKeyBindingListResponse
	err := c.sendRequest(ctx, "GET", "/console/api/api-key-auth/data-source", nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateDataSourceAPIKeyBinding creates a new API key binding for a data source.
func (c *client.Client) CreateDataSourceAPIKeyBinding(ctx context.Context, req *CreateDataSourceBindingRequest) (*types.StopResponse, error) {
	var result types.StopResponse
	err := c.sendRequest(ctx, "POST", "/console/api/api-key-auth/data-source/binding", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteDataSourceAPIKeyBinding deletes an API key binding for a data source.
func (c *client.Client) DeleteDataSourceAPIKeyBinding(ctx context.Context, bindingID string) error {
	path := fmt.Sprintf("/console/api/api-key-auth/data-source/%s", bindingID)
	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}

// GetDataSourceOAuthURL retrieves the OAuth authorization URL for a given provider.
func (c *client.Client) GetDataSourceOAuthURL(ctx context.Context, provider string) (string, error) {
	var result struct {
		Data string `json:"data"`
	}
	path := fmt.Sprintf("/console/api/oauth/data-source/%s", provider)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return "", err
	}
	return result.Data, nil
}

// BindDataSourceOAuth binds a data source using the OAuth code.
func (c *client.Client) BindDataSourceOAuth(ctx context.Context, provider, code string) (*types.StopResponse, error) {
	var result types.StopResponse
	path := fmt.Sprintf("/console/api/oauth/data-source/binding/%s?code=%s", provider, code)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// SyncDataSourceOAuth triggers a sync for an OAuth-bound data source.
func (c *client.Client) SyncDataSourceOAuth(ctx context.Context, provider, bindingID string) (*types.StopResponse, error) {
	var result types.StopResponse
	path := fmt.Sprintf("/console/api/oauth/data-source/%s/%s/sync", provider, bindingID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
