package service_api

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/types"
)

// WorkflowRun represents the result of a single workflow execution.
type WorkflowRun struct {
	ID          string                 `json:"id"`
	WorkflowID  string                 `json:"workflow_id"`
	Status      string                 `json:"status"`
	Inputs      map[string]interface{} `json:"inputs"`
	Outputs     map[string]interface{} `json:"outputs"`
	Error       string                 `json:"error"`
	TotalSteps  int                    `json:"total_steps"`
	TotalTokens int                    `json:"total_tokens"`
	CreatedAt   int64                  `json:"created_at"`
	FinishedAt  int64                  `json:"finished_at"`
	ElapsedTime float64                `json:"elapsed_time"`
}

// SimpleAccount represents a user account with basic details.
type SimpleAccount struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// SimpleEndUser represents an end user with basic details.
type SimpleEndUser struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	IsAnonymous bool   `json:"is_anonymous"`
}

// WorkflowRunForLog is a simplified workflow run for logging purposes.
type WorkflowRunForLog struct {
	ID              string  `json:"id"`
	Version         string  `json:"version"`
	Status          string  `json:"status"`
	Error           string  `json:"error"`
	ElapsedTime     float64 `json:"elapsed_time"`
	TotalTokens     int     `json:"total_tokens"`
	TotalSteps      int     `json:"total_steps"`
	CreatedAt       int64   `json:"created_at"`
	FinishedAt      int64   `json:"finished_at"`
	ExceptionsCount int     `json:"exceptions_count"`
}

// WorkflowLog represents a single log entry for a workflow.
type WorkflowLog struct {
	ID                string             `json:"id"`
	WorkflowRun       *WorkflowRunForLog `json:"workflow_run"`
	CreatedFrom       string             `json:"created_from"`
	CreatedByRole     string             `json:"created_by_role"`
	CreatedByAccount  *SimpleAccount     `json:"created_by_account"`
	CreatedByEndUser  *SimpleEndUser     `json:"created_by_end_user"`
	CreatedAt         int64              `json:"created_at"`
}

// WorkflowLogListResponse is the paginated response for workflow logs.
type WorkflowLogListResponse struct {
	Page    int           `json:"page"`
	Limit   int           `json:"limit"`
	Total   int           `json:"total"`
	HasMore bool          `json:"has_more"`
	Data    []WorkflowLog `json:"data"`
}

// RunWorkflow runs a workflow in blocking mode.
func (c *client.Client) RunWorkflow(ctx context.Context, req *types.WorkflowRunRequest) (*types.BlockingResponse, error) {
	req.ResponseMode = "blocking"
	var result types.BlockingResponse
	err := c.sendRequest(ctx, "POST", "/v1/workflows/run", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// RunWorkflowStream runs a workflow in streaming mode.
func (c *client.Client) RunWorkflowStream(ctx context.Context, req *types.WorkflowRunRequest) (<-chan *types.StreamEvent, error) {
	req.ResponseMode = "streaming"
	return c.handleStream(ctx, "POST", "/v1/workflows/run", req)
}

// GetWorkflowRun retrieves the details of a specific workflow run.
func (c *client.Client) GetWorkflowRun(ctx context.Context, runID string) (*WorkflowRun, error) {
	var result WorkflowRun
	path := fmt.Sprintf("/v1/workflows/run/%s", runID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// StopWorkflowTask stops a running workflow task.
func (c *client.Client) StopWorkflowTask(ctx context.Context, taskID, user string) (*types.StopResponse, error) {
	var result types.StopResponse
	payload := map[string]string{"user": user}
	path := fmt.Sprintf("/v1/workflows/tasks/%s/stop", taskID)
	err := c.sendRequest(ctx, "POST", path, payload, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetWorkflowLogs retrieves the logs for workflows.
func (c *client.Client) GetWorkflowLogs(ctx context.Context, page, limit int) (*WorkflowLogListResponse, error) {
	var result WorkflowLogListResponse
	path := fmt.Sprintf("/v1/workflows/logs?page=%d&limit=%d", page, limit)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
