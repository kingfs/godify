package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// WorkflowRun represents a single workflow run in a list.
type WorkflowRun struct {
	ID                string         `json:"id"`
	Version           string         `json:"version"`
	Status            string         `json:"status"`
	ElapsedTime       float64        `json:"elapsed_time"`
	TotalTokens       int            `json:"total_tokens"`
	TotalSteps        int            `json:"total_steps"`
	CreatedByAccount  *SimpleAccount `json:"created_by_account"`
	CreatedAt         int64          `json:"created_at"`
	FinishedAt        int64          `json:"finished_at"`
	ExceptionsCount   int            `json:"exceptions_count"`
}

// AdvancedChatWorkflowRun represents a workflow run in an advanced chat app.
type AdvancedChatWorkflowRun struct {
	WorkflowRun
	ConversationID string `json:"conversation_id"`
	MessageID      string `json:"message_id"`
}

// WorkflowRunDetail represents the detailed information of a workflow run.
type WorkflowRunDetail struct {
	WorkflowRun
	Graph   map[string]interface{} `json:"graph"`
	Inputs  map[string]interface{} `json:"inputs"`
	Outputs map[string]interface{} `json:"outputs"`
	Error   string                 `json:"error"`
}

// WorkflowNodeExecution represents a single node execution in a workflow run.
type WorkflowNodeExecution struct {
	ID                 string                 `json:"id"`
	Index              int                    `json:"index"`
	PredecessorNodeID  string                 `json:"predecessor_node_id"`
	NodeID             string                 `json:"node_id"`
	NodeType           string                 `json:"node_type"`
	Title              string                 `json:"title"`
	Inputs             map[string]interface{} `json:"inputs"`
	ProcessData        map[string]interface{} `json:"process_data"`
	Outputs            map[string]interface{} `json:"outputs"`
	Status             string                 `json:"status"`
	Error              string                 `json:"error"`
	ElapsedTime        float64                `json:"elapsed_time"`
	ExecutionMetadata  map[string]interface{} `json:"execution_metadata"`
	CreatedAt          int64                  `json:"created_at"`
	FinishedAt         int64                  `json:"finished_at"`
}

// WorkflowRunListResponse is the paginated response for listing workflow runs.
type WorkflowRunListResponse struct {
	Data    []WorkflowRun `json:"data"`
	HasMore bool          `json:"has_more"`
	Limit   int           `json:"limit"`
}

// AdvancedChatWorkflowRunListResponse is the paginated response for listing advanced chat workflow runs.
type AdvancedChatWorkflowRunListResponse struct {
	Data    []AdvancedChatWorkflowRun `json:"data"`
	HasMore bool                      `json:"has_more"`
	Limit   int                       `json:"limit"`
}

// WorkflowNodeExecutionListResponse is the response for listing node executions.
type WorkflowNodeExecutionListResponse struct {
	Data []WorkflowNodeExecution `json:"data"`
}

// GetAdvancedChatWorkflowRuns retrieves a list of workflow runs for an advanced chat app.
func (c *client.Client) GetAdvancedChatWorkflowRuns(ctx context.Context, appID, lastID string, limit int) (*AdvancedChatWorkflowRunListResponse, error) {
	var result AdvancedChatWorkflowRunListResponse
	path := fmt.Sprintf("/console/api/apps/%s/advanced-chat/workflow-runs?limit=%d", appID, limit)
	if lastID != "" {
		path += fmt.Sprintf("&last_id=%s", lastID)
	}
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetWorkflowRuns retrieves a list of workflow runs for an app.
func (c *client.Client) GetWorkflowRuns(ctx context.Context, appID, lastID string, limit int) (*WorkflowRunListResponse, error) {
	var result WorkflowRunListResponse
	path := fmt.Sprintf("/console/api/apps/%s/workflow-runs?limit=%d", appID, limit)
	if lastID != "" {
		path += fmt.Sprintf("&last_id=%s", lastID)
	}
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetWorkflowRunDetail retrieves the details of a specific workflow run.
func (c *client.Client) GetWorkflowRunDetail(ctx context.Context, appID, runID string) (*WorkflowRunDetail, error) {
	var result WorkflowRunDetail
	path := fmt.Sprintf("/console/api/apps/%s/workflow-runs/%s", appID, runID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetWorkflowRunNodeExecutions retrieves the node executions for a specific workflow run.
func (c *client.Client) GetWorkflowRunNodeExecutions(ctx context.Context, appID, runID string) (*WorkflowNodeExecutionListResponse, error) {
	var result WorkflowNodeExecutionListResponse
	path := fmt.Sprintf("/console/api/apps/%s/workflow-runs/%s/node-executions", appID, runID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
