package console

import (
	"context"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/models"
)

func (c *Client) GetAccountProfile(ctx context.Context, auth_token string) (*models.Account, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/account/profile",
		Headers: map[string]string{
			"Authorization": "Bearer " + auth_token,
		},
	}

	var result models.Account
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

func (c *Client) SetupAccount(ctx context.Context, email string, name string, password string) (*models.OperationResponse, error) {
	req := &client.Request{
		Method: "POST",
		Path:   "/setup",
		Body: map[string]interface{}{
			"email":    email,
			"name":     name,
			"password": password,
		},
	}

	var result models.OperationResponse
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}
