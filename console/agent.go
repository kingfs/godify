package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// AgentLogMeta represents the metadata of an agent log.
type AgentLogMeta struct {
	Status      string  `json:"status"`
	Executor    string  `json:"executor"`
	StartTime   string  `json:"start_time"`
	ElapsedTime float64 `json:"elapsed_time"`
	TotalTokens int     `json:"total_tokens"`
	AgentMode   string  `json:"agent_mode"`
	Iterations  int     `json:"iterations"`
}

// AgentLogToolCall represents a single tool call within an agent iteration.
type AgentLogToolCall struct {
	Status         string                 `json:"status"`
	Error          string                 `json:"error"`
	TimeCost       float64                `json:"time_cost"`
	ToolName       string                 `json:"tool_name"`
	ToolLabel      string                 `json:"tool_label"`
	ToolInput      map[string]interface{} `json:"tool_input"`
	ToolOutput     map[string]interface{} `json:"tool_output"`
	ToolParameters map[string]interface{} `json:"tool_parameters"`
	ToolIcon       string                 `json:"tool_icon"`
}

// AgentLogIteration represents a single iteration in an agent's thought process.
type AgentLogIteration struct {
	Tokens    int                `json:"tokens"`
	ToolCalls []AgentLogToolCall `json:"tool_calls"`
	ToolRaw   struct {
		Inputs  string `json:"inputs"`
		Outputs string `json:"outputs"`
	} `json:"tool_raw"`
	Thought   string   `json:"thought"`
	CreatedAt string   `json:"created_at"`
	Files     []string `json:"files"`
}

// AgentLogResponse is the response from the agent logs endpoint.
type AgentLogResponse struct {
	Meta       AgentLogMeta        `json:"meta"`
	Iterations []AgentLogIteration `json:"iterations"`
	Files      []interface{}       `json:"files"`
}

// GetAgentLogs retrieves the agent logs for a specific message in a conversation.
func (c *client.Client) GetAgentLogs(ctx context.Context, appID, conversationID, messageID string) (*AgentLogResponse, error) {
	var result AgentLogResponse
	path := fmt.Sprintf("/console/api/apps/%s/agent/logs?conversation_id=%s&message_id=%s", appID, conversationID, messageID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
