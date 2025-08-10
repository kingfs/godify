package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// AppMCPServer represents an MCP server configuration for an app.
type AppMCPServer struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ServerCode  string `json:"server_code"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Parameters  string `json:"parameters"` // JSON string
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

// CreateMCPServerRequest is the request to create an MCP server.
type CreateMCPServerRequest struct {
	Description string                 `json:"description,omitempty"`
	Parameters  map[string]interface{} `json:"parameters"`
}

// UpdateMCPServerRequest is the request to update an MCP server.
type UpdateMCPServerRequest struct {
	ID          string                 `json:"id"`
	Description string                 `json:"description,omitempty"`
	Parameters  map[string]interface{} `json:"parameters"`
	Status      string                 `json:"status,omitempty"`
}

// GetMCPServer retrieves the MCP server configuration for an app.
func (c *client.Client) GetMCPServer(ctx context.Context, appID string) (*AppMCPServer, error) {
	var result AppMCPServer
	path := fmt.Sprintf("/console/api/apps/%s/server", appID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateMCPServer creates an MCP server configuration for an app.
func (c *client.Client) CreateMCPServer(ctx context.Context, appID string, req *CreateMCPServerRequest) (*AppMCPServer, error) {
	var result AppMCPServer
	path := fmt.Sprintf("/console/api/apps/%s/server", appID)
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateMCPServer updates an MCP server configuration for an app.
func (c *client.Client) UpdateMCPServer(ctx context.Context, appID string, req *UpdateMCPServerRequest) (*AppMCPServer, error) {
	var result AppMCPServer
	path := fmt.Sprintf("/console/api/apps/%s/server", appID)
	err := c.sendRequest(ctx, "PUT", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// RefreshMCPServerCode refreshes the server code for an MCP server.
func (c *client.Client) RefreshMCPServerCode(ctx context.Context, serverID string) (*AppMCPServer, error) {
	var result AppMCPServer
	path := fmt.Sprintf("/console/api/apps/%s/server/refresh", serverID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
