package console

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
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
		Path:   pluginURLPrefix + "/list",
		Query:  query,
	}
	var resp models.PluginListResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

// PluginUploadPkg 上传插件包（pkg文件）
func (c *Client) UploadPluginPkg(ctx context.Context, filename string, fileData []byte) (string, error) {
	req := &client.Request{
		Path: pluginURLPrefix + "/upload/pkg",
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
		Path:   pluginURLPrefix + "/install/pkg",
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
		Path:   pluginURLPrefix + "/uninstall",
		Body:   map[string]any{"plugin_installation_id": pluginInstallationID},
	}
	var result any
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}
