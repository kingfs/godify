package service_api

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/types"
)

// ConversationVariable represents a variable associated with a conversation.
type ConversationVariable struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ValueType   string `json:"value_type"`
	Value       string `json:"value"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

// ConversationVariableListResponse is the response for listing conversation variables.
type ConversationVariableListResponse struct {
	Limit   int                    `json:"limit"`
	HasMore bool                   `json:"has_more"`
	Data    []ConversationVariable `json:"data"`
}

// GetConversations retrieves a list of conversations.
func (c *client.Client) GetConversations(ctx context.Context, user, lastID string, limit int) (*types.ConversationListResponse, error) {
	var result types.ConversationListResponse
	path := fmt.Sprintf("/v1/conversations?user=%s&limit=%d", user, limit)
	if lastID != "" {
		path += fmt.Sprintf("&last_id=%s", lastID)
	}
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// RenameConversation renames a conversation.
func (c *client.Client) RenameConversation(ctx context.Context, conversationID string, req *types.RenameConversationRequest) (*types.SimpleConversation, error) {
	var result types.SimpleConversation
	path := fmt.Sprintf("/v1/conversations/%s/name", conversationID)
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteConversation deletes a conversation.
func (c *client.Client) DeleteConversation(ctx context.Context, conversationID, user string) (*types.ConversationDeleteResponse, error) {
	var result types.ConversationDeleteResponse
	path := fmt.Sprintf("/v1/conversations/%s", conversationID)
	payload := map[string]string{"user": user}
	err := c.sendRequest(ctx, "DELETE", path, payload, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetConversationVariables retrieves variables for a specific conversation.
func (c *client.Client) GetConversationVariables(ctx context.Context, conversationID, user, lastID string, limit int) (*ConversationVariableListResponse, error) {
	var result ConversationVariableListResponse
	path := fmt.Sprintf("/v1/conversations/%s/variables?user=%s&limit=%d", conversationID, user, limit)
	if lastID != "" {
		path += fmt.Sprintf("&last_id=%s", lastID)
	}
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
