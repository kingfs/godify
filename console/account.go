package console

import (
	"context"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/models"
)

func (c *Client) GetAccountProfile(ctx context.Context) (*models.Account, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/account/profile",
	}

	var result models.Account
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}
