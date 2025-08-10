package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/types"
)

// InstalledAppGetConversations retrieves a list of conversations for an installed app.
func (c *client.Client) InstalledAppGetConversations(ctx context.Context, installedAppID, lastID string, limit int, pinned *bool) (*types.ConversationListResponse, error) {
	var result types.ConversationListResponse
	path := fmt.Sprintf("/console/api/installed-apps/%s/conversations?limit=%d", installedAppID, limit)
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

// InstalledAppRenameConversation renames a conversation for an installed app.
func (c *client.Client) InstalledAppRenameConversation(ctx context.Context, installedAppID, conversationID string, req *types.RenameConversationRequest) (*types.SimpleConversation, error) {
	var result types.SimpleConversation
	path := fmt.Sprintf("/console/api/installed-apps/%s/conversations/%s/name", installedAppID, conversationID)
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// InstalledAppDeleteConversation deletes a conversation for an installed app.
func (c *client.Client) InstalledAppDeleteConversation(ctx context.Context, installedAppID, conversationID string) error {
	path := fmt.Sprintf("/console/api/installed-apps/%s/conversations/%s", installedAppID, conversationID)
	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}

// InstalledAppPinConversation pins a conversation for an installed app.
func (c *client.Client) InstalledAppPinConversation(ctx context.Context, installedAppID, conversationID string) (*PinResponse, error) {
	var result PinResponse
	path := fmt.Sprintf("/console/api/installed-apps/%s/conversations/%s/pin", installedAppID, conversationID)
	err := c.sendRequest(ctx, "PATCH", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// InstalledAppUnpinConversation unpins a conversation for an installed app.
func (c *client.Client) InstalledAppUnpinConversation(ctx context.Context, installedAppID, conversationID string) (*PinResponse, error) {
	var result PinResponse
	path := fmt.Sprintf("/console/api/installed-apps/%s/conversations/%s/unpin", installedAppID, conversationID)
	err := c.sendRequest(ctx, "PATCH", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
