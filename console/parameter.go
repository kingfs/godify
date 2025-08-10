package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/types"
)

// GetInstalledAppParameters retrieves the parameters for an installed app.
func (c *client.Client) GetInstalledAppParameters(ctx context.Context, installedAppID string) (*types.AppParametersResponse, error) {
	var result types.AppParametersResponse
	path := fmt.Sprintf("/console/api/installed-apps/%s/parameters", installedAppID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetInstalledAppMeta retrieves the metadata for an installed app.
func (c *client.Client) GetInstalledAppMeta(ctx context.Context, installedAppID string) (*types.AppMetaResponse, error) {
	var result types.AppMetaResponse
	path := fmt.Sprintf("/console/api/installed-apps/%s/meta", installedAppID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
