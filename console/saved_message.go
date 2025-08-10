package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// ConsoleSavedMessage represents a saved message record in the console context.
type ConsoleSavedMessage struct {
	ID           string                 `json:"id"`
	Inputs       map[string]interface{} `json:"inputs"`
	Query        string                 `json:"query"`
	Answer       string                 `json:"answer"`
	MessageFiles []interface{}          `json:"message_files"` // Simplified
	Feedback     interface{}            `json:"feedback"`
	CreatedAt    int64                  `json:"created_at"`
}

// ConsoleSavedMessageListResponse is the paginated response for listing saved messages.
type ConsoleSavedMessageListResponse struct {
	Limit   int                   `json:"limit"`
	HasMore bool                  `json:"has_more"`
	Data    []ConsoleSavedMessage `json:"data"`
}

// ConsoleSaveMessageRequest is the request to save a message.
type ConsoleSaveMessageRequest struct {
	MessageID string `json:"message_id"`
}

// ConsoleSaveMessageResponse is the response after saving a message.
type ConsoleSaveMessageResponse struct {
	Result string `json:"result"`
}

// GetInstalledAppSavedMessages retrieves a list of saved messages for an installed app.
func (c *client.Client) GetInstalledAppSavedMessages(ctx context.Context, installedAppID, lastID string, limit int) (*ConsoleSavedMessageListResponse, error) {
	var result ConsoleSavedMessageListResponse
	path := fmt.Sprintf("/console/api/installed-apps/%s/saved-messages?limit=%d", installedAppID, limit)
	if lastID != "" {
		path += fmt.Sprintf("&last_id=%s", lastID)
	}
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// SaveInstalledAppMessage saves a message for an installed app.
func (c *client.Client) SaveInstalledAppMessage(ctx context.Context, installedAppID, messageID string) (*ConsoleSaveMessageResponse, error) {
	var result ConsoleSaveMessageResponse
	path := fmt.Sprintf("/console/api/installed-apps/%s/saved-messages", installedAppID)
	req := ConsoleSaveMessageRequest{MessageID: messageID}
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteInstalledAppSavedMessage deletes a saved message for an installed app.
func (c *client.Client) DeleteInstalledAppSavedMessage(ctx context.Context, installedAppID, messageID string) error {
	path := fmt.Sprintf("/console/api/installed-apps/%s/saved-messages/%s", installedAppID, messageID)
	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}
