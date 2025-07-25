package console

import (
	"fmt"
	"os"
	"path/filepath"

	"context"
	"encoding/json"
	"strconv"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/models"
)

var pluginURLPrefix string = "/workspaces/current/plugin"

// PluginList 获取插件列表
func (c *Client) GetPluginList(ctx context.Context, page, pageSize int) (*models.PluginListResponse, error) {
	query := map[string]string{
		"page":      strconv.Itoa(page),
		"page_size": strconv.Itoa(pageSize),
	}
	req := &client.Request{
		Method: "GET",
		Path:   "/list",
		Query:  query,
	}
	var resp models.PluginListResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

// PluginUploadPkg 上传插件包（pkg文件）
func (c *Client) UploadPluginPkg(ctx context.Context, filename string, fileData []byte) (string, error) {
	req := &client.Request{
		Path:   pluginURLPrefix + "/upload/pkg",
	}
	resp, err := c.baseClient.UploadFile(ctx, req.Path, "pkg", filename, fileData, nil)
	if err != nil {
		return "", err
	}
	var result struct {
		UniqueIdentifier string `json:"unique_identifier"`
	}
	err = json.Unmarshal(resp.Body, &result)
	if err != nil {
		return "", err
	}
	if result.UniqueIdentifier == "" {
		return "", fmt.Errorf("未找到 unique_identifier 字段或类型错误")
	}
	return result.UniqueIdentifier, nil
}

// PluginInstallFromPkg 安装插件（本地包）
func (c *Client) InstallPluginFromPkg(ctx context.Context, pkgPath string) (*models.PluginInstallResponse, error) {
	fileData, err := os.ReadFile(pkgPath)
	if err != nil {
		return nil, fmt.Errorf("读取测试插件包失败: %v", err)
	}

	filename := filepath.Base(pkgPath)
	uniqueIdentifier, err := c.UploadPluginPkg(ctx, filename, fileData)
	if err != nil {
		return nil, fmt.Errorf("上传插件包失败: %v", err)
	}
	body := map[string]any{
		"plugin_unique_identifiers": []any{uniqueIdentifier},
	}
	req := &client.Request{
		Method: "POST",
		Path:   "/install/pkg",
		Body:   body,
	}
	var result models.PluginInstallResponse
	err = c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// PluginUninstall 卸载插件
func (c *Client) UninstallPlugin(ctx context.Context, pluginInstallationID string) (*any, error) {
	req := &client.Request{
		Method: "POST",
		Path:   "/uninstall",
		Body:   map[string]any{"plugin_installation_id": pluginInstallationID},
	}
	var result any
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// PluginGetPermission 获取插件权限
func (c *Client) GetPluginPermission(ctx context.Context) (*any, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/permission/fetch",
	}
	var result any
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// PluginChangePermission 修改插件权限
func (c *Client) UpdatePluginPermission(ctx context.Context, installPermission, debugPermission string) (*any, error) {
	body := map[string]any{
		"install_permission": installPermission,
		"debug_permission":  debugPermission,
	}
	req := &client.Request{
		Method: "POST",
		Path:   "/permission/change",
		Body:   body,
	}
	var result any
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// PluginGetManifest 获取插件manifest
func (c *Client) GetPluginManifest(ctx context.Context, pluginUniqueIdentifier string) (*any, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/fetch-manifest",
		Query:  map[string]string{"plugin_unique_identifier": pluginUniqueIdentifier},
	}
	var result any
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// PluginGetIcon 获取插件图标
// 返回二进制数据和mimetype
func (c *Client) GetPluginIcon(ctx context.Context, tenantID, filename string) ([]byte, string, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/icon",
		Query:  map[string]string{"tenant_id": tenantID, "filename": filename},
	}
	resp, err := c.baseClient.Do(ctx, req)
	if err != nil {
		return nil, "", err
	}
	contentType := ""
	if resp.Headers != nil {
		contentType = resp.Headers.Get("Content-Type")
	}
	return resp.Body, contentType, nil
}

// PluginListLatestVersions 获取插件最新版本信息
func (c *Client) GetPluginListLatestVersions(ctx context.Context, pluginIDs []string) (*any, error) {
	body := map[string]any{"plugin_ids": pluginIDs}
	req := &client.Request{
		Method: "POST",
		Path:   "/list/latest-versions",
		Body:   body,
	}
	var result any
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// PluginListInstallationsFromIds 批量获取插件安装信息
func (c *Client) GetPluginListInstallationsFromIds(ctx context.Context, pluginIDs []string) (*any, error) {
	body := map[string]any{"plugin_ids": pluginIDs}
	req := &client.Request{
		Method: "POST",
		Path:   "/list/installations/ids",
		Body:   body,
	}
	var result any
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// PluginUploadFromGithub 从 Github 上传插件
func (c *Client) UploadPluginFromGithub(ctx context.Context, repo, version, pkg string) (*any, error) {
	body := map[string]any{
		"repo":    repo,
		"version": version,
		"package": pkg,
	}
	req := &client.Request{
		Method: "POST",
		Path:   "/upload/github",
		Body:   body,
	}
	var result any
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// PluginUploadFromBundle 上传插件 bundle
func (c *Client) UploadPluginFromBundle(ctx context.Context, fileName string, fileData []byte) (*any, error) {
	resp, err := c.baseClient.UploadFile(ctx, "/upload/bundle", "bundle", fileName, fileData, nil)
	if err != nil {
		return nil, err
	}
	var result any
	err = json.Unmarshal(resp.Body, &result)
	return &result, err
}

// PluginInstallFromGithub 从 Github 安装插件
func (c *Client) InstallPluginFromGithub(ctx context.Context, pluginUniqueIdentifier, repo, version, pkg string) (*any, error) {
	body := map[string]any{
		"plugin_unique_identifier": pluginUniqueIdentifier,
		"repo":    repo,
		"version": version,
		"package": pkg,
	}
	req := &client.Request{
		Method: "POST",
		Path:   "/install/github",
		Body:   body,
	}
	var result any
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// PluginInstallFromMarketplace 从市场安装插件
func (c *Client) InstallPluginFromMarketplace(ctx context.Context, pluginUniqueIdentifiers []string) (*any, error) {
	body := map[string]any{"plugin_unique_identifiers": pluginUniqueIdentifiers}
	req := &client.Request{
		Method: "POST",
		Path:   "/install/marketplace",
		Body:   body,
	}
	var result any
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// PluginFetchMarketplacePkg 获取市场插件包信息
func (c *Client) GetMarketplacePkg(ctx context.Context, pluginUniqueIdentifier string) (*any, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/marketplace/pkg",
		Query:  map[string]string{"plugin_unique_identifier": pluginUniqueIdentifier},
	}
	var result any
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// PluginFetchInstallTasks 获取插件安装任务列表
// PluginFetchInstallTasks 获取插件安装任务列表
func (c *Client) GetPluginFetchInstallTasks(ctx context.Context, page, pageSize int) (*any, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/tasks",
		Query:  map[string]string{"page": strconv.Itoa(page), "page_size": strconv.Itoa(pageSize)},
	}
	var resp any
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

// PluginFetchInstallTask 获取单个插件安装任务
func (c *Client) GetPluginFetchInstallTask(ctx context.Context, taskID string) (*any, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/tasks/" + taskID,
	}
	var result any
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// PluginDeleteInstallTask 删除插件安装任务
func (c *Client) DeletePluginInstallTask(ctx context.Context, taskID string) (*any, error) {
	req := &client.Request{
		Method: "POST",
		Path:   "/tasks/" + taskID + "/delete",
	}
	var result any
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// PluginDeleteAllInstallTaskItems 删除所有插件安装任务项
func (c *Client) DeleteAllPluginInstallTaskItems(ctx context.Context) (*any, error) {
	req := &client.Request{
		Method: "POST",
		Path:   "/tasks/delete_all",
	}
	var result any
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// PluginDeleteInstallTaskItem 删除指定插件安装任务项
func (c *Client) DeletePluginInstallTaskItem(ctx context.Context, taskID, identifier string) (*any, error) {
	path := "/tasks/" + taskID + "/delete/" + identifier
	req := &client.Request{
		Method: "POST",
		Path:   path,
	}
	var result any
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// PluginUpgradeFromMarketplace 从市场升级插件
func (c *Client) UpgradePluginFromMarketplace(ctx context.Context, originalPluginUniqueIdentifier, newPluginUniqueIdentifier string) (*any, error) {
	body := map[string]any{
		"original_plugin_unique_identifier": originalPluginUniqueIdentifier,
		"new_plugin_unique_identifier":      newPluginUniqueIdentifier,
	}
	req := &client.Request{
		Method: "POST",
		Path:   "/upgrade/marketplace",
		Body:   body,
	}
	var result any
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// PluginUpgradeFromGithub 从 Github 升级插件
func (c *Client) UpgradePluginFromGithub(ctx context.Context, originalPluginUniqueIdentifier, newPluginUniqueIdentifier, repo, version, pkg string) (*any, error) {
	body := map[string]any{
		"original_plugin_unique_identifier": originalPluginUniqueIdentifier,
		"new_plugin_unique_identifier":      newPluginUniqueIdentifier,
		"repo":    repo,
		"version": version,
		"package": pkg,
	}
	req := &client.Request{
		Method: "POST",
		Path:   "/upgrade/github",
		Body:   body,
	}
	var result any
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// PluginDebuggingKey 获取插件调试 key
func (c *Client) GetPluginDebuggingKey(ctx context.Context) (*any, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/debugging-key",
	}
	var result any
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// PluginFetchDynamicSelectOptions 获取插件动态参数选项
func (c *Client) GetPluginFetchDynamicSelectOptions(ctx context.Context, pluginID, provider, action, parameter, providerType string) (*any, error) {
	query := map[string]string{
		"plugin_id":     pluginID,
		"provider":      provider,
		"action":        action,
		"parameter":     parameter,
		"provider_type": providerType,
	}
	req := &client.Request{
		Method: "GET",
		Path:   "/parameters/dynamic-options",
		Query:  query,
	}
	var result any
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}
