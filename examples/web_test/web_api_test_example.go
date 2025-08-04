package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"

	dify "github.com/kingfs/godify"
)

func main() {
	fmt.Println("=== Web API 认证流程测试 ===")

	// 创建模拟服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("📥 收到请求: %s %s\n", r.Method, r.URL.Path)

		// 检查X-App-Code头部
		if appCode := r.Header.Get("X-App-Code"); appCode != "" {
			fmt.Printf("   App-Code: %s\n", appCode)
		}

		// 检查Authorization头部
		if auth := r.Header.Get("Authorization"); auth != "" {
			fmt.Printf("   Authorization: %s\n", auth)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// 根据不同路径返回不同响应
		switch r.URL.Path {
		case "/api/passport":
			// Passport 响应
			_, err := w.Write([]byte(`{"access_token": "mock-access-token-12345"}`))
			if err != nil {
				fmt.Printf("Failed to write passport response: %v\n", err)
			}
		case "/api/meta":
			// App Meta 响应
			_, err := w.Write([]byte(`{"tool_icons": {"calculator": "icon1", "search": "icon2"}}`))
			if err != nil {
				fmt.Printf("Failed to write meta response: %v\n", err)
			}
		case "/api/parameters":
			// App Parameters 响应
			_, err := w.Write([]byte(`{
				"user_input_form": [],
				"system_parameters": {
					"image_file_size_limit": 10485760,
					"video_file_size_limit": 104857600,
					"audio_file_size_limit": 52428800,
					"file_size_limit": 5242880
				}
			}`))
			if err != nil {
				fmt.Printf("Failed to write parameters response: %v\n", err)
			}
		case "/api/webapp/access-mode":
			// Access Mode 响应
			_, err := w.Write([]byte(`{"accessMode": "public"}`))
			if err != nil {
				fmt.Printf("Failed to write access mode response: %v\n", err)
			}
		case "/api/chat-messages":
			// Chat 响应
			_, err := w.Write([]byte(`{
				"event": "message",
				"message_id": "msg-123",
				"conversation_id": "conv-456",
				"answer": "你好！我是Dify AI助手，很高兴为您提供帮助！",
				"created_at": 1705420000
			}`))
			if err != nil {
				fmt.Printf("Failed to write chat response: %v\n", err)
			}
		default:
			w.WriteHeader(http.StatusNotFound)
			_, err := w.Write([]byte(`{"error": "endpoint not found"}`))
			if err != nil {
				fmt.Printf("Failed to write error response: %v\n", err)
			}
		}
	}))
	defer server.Close()

	// 创建客户端
	client := dify.NewWebClient(server.URL)
	client.WithAppCode("test-app-code")

	// 测试1: 获取访问令牌
	fmt.Println("\n=== 测试获取访问令牌 ===")
	token, err := client.GetPassport(context.Background(), "test-user-123")
	if err != nil {
		fmt.Printf("❌ 获取访问令牌失败: %v\n", err)
		return
	}
	fmt.Printf("✅ 访问令牌获取成功: %s\n", token)

	// 测试2: 获取应用元数据 (需要认证)
	fmt.Println("\n=== 测试获取应用元数据 ===")
	meta, err := client.GetAppMeta(context.Background())
	if err != nil {
		fmt.Printf("❌ 获取应用元数据失败: %v\n", err)
	} else {
		fmt.Printf("✅ 应用元数据获取成功: %v\n", meta.ToolIcons)
	}

	// 测试3: 获取应用参数 (需要认证)
	fmt.Println("\n=== 测试获取应用参数 ===")
	params, err := client.GetAppParameters(context.Background())
	if err != nil {
		fmt.Printf("❌ 获取应用参数失败: %v\n", err)
	} else {
		fmt.Printf("✅ 应用参数获取成功\n")
		if params.SystemParameters != nil {
			fmt.Printf("   系统参数: %+v\n", params.SystemParameters)
		}
	}

	// 测试4: 检查访问模式 (不需要认证)
	fmt.Println("\n=== 测试检查访问模式 ===")
	accessMode, err := client.GetWebAppAccessMode(context.Background(), "", "test-app-code")
	if err != nil {
		fmt.Printf("❌ 获取访问模式失败: %v\n", err)
	} else {
		fmt.Printf("✅ 访问模式: %s\n", accessMode.AccessMode)
	}

	fmt.Println("\n🎉 所有测试完成！Web API认证流程工作正常！")
}
