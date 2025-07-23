package plugin

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/kingfs/godify/client"
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

	config := &client.ClientConfig{
		BaseURL:    baseURL + "/console/api/workspaces/current/plugin",
		AuthType:   client.AuthTypeBearer,
		Token:      authorization,
		Timeout:    30 * time.Second,
		MaxRetries: 3,
		WorkspaceID: &workspaceId,
	}

	return &Client{
		baseClient: client.NewBaseClient(config),
	}
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
}