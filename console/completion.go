package console

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/types"
)

// ConsoleCompletionRequest is the request for creating a completion message from the console.
type ConsoleCompletionRequest struct {
	Inputs       map[string]interface{} `json:"inputs"`
	Query        string                 `json:"query,omitempty"`
	Files        []types.File           `json:"files,omitempty"`
	ModelConfig  map[string]interface{} `json:"model_config"`
	ResponseMode string                 `json:"response_mode"`
}

// ConsoleChatRequest is the request for creating a chat message from the console.
type ConsoleChatRequest struct {
	Inputs         map[string]interface{} `json:"inputs"`
	Query          string                 `json:"query"`
	Files          []types.File           `json:"files,omitempty"`
	ModelConfig    map[string]interface{} `json:"model_config"`
	ConversationID string                 `json:"conversation_id,omitempty"`
	ResponseMode   string                 `json:"response_mode"`
}

// ConsoleCreateCompletionMessage sends a request for a completion message in blocking mode.
func (c *client.Client) ConsoleCreateCompletionMessage(ctx context.Context, appID string, req *ConsoleCompletionRequest) (*types.BlockingResponse, error) {
	req.ResponseMode = "blocking"
	var result types.BlockingResponse
	path := fmt.Sprintf("/console/api/apps/%s/completion-messages", appID)
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ConsoleCreateChatMessage sends a request for a chat message in blocking mode.
func (c *client.Client) ConsoleCreateChatMessage(ctx context.Context, appID string, req *ConsoleChatRequest) (*types.BlockingResponse, error) {
	req.ResponseMode = "blocking"
	var result types.BlockingResponse
	path := fmt.Sprintf("/console/api/apps/%s/chat-messages", appID)
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ConsoleStopCompletionMessage stops a running completion task.
func (c *client.Client) ConsoleStopCompletionMessage(ctx context.Context, appID, taskID string) (*types.StopResponse, error) {
	var result types.StopResponse
	path := fmt.Sprintf("/console/api/apps/%s/completion-messages/%s/stop", appID, taskID)
	err := c.sendRequest(ctx, "POST", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ConsoleStopChatMessage stops a running chat task.
func (c *client.Client) ConsoleStopChatMessage(ctx context.Context, appID, taskID string) (*types.StopResponse, error) {
	var result types.StopResponse
	path := fmt.Sprintf("/console/api/apps/%s/chat-messages/%s/stop", appID, taskID)
	err := c.sendRequest(ctx, "POST", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client.Client) consoleHandleStream(ctx context.Context, method, path string, req interface{}) (<-chan *types.StreamEvent, error) {
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

// ConsoleCreateCompletionMessageStream sends a request for a completion message in streaming mode.
func (c *client.Client) ConsoleCreateCompletionMessageStream(ctx context.Context, appID string, req *ConsoleCompletionRequest) (<-chan *types.StreamEvent, error) {
	req.ResponseMode = "streaming"
	path := fmt.Sprintf("/console/api/apps/%s/completion-messages", appID)
	return c.consoleHandleStream(ctx, "POST", path, req)
}

// ConsoleCreateChatMessageStream sends a request for a chat message in streaming mode.
func (c *client.Client) ConsoleCreateChatMessageStream(ctx context.Context, appID string, req *ConsoleChatRequest) (<-chan *types.StreamEvent, error) {
	req.ResponseMode = "streaming"
	path := fmt.Sprintf("/console/api/apps/%s/chat-messages", appID)
	return c.consoleHandleStream(ctx, "POST", path, req)
}
