package plugin

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"context"
	"encoding/json"
	"strconv"

	"github.com/kingfs/godify/client"
)

type Client struct {
	baseClient *client.BaseClient
}

func PluginNewClient(accessToken, baseURL string) *Client {
	config := &client.ClientConfig{
		BaseURL:    baseURL + "/console/api/workspaces/current/plugin",
		AuthType:   client.AuthTypeBearer,
		Token:      accessToken,
		Timeout:    30 * time.Second,
		MaxRetries: 3,
	}

	return &Client{
		baseClient: client.NewBaseClient(config),
	}
}

// PluginGetDebuggingKey 获取插件调试密钥信息
func (c *Client) PluginGetDebuggingKey() {
	// TODO: 实现获取插件调试密钥信息的接口调用
}

// Plugin 插件基本信息
// 字段根据 Python 返回结构推断
// TODO: 可根据实际API返回补充字段
// 示例字段
// ID/Name/Description/Version/Status/CreatedAt/UpdatedAt
//
type Plugin struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Icon          string `json:"icon"`
	LatestVersion string `json:"latest_version"`
	Status        string `json:"status"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

// PluginListResponse 插件列表响应
//
type PluginListResponse struct {
	Plugins []Plugin `json:"plugins"`
	Total   int      `json:"total"`
}

// PluginPermissionResponse 插件权限响应
//
type PluginPermissionResponse struct {
	InstallPermission string `json:"install_permission"`
	DebugPermission   string `json:"debug_permission"`
}

// PluginManifestResponse 插件manifest响应
//
type PluginManifestResponse struct {
	Manifest map[string]interface{} `json:"manifest"`
}

// PluginIconResponse 插件图标响应
// 直接返回二进制数据，不需要结构体

// PluginList 获取插件列表
func (c *Client) PluginList(ctx context.Context, page, pageSize int) (*PluginListResponse, error) {
	query := map[string]string{
		"page":      strconv.Itoa(page),
		"page_size": strconv.Itoa(pageSize),
	}
	req := &client.Request{
		Method: "GET",
		Path:   "/list",
		Query:  query,
	}
	var resp PluginListResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

// PluginUploadPkg 上传插件包（pkg文件）
func (c *Client) PluginUploadPkg(ctx context.Context, filename string, fileData []byte) (string, error) {
	resp, err := c.baseClient.UploadFile(ctx, "/upload/pkg", "pkg", filename, fileData, nil)
	if err != nil {
		return "", err
	}
	var result map[string]interface{}
	err = json.Unmarshal(resp.Body, &result)
	if err != nil {
		return "", err
	}
	// 从返回值中提取 unique_identifier
	uniqueIdentifier, ok := result["unique_identifier"].(string)
	if !ok {
		return "", fmt.Errorf("未找到 unique_identifier 字段或类型错误")
	}
	return uniqueIdentifier, nil
}

// PluginInstallFromPkg 安装插件（本地包）
func (c *Client) PluginInstallFromPkg(ctx context.Context, pkgPath string) (map[string]interface{}, error) {
	fileData, err := os.ReadFile(pkgPath)
	if err != nil {
		return nil, fmt.Errorf("读取测试插件包失败: %v", err)
	}

	filename := filepath.Base(pkgPath)
	uniqueIdentifier, err := c.PluginUploadPkg(ctx, filename, fileData)
	if err != nil {
		return nil, fmt.Errorf("上传插件包失败: %v", err)
	}
	body := map[string]interface{}{
		"plugin_unique_identifiers": []interface{}{uniqueIdentifier},
	}
	req := &client.Request{
		Method: "POST",
		Path:   "/install/pkg",
		Body:   body,
	}
	var result map[string]interface{}
	err = c.baseClient.DoJSON(ctx, req, &result)
	return result, err
}

// PluginUninstall 卸载插件
func (c *Client) PluginUninstall(ctx context.Context, pluginInstallationID string) (map[string]interface{}, error) {
	body := map[string]interface{}{
		"plugin_installation_id": pluginInstallationID,
	}
	req := &client.Request{
		Method: "POST",
		Path:   "/uninstall",
		Body:   body,
	}
	var result map[string]interface{}
	err := c.baseClient.DoJSON(ctx, req, &result)
	return result, err
}

// PluginGetPermission 获取插件权限
func (c *Client) PluginGetPermission(ctx context.Context) (*PluginPermissionResponse, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/permission/fetch",
	}
	var resp PluginPermissionResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

// PluginChangePermission 修改插件权限
func (c *Client) PluginChangePermission(ctx context.Context, installPermission, debugPermission string) (map[string]interface{}, error) {
	body := map[string]interface{}{
		"install_permission": installPermission,
		"debug_permission":   debugPermission,
	}
	req := &client.Request{
		Method: "POST",
		Path:   "/permission/change",
		Body:   body,
	}
	var result map[string]interface{}
	err := c.baseClient.DoJSON(ctx, req, &result)
	return result, err
}

// PluginGetManifest 获取插件manifest
func (c *Client) PluginGetManifest(ctx context.Context, pluginUniqueIdentifier string) (*PluginManifestResponse, error) {
	query := map[string]string{
		"plugin_unique_identifier": pluginUniqueIdentifier,
	}
	req := &client.Request{
		Method: "GET",
		Path:   "/fetch-manifest",
		Query:  query,
	}
	var resp PluginManifestResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

// PluginGetIcon 获取插件图标
// 返回二进制数据和mimetype
func (c *Client) PluginGetIcon(ctx context.Context, tenantID, filename string) ([]byte, string, error) {
	query := map[string]string{
		"tenant_id": tenantID,
		"filename":  filename,
	}
	req := &client.Request{
		Method: "GET",
		Path:   "/icon",
		Query:  query,
	}
	resp, err := c.baseClient.Do(ctx, req)
	if err != nil {
		return nil, "", err
	}
	contentType := resp.Headers.Get("Content-Type")
	return resp.Body, contentType, nil
}

// PluginListLatestVersions 获取插件最新版本信息
func (c *Client) PluginListLatestVersions(pluginIDs []string) {
	// TODO: 实现获取插件最新版本信息的接口调用
}

// PluginGet 获取插件详情
func (c *Client) PluginGet(pluginID string) {
	// TODO: 实现获取插件详情的接口调用
}

// PluginInstall 安装插件
func (c *Client) PluginInstall(pluginID string, version string, params map[string]interface{}) {
	// TODO: 实现安装插件的接口调用
}

// PluginUpdate 更新插件
func (c *Client) PluginUpdate(pluginID string, version string, params map[string]interface{}) {
	// TODO: 实现更新插件的接口调用
}

// PluginGetParameter 获取插件参数
func (c *Client) PluginGetParameter(pluginID string) {
	// TODO: 实现获取插件参数的接口调用
}

// PluginUpdateParameter 更新插件参数
func (c *Client) PluginUpdateParameter(pluginID string, params map[string]interface{}) {
	// TODO: 实现更新插件参数的接口调用
}

// PluginUpdatePermission 更新插件权限
func (c *Client) PluginUpdatePermission(pluginID string, permission map[string]interface{}) {
	// TODO: 实现更新插件权限的接口调用
}

// PluginDownload 下载插件
func (c *Client) PluginDownload(pluginID string, version string) {
	// TODO: 实现下载插件的接口调用
}

// PluginUpload 上传插件
func (c *Client) PluginUpload(fileName string, fileData []byte) {
	// TODO: 实现上传插件的接口调用
}

// PluginGetSchema 获取插件schema
func (c *Client) PluginGetSchema(pluginID string) {
	// TODO: 实现获取插件schema的接口调用
}

// PluginGetLog 获取插件日志
func (c *Client) PluginGetLog(pluginID string, page int, pageSize int) {
	// TODO: 实现获取插件日志的接口调用
}

// PluginGetDebugLog 获取插件调试日志
func (c *Client) PluginGetDebugLog(pluginID string, page int, pageSize int) {
	// TODO: 实现获取插件调试日志的接口调用
}

// PluginGetDebugStatus 获取插件调试状态
func (c *Client) PluginGetDebugStatus(pluginID string) {
	// TODO: 实现获取插件调试状态的接口调用
}

// PluginStartDebug 启动插件调试
func (c *Client) PluginStartDebug(pluginID string) {
	// TODO: 实现启动插件调试的接口调用
}

// PluginStopDebug 停止插件调试
func (c *Client) PluginStopDebug(pluginID string) {
	// TODO: 实现停止插件调试的接口调用
}

// PluginGetDebugToken 获取插件调试token
func (c *Client) PluginGetDebugToken(pluginID string) {
	// TODO: 实现获取插件调试token的接口调用
}
