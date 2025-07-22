package main

import (
	"context"
	"fmt"
	"log"

	dify "github.com/kingfs/godify"
	"github.com/kingfs/godify/models"
)

func main() {
	// 替换为你的访问令牌和 Dify 服务器地址
	accessToken := "your-app-api-token"
	baseURL := "https://api.dify.ai"
	workspaceID := "your-workspace-id"

	// 创建 Console API 客户端
	client := dify.NewConsoleClient(accessToken, baseURL)
	client.WithWorkspaceID(workspaceID)

	// 示例1: 获取应用列表
	fmt.Println("=== 获取应用列表 ===")
	apps, err := client.GetApps(context.Background(), 1, 20, "all", "", nil, nil)
	if err != nil {
		log.Printf("获取应用列表失败: %v", err)
	} else {
		fmt.Printf("应用总数: %d\n", apps.Total)
		for i, app := range apps.Data {
			fmt.Printf("应用%d: %s (%s)\n", i+1, app.Name, app.Mode)
		}
	}

	// 示例2: 创建新应用
	fmt.Println("\n=== 创建新应用 ===")
	createReq := &models.CreateAppRequest{
		Name:        "测试应用",
		Description: "这是一个测试应用",
		Mode:        models.AppModeChat,
		IconType:    "emoji",
		Icon:        "🤖",
	}

	newApp, err := client.CreateApp(context.Background(), createReq)
	if err != nil {
		log.Printf("创建应用失败: %v", err)
	} else {
		fmt.Printf("创建成功，应用ID: %s\n", newApp.ID)
		fmt.Printf("应用名称: %s\n", newApp.Name)
		fmt.Printf("应用模式: %s\n", newApp.Mode)

		// 示例3: 更新应用信息
		fmt.Println("\n=== 更新应用信息 ===")
		updateReq := &models.UpdateAppRequest{
			Name:        "更新后的测试应用",
			Description: "这是一个更新后的测试应用",
			Icon:        "🚀",
		}

		updatedApp, err := client.UpdateApp(context.Background(), newApp.ID, updateReq)
		if err != nil {
			log.Printf("更新应用失败: %v", err)
		} else {
			fmt.Printf("更新成功，新名称: %s\n", updatedApp.Name)
		}

		// 示例4: 获取应用详情
		fmt.Println("\n=== 获取应用详情 ===")
		appDetail, err := client.GetApp(context.Background(), newApp.ID)
		if err != nil {
			log.Printf("获取应用详情失败: %v", err)
		} else {
			fmt.Printf("应用详情: %+v\n", appDetail)
		}

		// 示例5: 复制应用
		fmt.Println("\n=== 复制应用 ===")
		copyReq := &models.CopyAppRequest{
			Name:        "复制的测试应用",
			Description: "这是一个复制的应用",
		}

		copiedApp, err := client.CopyApp(context.Background(), newApp.ID, copyReq)
		if err != nil {
			log.Printf("复制应用失败: %v", err)
		} else {
			fmt.Printf("复制成功，新应用ID: %s\n", copiedApp.ID)
		}

		// 示例6: 导出应用
		fmt.Println("\n=== 导出应用 ===")
		exportData, err := client.ExportApp(context.Background(), newApp.ID, false)
		if err != nil {
			log.Printf("导出应用失败: %v", err)
		} else {
			fmt.Printf("导出数据长度: %d 字符\n", len(exportData.Data))
		}

		// 示例7: 创建API密钥
		fmt.Println("\n=== 创建API密钥 ===")
		apiKey, err := client.CreateAppAPIKey(context.Background(), newApp.ID)
		if err != nil {
			log.Printf("创建API密钥失败: %v", err)
		} else {
			fmt.Printf("API密钥ID: %s\n", apiKey.ID)
			fmt.Printf("API密钥: %s\n", apiKey.Token)
		}

		// 示例8: 获取API密钥列表
		fmt.Println("\n=== 获取API密钥列表 ===")
		apiKeys, err := client.GetAppAPIKeys(context.Background(), newApp.ID)
		if err != nil {
			log.Printf("获取API密钥列表失败: %v", err)
		} else {
			fmt.Printf("API密钥数量: %d\n", len(apiKeys.Data))
			for i, key := range apiKeys.Data {
				fmt.Printf("密钥%d: %s (类型: %s)\n", i+1, key.ID, key.Type)
			}
		}
	}

	// 示例9: 数据集管理
	fmt.Println("\n=== 数据集管理 ===")
	datasets, err := client.GetDatasets(context.Background(), 1, 10, "", nil, false)
	if err != nil {
		log.Printf("获取数据集列表失败: %v", err)
	} else {
		fmt.Printf("数据集总数: %d\n", datasets.Total)
		for i, dataset := range datasets.Data {
			fmt.Printf("数据集%d: %s (文档数: %d)\n", i+1, dataset.Name, dataset.DocumentCount)
		}
	}

	// 示例10: 创建数据集
	fmt.Println("\n=== 创建数据集 ===")
	createDatasetReq := &models.CreateDatasetRequest{
		Name:              "测试数据集",
		Description:       "这是一个测试数据集",
		Permission:        "only_me",
		IndexingTechnique: "high_quality",
	}

	newDataset, err := client.CreateDataset(context.Background(), createDatasetReq)
	if err != nil {
		log.Printf("创建数据集失败: %v", err)
	} else {
		fmt.Printf("创建数据集成功，ID: %s\n", newDataset.ID)
		fmt.Printf("数据集名称: %s\n", newDataset.Name)
	}
}
