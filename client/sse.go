package client

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

// SSEEvent SSE事件
type SSEEvent struct {
	Event string
	Data  string
	ID    string
	Retry string
}

// SSEHandler SSE事件处理器
type SSEHandler interface {
	OnEvent(event *SSEEvent) error
	OnError(err error)
	OnComplete()
}

// StreamResponse 流式响应处理
func (c *BaseClient) StreamResponse(ctx context.Context, req *Request, handler SSEHandler) error {
	// 执行请求
	resp, err := c.Do(ctx, req)
	if err != nil {
		return err
	}

	// 检查Content-Type是否为SSE
	contentType := resp.Headers.Get("Content-Type")
	if !strings.Contains(contentType, "text/event-stream") && !strings.Contains(contentType, "text/plain") {
		return fmt.Errorf("unexpected content type for streaming response: %s", contentType)
	}

	// 解析SSE流
	return c.parseSSEStream(resp.Body, handler)
}

// parseSSEStream 解析SSE数据流
func (c *BaseClient) parseSSEStream(data []byte, handler SSEHandler) error {
	defer handler.OnComplete()

	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	var event SSEEvent

	for scanner.Scan() {
		line := scanner.Text()

		// 空行表示事件结束
		if line == "" {
			if event.Data != "" || event.Event != "" {
				if err := handler.OnEvent(&event); err != nil {
					handler.OnError(err)
					return err
				}
			}
			event = SSEEvent{}
			continue
		}

		// 解析SSE字段
		if strings.HasPrefix(line, "data: ") {
			event.Data = strings.TrimPrefix(line, "data: ")
		} else if strings.HasPrefix(line, "event: ") {
			event.Event = strings.TrimPrefix(line, "event: ")
		} else if strings.HasPrefix(line, "id: ") {
			event.ID = strings.TrimPrefix(line, "id: ")
		} else if strings.HasPrefix(line, "retry: ") {
			event.Retry = strings.TrimPrefix(line, "retry: ")
		}
	}

	if err := scanner.Err(); err != nil {
		handler.OnError(err)
		return err
	}

	return nil
}

// JSONSSEHandler 将SSE事件解析为JSON的处理器
type JSONSSEHandler struct {
	OnEventFunc    func(eventType string, data map[string]interface{}) error
	OnErrorFunc    func(err error)
	OnCompleteFunc func()
}

// OnEvent 处理SSE事件
func (h *JSONSSEHandler) OnEvent(event *SSEEvent) error {
	if event.Data == "" {
		return nil
	}

	// 解析JSON数据
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(event.Data), &data); err != nil {
		// 如果不是JSON格式，作为原始数据处理
		data = map[string]interface{}{
			"raw_data": event.Data,
		}
	}

	if h.OnEventFunc != nil {
		return h.OnEventFunc(event.Event, data)
	}

	return nil
}

// OnError 处理错误
func (h *JSONSSEHandler) OnError(err error) {
	if h.OnErrorFunc != nil {
		h.OnErrorFunc(err)
	}
}

// OnComplete 处理完成
func (h *JSONSSEHandler) OnComplete() {
	if h.OnCompleteFunc != nil {
		h.OnCompleteFunc()
	}
}
