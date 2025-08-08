package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// Conversation represents a completion conversation.
type Conversation struct {
	ID        string                 `json:"id"`
	Status    string                 `json:"status"`
	FromSource string                `json:"from_source"`
	CreatedAt int64                  `json:"created_at"`
	Message   map[string]interface{} `json:"message"`
}

// ConversationListResponse is the paginated response for listing completion conversations.
type ConversationListResponse struct {
	Data    []Conversation `json:"data"`
	HasMore bool           `json:"has_more"`
	Limit   int            `json:"limit"`
	Total   int            `json:"total"`
	Page    int            `json:"page"`
}

// ChatConversation represents a chat conversation with a summary.
type ChatConversation struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Summary   string `json:"summary"`
	Status    string `json:"status"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// ChatConversationListResponse is the paginated response for listing chat conversations.
type ChatConversationListResponse struct {
	Data    []ChatConversation `json:"data"`
	HasMore bool               `json:"has_more"`
	Limit   int                `json:"limit"`
	Total   int                `json:"total"`
	Page    int                `json:"page"`
}

// ConversationDetail represents the detailed information of a conversation.
type ConversationDetail struct {
	ID          string                 `json:"id"`
	Status      string                 `json:"status"`
	FromSource  string                 `json:"from_source"`
	CreatedAt   int64                  `json:"created_at"`
	ModelConfig map[string]interface{} `json:"model_config"`
	Message     map[string]interface{} `json:"message"`
}

// GetCompletionConversations retrieves a list of completion conversations for an app.
func (c *client.Client) GetCompletionConversations(ctx context.Context, appID, keyword, start, end, annotationStatus string, page, limit int) (*ConversationListResponse, error) {
	var result ConversationListResponse
	path := fmt.Sprintf("/console/api/apps/%s/completion-conversations?page=%d&limit=%d&annotation_status=%s", appID, page, limit, annotationStatus)
	if keyword != "" {
		path += fmt.Sprintf("&keyword=%s", keyword)
	}
	// ... add other query params
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetCompletionConversationDetail retrieves the details of a completion conversation.
func (c *client.Client) GetCompletionConversationDetail(ctx context.Context, appID, conversationID string) (*ConversationDetail, error) {
	var result ConversationDetail
	path := fmt.Sprintf("/console/api/apps/%s/completion-conversations/%s", appID, conversationID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteCompletionConversation deletes a completion conversation.
func (c *client.Client) DeleteCompletionConversation(ctx context.Context, appID, conversationID string) error {
	path := fmt.Sprintf("/console/api/apps/%s/completion-conversations/%s", appID, conversationID)
	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}

// GetChatConversations retrieves a list of chat conversations for an app.
func (c *client.Client) GetChatConversations(ctx context.Context, appID string, page, limit int) (*ChatConversationListResponse, error) {
	var result ChatConversationListResponse
	path := fmt.Sprintf("/console/api/apps/%s/chat-conversations?page=%d&limit=%d", appID, page, limit)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetChatConversationDetail retrieves the details of a chat conversation.
func (c *client.Client) GetChatConversationDetail(ctx context.Context, appID, conversationID string) (*ConversationDetail, error) {
	var result ConversationDetail
	path := fmt.Sprintf("/console/api/apps/%s/chat-conversations/%s", appID, conversationID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteChatConversation deletes a chat conversation.
func (c *client.Client) DeleteChatConversation(ctx context.Context, appID, conversationID string) error {
	path := fmt.Sprintf("/console/api/apps/%s/chat-conversations/%s", appID, conversationID)
	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}
