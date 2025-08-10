package service_api

import (
	"context"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/types"
)

// GetAppParameters retrieves the parameters of the application.
func (c *client.Client) GetAppParameters(ctx context.Context) (*types.AppParametersResponse, error) {
	var result types.AppParametersResponse
	err := c.sendRequest(ctx, "GET", "/v1/parameters", nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAppMeta retrieves the metadata of the application.
func (c *client.Client) GetAppMeta(ctx context.Context) (*types.AppMetaResponse, error) {
	var result types.AppMetaResponse
	err := c.sendRequest(ctx, "GET", "/v1/meta", nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAppInfo retrieves the information of the application.
func (c *client.Client) GetAppInfo(ctx context.Context) (*types.AppInfoResponse, error) {
	var result types.AppInfoResponse
	err := c.sendRequest(ctx, "GET", "/v1/info", nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
