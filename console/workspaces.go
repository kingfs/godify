package console

// test pr
import (
	"context"
	"errors"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/models"
)

// GetCurrentTenant 获取当前租户
func (c *Client) GetCurrentTenant(ctx context.Context) (*models.Tenant, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/workspaces/current",
	}

	var result models.Tenant
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

func (c *Client) GetWorkspaces(ctx context.Context) (*models.WorkspacesApiResponse, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/workspaces",
	}
	var resp models.WorkspacesApiResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

func (c *Client) GetWorkspacesCurrent(ctx context.Context) (*models.Workspace, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/workspaces/current",
	}
	var resp models.Workspace
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

func (c *Client) GetWorkspacesCurrentMembers(ctx context.Context) (*models.WorkspacesCurrentMembersApiResponse, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/workspaces/current/members",
	}
	var resp models.WorkspacesCurrentMembersApiResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

// 邀请成员邮件接口
func (c *Client) CreateWorkspacesCurrentMembersInviteEmail(ctx context.Context, emails []string, role, language string) (*models.WorkspaceInviteEmailApiResponse, error) {
	role_types := []string{"normal", "editor", "admin"}
	// 如果role不在role_types里，报错
	valid := false
	for _, r := range role_types {
		if role == r {
			valid = true
			break
		}
	}
	if !valid {
		err := errors.New("参数错误：role 必须为 normal、editor 或 admin")
		return nil, err
	}
	req := &client.Request{
		Method: "POST",
		Path:   "/workspaces/current/members/invite-email",
		Body: map[string]interface{}{
			"emails":   emails,
			"role":     role,
			"language": language,
		},
	}
	var resp models.WorkspaceInviteEmailApiResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

// 取消成员邀请/移除成员接口
func (c *Client) DeleteWorkspacesCurrentMembers(ctx context.Context, memberID string) (*models.WorkspaceOperationResponse, error) {
	path := "/workspaces/current/members/" + memberID
	req := &client.Request{
		Method: "DELETE",
		Path:   path,
	}
	var resp models.WorkspaceOperationResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

// 更新成员角色接口
func (c *Client) UpdateWorkspacesCurrentMembersRole(ctx context.Context, memberID, role string) (*models.WorkspaceUpdateRoleResponse, error) {
	path := "/workspaces/current/members/" + memberID + "/update-role"
	req := &client.Request{
		Method: "PUT",
		Path:   path,
		Body: map[string]interface{}{
			"role": role,
		},
	}
	var resp models.WorkspaceUpdateRoleResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

// 获取数据集操作员成员列表接口
func (c *Client) GetWorkspacesCurrentDatasetOperators(ctx context.Context) (*models.WorkspacesCurrentDatasetOperatorsApiResponse, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/workspaces/current/dataset-operators",
	}
	var resp models.WorkspacesCurrentDatasetOperatorsApiResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

// model_providers相关接口
func (c *Client) GetModelProviderList(ctx context.Context, model_type string) (*models.ModelProvidersResponse, error) {
	switch models.ModelType(model_type) {
	case models.LLM, models.TEXT_EMBEDDING, models.RERANK, models.SPEECH2TEXT, models.MODERATION, models.TTS:
		// 合法，继续执行
	default:
		return nil, errors.New("参数错误：model_type 不合法")
	}
	req := &client.Request{
		Method: "GET",
		Path:   "/workspaces/current/model-providers",
		Query: map[string]string{
			"model_type": model_type,
		},
	}
	var resp models.ModelProvidersResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

// 获取指定 provider 下的所有模型
func (c *Client) GetModelProviderModels(ctx context.Context, provider string) (*models.ModelProviderModelsResponse, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/workspaces/current/model-providers/" + provider + "/models",
	}
	var resp models.ModelProviderModelsResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

// 新增/更新 provider 下的模型
func (c *Client) UpdateModelProviderModel(ctx context.Context, provider, model, modelType string, credentials, loadBalancing map[string]interface{}, configFrom *string) (map[string]interface{}, error) {
	reqBody := map[string]interface{}{
		"model":      model,
		"model_type": modelType,
	}
	if credentials != nil {
		reqBody["credentials"] = credentials
	}
	if loadBalancing != nil {
		reqBody["load_balancing"] = loadBalancing
	}
	if configFrom != nil {
		reqBody["config_from"] = configFrom
	}
	req := &client.Request{
		Method: "POST",
		Path:   "/workspaces/current/model-providers/" + provider + "/models",
		Body:   reqBody,
	}
	var resp map[string]interface{}
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return resp, err
}

func (c *Client) GetToolProviderList(ctx context.Context, toolType string) (*models.ToolProviderListDetailResponse, error) {
	tool_types := []string{"builtin", "model", "api", "workflow", "mcp"}
	if toolType != "" {
		valid := false
		for _, t := range tool_types {
			if toolType == t {
				valid = true
				break
			}
		}
		if !valid {
			return nil, errors.New("参数错误：tool_type 不合法，应为[builtin, model, api, workflow, mcp]")
		}
	}

	req := &client.Request{
		Method: "GET",
		Path:   "/workspaces/current/tool-providers",
		Query: map[string]string{
			"tool_type": toolType, // 可选参数，不传则返回所有工具提供者
		},
	}
	var resp models.ToolProviderListDetailResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

func (c *Client) GetToolBuiltinProviderListTools(ctx context.Context, provider string) (*models.BuiltinToolListResponse, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/workspaces/current/tool-provider/builtin/" + provider + "/tools",
	}
	var resp models.BuiltinToolListResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

// 获取内置工具提供商信息接口
func (c *Client) GetToolBuiltinProviderInfo(ctx context.Context, provider string) (*models.ToolProviderEntity, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/workspaces/current/tool-provider/builtin/" + provider + "/info",
	}
	var resp models.ToolProviderEntity
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

// 删除内置工具提供商接口
func (c *Client) DeleteToolBuiltinProvider(ctx context.Context, provider string) (any, error) {
	req := &client.Request{
		Method: "POST",
		Path:   "/workspaces/current/tool-provider/builtin/" + provider + "/delete",
	}
	var resp any
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return resp, err
}

// 获取内置工具提供商图标接口
func (c *Client) GetToolBuiltinProviderIcon(ctx context.Context, provider string) (any, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/workspaces/current/tool-provider/builtin/" + provider + "/icon",
	}
	var resp any
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return resp, err
}

// 更新内置工具提供商credentials接口
func (c *Client) UpdateToolBuiltinProvider(ctx context.Context, provider string, credentials map[string]string) (*models.OperationResponse, error) {
	req := &client.Request{
		Method: "POST",
		Path:   "/workspaces/current/tool-provider/builtin/" + provider + "/update",
		Body: map[string]map[string]string{
			"credentials": credentials,
		},
	}
	var resp models.OperationResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

// 获取内置工具提供商凭据接口
func (c *Client) GetToolBuiltinProviderCredentials(ctx context.Context, provider string) (map[string]interface{}, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/workspaces/current/tool-provider/builtin/" + provider + "/credentials",
	}
	var resp map[string]interface{}
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return resp, err
}

// 获取内置工具提供商凭据模式接口
func (c *Client) GetToolBuiltinProviderCredentialsSchema(ctx context.Context, provider string) (*models.ToolBuiltinProviderCredentialsSchemaResponse, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/workspaces/current/tool-provider/builtin/" + provider + "/credentials_schema",
	}
	var resp models.ToolBuiltinProviderCredentialsSchemaResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

// 添加API工具提供商接口
func (c *Client) AddToolApiProvider(ctx context.Context, provider string, credentials map[string]interface{}, icon map[string]interface{}, schemaType, schema string, labels []string, privacyPolicy *string, customDisclaimer *string) (*models.OperationResponse, error) {
	reqBody := map[string]interface{}{
		"provider":    provider,
		"credentials": credentials,
		"icon":        icon,
		"schema_type": schemaType,
		"schema":      schema,
		"labels":      labels,
	}
	if privacyPolicy != nil {
		reqBody["privacy_policy"] = *privacyPolicy
	}
	if customDisclaimer != nil {
		reqBody["custom_disclaimer"] = *customDisclaimer
	}

	req := &client.Request{
		Method: "POST",
		Path:   "/workspaces/current/tool-provider/api/add",
		Body:   reqBody,
	}
	var resp models.OperationResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

// 获取API工具提供商工具列表接口
func (c *Client) GetToolApiProviderListTools(ctx context.Context, provider string) (*models.ApiToolListResponse, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/workspaces/current/tool-provider/api/tools",
		Query: map[string]string{
			"provider": provider,
		},
	}
	var resp models.ApiToolListResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

// 更新API工具提供商接口
func (c *Client) UpdateToolApiProvider(ctx context.Context, credentials map[string]interface{}, schemaType, schema, provider, originalProvider string, icon map[string]interface{}, privacyPolicy, customDisclaimer string, labels []string) (*models.OperationResponse, error) {
	reqBody := map[string]interface{}{
		"credentials":       credentials,
		"schema_type":       schemaType,
		"schema":            schema,
		"provider":          provider,
		"original_provider": originalProvider,
		"icon":              icon,
		"privacy_policy":    privacyPolicy,
		"custom_disclaimer": customDisclaimer,
	}
	if labels != nil {
		reqBody["labels"] = labels
	}

	req := &client.Request{
		Method: "POST",
		Path:   "/workspaces/current/tool-provider/api/update",
		Body:   reqBody,
	}
	var resp models.OperationResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

// 删除API工具提供商接口
func (c *Client) DeleteToolApiProvider(ctx context.Context, provider string) (*models.OperationResponse, error) {
	reqBody := map[string]interface{}{
		"provider": provider,
	}
	req := &client.Request{
		Method: "POST",
		Path:   "/workspaces/current/tool-provider/api/delete",
		Body:   reqBody,
	}
	var resp models.OperationResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

// CreateMCPProvider 创建 MCP Provider
func (c *Client) CreateMCPProvider(ctx context.Context, headers map[string]string, serverURL, name string, icon any, iconType, iconBackground, serverIdentifier string, sseReadTimeout, timeout float64) (*models.ToolProviderEntity, error) {
	req := &client.Request{
		Method: "POST",
		Path:   "/workspaces/current/tool-provider/mcp",
		Body: map[string]interface{}{
			"headers":           headers,
			"server_url":        serverURL,
			"name":              name,
			"icon":              icon,
			"icon_type":         iconType,
			"icon_background":   iconBackground,
			"server_identifier": serverIdentifier,
			"sse_read_timeout":  sseReadTimeout,
			"timeout":           timeout,
		},
	}
	var resp models.ToolProviderEntity
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

// UpdateMCPProvider 更新 MCP Provider
func (c *Client) UpdateMCPProvider(ctx context.Context, providerID, serverURL, name string, icon any, iconType, iconBackground, serverIdentifier string) (*models.ToolProviderEntity, error) {
	req := &client.Request{
		Method: "PUT",
		Path:   "/workspaces/current/tool-provider/mcp",
		Body: map[string]interface{}{
			"provider_id":       providerID,
			"server_url":        serverURL,
			"name":              name,
			"icon":              icon,
			"icon_type":         iconType,
			"icon_background":   iconBackground,
			"server_identifier": serverIdentifier,
		},
	}
	var resp models.ToolProviderEntity
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

// DeleteMCPProvider 删除 MCP Provider
func (c *Client) DeleteMCPProvider(ctx context.Context, providerID string) (map[string]interface{}, error) {
	req := &client.Request{
		Method: "DELETE",
		Path:   "/workspaces/current/tool-provider/mcp",
		Body: map[string]interface{}{
			"provider_id": providerID,
		},
	}
	var resp map[string]interface{}
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return resp, err
}

// GetMCPProviderDetail 获取 MCP Provider 详情
func (c *Client) GetMCPProviderDetail(ctx context.Context, providerID string) (*models.ToolProviderEntity, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/workspaces/current/tool-provider/mcp/tools/" + providerID,
	}
	var resp models.ToolProviderEntity
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

// UpdateMCPProviderTools 拉取 MCP Provider 工具列表
func (c *Client) UpdateMCPProviderTools(ctx context.Context, providerID string) (*models.OperationResponse, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/workspaces/current/tool-provider/mcp/update/" + providerID,
	}
	var resp models.OperationResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

// AuthMCPProvider MCP Provider 认证
func (c *Client) AuthMCPProvider(ctx context.Context, providerID, authorizationCode string) (*models.OperationResponse, error) {
	req := &client.Request{
		Method: "POST",
		Path:   "/workspaces/current/tool-provider/mcp/auth",
		Body: map[string]interface{}{
			"provider_id": providerID,
		},
	}
	var resp models.OperationResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

func (c *Client) GetTenantList(ctx context.Context, auth_token string) (*models.TenantListResponse, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/workspaces",
		Headers: map[string]string{
			"Authorization": "Bearer " + auth_token,
		},
	}
	var resp models.TenantListResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

func (c *Client) SetModelProvider(ctx context.Context, modelProvider string, request *models.SetModelProvidersRequest) (*models.OperationResponse, error) {
	req := &client.Request{
		Method: "POST",
		Path:   "/workspaces/current/model-providers/" + modelProvider + "/models/credentials",
		Body:   request,
	}
	var resp models.OperationResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

func (c *Client) DeleteModelProvider(ctx context.Context, modelProvider, model, model_type string) (*models.OperationResponse, error) {
	req := &client.Request{
		Method: "DELETE",
		Path:   "/workspaces/current/model-providers/" + modelProvider + "/models",
		Body: map[string]interface{}{
			"model":      model,
			"model_type": model_type,
		},
	}
	var resp models.OperationResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

func (c *Client) UpdatePluginCredential(ctx context.Context, provider string, credential map[string]any, credentialType, credentialName string) (*models.OperationResponse, error) {
	req := &client.Request{
		Method: "POST",
		Path:   "/workspaces/current/tool-provider/builtin/" + provider + "/add",
		Body: map[string]any{
			"credentials": credential,
			"type":        credentialType,
			"name":        credentialName,
		},
	}
	var resp models.OperationResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}
