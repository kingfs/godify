package web

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/types"
)

// AccessModeResponse is the response for the GetWebAppAccessMode endpoint.
type AccessModeResponse struct {
	AccessMode string `json:"accessMode"`
}

// WebAppPermissionResponse is the response for the GetWebAppPermission endpoint.
type WebAppPermissionResponse struct {
	Result bool `json:"result"`
}

// GetAppParameters retrieves the parameters of the application for the web client.
func (c *client.Client) GetAppParameters(ctx context.Context, user string) (*types.AppParametersResponse, error) {
	var result types.AppParametersResponse
	path := fmt.Sprintf("/api/parameters?user=%s", user)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAppMeta retrieves the metadata of the application for the web client.
func (c *client.Client) GetAppMeta(ctx context.Context, user string) (*types.AppMetaResponse, error) {
	var result types.AppMetaResponse
	path := fmt.Sprintf("/api/meta?user=%s", user)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetWebAppAccessMode retrieves the access mode for a web application.
func (c *client.Client) GetWebAppAccessMode(ctx context.Context, appID, appCode string) (*AccessModeResponse, error) {
	var result AccessModeResponse
	path := "/api/webapp/access-mode"
	if appID != "" {
		path += fmt.Sprintf("?appId=%s", appID)
	} else if appCode != "" {
		path += fmt.Sprintf("?appCode=%s", appCode)
	}

	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetWebAppPermission checks if the current user has permission to access the web app.
func (c *client.Client) GetWebAppPermission(ctx context.Context, appID string) (*WebAppPermissionResponse, error) {
	var result WebAppPermissionResponse
	path := fmt.Sprintf("/api/webapp/permission?appId=%s", appID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
