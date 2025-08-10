package console

import (
	"context"

	"github.com/kingfs/godify/client"
)

// PingResponse is the response from the ping endpoint.
type PingResponse struct {
	Result string `json:"result"`
}

// Ping checks the health of the Dify API.
func (c *client.Client) Ping(ctx context.Context) (*PingResponse, error) {
	var result PingResponse
	err := c.sendRequest(ctx, "GET", "/console/api/ping", nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
