package console

import (
	"context"

	"github.com/kingfs/godify/client"
)

// SetupStatusResponse is the response from the GET /setup endpoint.
type SetupStatusResponse struct {
	Step     string `json:"step"`
	SetupAt  string `json:"setup_at,omitempty"`
}

// SetupRequest is the request to perform the initial setup.
type SetupRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// SetupResponse is the response after a successful setup.
type SetupResponse struct {
	Result string `json:"result"`
}

// GetSetupStatus checks the setup status of the Dify console.
func (c *client.Client) GetSetupStatus(ctx context.Context) (*SetupStatusResponse, error) {
	var result SetupStatusResponse
	err := c.sendRequest(ctx, "GET", "/console/api/setup", nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Setup performs the initial setup of the Dify console.
func (c *client.Client) Setup(ctx context.Context, req *SetupRequest) (*SetupResponse, error) {
	var result SetupResponse
	err := c.sendRequest(ctx, "POST", "/console/api/setup", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
