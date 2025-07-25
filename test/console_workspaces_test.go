package test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/kingfs/godify/console"
	"github.com/kingfs/godify/models"
)

// mock server，路径和返回内容严格参考workspaces.go接口实现
func SetupWorkspacesMockServer() *httptest.Server {
	handler := http.NewServeMux()

	// /workspaces/current/members
	handler.HandleFunc("/console/api/workspaces/current/members", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resp := models.WorkspacesCurrentMembersApiResponse{
			Accounts: []models.WorkspaceMember{
				{
					ID:           "m1",
					Name:         "mock member",
					Email:        "mock@mock.com",
					Role:         "admin",
					Status:       "active",
					Avatar:       nil,
					AvatarURL:    nil,
					CreatedAt:    1753153422,
					LastActiveAt: 1753347058,
					LastLoginAt:  1753153422,
				},
			},
		}
		json.NewEncoder(w).Encode(resp)
	})

	// /workspaces/current/members/invite-email
	handler.HandleFunc("/console/api/workspaces/current/members/invite-email", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resp := models.WorkspaceInviteEmailApiResponse{
			Result:   "success",
			InvitationResults: []models.WorkspaceInvitationResult{
				{Status: "invited", Email: "test1@example.com"},
				{Status: "invited", Email: "test2@example.com"},
			},
			TenantID: "tenant-1",
		}
		json.NewEncoder(w).Encode(resp)
	})

	// /workspaces/current/members/{memberID} (DELETE) 和 /workspaces/current/members/{memberID}/update-role (PUT)
	handler.HandleFunc("/console/api/workspaces/current/members/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "DELETE" && !strings.HasSuffix(r.URL.Path, "/update-role") {
			resp := models.WorkspaceOperationResponse{
				Result:   "success",
				TenantID: "tenant-1",
			}
			json.NewEncoder(w).Encode(resp)
			return
		}
		if r.Method == "PUT" && strings.HasSuffix(r.URL.Path, "/update-role") {
			resp := models.WorkspaceUpdateRoleResponse{
				Result: "success",
			}
			json.NewEncoder(w).Encode(resp)
			return
		}
	})

	// /workspaces/current/dataset-operators
	handler.HandleFunc("/console/api/workspaces/current/dataset-operators", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resp := models.WorkspacesCurrentDatasetOperatorsApiResponse{
			Result: "success",
			Accounts: []models.WorkspaceMember{
				{
					ID:           "op1",
					Name:         "mock operator",
					Email:        "operator@mock.com",
					Role:         "operator",
					Status:       "active",
					Avatar:       nil,
					AvatarURL:    nil,
					CreatedAt:    1753153422,
					LastActiveAt: 1753347058,
					LastLoginAt:  1753153422,
				},
			},
		}
		json.NewEncoder(w).Encode(resp)
	})

	return httptest.NewServer(handler)
}

// 支持自定义 baseURL
func NewClientWithBaseURL(baseURL string) *console.Client {
	_ = godotenv.Load("../.env")
	auth := os.Getenv("authorization")
	workspaceID := os.Getenv("workspace_id")
	client := console.NewClient(auth, baseURL)
	client.WithWorkspaceID(workspaceID)
	return client
}

func TestGetWorkspaces(t *testing.T) {
	// mockServer := SetupWorkspacesMockServer()
	// defer mockServer.Close()

	test_url := "http://localhost"
	client := NewClientWithBaseURL(test_url)
	resp, err := client.GetWorkspaces(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp: %+v", *resp)
}

func TestGetWorkspacesCurrent(t *testing.T) {
	// mockServer := SetupWorkspacesMockServer()
	// defer mockServer.Close()

	test_url := "http://localhost"
	client := NewClientWithBaseURL(test_url)
	resp, err := client.GetWorkspacesCurrent(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp: %+v", *resp)
}

func TestGetWorkspacesCurrentMembers(t *testing.T) {
	// mockServer := SetupWorkspacesMockServer()
	// defer mockServer.Close()

	test_url := "http://localhost"
	client := NewClientWithBaseURL(test_url)
	resp, err := client.GetWorkspacesCurrentMembers(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp: %+v", *resp)
}

func TestCreateWorkspacesCurrentMembersInviteEmail(t *testing.T) {
	// mockServer := SetupWorkspacesMockServer()
	// defer mockServer.Close()

	test_url := "http://localhost"

	client := NewClientWithBaseURL(test_url)

	emails := []string{"test1@example.com", "test2@example.com"}
	role := "normal"
	language := "zh-CN"
	resp, err := client.CreateWorkspacesCurrentMembersInviteEmail(context.Background(), emails, role, language)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp: %+v", *resp)
}

func TestDeleteWorkspacesCurrentMembers(t *testing.T) {
	// mockServer := SetupWorkspacesMockServer()
	// defer mockServer.Close()

	test_url := "http://localhost"
	client := NewClientWithBaseURL(test_url)

	memberID := "member-id-placeholder"
	resp, err := client.DeleteWorkspacesCurrentMembers(context.Background(), memberID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp: %+v", *resp)
}

func TestUpdateWorkspacesCurrentMembersRole(t *testing.T) {
	// mockServer := SetupWorkspacesMockServer()
	// defer mockServer.Close()

	test_url := "http://localhost"
	client := NewClientWithBaseURL(test_url)

	memberID := "member-id-placeholder"
	role := "admin"
	resp, err := client.UpdateWorkspacesCurrentMembersRole(context.Background(), memberID, role)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp: %+v", *resp)
}

func TestGetWorkspacesCurrentDatasetOperators(t *testing.T) {
	// mockServer := SetupWorkspacesMockServer()
	// defer mockServer.Close()

	test_url := "http://localhost"
	client := NewClientWithBaseURL(test_url)
	resp, err := client.GetWorkspacesCurrentDatasetOperators(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp: %+v", *resp)
}