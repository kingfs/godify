package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// WorkflowDraftVariable represents a variable in a draft workflow.
type WorkflowDraftVariable struct {
	ID          string      `json:"id"`
	Type        string      `json:"type"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Selector    []string    `json:"selector"`
	ValueType   string      `json:"value_type"`
	Edited      bool        `json:"edited"`
	Visible     bool        `json:"visible"`
	Value       interface{} `json:"value,omitempty"`
}

// WorkflowDraftVariableListResponse is the paginated response for listing workflow variables.
type WorkflowDraftVariableListResponse struct {
	Items []WorkflowDraftVariable `json:"items"`
	Total int                     `json:"total,omitempty"`
}

// UpdateVariableRequest is the request to update a variable.
type UpdateVariableRequest struct {
	Name  string      `json:"name,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// GetWorkflowVariables retrieves the variables for a draft workflow.
func (c *client.Client) GetWorkflowVariables(ctx context.Context, appID string, page, limit int) (*WorkflowDraftVariableListResponse, error) {
	var result WorkflowDraftVariableListResponse
	path := fmt.Sprintf("/console/api/apps/%s/workflows/draft/variables?page=%d&limit=%d", appID, page, limit)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteWorkflowVariables deletes all variables for a draft workflow.
func (c *client.Client) DeleteWorkflowVariables(ctx context.Context, appID string) error {
	path := fmt.Sprintf("/console/api/apps/%s/workflows/draft/variables", appID)
	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}

// GetNodeVariables retrieves the variables for a specific node in a draft workflow.
func (c *client.Client) GetNodeVariables(ctx context.Context, appID, nodeID string) (*WorkflowDraftVariableListResponse, error) {
	var result WorkflowDraftVariableListResponse
	path := fmt.Sprintf("/console/api/apps/%s/workflows/draft/nodes/%s/variables", appID, nodeID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteNodeVariables deletes all variables for a specific node in a draft workflow.
func (c *client.Client) DeleteNodeVariables(ctx context.Context, appID, nodeID string) error {
	path := fmt.Sprintf("/console/api/apps/%s/workflows/draft/nodes/%s/variables", appID, nodeID)
	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}

// GetVariable retrieves a single workflow variable.
func (c *client.Client) GetVariable(ctx context.Context, appID, variableID string) (*WorkflowDraftVariable, error) {
	var result WorkflowDraftVariable
	path := fmt.Sprintf("/console/api/apps/%s/workflows/draft/variables/%s", appID, variableID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateVariable updates a workflow variable.
func (c *client.Client) UpdateVariable(ctx context.Context, appID, variableID string, req *UpdateVariableRequest) (*WorkflowDraftVariable, error) {
	var result WorkflowDraftVariable
	path := fmt.Sprintf("/console/api/apps/%s/workflows/draft/variables/%s", appID, variableID)
	err := c.sendRequest(ctx, "PATCH", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteVariable deletes a single workflow variable.
func (c *client.Client) DeleteVariable(ctx context.Context, appID, variableID string) error {
	path := fmt.Sprintf("/console/api/apps/%s/workflows/draft/variables/%s", appID, variableID)
	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}

// ResetVariable resets a workflow variable to its default value.
func (c *client.Client) ResetVariable(ctx context.Context, appID, variableID string) (*WorkflowDraftVariable, error) {
	var result WorkflowDraftVariable
	path := fmt.Sprintf("/console/api/apps/%s/workflows/draft/variables/%s/reset", appID, variableID)
	err := c.sendRequest(ctx, "PUT", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDraftConversationVariables retrieves the conversation variables for a draft workflow.
func (c *client.Client) GetDraftConversationVariables(ctx context.Context, appID string) (*WorkflowDraftVariableListResponse, error) {
	var result WorkflowDraftVariableListResponse
	path := fmt.Sprintf("/console/api/apps/%s/workflows/draft/conversation-variables", appID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDraftSystemVariables retrieves the system variables for a draft workflow.
func (c *client.Client) GetDraftSystemVariables(ctx context.Context, appID string) (*WorkflowDraftVariableListResponse, error) {
	var result WorkflowDraftVariableListResponse
	path := fmt.Sprintf("/console/api/apps/%s/workflows/draft/system-variables", appID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDraftEnvironmentVariables retrieves the environment variables for a draft workflow.
func (c *client.Client) GetDraftEnvironmentVariables(ctx context.Context, appID string) (*WorkflowDraftVariableListResponse, error) {
	var result WorkflowDraftVariableListResponse
	path := fmt.Sprintf("/console/api/apps/%s/workflows/draft/environment-variables", appID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
