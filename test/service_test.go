package test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/kingfs/godify/service"
)

var (
	testUser = "test_user"
)

// 每个测试函数中单独加载godotenv
func getServiceClient(t *testing.T) *service.Client {
	baseURL := os.Getenv("GODIFY_BASE_URL")
	appToken := os.Getenv("GODIFY_APP_TOKEN")
	return service.NewClient(appToken, baseURL)
}

func TestService_GetAppParameters(t *testing.T) {
	client := getServiceClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.GetAppParameters(ctx)
	if err != nil {
		t.Fatalf("GetAppParameters error: %v", err)
	}
	if resp == nil {
		t.Fatal("GetAppParameters resp is nil")
	}
	t.Logf("GetAppParameters resp: %+v", resp)
}

func TestService_GetAppMeta(t *testing.T) {
	client := getServiceClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.GetAppMeta(ctx)
	if err != nil {
		t.Fatalf("GetAppMeta error: %v", err)
	}
	if resp == nil {
		t.Fatal("GetAppMeta resp is nil")
	}
	t.Logf("GetAppMeta resp: %+v", resp)
}

func TestService_GetAppInfo(t *testing.T) {
	client := getServiceClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.GetAppInfo(ctx)
	if err != nil {
		t.Fatalf("GetAppInfo error: %v", err)
	}
	if resp == nil {
		t.Fatal("GetAppInfo resp is nil")
	}
	t.Logf("GetAppInfo resp: %+v", resp)
}

func TestService_Completion(t *testing.T) {
	client := getServiceClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	req := &service.CompletionRequest{
		Inputs: map[string]interface{}{
			"input": "你好，世界",
		},
		Query:        "你好，世界",
		ResponseMode: "",
		User:         testUser,
	}

	resp, err := client.Completion(ctx, req)
	if err != nil {
		t.Fatalf("Completion error: %v", err)
	}
	if resp == nil {
		t.Fatal("Completion resp is nil")
	}
	t.Logf("Completion resp: %+v", resp)
}
