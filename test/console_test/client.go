package console_test

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kingfs/godify/console"
)

func TestNewClientWithBaseURL(baseURL string, use_real_url bool) *console.Client {
	var auth string
	var workspaceID string
	var csrfToken string
	// 是否使用真实测试url
	if use_real_url {
		_ = godotenv.Load("../../.env")
		baseURL = os.Getenv("base_url")
		auth = os.Getenv("authorization")
		workspaceID = os.Getenv("workspace_id")
		csrfToken = os.Getenv("csrf_token")
	} else {
		auth = "1234567890"
		workspaceID = "1234567890"
	}

	client := console.NewClient(auth, baseURL)
	client.WithWorkspaceID(workspaceID)
	client.WithCookies(map[string]string{
		"csrf_token": csrfToken,
	})
	return client
}
