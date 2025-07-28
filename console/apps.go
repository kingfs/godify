package console

import (
	"context"
	"strconv"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/models"
)

// api.add_resource(CompletionMessageApi, "/apps/<uuid:app_id>/completion-messages")
// api.add_resource(CompletionMessageStopApi, "/apps/<uuid:app_id>/completion-messages/<string:task_id>/stop")
// api.add_resource(ChatMessageApi, "/apps/<uuid:app_id>/chat-messages")
// api.add_resource(ChatMessageStopApi, "/apps/<uuid:app_id>/chat-messages/<string:task_id>/stop")

// api.add_resource(MessageApi, "/apps/<uuid:app_id>/messages/<uuid:message_id>", endpoint="console_message")
// api.add_resource(ChatMessageListApi, "/apps/<uuid:app_id>/chat-messages", endpoint="console_chat_messages")

// ============ 应用管理 ============

// GetApps 获取应用列表
func (c *Client) GetApps(ctx context.Context, page, limit int, mode, name string, tagIDs []string, isCreatedByMe *bool) (*models.ConsoleAppListResponse, error) {
	query := make(map[string]string)
	if page > 0 {
		query["page"] = strconv.Itoa(page)
	}
	if limit > 0 {
		query["limit"] = strconv.Itoa(limit)
	}
	if mode != "" {
		query["mode"] = mode
	}
	if name != "" {
		query["name"] = name
	}
	if len(tagIDs) > 0 {
		// tagIDs需要用逗号分隔
		tagIDsStr := ""
		for i, id := range tagIDs {
			if i > 0 {
				tagIDsStr += ","
			}
			tagIDsStr += id
		}
		query["tag_ids"] = tagIDsStr
	}
	if isCreatedByMe != nil {
		query["is_created_by_me"] = strconv.FormatBool(*isCreatedByMe)
	}

	req := &client.Request{
		Method: "GET",
		Path:   "/apps",
		Query:  query,
	}

	var result models.ConsoleAppListResponse
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// CreateApp 创建应用
func (c *Client) CreateApp(ctx context.Context, req *models.CreateAppRequest) (*models.ConsoleApp, error) {
	httpReq := &client.Request{
		Method: "POST",
		Path:   "/apps",
		Body:   req,
	}

	var result models.ConsoleApp
	err := c.baseClient.DoJSON(ctx, httpReq, &result)
	return &result, err
}

// GetApp 获取应用详情
func (c *Client) GetApp(ctx context.Context, appID string) (*models.ConsoleAppDetail, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/apps/" + appID,
	}

	var result models.ConsoleAppDetail
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// UpdateApp 更新应用
func (c *Client) UpdateApp(ctx context.Context, appID string, req *models.UpdateAppRequest) (*models.ConsoleApp, error) {
	httpReq := &client.Request{
		Method: "PUT",
		Path:   "/apps/" + appID,
		Body:   req,
	}

	var result models.ConsoleApp
	err := c.baseClient.DoJSON(ctx, httpReq, &result)
	return &result, err
}

// DeleteApp 删除应用
func (c *Client) DeleteApp(ctx context.Context, appID string) error {
	req := &client.Request{
		Method: "DELETE",
		Path:   "/apps/" + appID,
	}

	var result models.OperationResponse
	return c.baseClient.DoJSON(ctx, req, &result)
}

// CopyApp 复制应用
func (c *Client) CopyApp(ctx context.Context, appID string, req *models.CopyAppRequest) (*models.ConsoleApp, error) {
	httpReq := &client.Request{
		Method: "POST",
		Path:   "/apps/" + appID + "/copy",
		Body:   req,
	}

	var result models.ConsoleApp
	err := c.baseClient.DoJSON(ctx, httpReq, &result)
	return &result, err
}

// ExportApp 导出应用
func (c *Client) ExportApp(ctx context.Context, appID string, includeSecret bool) (*models.AppExportResponse, error) {
	query := map[string]string{
		"include_secret": strconv.FormatBool(includeSecret),
	}

	req := &client.Request{
		Method: "GET",
		Path:   "/apps/" + appID + "/export",
		Query:  query,
	}

	var result models.AppExportResponse
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// UpdateAppName 更新应用名称
func (c *Client) UpdateAppName(ctx context.Context, appID string, req *models.UpdateAppNameRequest) (*models.ConsoleApp, error) {
	httpReq := &client.Request{
		Method: "POST",
		Path:   "/apps/" + appID + "/name",
		Body:   req,
	}

	var result models.ConsoleApp
	err := c.baseClient.DoJSON(ctx, httpReq, &result)
	return &result, err
}

// UpdateAppIcon 更新应用图标
func (c *Client) UpdateAppIcon(ctx context.Context, appID string, req *models.UpdateAppIconRequest) (*models.ConsoleApp, error) {
	httpReq := &client.Request{
		Method: "POST",
		Path:   "/apps/" + appID + "/icon",
		Body:   req,
	}

	var result models.ConsoleApp
	err := c.baseClient.DoJSON(ctx, httpReq, &result)
	return &result, err
}

// UpdateAppSiteStatus 更新应用站点状态
func (c *Client) UpdateAppSiteStatus(ctx context.Context, appID string, req *models.UpdateAppSiteStatusRequest) (*models.ConsoleApp, error) {
	httpReq := &client.Request{
		Method: "POST",
		Path:   "/apps/" + appID + "/site-enable",
		Body:   req,
	}

	var result models.ConsoleApp
	err := c.baseClient.DoJSON(ctx, httpReq, &result)
	return &result, err
}

// UpdateAppAPIStatus 更新应用API状态
func (c *Client) UpdateAppAPIStatus(ctx context.Context, appID string, req *models.UpdateAppAPIStatusRequest) (*models.ConsoleApp, error) {
	httpReq := &client.Request{
		Method: "POST",
		Path:   "/apps/" + appID + "/api-enable",
		Body:   req,
	}

	var result models.ConsoleApp
	err := c.baseClient.DoJSON(ctx, httpReq, &result)
	return &result, err
}

// GetAppTrace 获取应用追踪配置
func (c *Client) GetAppTrace(ctx context.Context, appID string) (*models.AppTraceConfig, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/apps/" + appID + "/trace",
	}

	var result models.AppTraceConfig
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// UpdateAppTrace 更新应用追踪配置
func (c *Client) UpdateAppTrace(ctx context.Context, appID string, req *models.UpdateAppTraceRequest) error {
	httpReq := &client.Request{
		Method: "POST",
		Path:   "/apps/" + appID + "/trace",
		Body:   req,
	}

	var result models.OperationResponse
	return c.baseClient.DoJSON(ctx, httpReq, &result)
}

func (c *Client) GetAppsChatMessageList(ctx context.Context, appID string, conversationID string, firstID, limit *int) (*models.AppsChatMessageListApiResponse, error) {
	reqQuery := map[string]string{"conversation_id": conversationID}
	if firstID != nil {
		reqQuery["first_id"] = strconv.Itoa(*firstID)
	}
	if limit != nil {
		reqQuery["limit"] = strconv.Itoa(*limit)
	}
	req := &client.Request{
		Method: "GET",
		Path:   "/apps/" + appID + "/chat-messages",
		Query:  reqQuery,
	}
	var resp models.AppsChatMessageListApiResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

func (c *Client) GetAppsMessage(ctx context.Context, appID string, messageID string) (*models.AppsMessageApiResponse, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/apps/" + appID + "/messages/" + messageID,
	}
	var resp models.AppsMessageApiResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

// AppImport 导入应用
// POST /apps/import
func (c *Client) AppImport(ctx context.Context, mode string, yamlContent string) (*models.AppImportResponse, error) {
	httpReq := &client.Request{
		Method: "POST",
		Path:   "/apps/imports",
		Body: map[string]interface{}{
			"mode":         mode,
			"yaml_content": yamlContent,
		},
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	var resp models.AppImportResponse
	err := c.baseClient.DoJSON(ctx, httpReq, &resp)
	return &resp, err
}

// ============ API Key管理 ============

// GetAppAPIKeys 获取应用API密钥列表
func (c *Client) GetAppAPIKeys(ctx context.Context, appID string) (*models.APIKeyListResponse, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/apps/" + appID + "/api-keys",
	}

	var result models.APIKeyListResponse
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// CreateAppAPIKey 创建应用API密钥
func (c *Client) CreateAppAPIKey(ctx context.Context, appID string) (*models.APIKey, error) {
	req := &client.Request{
		Method: "POST",
		Path:   "/apps/" + appID + "/api-keys",
	}

	var result models.APIKey
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// DeleteAppAPIKey 删除应用API密钥
func (c *Client) DeleteAppAPIKey(ctx context.Context, appID, keyID string) error {
	req := &client.Request{
		Method: "DELETE",
		Path:   "/apps/" + appID + "/api-keys/" + keyID,
	}

	var result models.OperationResponse
	return c.baseClient.DoJSON(ctx, req, &result)
}
