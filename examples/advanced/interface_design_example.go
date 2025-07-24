package main

import (
	"context"
	"fmt"
	"time"

	"github.com/kingfs/godify"
	"github.com/kingfs/godify/service"
	"github.com/kingfs/godify/models"
)

// ChatClient 聊天客户端接口
type ChatClient interface {
	Chat(ctx context.Context, req *service.ChatRequest) (*models.GenerateResponse, error)
	ChatStream(ctx context.Context, req *service.ChatRequest, handler service.SSEHandler) error
}

// CompletionClient 补全客户端接口
type CompletionClient interface {
	Completion(ctx context.Context, req *service.CompletionRequest) (*models.GenerateResponse, error)
	CompletionStream(ctx context.Context, req *service.CompletionRequest, handler service.SSEHandler) error
}

// AIProvider 统一的AI服务提供者接口
type AIProvider interface {
	ChatClient
	CompletionClient
	GetAppInfo(ctx context.Context) (*models.AppInfo, error)
}

// DifyProvider Dify平台实现
type DifyProvider struct {
	client *service.Client
}

func NewDifyProvider(token, baseURL string) *DifyProvider {
	return &DifyProvider{
		client: dify.NewServiceClient(token, baseURL),
	}
}

func (d *DifyProvider) Chat(ctx context.Context, req *service.ChatRequest) (*models.GenerateResponse, error) {
	return d.client.Chat(ctx, req)
}

func (d *DifyProvider) ChatStream(ctx context.Context, req *service.ChatRequest, handler service.SSEHandler) error {
	return d.client.ChatStream(ctx, req, handler)
}

func (d *DifyProvider) Completion(ctx context.Context, req *service.CompletionRequest) (*models.GenerateResponse, error) {
	return d.client.Completion(ctx, req)
}

func (d *DifyProvider) CompletionStream(ctx context.Context, req *service.CompletionRequest, handler service.SSEHandler) error {
	return d.client.CompletionStream(ctx, req, handler)
}

func (d *DifyProvider) GetAppInfo(ctx context.Context) (*models.AppInfo, error) {
	return d.client.GetAppInfo(ctx)
}

// MockProvider 用于测试的模拟实现
type MockProvider struct {
	responses map[string]*models.GenerateResponse
	errors    map[string]error
}

func NewMockProvider() *MockProvider {
	return &MockProvider{
		responses: make(map[string]*models.GenerateResponse),
		errors:    make(map[string]error),
	}
}

func (m *MockProvider) SetResponse(query string, response *models.GenerateResponse) {
	m.responses[query] = response
}

func (m *MockProvider) SetError(query string, err error) {
	m.errors[query] = err
}

func (m *MockProvider) Chat(ctx context.Context, req *service.ChatRequest) (*models.GenerateResponse, error) {
	if err, exists := m.errors[req.Query]; exists {
		return nil, err
	}
	if resp, exists := m.responses[req.Query]; exists {
		return resp, nil
	}
	return &models.GenerateResponse{
		Answer: "Mock response for: " + req.Query,
	}, nil
}

func (m *MockProvider) ChatStream(ctx context.Context, req *service.ChatRequest, handler service.SSEHandler) error {
	// 模拟流式响应
	handler.OnEvent("message", map[string]interface{}{
		"answer": "Mock streaming response",
	})
	handler.OnComplete()
	return nil
}

func (m *MockProvider) Completion(ctx context.Context, req *service.CompletionRequest) (*models.GenerateResponse, error) {
	return &models.GenerateResponse{
		Answer: "Mock completion response",
	}, nil
}

func (m *MockProvider) CompletionStream(ctx context.Context, req *service.CompletionRequest, handler service.SSEHandler) error {
	handler.OnEvent("message", map[string]interface{}{
		"answer": "Mock completion streaming",
	})
	return nil
}

func (m *MockProvider) GetAppInfo(ctx context.Context) (*models.AppInfo, error) {
	return &models.AppInfo{
		Name: "Mock App",
		Mode: models.AppModeChat,
	}, nil
}

// RetryableProvider 带重试功能的包装器
type RetryableProvider struct {
	provider AIProvider
	maxRetries int
	backoff    time.Duration
}

func NewRetryableProvider(provider AIProvider, maxRetries int, backoff time.Duration) *RetryableProvider {
	return &RetryableProvider{
		provider:   provider,
		maxRetries: maxRetries,
		backoff:    backoff,
	}
}

func (r *RetryableProvider) Chat(ctx context.Context, req *service.ChatRequest) (*models.GenerateResponse, error) {
	var lastErr error
	
	for attempt := 0; attempt <= r.maxRetries; attempt++ {
		if attempt > 0 {
			time.Sleep(r.backoff * time.Duration(attempt))
		}
		
		resp, err := r.provider.Chat(ctx, req)
		if err == nil {
			return resp, nil
		}
		
		lastErr = err
	}
	
	return nil, fmt.Errorf("达到最大重试次数: %w", lastErr)
}

func (r *RetryableProvider) ChatStream(ctx context.Context, req *service.ChatRequest, handler service.SSEHandler) error {
	return r.provider.ChatStream(ctx, req, handler)
}

func (r *RetryableProvider) Completion(ctx context.Context, req *service.CompletionRequest) (*models.GenerateResponse, error) {
	var lastErr error
	
	for attempt := 0; attempt <= r.maxRetries; attempt++ {
		if attempt > 0 {
			time.Sleep(r.backoff * time.Duration(attempt))
		}
		
		resp, err := r.provider.Completion(ctx, req)
		if err == nil {
			return resp, nil
		}
		
		lastErr = err
	}
	
	return nil, fmt.Errorf("达到最大重试次数: %w", lastErr)
}

func (r *RetryableProvider) CompletionStream(ctx context.Context, req *service.CompletionRequest, handler service.SSEHandler) error {
	return r.provider.CompletionStream(ctx, req, handler)
}

func (r *RetryableProvider) GetAppInfo(ctx context.Context) (*models.AppInfo, error) {
	return r.provider.GetAppInfo(ctx)
}

// CachedProvider 带缓存功能的包装器
type CachedProvider struct {
	provider AIProvider
	cache    map[string]*models.GenerateResponse
}

func NewCachedProvider(provider AIProvider) *CachedProvider {
	return &CachedProvider{
		provider: provider,
		cache:    make(map[string]*models.GenerateResponse),
	}
}

func (c *CachedProvider) Chat(ctx context.Context, req *service.ChatRequest) (*models.GenerateResponse, error) {
	// 简单的缓存键
	cacheKey := req.Query + req.User
	
	if cached, exists := c.cache[cacheKey]; exists {
		return cached, nil
	}
	
	resp, err := c.provider.Chat(ctx, req)
	if err == nil {
		c.cache[cacheKey] = resp
	}
	
	return resp, err
}

func (c *CachedProvider) ChatStream(ctx context.Context, req *service.ChatRequest, handler service.SSEHandler) error {
	return c.provider.ChatStream(ctx, req, handler)
}

func (c *CachedProvider) Completion(ctx context.Context, req *service.CompletionRequest) (*models.GenerateResponse, error) {
	return c.provider.Completion(ctx, req)
}

func (c *CachedProvider) CompletionStream(ctx context.Context, req *service.CompletionRequest, handler service.SSEHandler) error {
	return c.provider.CompletionStream(ctx, req, handler)
}

func (c *CachedProvider) GetAppInfo(ctx context.Context) (*models.AppInfo, error) {
	return c.provider.GetAppInfo(ctx)
}

// AIService 业务服务层
type AIService struct {
	provider AIProvider
}

func NewAIService(provider AIProvider) *AIService {
	return &AIService{
		provider: provider,
	}
}

func (s *AIService) ProcessUserQuery(ctx context.Context, query, user string) (*models.GenerateResponse, error) {
	req := &service.ChatRequest{
		Query: query,
		User:  user,
	}
	
	return s.provider.Chat(ctx, req)
}

func (s *AIService) ProcessBatchQueries(ctx context.Context, queries []string, user string) ([]*models.GenerateResponse, error) {
	responses := make([]*models.GenerateResponse, len(queries))
	
	for i, query := range queries {
		resp, err := s.ProcessUserQuery(ctx, query, user)
		if err != nil {
			return nil, err
		}
		responses[i] = resp
	}
	
	return responses, nil
}

// 演示接口解耦的使用
func demonstrateInterfaceDesign() {
	fmt.Println("=== 接口设计示例 ===")
	
	// 1. 使用真实Dify提供者
	difyProvider := NewDifyProvider("your-token", "https://api.dify.ai")
	
	// 2. 使用模拟提供者进行测试
	mockProvider := NewMockProvider()
	mockProvider.SetResponse("test query", &models.GenerateResponse{
		Answer: "Mock response for test query",
	})
	
	// 3. 包装重试功能
	retryableProvider := NewRetryableProvider(difyProvider, 3, time.Second)
	
	// 4. 包装缓存功能
	cachedProvider := NewCachedProvider(difyProvider)
	
	// 5. 创建业务服务
	service := NewAIService(cachedProvider)
	
	// 使用服务
	ctx := context.Background()
	resp, err := service.ProcessUserQuery(ctx, "Hello", "user123")
	if err != nil {
		fmt.Printf("处理查询失败: %v\n", err)
		return
	}
	
	fmt.Printf("响应: %s\n", resp.Answer)
	
	// 批量处理
	queries := []string{"Query 1", "Query 2", "Query 3"}
	responses, err := service.ProcessBatchQueries(ctx, queries, "user123")
	if err != nil {
		fmt.Printf("批量处理失败: %v\n", err)
		return
	}
	
	for i, resp := range responses {
		fmt.Printf("查询 %d: %s\n", i+1, resp.Answer)
	}
}

func main() {
	demonstrateInterfaceDesign()
}