package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kingfs/godify/errors"
	"github.com/kingfs/godify/models"
	"github.com/kingfs/godify/service"
)

// 自定义错误处理器
type ErrorHandler struct {
	maxRetries int
	backoff    time.Duration
}

func NewErrorHandler(maxRetries int, backoff time.Duration) *ErrorHandler {
	return &ErrorHandler{
		maxRetries: maxRetries,
		backoff:    backoff,
	}
}

func (h *ErrorHandler) HandleWithRetry(ctx context.Context, client *service.Client, req *service.ChatRequest) (*models.GenerateResponse, error) {
	var lastErr error

	for attempt := 0; attempt <= h.maxRetries; attempt++ {
		if attempt > 0 {
			log.Printf("重试第 %d 次...", attempt)
			time.Sleep(h.backoff * time.Duration(attempt))
		}

		resp, err := client.Chat(ctx, req)
		if err == nil {
			return resp, nil
		}

		lastErr = err

		// 检查错误类型，决定是否重试
		if !h.shouldRetry(err) {
			return nil, err
		}
	}

	return nil, fmt.Errorf("达到最大重试次数: %w", lastErr)
}

func (h *ErrorHandler) shouldRetry(err error) bool {
	if !errors.IsAPIError(err) {
		return false
	}

	apiErr := errors.GetAPIError(err)
	switch apiErr.Code {
	case "rate_limit_exceeded":
		return true
	case "service_unavailable":
		return true
	case "timeout":
		return true
	default:
		return false
	}
}

// 并发处理示例
func concurrentChatExample() {
	client := dify.NewServiceClient("your-token", "https://api.dify.ai")
	handler := NewErrorHandler(3, time.Second)

	requests := []*service.ChatRequest{
		{Query: "第一个问题", User: "user1"},
		{Query: "第二个问题", User: "user2"},
		{Query: "第三个问题", User: "user3"},
	}

	results := make(chan *models.GenerateResponse, len(requests))
	errors := make(chan error, len(requests))

	// 并发处理请求
	for _, req := range requests {
		go func(r *service.ChatRequest) {
			resp, err := handler.HandleWithRetry(context.Background(), client, r)
			if err != nil {
				errors <- err
				return
			}
			results <- resp
		}(req)
	}

	// 收集结果
	for i := 0; i < len(requests); i++ {
		select {
		case resp := <-results:
			fmt.Printf("成功: %s\n", resp.Answer)
		case err := <-errors:
			fmt.Printf("失败: %v\n", err)
		}
	}
}

// 性能监控示例
func performanceMonitoringExample() {
	client := dify.NewServiceClient("your-token", "https://api.dify.ai")

	// 启动性能监控
	go func() {
		ctx := context.Background()
		client.Metrics.StartMetricsServer(ctx, 8080)
	}()

	// 执行请求并监控性能
	start := time.Now()
	resp, err := client.Chat(context.Background(), &service.ChatRequest{
		Query: "性能测试",
		User:  "test-user",
	})

	if err != nil {
		log.Printf("请求失败: %v", err)
		return
	}

	duration := time.Since(start)
	fmt.Printf("请求耗时: %v\n", duration)
	fmt.Printf("响应: %s\n", resp.Answer)

	// 获取统计信息
	stats := client.Metrics.GetStats()
	fmt.Printf("成功率: %.2f%%\n", stats["success_rate"])
	fmt.Printf("平均响应时间: %.2fms\n", stats["avg_response_time"])
}

func main() {
	fmt.Println("=== 错误处理最佳实践示例 ===")

	// 基础错误处理
	client := dify.NewServiceClient("your-token", "https://api.dify.ai")

	resp, err := client.Chat(context.Background(), &service.ChatRequest{
		Query: "Hello",
		User:  "test-user",
	})

	if err != nil {
		if errors.IsAPIError(err) {
			apiErr := errors.GetAPIError(err)
			switch apiErr.Code {
			case "app_unavailable":
				log.Printf("应用不可用: %s", apiErr.Message)
			case "quota_exceeded":
				log.Printf("配额超限: %s", apiErr.Message)
			case "rate_limit_exceeded":
				log.Printf("速率限制: %s", apiErr.Message)
			default:
				log.Printf("API错误: %s", apiErr.Message)
			}
		} else {
			log.Printf("网络错误: %v", err)
		}
		return
	}

	fmt.Printf("响应: %s\n", resp.Answer)

	// 演示并发处理
	fmt.Println("\n=== 并发处理示例 ===")
	concurrentChatExample()

	// 演示性能监控
	fmt.Println("\n=== 性能监控示例 ===")
	performanceMonitoringExample()
}
