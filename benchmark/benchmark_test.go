package main

import (
	"context"
	"testing"
	"time"

	"github.com/kingfs/godify"
	"github.com/kingfs/godify/service"
)

var (
	benchClient *service.Client
	benchCtx    context.Context
)

func init() {
	// 初始化测试客户端
	benchClient = dify.NewServiceClient("benchmark-token", "https://api.dify.ai")
	benchCtx = context.Background()
}

// BenchmarkSingleRequest 测试单次请求性能
func BenchmarkSingleRequest(b *testing.B) {
	req := &service.ChatRequest{
		Query: "Hello, this is a benchmark test",
		User:  "benchmark-user",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := benchClient.Chat(benchCtx, req)
		if err != nil {
			b.Logf("Request failed: %v", err)
		}
	}
}

// BenchmarkConcurrentRequests 测试并发请求性能
func BenchmarkConcurrentRequests(b *testing.B) {
	req := &service.ChatRequest{
		Query: "Concurrent benchmark test",
		User:  "benchmark-user",
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := benchClient.Chat(benchCtx, req)
			if err != nil {
				b.Logf("Concurrent request failed: %v", err)
			}
		}
	})
}

// BenchmarkStreamingResponse 测试流式响应性能
func BenchmarkStreamingResponse(b *testing.B) {
	req := &service.ChatRequest{
		Query:        "Streaming benchmark test",
		User:         "benchmark-user",
		ResponseMode: "streaming",
	}

	handler := &service.JSONSSEHandler{
		OnEventFunc: func(eventType string, data map[string]interface{}) error {
			return nil
		},
		OnErrorFunc: func(err error) {
			b.Logf("Streaming error: %v", err)
		},
		OnCompleteFunc: func() {
			// 流式响应完成
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := benchClient.ChatStream(benchCtx, req, handler)
		if err != nil {
			b.Logf("Streaming request failed: %v", err)
		}
	}
}

// BenchmarkErrorHandling 测试错误处理性能
func BenchmarkErrorHandling(b *testing.B) {
	// 使用无效的token来测试错误处理
	invalidClient := dify.NewServiceClient("invalid-token", "https://api.dify.ai")
	req := &service.ChatRequest{
		Query: "Error handling benchmark",
		User:  "benchmark-user",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := invalidClient.Chat(benchCtx, req)
		if err == nil {
			b.Fatal("Expected error but got none")
		}
	}
}

// BenchmarkRetryMechanism 测试重试机制性能
func BenchmarkRetryMechanism(b *testing.B) {
	// 模拟需要重试的场景
	req := &service.ChatRequest{
		Query: "Retry mechanism benchmark",
		User:  "benchmark-user",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 这里可以模拟网络错误或临时服务不可用
		_, err := benchClient.Chat(benchCtx, req)
		if err != nil {
			// 模拟重试逻辑
			time.Sleep(10 * time.Millisecond)
			_, err = benchClient.Chat(benchCtx, req)
		}
	}
}

// BenchmarkMemoryUsage 测试内存使用情况
func BenchmarkMemoryUsage(b *testing.B) {
	req := &service.ChatRequest{
		Query: "Memory usage benchmark test with a longer message to simulate real usage patterns",
		User:  "benchmark-user",
		Inputs: map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
			"key3": map[string]interface{}{
				"nested": "value",
			},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := benchClient.Chat(benchCtx, req)
		if err != nil {
			b.Logf("Memory benchmark request failed: %v", err)
		}
		// 模拟处理响应
		_ = resp
	}
}

// BenchmarkConnectionPool 测试连接池性能
func BenchmarkConnectionPool(b *testing.B) {
	// 创建配置了连接池的客户端
	config := &service.ClientConfig{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     time.Minute,
	}
	
	poolClient := service.NewClientWithConfig("benchmark-token", "https://api.dify.ai", config)
	req := &service.ChatRequest{
		Query: "Connection pool benchmark",
		User:  "benchmark-user",
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := poolClient.Chat(benchCtx, req)
			if err != nil {
				b.Logf("Connection pool request failed: %v", err)
			}
		}
	})
}

// BenchmarkLargePayload 测试大负载性能
func BenchmarkLargePayload(b *testing.B) {
	// 创建包含大量数据的请求
	largeInputs := make(map[string]interface{})
	for i := 0; i < 100; i++ {
		largeInputs[fmt.Sprintf("key%d", i)] = fmt.Sprintf("value%d", i)
	}

	req := &service.ChatRequest{
		Query:  "Large payload benchmark test",
		User:   "benchmark-user",
		Inputs: largeInputs,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := benchClient.Chat(benchCtx, req)
		if err != nil {
			b.Logf("Large payload request failed: %v", err)
		}
	}
}

// BenchmarkMetricsCollection 测试指标收集性能
func BenchmarkMetricsCollection(b *testing.B) {
	// 启用指标收集的客户端
	metricsClient := dify.NewServiceClientWithMetrics("benchmark-token", "https://api.dify.ai", true)
	req := &service.ChatRequest{
		Query: "Metrics collection benchmark",
		User:  "benchmark-user",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := metricsClient.Chat(benchCtx, req)
		if err != nil {
			b.Logf("Metrics collection request failed: %v", err)
		}
	}
}

// BenchmarkLogging 测试日志记录性能
func BenchmarkLogging(b *testing.B) {
	// 启用详细日志的客户端
	loggingClient := dify.NewServiceClientWithLogging("benchmark-token", "https://api.dify.ai", "debug")
	req := &service.ChatRequest{
		Query: "Logging benchmark test",
		User:  "benchmark-user",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := loggingClient.Chat(benchCtx, req)
		if err != nil {
			b.Logf("Logging request failed: %v", err)
		}
	}
}