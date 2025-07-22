package main

import (
	"context"
	"fmt"
	"log"

	dify "github.com/kingfs/godify"
	"github.com/kingfs/godify/service"
)

func main() {
	// 替换为你的 App API Token 和 Dify 服务器地址
	appToken := "your-app-api-token"
	baseURL := "https://api.dify.ai"

	// 创建 Service API 客户端
	client := dify.NewServiceClient(appToken, baseURL)

	// 示例1: 获取应用信息
	fmt.Println("=== 获取应用信息 ===")
	appInfo, err := client.GetAppInfo(context.Background())
	if err != nil {
		log.Printf("获取应用信息失败: %v", err)
	} else {
		fmt.Printf("应用名称: %s\n", appInfo.Name)
		fmt.Printf("应用描述: %s\n", appInfo.Description)
		fmt.Printf("应用模式: %s\n", appInfo.Mode)
		fmt.Printf("作者: %s\n", appInfo.AuthorName)
		fmt.Printf("标签: %v\n", appInfo.Tags)
	}

	// 示例2: 获取应用参数
	fmt.Println("\n=== 获取应用参数 ===")
	params, err := client.GetAppParameters(context.Background())
	if err != nil {
		log.Printf("获取应用参数失败: %v", err)
	} else {
		fmt.Printf("用户输入表单项数量: %d\n", len(params.UserInputForm))
		fmt.Printf("系统参数: %v\n", params.SystemParameters)
	}

	// 示例3: 文本补全 (适用于 completion 模式应用)
	fmt.Println("\n=== 文本补全示例 ===")
	completionReq := &service.CompletionRequest{
		Inputs: map[string]interface{}{
			"query": "什么是人工智能？",
		},
		Query: "什么是人工智能？",
		User:  "user-123",
	}

	completionResp, err := client.Completion(context.Background(), completionReq)
	if err != nil {
		log.Printf("文本补全失败: %v", err)
	} else {
		fmt.Printf("补全结果: %s\n", completionResp.Answer)
		fmt.Printf("消息ID: %s\n", completionResp.MessageID)
		fmt.Printf("任务ID: %s\n", completionResp.TaskID)
	}

	// 示例4: 聊天对话 (适用于 chat 模式应用)
	fmt.Println("\n=== 聊天对话示例 ===")
	chatReq := &service.ChatRequest{
		Inputs:           map[string]interface{}{},
		Query:            "你好，我想了解一下你的功能",
		User:             "user-123",
		AutoGenerateName: true,
	}

	chatResp, err := client.Chat(context.Background(), chatReq)
	if err != nil {
		log.Printf("聊天对话失败: %v", err)
	} else {
		fmt.Printf("聊天回复: %s\n", chatResp.Answer)
		fmt.Printf("对话ID: %s\n", chatResp.ConversationID)
		fmt.Printf("消息ID: %s\n", chatResp.MessageID)
	}

	// 示例5: 继续对话
	if chatResp != nil && chatResp.ConversationID != "" {
		fmt.Println("\n=== 继续对话示例 ===")
		followUpReq := &service.ChatRequest{
			Inputs:         map[string]interface{}{},
			Query:          "能给我更详细的解释吗？",
			User:           "user-123",
			ConversationID: chatResp.ConversationID,
		}

		followUpResp, err := client.Chat(context.Background(), followUpReq)
		if err != nil {
			log.Printf("继续对话失败: %v", err)
		} else {
			fmt.Printf("继续对话回复: %s\n", followUpResp.Answer)
		}
	}
}
