package service_api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/types"
)

// CreateCompletionMessage sends a request for a completion message in blocking mode.
func (c *client.Client) CreateCompletionMessage(ctx context.Context, req *types.CompletionRequest) (*types.BlockingResponse, error) {
	req.ResponseMode = "blocking"
	var result types.BlockingResponse
	err := c.sendRequest(ctx, "POST", "/v1/completion-messages", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateChatMessage sends a request for a chat message in blocking mode.
func (c *client.Client) CreateChatMessage(ctx context.Context, req *types.ChatRequest) (*types.BlockingResponse, error) {
	req.ResponseMode = "blocking"
	var result types.BlockingResponse
	err := c.sendRequest(ctx, "POST", "/v1/chat-messages", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// StopCompletionMessage stops a running completion task.
func (c *client.Client) StopCompletionMessage(ctx context.Context, taskID, user string) (*types.StopResponse, error) {
	var result types.StopResponse
	payload := map[string]string{"user": user}
	path := fmt.Sprintf("/v1/completion-messages/%s/stop", taskID)
	err := c.sendRequest(ctx, "POST", path, payload, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// StopChatMessage stops a running chat task.
func (c *client.Client) StopChatMessage(ctx context.Context, taskID, user string) (*types.StopResponse, error) {
	var result types.StopResponse
	payload := map[string]string{"user": user}
	path := fmt.Sprintf("/v1/chat-messages/%s/stop", taskID)
	err := c.sendRequest(ctx, "POST", path, payload, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client.Client) handleStream(ctx context.Context, method, path string, req interface{}) (<-chan *types.StreamEvent, error) {
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

// CreateCompletionMessageStream sends a request for a completion message in streaming mode.
func (c *client.Client) CreateCompletionMessageStream(ctx context.Context, req *types.CompletionRequest) (<-chan *types.StreamEvent, error) {
	req.ResponseMode = "streaming"
	return c.handleStream(ctx, "POST", "/v1/completion-messages", req)
}

// CreateChatMessageStream sends a request for a chat message in streaming mode.
func (c *client.Client) CreateChatMessageStream(ctx context.Context, req *types.ChatRequest) (<-chan *types.StreamEvent, error) {
	req.ResponseMode = "streaming"
	return c.handleStream(ctx, "POST", "/v1/chat-messages", req)
}
