package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/kingfs/godify/errors"
)

// AuthType 认证类型
type AuthType string

const (
	AuthTypeBearer AuthType = "Bearer"
	AuthTypeAPIKey AuthType = "ApiKey"
)

// ClientConfig 客户端配置
type ClientConfig struct {
	BaseURL    string
	AuthType   AuthType
	Token      string
	HTTPClient *http.Client
	Timeout    time.Duration
	MaxRetries int

	WorkspaceID *string
}

// BaseClient 基础HTTP客户端
type BaseClient struct {
	config     *ClientConfig
	httpClient *http.Client
}

// NewBaseClient 创建基础客户端
func NewBaseClient(config *ClientConfig) *BaseClient {
	if config.HTTPClient == nil {
		config.HTTPClient = &http.Client{
			Timeout: config.Timeout,
		}
	}

	if config.Timeout == 0 {
		config.Timeout = 30 * time.Second
	}

	if config.MaxRetries == 0 {
		config.MaxRetries = 3
	}

	return &BaseClient{
		config:     config,
		httpClient: config.HTTPClient,
	}
}

// WithWorkspaceID 设置工作空间ID
func (c *BaseClient) WithWorkspaceID(workspaceID string) *BaseClient {
	c.config.WorkspaceID = &workspaceID
	return c
}

// WithToken 设置认证token
func (c *BaseClient) WithToken(token string) *BaseClient {
	c.config.Token = token
	return c
}

// Request 请求参数
type Request struct {
	Method  string
	Path    string
	Headers map[string]string
	Query   map[string]string
	Body    interface{}
}

// Response 响应
type Response struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
}

// Do 执行HTTP请求
func (c *BaseClient) Do(ctx context.Context, req *Request) (*Response, error) {
	// 构建URL
	u, err := url.Parse(c.config.BaseURL + req.Path)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	// 添加查询参数
	if req.Query != nil {
		q := u.Query()
		for k, v := range req.Query {
			q.Set(k, v)
		}
		u.RawQuery = q.Encode()
	}

	// 准备请求体
	var body io.Reader
	var contentType string

	if req.Body != nil {
		switch v := req.Body.(type) {
		case string:
			body = strings.NewReader(v)
			contentType = "text/plain"
		case []byte:
			body = bytes.NewReader(v)
			contentType = "application/octet-stream"
		default:
			jsonData, err := json.Marshal(v)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal request body: %w", err)
			}
			body = bytes.NewReader(jsonData)
			contentType = "application/json"
		}
	}

	// 创建HTTP请求
	httpReq, err := http.NewRequestWithContext(ctx, req.Method, u.String(), body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// 设置Content-Type
	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	// 设置认证头
	switch c.config.AuthType {
	case AuthTypeBearer:
		httpReq.Header.Set("Authorization", "Bearer "+c.config.Token)
	case AuthTypeAPIKey:
		httpReq.Header.Set("Authorization", "Bearer "+c.config.Token)
	}

	if c.config.WorkspaceID != nil {
		httpReq.Header.Set("X-Workspace-Id", *c.config.WorkspaceID)
	}

	// 设置自定义头
	for k, v := range req.Headers {
		httpReq.Header.Set(k, v)
	}

	// 执行请求（带重试）
	var resp *http.Response
	for i := 0; i <= c.config.MaxRetries; i++ {
		resp, err = c.httpClient.Do(httpReq)
		if err == nil && resp.StatusCode < 500 {
			break
		}

		if i < c.config.MaxRetries {
			time.Sleep(time.Duration(i+1) * time.Second)
		}
	}

	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			// 记录关闭错误，但不影响主流程
			fmt.Printf("Warning: failed to close response body: %v\n", closeErr)
		}
	}()

	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	response := &Response{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       respBody,
	}

	// 检查错误响应
	if resp.StatusCode >= 400 {
		return response, c.parseError(response)
	}

	return response, nil
}

// DoJSON 执行JSON请求并解析响应
func (c *BaseClient) DoJSON(ctx context.Context, req *Request, result interface{}) error {
	resp, err := c.Do(ctx, req)
	if err != nil {
		return err
	}

	if result != nil && len(resp.Body) > 0 {
		if err := json.Unmarshal(resp.Body, result); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return nil
}

// UploadFile 上传文件
func (c *BaseClient) UploadFile(ctx context.Context, path string, fieldName string, filename string, fileData []byte, extraFields map[string]string) (*Response, error) {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// 添加文件字段
	part, err := writer.CreateFormFile(fieldName, filename)
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}

	if _, err := part.Write(fileData); err != nil {
		return nil, fmt.Errorf("failed to write file data: %w", err)
	}

	// 添加额外字段
	for key, value := range extraFields {
		if err := writer.WriteField(key, value); err != nil {
			return nil, fmt.Errorf("failed to write field %s: %w", key, err)
		}
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close multipart writer: %w", err)
	}

	req := &Request{
		Method: "POST",
		Path:   path,
		Headers: map[string]string{
			"Content-Type": writer.FormDataContentType(),
		},
		Body: buf.Bytes(),
	}

	return c.Do(ctx, req)
}

// parseError 解析错误响应
func (c *BaseClient) parseError(resp *Response) error {
	var errResp errors.ErrorResponse
	if err := json.Unmarshal(resp.Body, &errResp); err != nil {
		// 如果无法解析为结构化错误，返回通用错误
		return &errors.APIError{
			StatusCode: resp.StatusCode,
			Message:    string(resp.Body),
		}
	}

	return &errors.APIError{
		StatusCode: resp.StatusCode,
		Code:       errResp.Code,
		Message:    errResp.Message,
		Details:    errResp.Details,
	}
}
