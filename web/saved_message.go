package web

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// SavedMessage represents a saved message record.
type SavedMessage struct {
	ID           string                 `json:"id"`
	Inputs       map[string]interface{} `json:"inputs"`
	Query        string                 `json:"query"`
	Answer       string                 `json:"answer"`
	MessageFiles []interface{}          `json:"message_files"` // Simplified
	Feedback     *Feedback              `json:"feedback"`
	CreatedAt    int64                  `json:"created_at"`
}

// SavedMessageListResponse is the paginated response for listing saved messages.
type SavedMessageListResponse struct {
	Limit   int            `json:"limit"`
	HasMore bool           `json:"has_more"`
	Data    []SavedMessage `json:"data"`
}

// SaveMessageRequest is the request to save a message.
type SaveMessageRequest struct {
	MessageID string `json:"message_id"`
	User      string `json:"user"`
}

// SaveMessageResponse is the response after saving a message.
type SaveMessageResponse struct {
	Result string `json:"result"`
}

// DeleteSavedMessageResponse is the response after deleting a saved message.
type DeleteSavedMessageResponse struct {
	Result string `json:"result"`
}

// GetSavedMessages retrieves a list of saved messages.
func (c *client.Client) GetSavedMessages(ctx context.Context, user, lastID string, limit int) (*SavedMessageListResponse, error) {
	var result SavedMessageListResponse
	path := fmt.Sprintf("/api/saved-messages?user=%s&limit=%d", user, limit)
	if lastID != "" {
		path += fmt.Sprintf("&last_id=%s", lastID)
	}
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// SaveMessage saves a message to the user's saved list.
func (c *client.Client) SaveMessage(ctx context.Context, req *SaveMessageRequest) (*SaveMessageResponse, error) {
	var result SaveMessageResponse
	err := c.sendRequest(ctx, "POST", "/api/saved-messages", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteSavedMessage deletes a message from the user's saved list.
func (c *client.Client) DeleteSavedMessage(ctx context.Context, messageID, user string) (*DeleteSavedMessageResponse, error) {
	var result DeleteSavedMessageResponse
	path := fmt.Sprintf("/api/saved-messages/%s", messageID)
	payload := map[string]string{"user": user}
	err := c.sendRequest(ctx, "DELETE", path, payload, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
