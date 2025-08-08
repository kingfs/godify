package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// ConversationVariable represents a variable in a conversation.
type ConversationVariable struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ValueType   string `json:"value_type"`
	Value       string `json:"value"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

// ConversationVariableListResponse is the paginated response for listing conversation variables.
type ConversationVariableListResponse struct {
	Page    int                    `json:"page"`
	Limit   int                    `json:"limit"`
	Total   int                    `json:"total"`
	HasMore bool                   `json:"has_more"`
	Data    []ConversationVariable `json:"data"`
}

// GetConversationVariables retrieves the variables for a specific conversation within an app.
func (c *client.Client) GetAppConversationVariables(ctx context.Context, appID, conversationID string) (*ConversationVariableListResponse, error) {
	var result ConversationVariableListResponse
	path := fmt.Sprintf("/console/api/apps/%s/conversation-variables?conversation_id=%s", appID, conversationID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
