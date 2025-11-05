package main

import (
	"context"
	"fmt"
	"log"

	dify "github.com/kingfs/godify"
	"github.com/kingfs/godify/models"
	"github.com/kingfs/godify/web"
)

func main() {
	// 替换为你的 API Code 和 Dify 服务器地址
	appCode := "your-web-api-code"

	// 创建 Web API 客户端
	client := dify.NewWebClient(appCode)

	// 首先获取访问令牌 (可选，会自动调用)
	fmt.Println("=== 获取访问令牌 ===")
	_, err := client.GetPassport(context.Background(), "")
	if err != nil {
		log.Printf("获取访问令牌失败: %v", err)
		return
	}
	fmt.Println("访问令牌获取成功")

	// 示例1: 获取应用元数据
	fmt.Println("\n=== 获取应用元数据 ===")
	meta, err := client.GetAppMeta(context.Background())
	if err != nil {
		log.Printf("获取应用元数据失败: %v", err)
	} else {
		fmt.Printf("工具图标: %v\n", meta.ToolIcons)
	}

	// 示例2: 获取应用参数
	fmt.Println("\n=== 获取应用参数 ===")
	params, err := client.GetAppParameters(context.Background())
	if err != nil {
		log.Printf("获取应用参数失败: %v", err)
	} else {
		fmt.Printf("用户输入表单项数量: %d\n", len(params.UserInputForm))
		if params.SystemParameters != nil {
			fmt.Printf("系统参数: %+v\n", params.SystemParameters)
		}
	}

	// 示例3: 检查Web应用访问模式 (不需要认证)
	fmt.Println("\n=== 检查访问模式 ===")
	accessMode, err := client.GetWebAppAccessMode(context.Background(), "", appCode)
	if err != nil {
		log.Printf("获取访问模式失败: %v", err)
	} else {
		fmt.Printf("访问模式: %s\n", accessMode.AccessMode)
	}

	// 示例4: 聊天对话
	fmt.Println("\n=== 聊天对话示例 ===")
	chatReq := &web.ChatRequest{
		Inputs: map[string]interface{}{},
		Query:  "你好，我想了解一下你的功能",
	}

	chatResp, err := client.Chat(context.Background(), chatReq)
	if err != nil {
		log.Printf("聊天对话失败: %v", err)
	} else {
		fmt.Printf("聊天回复: %s\n", chatResp.Answer)
		fmt.Printf("对话ID: %s\n", chatResp.ConversationID)
		fmt.Printf("消息ID: %s\n", chatResp.MessageID)
	}

	var conversationID string
	if chatResp != nil {
		conversationID = chatResp.ConversationID
	}

	// 示例5: 获取对话列表
	if conversationID != "" {
		fmt.Println("\n=== 获取对话列表 ===")
		conversations, err := client.GetConversations(context.Background(), "", 20, nil, "-updated_at")
		if err != nil {
			log.Printf("获取对话列表失败: %v", err)
		} else {
			fmt.Printf("对话数量: %d\n", len(conversations.Data))
			for i, conv := range conversations.Data {
				fmt.Printf("对话%d: ID=%s, 名称=%s\n", i+1, conv.ID, conv.Name)
			}
		}

		// 示例6: 重命名对话
		fmt.Println("\n=== 重命名对话 ===")
		newName := "我的测试对话"
		renameReq := &models.ConversationRenameRequest{
			Name: &newName,
		}

		renamedConv, err := client.RenameConversation(context.Background(), conversationID, renameReq)
		if err != nil {
			log.Printf("重命名对话失败: %v", err)
		} else {
			fmt.Printf("重命名成功: %s\n", renamedConv.Name)
		}

		// 示例7: 置顶对话
		fmt.Println("\n=== 置顶对话 ===")
		err = client.PinConversation(context.Background(), conversationID)
		if err != nil {
			log.Printf("置顶对话失败: %v", err)
		} else {
			fmt.Println("对话置顶成功")
		}

		// 示例8: 获取消息列表
		fmt.Println("\n=== 获取消息列表 ===")
		messages, err := client.GetMessages(context.Background(), conversationID, "", 20)
		if err != nil {
			log.Printf("获取消息列表失败: %v", err)
		} else {
			fmt.Printf("消息数量: %d\n", len(messages.Data))
			for i, msg := range messages.Data {
				fmt.Printf("消息%d: Query=%s, Answer=%s\n", i+1, msg.Query, msg.Answer)
			}

			// 示例9: 发送消息反馈
			if len(messages.Data) > 0 {
				fmt.Println("\n=== 发送消息反馈 ===")
				messageID := messages.Data[0].ID
				rating := "like"
				content := "这个回答很有帮助"

				feedback := &models.MessageFeedbackRequest{
					Rating:  &rating,
					Content: &content,
				}

				err = client.SendMessageFeedback(context.Background(), messageID, feedback)
				if err != nil {
					log.Printf("发送反馈失败: %v", err)
				} else {
					fmt.Println("反馈发送成功")
				}

				// 示例10: 获取建议问题
				fmt.Println("\n=== 获取建议问题 ===")
				suggestions, err := client.GetSuggestedQuestions(context.Background(), messageID)
				if err != nil {
					log.Printf("获取建议问题失败: %v", err)
				} else {
					fmt.Printf("建议问题数量: %d\n", len(suggestions.Data))
					for i, question := range suggestions.Data {
						fmt.Printf("建议问题%d: %s\n", i+1, question)
					}
				}
			}
		}
	}
}
