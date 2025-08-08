package web

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/types"
)

// WebCreateCompletionMessage sends a request for a completion message in blocking mode.
func (c *client.Client) WebCreateCompletionMessage(ctx context.Context, req *types.CompletionRequest) (*types.BlockingResponse, error) {
	req.ResponseMode = "blocking"
	var result types.BlockingResponse
	err := c.sendRequest(ctx, "POST", "/api/completion-messages", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// WebCreateChatMessage sends a request for a chat message in blocking mode.
func (c *client.Client) WebCreateChatMessage(ctx context.Context, req *types.ChatRequest) (*types.BlockingResponse, error) {
	req.ResponseMode = "blocking"
	var result types.BlockingResponse
	err := c.sendRequest(ctx, "POST", "/api/chat-messages", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// WebStopCompletionMessage stops a running completion task.
func (c *client.Client) WebStopCompletionMessage(ctx context.Context, taskID, user string) (*types.StopResponse, error) {
	var result types.StopResponse
	payload := map[string]string{"user": user}
	path := fmt.Sprintf("/api/completion-messages/%s/stop", taskID)
	err := c.sendRequest(ctx, "POST", path, payload, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// WebStopChatMessage stops a running chat task.
func (c *client.Client) WebStopChatMessage(ctx context.Context, taskID, user string) (*types.StopResponse, error) {
	var result types.StopResponse
	payload := map[string]string{"user": user}
	path := fmt.Sprintf("/api/chat-messages/%s/stop", taskID)
	err := c.sendRequest(ctx, "POST", path, payload, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client.Client) webHandleStream(ctx context.Context, method, path string, req interface{}) (<-chan *types.StreamEvent, error) {
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

// WebCreateCompletionMessageStream sends a request for a completion message in streaming mode.
func (c *client.Client) WebCreateCompletionMessageStream(ctx context.Context, req *types.CompletionRequest) (<-chan *types.StreamEvent, error) {
	req.ResponseMode = "streaming"
	return c.webHandleStream(ctx, "POST", "/api/completion-messages", req)
}

// WebCreateChatMessageStream sends a request for a chat message in streaming mode.
func (c *client.Client) WebCreateChatMessageStream(ctx context.Context, req *types.ChatRequest) (<-chan *types.StreamEvent, error) {
	req.ResponseMode = "streaming"
	return c.webHandleStream(ctx, "POST", "/api/chat-messages", req)
}
