package console

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/types"
)

// InstalledAppCompletionRequest is the request for creating a completion message for an installed app.
type InstalledAppCompletionRequest struct {
	Inputs       map[string]interface{} `json:"inputs"`
	Query        string                 `json:"query,omitempty"`
	Files        []types.File           `json:"files,omitempty"`
	ResponseMode string                 `json:"response_mode"`
}

// InstalledAppChatRequest is the request for creating a chat message for an installed app.
type InstalledAppChatRequest struct {
	Inputs         map[string]interface{} `json:"inputs"`
	Query          string                 `json:"query"`
	Files          []types.File           `json:"files,omitempty"`
	ConversationID string                 `json:"conversation_id,omitempty"`
	ResponseMode   string                 `json:"response_mode"`
}

// InstalledAppCreateCompletionMessage sends a request for a completion message in blocking mode.
func (c *client.Client) InstalledAppCreateCompletionMessage(ctx context.Context, installedAppID string, req *InstalledAppCompletionRequest) (*types.BlockingResponse, error) {
	req.ResponseMode = "blocking"
	var result types.BlockingResponse
	path := fmt.Sprintf("/console/api/installed-apps/%s/completion-messages", installedAppID)
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// InstalledAppCreateChatMessage sends a request for a chat message in blocking mode.
func (c *client.Client) InstalledAppCreateChatMessage(ctx context.Context, installedAppID string, req *InstalledAppChatRequest) (*types.BlockingResponse, error) {
	req.ResponseMode = "blocking"
	var result types.BlockingResponse
	path := fmt.Sprintf("/console/api/installed-apps/%s/chat-messages", installedAppID)
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// InstalledAppStopCompletionMessage stops a running completion task.
func (c *client.Client) InstalledAppStopCompletionMessage(ctx context.Context, installedAppID, taskID string) (*types.StopResponse, error) {
	var result types.StopResponse
	path := fmt.Sprintf("/console/api/installed-apps/%s/completion-messages/%s/stop", installedAppID, taskID)
	err := c.sendRequest(ctx, "POST", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// InstalledAppStopChatMessage stops a running chat task.
func (c *client.Client) InstalledAppStopChatMessage(ctx context.Context, installedAppID, taskID string) (*types.StopResponse, error) {
	var result types.StopResponse
	path := fmt.Sprintf("/console/api/installed-apps/%s/chat-messages/%s/stop", installedAppID, taskID)
	err := c.sendRequest(ctx, "POST", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client.Client) installedAppHandleStream(ctx context.Context, method, path string, req interface{}) (<-chan *types.StreamEvent, error) {
	rawEvents, err := c.SendSSEQuest(ctx, method, path, req)
	if err != nil {
		return nil, err
	}

	events := make(chan *types.StreamEvent)
	go func() {
		defer close(events)
		for rawEvent := range rawEvents {
			var event types.StreamEvent
			if err := json.Unmarshal(rawEvent, &event); err == nil {
				events <- &event
			}
		}
	}()

	return events, nil
}

// InstalledAppCreateCompletionMessageStream sends a request for a completion message in streaming mode.
func (c *client.Client) InstalledAppCreateCompletionMessageStream(ctx context.Context, installedAppID string, req *InstalledAppCompletionRequest) (<-chan *types.StreamEvent, error) {
	req.ResponseMode = "streaming"
	path := fmt.Sprintf("/console/api/installed-apps/%s/completion-messages", installedAppID)
	return c.installedAppHandleStream(ctx, "POST", path, req)
}

// InstalledAppCreateChatMessageStream sends a request for a chat message in streaming mode.
func (c *client.Client) InstalledAppCreateChatMessageStream(ctx context.Context, installedAppID string, req *InstalledAppChatRequest) (<-chan *types.StreamEvent, error) {
	req.ResponseMode = "streaming"
	path := fmt.Sprintf("/console/api/installed-apps/%s/chat-messages", installedAppID)
	return c.installedAppHandleStream(ctx, "POST", path, req)
}
