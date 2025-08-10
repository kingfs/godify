package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/types"
)

// UpdateModelConfigRequest is the request to update the model configuration of an app.
// The ModelConfig struct is defined in console/app.go and is being reused here.
// This highlights the need for a more comprehensive types package.
type UpdateModelConfigRequest ModelConfig

// UpdateModelConfig updates the model configuration for an application.
func (c *client.Client) UpdateModelConfig(ctx context.Context, appID string, req *UpdateModelConfigRequest) (*types.StopResponse, error) {
	var result types.StopResponse
	path := fmt.Sprintf("/console/api/apps/%s/model-config", appID)
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
