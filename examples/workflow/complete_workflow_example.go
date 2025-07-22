package main

import (
	"context"
	"fmt"
	"log"

	dify "github.com/kingfs/godify"
	"github.com/kingfs/godify/models"
)

func main() {
	// 示例：完整的工作流使用流程
	apiKey := "your-web-api-key"
	baseURL := "https://api.dify.ai"

	// 1. 使用Web API运行工作流
	fmt.Println("=== Web API工作流示例 ===")
	webClient := dify.NewWebClient(apiKey, baseURL)

	workflowReq := &models.WorkflowRunRequest{
		Inputs: map[string]interface{}{
			"user_input": "请帮我分析这段文本的情感",
			"text":       "今天是个美好的一天，我感到非常开心！",
		},
		Files: []map[string]interface{}{
			{
				"type":            "image",
				"transfer_method": "remote_url",
				"url":             "https://example.com/image.png",
			},
		},
	}

	workflowResp, err := webClient.RunWorkflow(context.Background(), workflowReq)
	if err != nil {
		log.Printf("运行工作流失败: %v", err)
	} else {
		fmt.Printf("工作流运行成功，任务ID: %s\n", workflowResp.TaskID)
		fmt.Printf("工作流运行ID: %s\n", workflowResp.WorkflowRunID)
		fmt.Printf("事件类型: %s\n", workflowResp.Event)

		// 如果需要停止工作流
		if workflowResp.TaskID != "" {
			fmt.Println("\n=== 停止工作流任务 ===")
			stopResp, err := webClient.StopWorkflowTask(context.Background(), workflowResp.TaskID)
			if err != nil {
				log.Printf("停止工作流失败: %v", err)
			} else {
				fmt.Printf("停止结果: %s\n", stopResp.Result)
			}
		}
	}

	// 2. 使用Dataset API进行数据集管理
	fmt.Println("\n=== Dataset API数据集管理示例 ===")
	datasetClient := dify.NewDatasetClient(apiKey, baseURL)

	// 获取数据集列表
	datasets, err := datasetClient.GetDatasets(context.Background(), 1, 10, "", nil, false)
	if err != nil {
		log.Printf("获取数据集失败: %v", err)
	} else {
		fmt.Printf("数据集数量: %d\n", len(datasets.Data))

		if len(datasets.Data) > 0 {
			datasetID := datasets.Data[0].ID
			fmt.Printf("使用数据集: %s\n", datasets.Data[0].Name)

			// 获取数据集文档
			documents, err := datasetClient.GetDatasetDocuments(context.Background(), datasetID, 1, 5, "")
			if err != nil {
				log.Printf("获取数据集文档失败: %v", err)
			} else {
				fmt.Printf("文档数量: %d\n", len(documents.Data))
			}

			// 进行命中测试
			hitTestReq := &models.HitTestingRequest{
				Query: "什么是人工智能？",
				RetrievalModel: map[string]interface{}{
					"search_method":    "semantic_search",
					"reranking_enable": false,
					"reranking_model": map[string]interface{}{
						"reranking_provider_name": "",
						"reranking_model_name":    "",
					},
					"top_k":                   2,
					"score_threshold_enabled": false,
				},
			}

			hitTestResp, err := datasetClient.HitTestDataset(context.Background(), datasetID, hitTestReq)
			if err != nil {
				log.Printf("命中测试失败: %v", err)
			} else {
				fmt.Printf("命中测试查询: %s\n", hitTestResp.Query)
				fmt.Printf("命中结果数量: %d\n", len(hitTestResp.Records))
				for i, record := range hitTestResp.Records {
					fmt.Printf("结果%d: %s (得分: %.2f)\n", i+1, record.Title, record.Score)
				}
			}
		}
	}

	// 3. 使用Console API进行应用管理
	fmt.Println("\n=== Console API应用管理示例 ===")
	consoleClient := dify.NewConsoleClient(apiKey, baseURL)

	// 获取应用列表
	apps, err := consoleClient.GetApps(context.Background(), 1, 5, "workflow", "", nil, nil)
	if err != nil {
		log.Printf("获取应用列表失败: %v", err)
	} else {
		fmt.Printf("工作流应用数量: %d\n", len(apps.Data))

		if len(apps.Data) > 0 {
			appID := apps.Data[0].ID
			fmt.Printf("管理应用: %s\n", apps.Data[0].Name)

			// 更新应用状态
			statusReq := &models.UpdateAppAPIStatusRequest{
				EnableAPI: true,
			}

			updatedApp, err := consoleClient.UpdateAppAPIStatus(context.Background(), appID, statusReq)
			if err != nil {
				log.Printf("更新应用API状态失败: %v", err)
			} else {
				fmt.Printf("API状态已更新: %t\n", updatedApp.EnableAPI)
			}

			// 获取应用追踪配置
			traceConfig, err := consoleClient.GetAppTrace(context.Background(), appID)
			if err != nil {
				log.Printf("获取追踪配置失败: %v", err)
			} else {
				fmt.Printf("追踪状态: %t\n", traceConfig.Enabled)
				fmt.Printf("追踪提供商: %s\n", traceConfig.TracingProvider)
			}
		}
	}

	// 4. 使用MCP API (示例)
	fmt.Println("\n=== MCP API示例 ===")
	mcpClient := dify.NewMCPClient("https://api.dify.ai")

	// 列出可用工具
	toolsResp, err := mcpClient.ListTools(context.Background())
	if err != nil {
		log.Printf("获取MCP工具列表失败: %v", err)
	} else {
		fmt.Printf("MCP工具响应: %+v\n", toolsResp.Result)
	}

	// 调用工具
	toolCallResp, err := mcpClient.CallTool(context.Background(), "search", map[string]interface{}{
		"query": "人工智能",
		"limit": 5,
	})
	if err != nil {
		log.Printf("调用MCP工具失败: %v", err)
	} else {
		fmt.Printf("工具调用结果: %+v\n", toolCallResp.Result)
	}

	fmt.Println("\n=== 所有示例执行完成 ===")
}
