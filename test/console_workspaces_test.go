package test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kingfs/godify/models"
)

var workspaces_use_real_url = false

// mock server，路径和返回内容严格参考workspaces.go接口实现
func SetupWorkspacesMockServer() *httptest.Server {
	handler := http.NewServeMux()
	// /workspaces
	handler.HandleFunc("/console/api/workspaces", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resp := models.WorkspacesApiResponse{
			Workspaces: []models.Workspace{
				{
					ID:   "w1",
					Name: "mock workspace",
				},
			},
		}
		json.NewEncoder(w).Encode(resp)
	})

	// /workspaces/current
	handler.HandleFunc("/console/api/workspaces/current", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resp := models.Workspace{
			ID:   "w1",
			Name: "mock workspace",
		}
		json.NewEncoder(w).Encode(resp)
	})

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
			Result: "success",
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

	handler.HandleFunc("/console/api/workspaces/current/model-providers", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resp := models.ModelProvidersResponse{
			Data: []models.ModelProvider{
				{
					Background:         nil,
					ConfigurateMethods: []string{"customizable-model"},
					CustomConfiguration: models.CustomConfiguration{
						Status: "active",
					},
					Description: map[string]string{
						"en_US":   "Model providers compatible with OpenAI's API standard, such as LM Studio.",
						"zh_Hans": "兼容 OpenAI API 的模型供应商，例如 LM Studio 。",
					},
					Help:      nil,
					IconLarge: nil,
					IconSmall: map[string]string{
						"en_US":   "/console/api/workspaces/xxx/model-providers/langgenius/openai_api_compatible/openai_api_compatible/icon_small/en_US",
						"zh_Hans": "/console/api/workspaces/xxx/model-providers/langgenius/openai_api_compatible/openai_api_compatible/icon_small/zh_Hans",
					},
					Label: map[string]string{
						"en_US":   "OpenAI-API-compatible",
						"zh_Hans": "OpenAI-API-compatible",
					},
					ModelCredentialSchema: models.ModelCredentialSchema{
						CredentialFormSchemas: []models.CredentialFormSchema{
							{
								Default: nil,
								Label: map[string]string{
									"en_US":   "Model display name",
									"zh_Hans": "模型显示名称",
								},
								MaxLength: 0,
								Options:   nil,
								Placeholder: map[string]string{
									"en_US":   "The display name of the model in the interface.",
									"zh_Hans": "模型在界面的显示名称",
								},
								Required: false,
								ShowOn:   nil,
								Type:     "text-input",
								Variable: "display_name",
							},
							// 可继续添加其他字段
						},
					},
					Model: models.ModelInfo{
						Label: map[string]string{
							"en_US":   "Model Name",
							"zh_Hans": "模型名称",
						},
						Placeholder: map[string]string{
							"en_US":   "Enter full model name",
							"zh_Hans": "输入模型全称",
						},
					},
					PreferredProviderType:    "custom",
					Provider:                 "langgenius/openai_api_compatible/openai_api_compatible",
					ProviderCredentialSchema: nil,
					SupportedModelTypes:      []string{"llm", "rerank", "text-embedding", "speech2text", "tts"},
					SystemConfiguration: models.SystemConfiguration{
						CurrentQuotaType:    nil,
						Enabled:             false,
						QuotaConfigurations: []interface{}{},
					},
					TenantID: "8b72a5d2-31cb-4ca5-a5e2-b7e4b79064b3",
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})
	// /workspaces/current/dataset-operators
	handler.HandleFunc("/console/api/workspaces/current/dataset-operators", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resp := models.WorkspacesCurrentDatasetOperatorsApiResponse{
			Accounts: []models.WorkspaceMember{
				{
					ID:           "op1",
					Name:         "dataset operator",
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

	// /workspaces/current/model-providers/{provider}/models
	handler.HandleFunc("/console/api/workspaces/current/model-providers/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		path := r.URL.Path
		// 判断是否是/models结尾
		if strings.HasSuffix(path, "/models") && r.Method == "GET" {
			resp := models.ModelProviderModelsResponse{
				Data: []models.ModelProviderModel{
					{
						Model:     "gpt-3.5-turbo",
						Label:     map[string]string{"en_US": "GPT-3.5 Turbo", "zh_Hans": "GPT-3.5 Turbo"},
						ModelType: "llm",
						Features:  []string{"chat", "completion"},
						FetchFrom: "openai",
						ModelProperties: map[string]interface{}{
							"context_length": 4096,
						},
						Deprecated:           false,
						Status:               "active",
						LoadBalancingEnabled: false,
						Provider: models.ModelProviderInfo{
							Provider: "openai",
							Label: map[string]string{
								"en_US":   "OpenAI",
								"zh_Hans": "OpenAI",
							},
							IconSmall: map[string]string{
								"en_US":   "https://openai.com/icon.png",
								"zh_Hans": "https://openai.com/icon.png",
							},
							IconLarge:           nil,
							SupportedModelTypes: []string{"llm", "text-embedding"},
							Models:              []string{"gpt-3.5-turbo", "text-embedding-ada-002"},
							TenantID:            "tenant-1",
						},
					},
				},
			}
			json.NewEncoder(w).Encode(resp)
			return
		}
		// 判断是否是/models结尾且POST（新增/更新模型）
		if strings.HasSuffix(path, "/models") && r.Method == "POST" {
			var reqBody map[string]interface{}
			json.NewDecoder(r.Body).Decode(&reqBody)
			resp := map[string]interface{}{
				"result":   "success",
				"provider": strings.Split(strings.TrimPrefix(path, "/console/api/workspaces/current/model-providers/"), "/")[0],
				"model":    reqBody["model"],
			}
			json.NewEncoder(w).Encode(resp)
			return
		}
		// 其他情况返回404
		http.NotFound(w, r)
	})
	return httptest.NewServer(handler)
}

func TestGetWorkspaces(t *testing.T) {
	mockServer := SetupWorkspacesMockServer()
	defer mockServer.Close()

	client := TestNewClientWithBaseURL(mockServer.URL, workspaces_use_real_url)
	resp, err := client.GetWorkspaces(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp: %+v", *resp)
}

func TestGetWorkspacesCurrent(t *testing.T) {
	mockServer := SetupWorkspacesMockServer()
	defer mockServer.Close()

	client := TestNewClientWithBaseURL(mockServer.URL, workspaces_use_real_url)
	resp, err := client.GetWorkspacesCurrent(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp: %+v", *resp)
}

func TestGetWorkspacesCurrentMembers(t *testing.T) {
	mockServer := SetupWorkspacesMockServer()
	defer mockServer.Close()

	client := TestNewClientWithBaseURL(mockServer.URL, workspaces_use_real_url)
	resp, err := client.GetWorkspacesCurrentMembers(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp: %+v", *resp)
}

func TestCreateWorkspacesCurrentMembersInviteEmail(t *testing.T) {
	mockServer := SetupWorkspacesMockServer()
	defer mockServer.Close()

	client := TestNewClientWithBaseURL(mockServer.URL, workspaces_use_real_url)

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
	mockServer := SetupWorkspacesMockServer()
	defer mockServer.Close()

	client := TestNewClientWithBaseURL(mockServer.URL, workspaces_use_real_url)

	memberID := "4cd7b103-3a7b-4784-ba15-2be1911554e4"
	resp, err := client.DeleteWorkspacesCurrentMembers(context.Background(), memberID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp: %+v", *resp)
}

func TestUpdateWorkspacesCurrentMembersRole(t *testing.T) {
	mockServer := SetupWorkspacesMockServer()
	defer mockServer.Close()

	client := TestNewClientWithBaseURL(mockServer.URL, workspaces_use_real_url)

	memberID := "3780d271-1f6b-42b3-95b5-b90dd6893764"
	role := "editor"
	resp, err := client.UpdateWorkspacesCurrentMembersRole(context.Background(), memberID, role)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp: %+v", *resp)
}

func TestGetWorkspacesCurrentDatasetOperators(t *testing.T) {
	mockServer := SetupWorkspacesMockServer()
	defer mockServer.Close()

	client := TestNewClientWithBaseURL(mockServer.URL, workspaces_use_real_url)
	resp, err := client.GetWorkspacesCurrentDatasetOperators(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp: %+v", *resp)
}

func TestGetModelProviderList(t *testing.T) {
	mockServer := SetupWorkspacesMockServer()
	defer mockServer.Close()

	client := TestNewClientWithBaseURL(mockServer.URL, workspaces_use_real_url)
	resp, err := client.GetModelProviderList(context.Background(), "llm")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp: %+v", *resp)
}

func TestGetModelProviderModels(t *testing.T) {
	mockServer := SetupWorkspacesMockServer()
	defer mockServer.Close()

	client := TestNewClientWithBaseURL(mockServer.URL, workspaces_use_real_url)
	resp, err := client.GetModelProviderModels(context.Background(), "langgenius/openai_api_compatible/openai_api_compatible")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp: %+v", *resp)
}

func TestUpdateModelProviderModel(t *testing.T) {
	mockServer := SetupWorkspacesMockServer()
	defer mockServer.Close()

	client := TestNewClientWithBaseURL(mockServer.URL, workspaces_use_real_url)
	model := "qwen3-14b"
	modelType := "llm"
	credentials := map[string]interface{}{
		"mode":                      "chat",
		"context_size":              "4096",
		"max_tokens_to_sample":      "4096",
		"agent_though_support":      "not_supported",
		"function_calling_type":     "no_call",
		"stream_function_calling":   "not_supported",
		"vision_support":            "no_support",
		"structured_output_support": "not_supported",
		"stream_mode_delimiter":     "\\n\\n",
		"voices":                    "alloy",
		"api_key":                   "sk-yGO8ahcqnCQ7jjXIlWiPRYB8OzxwulqefllfQnrTgFBM17zo",
		"endpoint_url":              "https://aiapi.chaitin.net/v1",
	}
	loadBalancing := map[string]interface{}{
		"enabled": false,
		"configs": []interface{}{},
	}
	resp, err := client.UpdateModelProviderModel(context.Background(), "langgenius/openai_api_compatible/openai_api_compatible", model, modelType, credentials, loadBalancing, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp: %+v", resp)
}
