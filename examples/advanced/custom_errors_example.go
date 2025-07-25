package main

import (
	"context"
	"fmt"
	"time"

	"github.com/kingfs/godify"
	"github.com/kingfs/godify/errors"
	"github.com/kingfs/godify/service"
)

// CustomError 自定义错误类型
type CustomError struct {
	Code       string
	Message    string
	Details    map[string]interface{}
	Retryable  bool
	StatusCode int
	Original   error
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func (e *CustomError) Unwrap() error {
	return e.Original
}

// BusinessError 业务错误
type BusinessError struct {
	*CustomError
	BusinessCode string
	UserID       string
	Timestamp    time.Time
}

func NewBusinessError(code, message, businessCode, userID string, original error) *BusinessError {
	return &BusinessError{
		CustomError: &CustomError{
			Code:      code,
			Message:   message,
			Retryable: false,
			Original:  original,
		},
		BusinessCode: businessCode,
		UserID:       userID,
		Timestamp:    time.Now(),
	}
}

// RateLimitError 速率限制错误
type RateLimitError struct {
	*CustomError
	RetryAfter time.Duration
	Limit      int
	Remaining  int
}

func NewRateLimitError(retryAfter time.Duration, limit, remaining int, original error) *RateLimitError {
	return &RateLimitError{
		CustomError: &CustomError{
			Code:      "rate_limit_exceeded",
			Message:   fmt.Sprintf("Rate limit exceeded. Retry after %v", retryAfter),
			Retryable: true,
			Original:  original,
		},
		RetryAfter: retryAfter,
		Limit:      limit,
		Remaining:  remaining,
	}
}

// QuotaError 配额错误
type QuotaError struct {
	*CustomError
	QuotaType string
	Used      int
	Limit     int
	ResetTime time.Time
}

func NewQuotaError(quotaType string, used, limit int, resetTime time.Time, original error) *QuotaError {
	return &QuotaError{
		CustomError: &CustomError{
			Code:      "quota_exceeded",
			Message:   fmt.Sprintf("Quota exceeded for %s: %d/%d", quotaType, used, limit),
			Retryable: false,
			Original:  original,
		},
		QuotaType: quotaType,
		Used:      used,
		Limit:     limit,
		ResetTime: resetTime,
	}
}

// NetworkError 网络错误
type NetworkError struct {
	*CustomError
	Endpoint    string
	RetryCount  int
	MaxRetries  int
}

func NewNetworkError(endpoint string, retryCount, maxRetries int, original error) *NetworkError {
	return &NetworkError{
		CustomError: &CustomError{
			Code:      "network_error",
			Message:   fmt.Sprintf("Network error for %s (attempt %d/%d)", endpoint, retryCount, maxRetries),
			Retryable: retryCount < maxRetries,
			Original:  original,
		},
		Endpoint:   endpoint,
		RetryCount: retryCount,
		MaxRetries: maxRetries,
	}
}

// ErrorHandler 错误处理器
type ErrorHandler struct {
	logger func(format string, args ...interface{})
}

func NewErrorHandler(logger func(format string, args ...interface{})) *ErrorHandler {
	return &ErrorHandler{
		logger: logger,
	}
}

func (h *ErrorHandler) HandleError(err error, context map[string]interface{}) {
	switch e := err.(type) {
	case *BusinessError:
		h.handleBusinessError(e, context)
	case *RateLimitError:
		h.handleRateLimitError(e, context)
	case *QuotaError:
		h.handleQuotaError(e, context)
	case *NetworkError:
		h.handleNetworkError(e, context)
	default:
		h.handleGenericError(err, context)
	}
}

func (h *ErrorHandler) handleBusinessError(err *BusinessError, context map[string]interface{}) {
	h.logger("业务错误: [%s] %s, 用户: %s, 业务代码: %s", 
		err.Code, err.Message, err.UserID, err.BusinessCode)
	
	// 记录到业务日志
	h.logger("错误详情: %+v", context)
}

func (h *ErrorHandler) handleRateLimitError(err *RateLimitError, context map[string]interface{}) {
	h.logger("速率限制: 剩余 %d/%d, 重试时间: %v", 
		err.Remaining, err.Limit, err.RetryAfter)
	
	// 实现指数退避
	if err.Retryable {
		backoff := time.Duration(err.RetryAfter.Seconds()) * time.Second
		time.Sleep(backoff)
	}
}

func (h *ErrorHandler) handleQuotaError(err *QuotaError, context map[string]interface{}) {
	h.logger("配额超限: %s %d/%d, 重置时间: %v", 
		err.QuotaType, err.Used, err.Limit, err.ResetTime)
	
	// 检查是否需要升级套餐
	if err.Used >= err.Limit {
		h.logger("建议升级套餐或等待配额重置")
	}
}

func (h *ErrorHandler) handleNetworkError(err *NetworkError, context map[string]interface{}) {
	h.logger("网络错误: %s (尝试 %d/%d)", 
		err.Endpoint, err.RetryCount, err.MaxRetries)
	
	if err.Retryable {
		backoff := time.Duration(err.RetryCount) * time.Second
		time.Sleep(backoff)
	}
}

func (h *ErrorHandler) handleGenericError(err error, context map[string]interface{}) {
	h.logger("通用错误: %v", err)
	h.logger("错误上下文: %+v", context)
}

// ErrorRecovery 错误恢复策略
type ErrorRecovery struct {
	handler *ErrorHandler
}

func NewErrorRecovery(handler *ErrorHandler) *ErrorRecovery {
	return &ErrorRecovery{
		handler: handler,
	}
}

func (r *ErrorRecovery) ExecuteWithRecovery(ctx context.Context, fn func() error, context map[string]interface{}) error {
	err := fn()
	if err != nil {
		r.handler.HandleError(err, context)
		
		// 检查是否可重试
		if customErr, ok := err.(*CustomError); ok && customErr.Retryable {
			return r.retryWithBackoff(ctx, fn, context)
		}
	}
	return err
}

func (r *ErrorRecovery) retryWithBackoff(ctx context.Context, fn func() error, context map[string]interface{}) error {
	maxRetries := 3
	baseDelay := time.Second
	
	for attempt := 1; attempt <= maxRetries; attempt++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		
		// 指数退避
		delay := baseDelay * time.Duration(attempt)
		time.Sleep(delay)
		
		err := fn()
		if err == nil {
			return nil
		}
		
		r.handler.HandleError(err, map[string]interface{}{
			"attempt": attempt,
			"max_retries": maxRetries,
		})
		
		// 检查是否仍然可重试
		if customErr, ok := err.(*CustomError); !ok || !customErr.Retryable {
			return err
		}
	}
	
	return fmt.Errorf("达到最大重试次数")
}

// 演示自定义错误处理
func demonstrateCustomErrors() {
	fmt.Println("=== 自定义错误处理示例 ===")
	
	// 创建错误处理器
	handler := NewErrorHandler(func(format string, args ...interface{}) {
		fmt.Printf("[ERROR] "+format+"\n", args...)
	})
	
	// 创建错误恢复策略
	recovery := NewErrorRecovery(handler)
	
	// 模拟不同的错误场景
	ctx := context.Background()
	
	// 1. 业务错误
	businessErr := NewBusinessError(
		"invalid_input",
		"用户输入无效",
		"USER_INPUT_ERROR",
		"user123",
		fmt.Errorf("validation failed"),
	)
	handler.HandleError(businessErr, map[string]interface{}{
		"input": "invalid data",
		"field": "email",
	})
	
	// 2. 速率限制错误
	rateLimitErr := NewRateLimitError(
		30*time.Second,
		100,
		0,
		fmt.Errorf("rate limit exceeded"),
	)
	handler.HandleError(rateLimitErr, map[string]interface{}{
		"endpoint": "/v1/chat",
		"user_id":  "user123",
	})
	
	// 3. 配额错误
	quotaErr := NewQuotaError(
		"daily_requests",
		1000,
		1000,
		time.Now().Add(24*time.Hour),
		fmt.Errorf("daily quota exceeded"),
	)
	handler.HandleError(quotaErr, map[string]interface{}{
		"plan": "basic",
		"user_id": "user123",
	})
	
	// 4. 网络错误
	networkErr := NewNetworkError(
		"https://api.dify.ai/v1/chat",
		2,
		3,
		fmt.Errorf("connection timeout"),
	)
	handler.HandleError(networkErr, map[string]interface{}{
		"timeout": "30s",
		"retry_count": 2,
	})
	
	// 5. 使用错误恢复策略
	err := recovery.ExecuteWithRecovery(ctx, func() error {
		// 模拟可能失败的操作
		return fmt.Errorf("simulated error")
	}, map[string]interface{}{
		"operation": "chat_request",
		"user_id":   "user123",
	})
	
	if err != nil {
		fmt.Printf("操作最终失败: %v\n", err)
	}
}

// 扩展的错误处理客户端
type ErrorAwareClient struct {
	client  *service.Client
	handler *ErrorHandler
	recovery *ErrorRecovery
}

func NewErrorAwareClient(token, baseURL string) *ErrorAwareClient {
	handler := NewErrorHandler(func(format string, args ...interface{}) {
		fmt.Printf("[CLIENT] "+format+"\n", args...)
	})
	
	return &ErrorAwareClient{
		client:   dify.NewServiceClient(token, baseURL),
		handler:  handler,
		recovery: NewErrorRecovery(handler),
	}
}

func (c *ErrorAwareClient) Chat(ctx context.Context, req *service.ChatRequest) (*models.GenerateResponse, error) {
	var result *models.GenerateResponse
	
	err := c.recovery.ExecuteWithRecovery(ctx, func() error {
		var chatErr error
		result, chatErr = c.client.Chat(ctx, req)
		return chatErr
	}, map[string]interface{}{
		"operation": "chat",
		"user":      req.User,
		"query":     req.Query,
	})
	
	return result, err
}

func main() {
	demonstrateCustomErrors()
}