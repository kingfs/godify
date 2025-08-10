package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// Workflow represents a Dify workflow.
type Workflow struct {
	ID                   string                  `json:"id"`
	Graph                map[string]interface{}  `json:"graph"`
	Features             map[string]interface{}  `json:"features"`
	Hash                 string                  `json:"hash"`
	Version              string                  `json:"version"`
	MarkedName           string                  `json:"marked_name"`
	MarkedComment        string                  `json:"marked_comment"`
	CreatedBy            *SimpleAccount          `json:"created_by"`
	CreatedAt            int64                   `json:"created_at"`
	UpdatedBy            *SimpleAccount          `json:"updated_by"`
	UpdatedAt            int64                   `json:"updated_at"`
	ToolPublished        bool                    `json:"tool_published"`
	EnvironmentVariables []interface{}           `json:"environment_variables"`
	ConversationVariables []ConversationVariable `json:"conversation_variables"`
}

// WorkflowListResponse is the paginated response for listing workflows.
type WorkflowListResponse struct {
	Items   []Workflow `json:"items"`
	Page    int        `json:"page"`
	Limit   int        `json:"limit"`
	HasMore bool       `json:"has_more"`
}

// SyncDraftWorkflowRequest is the request to sync a draft workflow.
type SyncDraftWorkflowRequest struct {
	Graph                 map[string]interface{} `json:"graph"`
	Features              map[string]interface{} `json:"features"`
	Hash                  string                 `json:"hash,omitempty"`
	EnvironmentVariables  []interface{}          `json:"environment_variables"`
	ConversationVariables []interface{}          `json:"conversation_variables,omitempty"`
}

// SyncDraftWorkflowResponse is the response after syncing a draft workflow.
type SyncDraftWorkflowResponse struct {
	Result    string `json:"result"`
	Hash      string `json:"hash"`
	UpdatedAt int64  `json:"updated_at"`
}

// PublishWorkflowRequest is the request to publish a workflow.
type PublishWorkflowRequest struct {
	MarkedName    string `json:"marked_name,omitempty"`
	MarkedComment string `json:"marked_comment,omitempty"`
}

// PublishWorkflowResponse is the response after publishing a workflow.
type PublishWorkflowResponse struct {
	Result    string `json:"result"`
	CreatedAt int64  `json:"created_at"`
}

// ConvertToWorkflowResponse is the response after converting an app to a workflow app.
type ConvertToWorkflowResponse struct {
	NewAppID string `json:"new_app_id"`
}

// GetDraftWorkflow retrieves the draft workflow for an app.
func (c *client.Client) GetDraftWorkflow(ctx context.Context, appID string) (*Workflow, error) {
	var result Workflow
	path := fmt.Sprintf("/console/api/apps/%s/workflows/draft", appID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// SyncDraftWorkflow syncs the draft workflow for an app.
func (c *client.Client) SyncDraftWorkflow(ctx context.Context, appID string, req *SyncDraftWorkflowRequest) (*SyncDraftWorkflowResponse, error) {
	var result SyncDraftWorkflowResponse
	path := fmt.Sprintf("/console/api/apps/%s/workflows/draft", appID)
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPublishedWorkflow retrieves the published workflow for an app.
func (c *client.Client) GetPublishedWorkflow(ctx context.Context, appID string) (*Workflow, error) {
	var result Workflow
	path := fmt.Sprintf("/console/api/apps/%s/workflows/publish", appID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// PublishWorkflow publishes the draft workflow for an app.
func (c *client.Client) PublishWorkflow(ctx context.Context, appID string, req *PublishWorkflowRequest) (*PublishWorkflowResponse, error) {
	var result PublishWorkflowResponse
	path := fmt.Sprintf("/console/api/apps/%s/workflows/publish", appID)
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetWorkflows retrieves a list of all published workflows for an app.
func (c *client.Client) GetWorkflows(ctx context.Context, appID string, page, limit int) (*WorkflowListResponse, error) {
	var result WorkflowListResponse
	path := fmt.Sprintf("/console/api/apps/%s/workflows?page=%d&limit=%d", appID, page, limit)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateWorkflow updates a published workflow.
func (c *client.Client) UpdateWorkflow(ctx context.Context, appID, workflowID string, req *PublishWorkflowRequest) (*Workflow, error) {
	var result Workflow
	path := fmt.Sprintf("/console/api/apps/%s/workflows/%s", appID, workflowID)
	err := c.sendRequest(ctx, "PATCH", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteWorkflow deletes a published workflow.
func (c *client.Client) DeleteWorkflow(ctx context.Context, appID, workflowID string) error {
	path := fmt.Sprintf("/console/api/apps/%s/workflows/%s", appID, workflowID)
	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}

// ConvertToWorkflow converts a simple app to a workflow app.
func (c *client.Client) ConvertToWorkflow(ctx context.Context, appID string) (*ConvertToWorkflowResponse, error) {
	var result ConvertToWorkflowResponse
	path := fmt.Sprintf("/console/api/apps/%s/convert-to-workflow", appID)
	err := c.sendRequest(ctx, "POST", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
