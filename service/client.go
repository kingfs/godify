package service

import (
	"context"
	"time"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/models"
)

// Client Service API客户端
type Client struct {
	baseClient *client.BaseClient
}

// NewClient 创建Service API客户端
func NewClient(appToken, baseURL string) *Client {
	config := &client.ClientConfig{
		BaseURL:    baseURL + "/v1",
		AuthType:   client.AuthTypeBearer,
		Token:      appToken,
		Timeout:    30 * time.Second,
		MaxRetries: 3,
	}

	return &Client{
		baseClient: client.NewBaseClient(config),
	}
}

// GetAppParameters 获取应用参数
func (c *Client) GetAppParameters(ctx context.Context) (*models.AppParameters, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/parameters",
	}

	var result models.AppParameters
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// GetAppMeta 获取应用元数据
func (c *Client) GetAppMeta(ctx context.Context) (*models.AppMeta, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/meta",
	}

	var result models.AppMeta
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// GetAppInfo 获取应用信息
func (c *Client) GetAppInfo(ctx context.Context) (*models.AppInfo, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/info",
	}

	var result models.AppInfo
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// CompletionRequest 文本补全请求
type CompletionRequest struct {
	Inputs        map[string]interface{}   `json:"inputs"`
	Query         string                   `json:"query,omitempty"`
	Files         []map[string]interface{} `json:"files,omitempty"`
	ResponseMode  models.ResponseMode      `json:"response_mode,omitempty"`
	User          string                   `json:"user"`
	RetrieverFrom string                   `json:"retriever_from,omitempty"`
}

// Completion 文本补全
func (c *Client) Completion(ctx context.Context, req *CompletionRequest) (*models.GenerateResponse, error) {
	if req.ResponseMode == "" {
		req.ResponseMode = models.ResponseModeBlocking
	}
	if req.RetrieverFrom == "" {
		req.RetrieverFrom = "dev"
	}

	httpReq := &client.Request{
		Method: "POST",
		Path:   "/completion-messages",
		Body:   req,
	}

	var result models.GenerateResponse
	err := c.baseClient.DoJSON(ctx, httpReq, &result)
	return &result, err
}

// CompletionStop 停止文本补全
func (c *Client) CompletionStop(ctx context.Context, taskID string, user string) error {
	httpReq := &client.Request{
		Method: "POST",
		Path:   "/completion-messages/" + taskID + "/stop",
		Body:   map[string]string{"user": user},
	}

	var result map[string]string
	return c.baseClient.DoJSON(ctx, httpReq, &result)
}

// ChatRequest 聊天请求
type ChatRequest struct {
	Inputs           map[string]interface{}   `json:"inputs"`
	Query            string                   `json:"query"`
	Files            []map[string]interface{} `json:"files,omitempty"`
	ResponseMode     models.ResponseMode      `json:"response_mode,omitempty"`
	ConversationID   string                   `json:"conversation_id,omitempty"`
	User             string                   `json:"user"`
	RetrieverFrom    string                   `json:"retriever_from,omitempty"`
	AutoGenerateName bool                     `json:"auto_generate_name,omitempty"`
}

// Chat 聊天对话
func (c *Client) Chat(ctx context.Context, req *ChatRequest) (*models.GenerateResponse, error) {
	if req.ResponseMode == "" {
		req.ResponseMode = models.ResponseModeBlocking
	}
	if req.RetrieverFrom == "" {
		req.RetrieverFrom = "dev"
	}

	httpReq := &client.Request{
		Method: "POST",
		Path:   "/chat-messages",
		Body:   req,
	}

	var result models.GenerateResponse
	err := c.baseClient.DoJSON(ctx, httpReq, &result)
	return &result, err
}

// ChatStop 停止聊天
func (c *Client) ChatStop(ctx context.Context, taskID string, user string) error {
	httpReq := &client.Request{
		Method: "POST",
		Path:   "/chat-messages/" + taskID + "/stop",
		Body:   map[string]string{"user": user},
	}

	var result map[string]string
	return c.baseClient.DoJSON(ctx, httpReq, &result)
}

// ChatStream 流式聊天对话
func (c *Client) ChatStream(ctx context.Context, req *ChatRequest, handler client.SSEHandler) error {
	// 强制设置为流式模式
	req.ResponseMode = models.ResponseModeStreaming

	if req.RetrieverFrom == "" {
		req.RetrieverFrom = "dev"
	}

	httpReq := &client.Request{
		Method: "POST",
		Path:   "/chat-messages",
		Body:   req,
	}

	return c.baseClient.StreamResponse(ctx, httpReq, handler)
}

// CompletionStream 流式文本补全
func (c *Client) CompletionStream(ctx context.Context, req *CompletionRequest, handler client.SSEHandler) error {
	// 强制设置为流式模式
	req.ResponseMode = models.ResponseModeStreaming

	if req.RetrieverFrom == "" {
		req.RetrieverFrom = "dev"
	}

	httpReq := &client.Request{
		Method: "POST",
		Path:   "/completion-messages",
		Body:   req,
	}

	return c.baseClient.StreamResponse(ctx, httpReq, handler)
}
