package web

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/types"
)

// WebRunWorkflow runs a workflow in blocking mode for the web client.
func (c *client.Client) WebRunWorkflow(ctx context.Context, req *types.WorkflowRunRequest) (*types.BlockingResponse, error) {
	req.ResponseMode = "blocking"
	var result types.BlockingResponse
	err := c.sendRequest(ctx, "POST", "/api/workflows/run", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// WebRunWorkflowStream runs a workflow in streaming mode for the web client.
func (c *client.Client) WebRunWorkflowStream(ctx context.Context, req *types.WorkflowRunRequest) (<-chan *types.StreamEvent, error) {
	req.ResponseMode = "streaming"
	return c.webHandleStream(ctx, "POST", "/api/workflows/run", req)
}

// WebStopWorkflowTask stops a running workflow task for the web client.
func (c *client.Client) WebStopWorkflowTask(ctx context.Context, taskID, user string) (*types.StopResponse, error) {
	var result types.StopResponse
	payload := map[string]string{"user": user}
	path := fmt.Sprintf("/api/workflows/tasks/%s/stop", taskID)
	err := c.sendRequest(ctx, "POST", path, payload, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
