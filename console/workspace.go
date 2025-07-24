package console

import (
	"context"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/models"
)

// GetCurrentTenant 获取当前租户
func (c *Client) GetCurrentTenant(ctx context.Context) (*models.Tenant, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/workspaces/current",
	}

	var result models.Tenant
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}
