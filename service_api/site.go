package service_api

import (
	"context"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/types"
)

// GetSite retrieves the site configuration for the application.
func (c *client.Client) GetSite(ctx context.Context) (*types.Site, error) {
	var result types.Site
	err := c.sendRequest(ctx, "GET", "/v1/site", nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
