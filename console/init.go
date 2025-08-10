package console

import (
	"context"

	"github.com/kingfs/godify/client"
)

// InitStatusResponse is the response from the GET /init endpoint.
type InitStatusResponse struct {
	Status string `json:"status"`
}

// InitValidateRequest is the request to validate the initial setup.
type InitValidateRequest struct {
	Password string `json:"password"`
}

// InitValidateResponse is the response after a successful validation.
type InitValidateResponse struct {
	Result string `json:"result"`
}

// GetInitStatus checks if the Dify console has been initialized.
func (c *client.Client) GetInitStatus(ctx context.Context) (*InitStatusResponse, error) {
	var result InitStatusResponse
	err := c.sendRequest(ctx, "GET", "/console/api/init", nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ValidateInit performs the initial setup validation.
func (c *client.Client) ValidateInit(ctx context.Context, password string) (*InitValidateResponse, error) {
	var result InitValidateResponse
	req := InitValidateRequest{Password: password}
	err := c.sendRequest(ctx, "POST", "/console/api/init", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
