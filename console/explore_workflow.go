package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/types"
)

// InstalledAppRunWorkflow runs a workflow for an installed app.
func (c *client.Client) InstalledAppRunWorkflow(ctx context.Context, installedAppID string, req *types.WorkflowRunRequest) (<-chan *types.StreamEvent, error) {
	req.ResponseMode = "streaming"
	path := fmt.Sprintf("/console/api/installed-apps/%s/workflows/run", installedAppID)
	return c.installedAppHandleStream(ctx, "POST", path, req)
}

// InstalledAppStopWorkflowTask stops a running workflow task for an installed app.
func (c *client.Client) InstalledAppStopWorkflowTask(ctx context.Context, installedAppID, taskID string) (*types.StopResponse, error) {
	var result types.StopResponse
	path := fmt.Sprintf("/console/api/installed-apps/%s/workflows/tasks/%s/stop", installedAppID, taskID)
	err := c.sendRequest(ctx, "POST", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
