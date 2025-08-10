package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/types"
)

// TraceConfig represents the tracing configuration for a specific provider.
type TraceConfig map[string]interface{}

// UpdateTraceConfigRequest is the request to create or update a trace configuration.
type UpdateTraceConfigRequest struct {
	TracingProvider string      `json:"tracing_provider"`
	TracingConfig   TraceConfig `json:"tracing_config"`
}

// GetTraceConfig retrieves the tracing configuration for a specific provider.
func (c *client.Client) GetTraceConfig(ctx context.Context, appID, tracingProvider string) (TraceConfig, error) {
	var result TraceConfig
	path := fmt.Sprintf("/console/api/apps/%s/trace-config?tracing_provider=%s", appID, tracingProvider)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateTraceConfig creates a new tracing configuration.
func (c *client.Client) CreateTraceConfig(ctx context.Context, appID string, req *UpdateTraceConfigRequest) (TraceConfig, error) {
	var result TraceConfig
	path := fmt.Sprintf("/console/api/apps/%s/trace-config", appID)
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateTraceConfig updates an existing tracing configuration.
func (c *client.Client) UpdateTraceConfig(ctx context.Context, appID string, req *UpdateTraceConfigRequest) (*types.StopResponse, error) {
	var result types.StopResponse
	path := fmt.Sprintf("/console/api/apps/%s/trace-config", appID)
	err := c.sendRequest(ctx, "PATCH", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteTraceConfig deletes a tracing configuration.
func (c *client.Client) DeleteTraceConfig(ctx context.Context, appID, tracingProvider string) error {
	path := fmt.Sprintf("/console/api/apps/%s/trace-config?tracing_provider=%s", appID, tracingProvider)
	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}
