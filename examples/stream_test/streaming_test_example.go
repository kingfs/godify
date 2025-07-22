package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"

	dify "github.com/kingfs/godify"
	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/service"
)

func main() {
	fmt.Println("=== SSE 流式处理测试 ===")

	// 创建模拟服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 设置SSE响应头
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		// 模拟流式聊天响应
		sseData := `data: {"event": "message", "id": "msg-123", "conversation_id": "conv-456"}

data: {"event": "message_delta", "delta": "你好"}

data: {"event": "message_delta", "delta": "！我是"}

data: {"event": "message_delta", "delta": "Dify AI"}

data: {"event": "message_delta", "delta": "助手。"}

data: {"event": "message_end", "id": "msg-123", "conversation_id": "conv-456", "answer": "你好！我是Dify AI助手。"}

`

		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(sseData))
		if err != nil {
			fmt.Printf("Failed to write SSE data: %v\n", err)
		}
	}))
	defer server.Close()

	// 创建客户端
	client := dify.NewServiceClient("test-token", server.URL)

	// 创建聊天请求
	req := &service.ChatRequest{
		Inputs: map[string]interface{}{},
		Query:  "你好",
		User:   "test-user",
	}

	// 创建测试处理器
	handler := &TestStreamHandler{}

	// 执行流式聊天
	err := client.ChatStream(context.Background(), req, handler)
	if err != nil {
		fmt.Printf("❌ 流式测试失败: %v\n", err)
	} else {
		fmt.Printf("✅ 流式测试成功完成!\n")
		fmt.Printf("📊 处理统计:\n")
		fmt.Printf("  - 接收到 %d 个事件\n", handler.EventCount)
		fmt.Printf("  - message_delta 事件: %d 个\n", handler.DeltaCount)
		fmt.Printf("  - 完整内容: \"%s\"\n", handler.FullContent)
	}
}

// TestStreamHandler 测试流式处理器
type TestStreamHandler struct {
	EventCount  int
	DeltaCount  int
	FullContent string
}

// OnEvent 处理SSE事件
func (h *TestStreamHandler) OnEvent(event *client.SSEEvent) error {
	h.EventCount++

	if event.Data == "" {
		return nil
	}

	fmt.Printf("📥 收到事件: %s\n", event.Data)

	// 简单的JSON解析来提取事件类型和delta
	if contains(event.Data, `"event": "message_delta"`) {
		h.DeltaCount++
		// 提取delta内容
		if delta := extractDelta(event.Data); delta != "" {
			h.FullContent += delta
			fmt.Printf("   → 增量内容: \"%s\"\n", delta)
		}
	} else if contains(event.Data, `"event": "message"`) {
		fmt.Printf("   → 消息开始\n")
	} else if contains(event.Data, `"event": "message_end"`) {
		fmt.Printf("   → 消息结束\n")
	}

	return nil
}

// OnError 处理错误
func (h *TestStreamHandler) OnError(err error) {
	fmt.Printf("❌ 测试处理器错误: %v\n", err)
}

// OnComplete 处理完成
func (h *TestStreamHandler) OnComplete() {
	fmt.Printf("🎉 测试处理器完成!\n")
}

// 辅助函数

func contains(s, substr string) bool {
	return len(s) >= len(substr) && findSubstring(s, substr) != -1
}

func findSubstring(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

func extractDelta(data string) string {
	// 简单的delta提取 - 查找 "delta": "内容"
	start := findSubstring(data, `"delta": "`)
	if start == -1 {
		return ""
	}
	start += len(`"delta": "`)

	end := start
	for end < len(data) && data[end] != '"' {
		end++
	}

	if end < len(data) {
		return data[start:end]
	}
	return ""
}
