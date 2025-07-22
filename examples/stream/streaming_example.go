package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	dify "github.com/kingfs/godify"
	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/service"
)

func main() {
	// 替换为你的 App API Token 和 Dify 服务器地址
	appToken := "your-app-api-token"
	baseURL := "https://api.dify.ai"

	// 创建 Service API 客户端
	serviceClient := dify.NewServiceClient(appToken, baseURL)

	// 流式聊天示例
	fmt.Println("=== 流式聊天示例 ===")
	err := streamingChatExample(serviceClient)
	if err != nil {
		log.Printf("流式聊天失败: %v", err)
	}

	// 流式文本补全示例
	fmt.Println("\n=== 流式文本补全示例 ===")
	err = streamingCompletionExample(serviceClient)
	if err != nil {
		log.Printf("流式补全失败: %v", err)
	}
}

// streamingChatExample 流式聊天示例
func streamingChatExample(client *service.Client) error {
	// 创建聊天请求
	req := &service.ChatRequest{
		Inputs: map[string]interface{}{},
		Query:  "请写一首关于春天的诗，要求逐字生成",
		User:   "user-123",
	}

	// 创建流式处理器
	handler := &ChatStreamHandler{
		conversationID: "",
		messageID:      "",
		fullContent:    "",
	}

	// 执行流式聊天
	return client.ChatStream(context.Background(), req, handler)
}

// streamingCompletionExample 流式补全示例
func streamingCompletionExample(client *service.Client) error {
	// 创建补全请求
	req := &service.CompletionRequest{
		Inputs: map[string]interface{}{
			"query": "解释一下什么是机器学习，要求详细说明",
		},
		Query: "解释一下什么是机器学习，要求详细说明",
		User:  "user-123",
	}

	// 创建流式处理器
	handler := &CompletionStreamHandler{
		messageID:   "",
		fullContent: "",
	}

	// 执行流式补全
	return client.CompletionStream(context.Background(), req, handler)
}

// ChatStreamHandler 聊天流式处理器
type ChatStreamHandler struct {
	conversationID string
	messageID      string
	fullContent    string
}

// OnEvent 处理SSE事件
func (h *ChatStreamHandler) OnEvent(event *client.SSEEvent) error {
	if event.Data == "" {
		return nil
	}

	// 解析事件数据
	var data map[string]interface{}
	if err := parseJSONEvent(event.Data, &data); err != nil {
		fmt.Printf("⚠️ 解析事件数据失败: %v\n", err)
		return nil
	}

	// 根据事件类型处理
	eventType := getEventType(data)
	switch eventType {
	case "message":
		return h.handleMessage(data)
	case "message_delta":
		return h.handleMessageDelta(data)
	case "message_end":
		return h.handleMessageEnd(data)
	case "error":
		return h.handleError(data)
	default:
		fmt.Printf("📄 收到事件 [%s]: %v\n", eventType, data)
	}

	return nil
}

func (h *ChatStreamHandler) handleMessage(data map[string]interface{}) error {
	h.conversationID = getString(data, "conversation_id")
	h.messageID = getString(data, "id")

	fmt.Printf("💬 开始新消息 [对话ID: %s, 消息ID: %s]\n",
		truncateString(h.conversationID, 8),
		truncateString(h.messageID, 8))

	return nil
}

func (h *ChatStreamHandler) handleMessageDelta(data map[string]interface{}) error {
	delta := getString(data, "delta")
	if delta != "" {
		fmt.Print(delta)
		h.fullContent += delta
	}
	return nil
}

func (h *ChatStreamHandler) handleMessageEnd(data map[string]interface{}) error {
	fmt.Printf("\n✅ 消息完成! 完整内容长度: %d 字符\n", len(h.fullContent))
	fmt.Printf("📝 完整回复:\n%s\n", h.fullContent)
	return nil
}

func (h *ChatStreamHandler) handleError(data map[string]interface{}) error {
	errorMsg := getString(data, "message")
	fmt.Printf("❌ 发生错误: %s\n", errorMsg)
	return fmt.Errorf("chat error: %s", errorMsg)
}

// OnError 处理错误
func (h *ChatStreamHandler) OnError(err error) {
	fmt.Printf("❌ 流处理错误: %v\n", err)
}

// OnComplete 处理完成
func (h *ChatStreamHandler) OnComplete() {
	fmt.Printf("🎉 聊天流处理完成!\n")
}

// CompletionStreamHandler 补全流式处理器
type CompletionStreamHandler struct {
	messageID   string
	fullContent string
}

// OnEvent 处理SSE事件
func (h *CompletionStreamHandler) OnEvent(event *client.SSEEvent) error {
	if event.Data == "" {
		return nil
	}

	// 解析事件数据
	var data map[string]interface{}
	if err := parseJSONEvent(event.Data, &data); err != nil {
		fmt.Printf("⚠️ 解析事件数据失败: %v\n", err)
		return nil
	}

	// 根据事件类型处理
	eventType := getEventType(data)
	switch eventType {
	case "message":
		return h.handleMessage(data)
	case "message_delta":
		return h.handleMessageDelta(data)
	case "message_end":
		return h.handleMessageEnd(data)
	case "error":
		return h.handleError(data)
	default:
		fmt.Printf("📄 收到事件 [%s]: %v\n", eventType, data)
	}

	return nil
}

func (h *CompletionStreamHandler) handleMessage(data map[string]interface{}) error {
	h.messageID = getString(data, "id")
	fmt.Printf("📝 开始文本补全 [消息ID: %s]\n", truncateString(h.messageID, 8))
	return nil
}

func (h *CompletionStreamHandler) handleMessageDelta(data map[string]interface{}) error {
	delta := getString(data, "delta")
	if delta != "" {
		fmt.Print(delta)
		h.fullContent += delta
	}
	return nil
}

func (h *CompletionStreamHandler) handleMessageEnd(data map[string]interface{}) error {
	fmt.Printf("\n✅ 补全完成! 完整内容长度: %d 字符\n", len(h.fullContent))
	fmt.Printf("📝 完整内容:\n%s\n", h.fullContent)
	return nil
}

func (h *CompletionStreamHandler) handleError(data map[string]interface{}) error {
	errorMsg := getString(data, "message")
	fmt.Printf("❌ 发生错误: %s\n", errorMsg)
	return fmt.Errorf("completion error: %s", errorMsg)
}

// OnError 处理错误
func (h *CompletionStreamHandler) OnError(err error) {
	fmt.Printf("❌ 流处理错误: %v\n", err)
}

// OnComplete 处理完成
func (h *CompletionStreamHandler) OnComplete() {
	fmt.Printf("🎉 补全流处理完成!\n")
}

// 辅助函数

func parseJSONEvent(data string, result interface{}) error {
	return json.Unmarshal([]byte(data), result)
}

func getEventType(data map[string]interface{}) string {
	if event, ok := data["event"].(string); ok {
		return event
	}
	return ""
}

func getString(data map[string]interface{}, key string) string {
	if val, ok := data[key].(string); ok {
		return val
	}
	return ""
}

func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
