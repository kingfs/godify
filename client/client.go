package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/kingfs/godify/config"
	"github.com/kingfs/godify/errors"
	"github.com/kingfs/godify/metrics"
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
	Cookies     map[string]string

	// 监控配置
	Metrics *metrics.Metrics

	// 连接池配置
	MaxIdleConns        int
	MaxIdleConnsPerHost int
	IdleConnTimeout     time.Duration
}

// BaseClient 基础HTTP客户端
type BaseClient struct {
	config     *ClientConfig
	httpClient *http.Client
	logger     *slog.Logger
	metrics    *metrics.Metrics
}

// NewBaseClient 创建基础客户端
func NewBaseClient(config *ClientConfig) *BaseClient {
	if config.HTTPClient == nil {
		// 配置连接池
		transport := &http.Transport{
			MaxIdleConns:        config.MaxIdleConns,
			MaxIdleConnsPerHost: config.MaxIdleConnsPerHost,
			IdleConnTimeout:     config.IdleConnTimeout,
		}

		config.HTTPClient = &http.Client{
			Timeout:   config.Timeout,
			Transport: transport,
		}
	}

	if config.Timeout == 0 {
		config.Timeout = 30 * time.Second
	}

	if config.MaxRetries == 0 {
		config.MaxRetries = 3
	}

	// 设置默认监控
	if config.Metrics == nil {
		config.Metrics = metrics.NewMetrics(false) // 默认关闭监控
	}

	return &BaseClient{
		config:     config,
		httpClient: config.HTTPClient,
		logger:     slog.Default(),
		metrics:    config.Metrics,
	}
}

// WithWorkspaceID 设置工作空间ID
func (c *BaseClient) WithWorkspaceID(workspaceID string) *BaseClient {
	c.config.WorkspaceID = &workspaceID
	return c
}

// WithToken 设置认证token
func (c *BaseClient) WithToken(token string) *BaseClient {
	c.config.Cookies["access_token"] = token
	return c
}

// WithLogger 设置日志器
func (c *BaseClient) WithLogger(logger *slog.Logger) *BaseClient {
	c.logger = logger
	return c
}

// WithCookies 设置cookies
func (c *BaseClient) WithCookies(cookies map[string]string) *BaseClient {
	for k, v := range cookies {
		c.config.Cookies[k] = v
	}
	return c
}

// GetMetrics 获取监控指标
func (c *BaseClient) GetMetrics() *metrics.Metrics {
	return c.metrics
}

// NewClientFromConfig 从配置创建客户端
func NewClientFromConfig(configPath string) (*BaseClient, error) {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	// 创建监控
	metrics := metrics.NewMetrics(cfg.EnableMetrics)

	// 创建客户端配置
	clientConfig := &ClientConfig{
		BaseURL:             cfg.BaseURL,
		AuthType:            AuthType(cfg.AuthType),
		Token:               cfg.Token,
		Timeout:             cfg.Timeout,
		MaxRetries:          cfg.MaxRetries,
		WorkspaceID:         &cfg.WorkspaceID,
		Metrics:             metrics,
		MaxIdleConns:        cfg.MaxIdleConns,
		MaxIdleConnsPerHost: cfg.MaxIdleConnsPerHost,
		IdleConnTimeout:     cfg.IdleConnTimeout,
	}

	return NewBaseClient(clientConfig), nil
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
	Cookies    map[string]string // Cookie 名称到值的映射
}

// Do 执行HTTP请求
func (c *BaseClient) Do(ctx context.Context, req *Request) (*Response, error) {
	startTime := time.Now()

	// 记录请求开始
	c.logger.DebugContext(ctx, "Starting HTTP request", "method", req.Method, "path", req.Path, "url", c.config.BaseURL+req.Path)

	// 构建URL
	u, err := url.Parse(c.config.BaseURL + req.Path)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to parse URL", "error", err)
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

	// // 设置认证头
	// switch c.config.AuthType {
	// case AuthTypeBearer:
	// 	httpReq.Header.Set("Authorization", "Bearer "+c.config.Token)
	// case AuthTypeAPIKey:
	// 	httpReq.Header.Set("Authorization", "Bearer "+c.config.Token)
	// }

	// 设置cookies
	if c.config.Cookies != nil {
		// 将c.config.Cookies中的所有键值对拼成string
		cookieString := ""
		for k, v := range c.config.Cookies {
			cookieString += k + "=" + v + "; "
		}
		httpReq.Header.Add("Cookie", cookieString)
		// csrf_token
		if c.config.Cookies["csrf_token"] != "" {
			httpReq.Header.Add("X-CSRF-Token", c.config.Cookies["csrf_token"])
		}
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
	var lastErr error
	for i := 0; i <= c.config.MaxRetries; i++ {
		resp, err = c.httpClient.Do(httpReq)
		if err == nil && resp.StatusCode < 500 {
			break
		}

		lastErr = err
		c.logger.WarnContext(ctx, "Request failed, retrying...", "attempt", i+1, "max_retries", c.config.MaxRetries, "error", err)

		if i < c.config.MaxRetries {
			time.Sleep(time.Duration(i+1) * time.Second)
		}
	}

	if lastErr != nil {
		c.logger.ErrorContext(ctx, "Request failed after all retries", "error", lastErr)
		return nil, fmt.Errorf("request failed: %w", lastErr)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			// 记录关闭错误，但不影响主流程
			c.logger.WarnContext(ctx, "Failed to close response body", "error", closeErr)
		}
	}()

	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// 提取 cookies（如果有）
	var cookies map[string]string
	if respCookies := resp.Cookies(); len(respCookies) > 0 {
		cookies = make(map[string]string, len(respCookies))
		for _, cookie := range respCookies {
			cookies[cookie.Name] = cookie.Value
		}
	}

	response := &Response{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       respBody,
		Cookies:    cookies,
	}

	// 记录请求完成
	duration := time.Since(startTime)
	c.logger.DebugContext(ctx, "HTTP request completed", "status_code", resp.StatusCode, "duration_ms", duration.Milliseconds(), "body_size", len(respBody))

	// 记录监控指标
	c.metrics.RecordRequest(resp.StatusCode < 400, duration)

	// 检查错误响应
	if resp.StatusCode >= 400 {
		c.logger.ErrorContext(ctx, "HTTP request failed", "status_code", resp.StatusCode, "body", string(respBody))
		return response, c.parseError(ctx, response)
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
func (c *BaseClient) parseError(ctx context.Context, resp *Response) error {
	var errResp errors.ErrorResponse
	if err := json.Unmarshal(resp.Body, &errResp); err != nil {
		// 如果无法解析为结构化错误，返回通用错误
		c.logger.WarnContext(ctx, "Failed to parse structured error response", "status_code", resp.StatusCode, "body", string(resp.Body), "error", err)

		// 记录错误类型
		c.metrics.RecordError("parse_error")

		return &errors.APIError{
			StatusCode: resp.StatusCode,
			Message:    string(resp.Body),
		}
	}

	// 记录特定错误类型
	c.metrics.RecordError(errResp.Code)

	c.logger.ErrorContext(ctx, "API error occurred", "status_code", resp.StatusCode, "error_code", errResp.Code, "message", errResp.Message)

	return &errors.APIError{
		StatusCode: resp.StatusCode,
		Code:       errResp.Code,
		Message:    errResp.Message,
		Details:    errResp.Details,
	}
}
