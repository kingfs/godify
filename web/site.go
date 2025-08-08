package web

import (
	"context"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/types"
)

// GetSite retrieves the site configuration for the web application.
func (c *client.Client) GetSite(ctx context.Context, user string) (*types.Site, error) {
	var result types.Site
	// The user is passed via the WebApiResource wrapper, so it's implicitly available.
	// I'll add it as a query param for consistency, though it might not be strictly needed.
	path := "/api/site?user=" + user
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
