package mcp

import (
	"context"
	"time"

	"github.com/kingfs/godify/client"
)

// Client MCP API客户端 (Model Context Protocol)
type Client struct {
	baseClient *client.BaseClient
}

// NewClient 创建MCP API客户端
func NewClient(baseURL string) *Client {
	config := &client.ClientConfig{
		BaseURL:    baseURL + "/mcp",
		Timeout:    30 * time.Second,
		MaxRetries: 3,
	}

	return &Client{
		baseClient: client.NewBaseClient(config),
	}
}

// MCPRequest MCP请求
type MCPRequest struct {
	Method string                 `json:"method"`
	Params map[string]interface{} `json:"params,omitempty"`
}

// MCPResponse MCP响应
type MCPResponse struct {
	ID     string                 `json:"id,omitempty"`
	Result map[string]interface{} `json:"result,omitempty"`
	Error  *MCPError              `json:"error,omitempty"`
}

// MCPError MCP错误
type MCPError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// SendMCPRequest 发送MCP请求
func (c *Client) SendMCPRequest(ctx context.Context, req *MCPRequest) (*MCPResponse, error) {
	httpReq := &client.Request{
		Method: "POST",
		Path:   "/",
		Body:   req,
	}

	var result MCPResponse
	err := c.baseClient.DoJSON(ctx, httpReq, &result)
	return &result, err
}

// ListResources 列出可用资源
func (c *Client) ListResources(ctx context.Context) (*MCPResponse, error) {
	req := &MCPRequest{
		Method: "resources/list",
	}
	return c.SendMCPRequest(ctx, req)
}

// ListTools 列出可用工具
func (c *Client) ListTools(ctx context.Context) (*MCPResponse, error) {
	req := &MCPRequest{
		Method: "tools/list",
	}
	return c.SendMCPRequest(ctx, req)
}

// CallTool 调用工具
func (c *Client) CallTool(ctx context.Context, name string, arguments map[string]interface{}) (*MCPResponse, error) {
	req := &MCPRequest{
		Method: "tools/call",
		Params: map[string]interface{}{
			"name":      name,
			"arguments": arguments,
		},
	}
	return c.SendMCPRequest(ctx, req)
}

// ReadResource 读取资源
func (c *Client) ReadResource(ctx context.Context, uri string) (*MCPResponse, error) {
	req := &MCPRequest{
		Method: "resources/read",
		Params: map[string]interface{}{
			"uri": uri,
		},
	}
	return c.SendMCPRequest(ctx, req)
}
