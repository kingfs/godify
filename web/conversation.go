package web

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/types"
)

// PinResponse is the response for pinning or unpinning a conversation.
type PinResponse struct {
	Result string `json:"result"`
}

// WebGetConversations retrieves a list of conversations for the web client.
func (c *client.Client) WebGetConversations(ctx context.Context, user, lastID string, limit int, pinned *bool) (*types.ConversationListResponse, error) {
	var result types.ConversationListResponse
	path := fmt.Sprintf("/api/conversations?user=%s&limit=%d", user, limit)
	if lastID != "" {
		path += fmt.Sprintf("&last_id=%s", lastID)
	}
	if pinned != nil {
		path += fmt.Sprintf("&pinned=%t", *pinned)
	}
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// WebRenameConversation renames a conversation for the web client.
func (c *client.Client) WebRenameConversation(ctx context.Context, conversationID string, req *types.RenameConversationRequest) (*types.SimpleConversation, error) {
	var result types.SimpleConversation
	path := fmt.Sprintf("/api/conversations/%s/name", conversationID)
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// WebDeleteConversation deletes a conversation for the web client.
func (c *client.Client) WebDeleteConversation(ctx context.Context, conversationID, user string) (*types.ConversationDeleteResponse, error) {
	var result types.ConversationDeleteResponse
	path := fmt.Sprintf("/api/conversations/%s", conversationID)
	payload := map[string]string{"user": user}
	err := c.sendRequest(ctx, "DELETE", path, payload, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// PinConversation pins a conversation.
func (c *client.Client) PinConversation(ctx context.Context, conversationID, user string) (*PinResponse, error) {
	var result PinResponse
	path := fmt.Sprintf("/api/conversations/%s/pin", conversationID)
	payload := map[string]string{"user": user}
	err := c.sendRequest(ctx, "PATCH", path, payload, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UnpinConversation unpins a conversation.
func (c *client.Client) UnpinConversation(ctx context.Context, conversationID, user string) (*PinResponse, error) {
	var result PinResponse
	path := fmt.Sprintf("/api/conversations/%s/unpin", conversationID)
	payload := map[string]string{"user": user}
	err := c.sendRequest(ctx, "PATCH", path, payload, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
