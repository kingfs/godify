package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// WorkflowAppLog represents a single log entry for a workflow app.
type WorkflowAppLog struct {
	ID                string             `json:"id"`
	WorkflowRun       *WorkflowRunForLog `json:"workflow_run"`
	CreatedFrom       string             `json:"created_from"`
	CreatedByRole     string             `json:"created_by_role"`
	CreatedByAccount  *SimpleAccount     `json:"created_by_account"`
	CreatedByEndUser  *SimpleEndUser     `json:"created_by_end_user"`
	CreatedAt         int64              `json:"created_at"`
}

// WorkflowAppLogListResponse is the paginated response for workflow app logs.
type WorkflowAppLogListResponse struct {
	Page    int              `json:"page"`
	Limit   int              `json:"limit"`
	Total   int              `json:"total"`
	HasMore bool             `json:"has_more"`
	Data    []WorkflowAppLog `json:"data"`
}

// GetWorkflowAppLogs retrieves the logs for a workflow app.
func (c *client.Client) GetWorkflowAppLogs(ctx context.Context, appID string, page, limit int) (*WorkflowAppLogListResponse, error) {
	var result WorkflowAppLogListResponse
	path := fmt.Sprintf("/console/api/apps/%s/workflow-app-logs?page=%d&limit=%d", appID, page, limit)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
