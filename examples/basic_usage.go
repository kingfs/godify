package main

import (
	"context"
	"fmt"
	"time"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/logger"
	"github.com/kingfs/godify/metrics"
)

func main() {
	// 创建日志器
	logConfig := &logger.Config{
		Level:  logger.InfoLevel,
		Format: "json",
		Output: "stdout",
	}
	
	log, err := logger.NewLogger(logConfig)
	if err != nil {
		panic(fmt.Sprintf("Failed to create logger: %v", err))
	}

	// 创建监控
	metrics := metrics.NewMetrics(true)

	// 创建客户端配置
	config := &client.ClientConfig{
		BaseURL:             "https://api.dify.ai",
		AuthType:            client.AuthTypeBearer,
		Token:               "your-token-here",
		Timeout:             30 * time.Second,
		MaxRetries:          3,
		Logger:              log,
		Metrics:             metrics,
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     90 * time.Second,
	}

	// 创建客户端
	client := client.NewBaseClient(config)

	// 创建请求
	req := &client.Request{
		Method: "POST",
		Path:   "/v1/chat-messages",
		Body: map[string]interface{}{
			"inputs": map[string]interface{}{},
			"query":  "Hello, how are you?",
			"user":   "user123",
		},
	}

	// 执行请求
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp, err := client.Do(ctx, req)
	if err != nil {
		fmt.Printf("Request failed: %v\n", err)
		return
	}

	// 处理响应
	fmt.Printf("Response Status: %d\n", resp.StatusCode)
	fmt.Printf("Response Body: %s\n", string(resp.Body))

	// 获取监控统计
	stats := metrics.GetStats()
	fmt.Printf("Metrics: %+v\n", stats)
}