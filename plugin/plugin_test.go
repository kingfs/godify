package plugin

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic("加载.env文件失败: " + err.Error())
	}
}


func NewTestClient() *Client {
	loadEnv()
	baseURL       := os.Getenv("base_url")
	authorization := os.Getenv("authorization")
	workspaceId   := os.Getenv("workspace_id")
	fmt.Println("baseURL", baseURL)
	fmt.Println("authorization", authorization)
	fmt.Println("workspaceId", workspaceId)

	return PluginNewClient(baseURL, authorization, workspaceId)
}

func TestPluginUploadPkg(t *testing.T) {
	client := NewTestClient()
	if client == nil {
		t.Fatal("未能初始化Client，请实现NewTestClient")
	}
	ctx := context.Background()

	// 假设有一个测试用的pkg文件
	testPkgPath := "./pkgs/file_analyse.difypkg"
	fileData, err := os.ReadFile(testPkgPath)
	if err != nil {
		t.Fatalf("读取测试插件包失败: %v", err)
	}

	filename := filepath.Base(testPkgPath)
	uniqueIdentifier, err := client.PluginUploadPkg(ctx, filename, fileData)
	fmt.Println("uniqueIdentifier", uniqueIdentifier)
	if err != nil {
		t.Fatalf("上传插件包失败: %v", err)
	}
	if strings.TrimSpace(uniqueIdentifier) == "" {
		t.Fatalf("uniqueIdentifier 为空")
	}
}

func TestPluginInstallFromPkg(t *testing.T) {
	client := NewTestClient()
	if client == nil {
		t.Fatal("未能初始化Client，请实现NewTestClient")
	}

	ctx := context.Background()

	installResp, err := client.PluginInstallFromPkg(ctx, "./pkgs/file_analyse.difypkg")
	if err != nil {
		t.Fatalf("安装插件失败: %v", err)
	}
	if installResp == nil {
		t.Fatalf("安装插件返回结果为空")
	}
	fmt.Printf("安装完成！%+v\n", installResp)
}

func TestPluginList(t *testing.T) {
	client := NewTestClient()
	if client == nil {
		t.Fatal("未能初始化Client，请实现NewTestClient")
	}

	ctx := context.Background()
	page := 1
	pageSize := 10

	resp, err := client.PluginList(ctx, page, pageSize)
	if err != nil {
		t.Fatalf("获取插件列表失败: %v", err)
	}
	if resp == nil {
		t.Fatalf("插件列表响应为空")
	}
	fmt.Printf("插件列表: %+v\n", *resp)
}

func TestPluginFetchInstallTasks(t *testing.T) {
	client := NewTestClient()
	if client == nil {
		t.Fatal("未能初始化Client，请实现NewTestClient")
	}

	ctx := context.Background()
	page := 1
	pageSize := 10

	resp, err := client.PluginFetchInstallTasks(ctx, page, pageSize)
	if err != nil {
		t.Fatalf("获取插件安装任务列表失败: %v", err)
	}
	fmt.Printf("插件安装任务列表: %+v\n", *resp)
}

func TestPluginUninstall(t *testing.T) {
	client := NewTestClient()
	if client == nil {
		t.Fatal("未能初始化Client，请实现NewTestClient")
	}
	ctx := context.Background()
	// pluginInstallationID就是list接口返回的插件的ID
	pluginInstallationID := "b27eb75f-1fc4-46c8-97c4-6f5dd67a702f"
	resp, err := client.PluginUninstall(ctx, pluginInstallationID)
	if err != nil {
		t.Fatalf("卸载插件失败: %v", err)
	}
	fmt.Printf("卸载插件完成！%+v\n", *resp)
}
