package console

import (
	"context"
	"strconv"
	"time"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/models"
)

// Client Console API客户端 (管理员API)
type Client struct {
	baseClient *client.BaseClient
}

// NewClient 创建Console API客户端
// 注意: Console API通常需要session认证，这里使用Bearer token作为临时方案
func NewClient(accessToken, baseURL string) *Client {
	config := &client.ClientConfig{
		BaseURL:    baseURL + "/console/api",
		AuthType:   client.AuthTypeBearer,
		Token:      accessToken,
		Timeout:    30 * time.Second,
		MaxRetries: 3,
	}

	return &Client{
		baseClient: client.NewBaseClient(config),
	}
}

// NewClientWithSession 使用Session Cookie创建Console API客户端
func NewClientWithSession(sessionCookie, baseURL string) *Client {
	config := &client.ClientConfig{
		BaseURL:    baseURL + "/console/api",
		Timeout:    30 * time.Second,
		MaxRetries: 3,
	}

	baseClient := client.NewBaseClient(config)
	// 这里需要自定义HTTP客户端来处理session cookie
	// 实际实现中需要设置Cookie头

	return &Client{
		baseClient: baseClient,
	}
}

func (c *Client) WithWorkspaceID(workspaceID string) *Client {
	c.baseClient.WithWorkspaceID(workspaceID)
	return c
}

// ============ 认证相关 ============

// Login 用户登录
func (c *Client) Login(ctx context.Context, req *models.LoginRequest) (*models.LoginResponse, error) {
	httpReq := &client.Request{
		Method: "POST",
		Path:   "/auth/login",
		Body:   req,
	}

	var result models.LoginResponse
	err := c.baseClient.DoJSON(ctx, httpReq, &result)
	return &result, err
}

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

// ============ 数据集管理 ============

// GetDatasets 获取数据集列表
func (c *Client) GetDatasets(ctx context.Context, page, limit int, keyword string, tagIDs []string, includeAll bool) (*models.DatasetListResponse, error) {
	query := make(map[string]string)
	if page > 0 {
		query["page"] = strconv.Itoa(page)
	}
	if limit > 0 {
		query["limit"] = strconv.Itoa(limit)
	}
	if keyword != "" {
		query["keyword"] = keyword
	}
	if len(tagIDs) > 0 {
		// 使用多个tag_ids参数
		for _, tagID := range tagIDs {
			// 注意：实际实现中需要支持重复参数名
			query["tag_ids"] = tagID
		}
	}
	if includeAll {
		query["include_all"] = "true"
	}

	req := &client.Request{
		Method: "GET",
		Path:   "/datasets",
		Query:  query,
	}

	var result models.DatasetListResponse
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// CreateDataset 创建数据集
func (c *Client) CreateDataset(ctx context.Context, req *models.CreateDatasetRequest) (*models.Dataset, error) {
	httpReq := &client.Request{
		Method: "POST",
		Path:   "/datasets",
		Body:   req,
	}

	var result models.Dataset
	err := c.baseClient.DoJSON(ctx, httpReq, &result)
	return &result, err
}

// GetDataset 获取数据集详情
func (c *Client) GetDataset(ctx context.Context, datasetID string) (*models.Dataset, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/datasets/" + datasetID,
	}

	var result models.Dataset
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// UpdateDataset 更新数据集
func (c *Client) UpdateDataset(ctx context.Context, datasetID string, req *models.UpdateDatasetRequest) (*models.Dataset, error) {
	httpReq := &client.Request{
		Method: "PATCH",
		Path:   "/datasets/" + datasetID,
		Body:   req,
	}

	var result models.Dataset
	err := c.baseClient.DoJSON(ctx, httpReq, &result)
	return &result, err
}

// DeleteDataset 删除数据集
func (c *Client) DeleteDataset(ctx context.Context, datasetID string) error {
	req := &client.Request{
		Method: "DELETE",
		Path:   "/datasets/" + datasetID,
	}

	var result models.OperationResponse
	return c.baseClient.DoJSON(ctx, req, &result)
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

// GetDatasetAPIKeys 获取数据集API密钥列表
func (c *Client) GetDatasetAPIKeys(ctx context.Context, datasetID string) (*models.APIKeyListResponse, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/datasets/" + datasetID + "/api-keys",
	}

	var result models.APIKeyListResponse
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// CreateDatasetAPIKey 创建数据集API密钥
func (c *Client) CreateDatasetAPIKey(ctx context.Context, datasetID string) (*models.APIKey, error) {
	req := &client.Request{
		Method: "POST",
		Path:   "/datasets/" + datasetID + "/api-keys",
	}

	var result models.APIKey
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// DeleteDatasetAPIKey 删除数据集API密钥
func (c *Client) DeleteDatasetAPIKey(ctx context.Context, datasetID, keyID string) error {
	req := &client.Request{
		Method: "DELETE",
		Path:   "/datasets/" + datasetID + "/api-keys/" + keyID,
	}

	var result models.OperationResponse
	return c.baseClient.DoJSON(ctx, req, &result)
}
