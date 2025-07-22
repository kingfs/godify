// Package dify provides a Go client library for Dify AI platform APIs
package dify

import (
	"github.com/kingfs/godify/console"
	"github.com/kingfs/godify/dataset"
	"github.com/kingfs/godify/files"
	"github.com/kingfs/godify/mcp"
	"github.com/kingfs/godify/service"
	"github.com/kingfs/godify/web"
)

// NewServiceClient 创建 Service API 客户端
// appToken: 应用 API Token
// baseURL: Dify 服务器地址，例如 "https://api.dify.ai"
func NewServiceClient(appToken, baseURL string) *service.Client {
	return service.NewClient(appToken, baseURL)
}

// NewWebClient 创建 Web API 客户端
// appCode: 应用代码，可以从Dify控制台获取
// baseURL: Dify 服务器地址，例如 "https://api.dify.ai"
func NewWebClient(appCode, baseURL string) *web.Client {
	return web.NewClient(appCode, baseURL)
}

// NewConsoleClient 创建 Console API 客户端 (管理员API)
// accessToken: 访问令牌或会话令牌
// baseURL: Dify 服务器地址，例如 "https://api.dify.ai"
func NewConsoleClient(accessToken, baseURL string) *console.Client {
	return console.NewClient(accessToken, baseURL)
}

// NewConsoleClientWithSession 使用Session Cookie创建 Console API 客户端
// sessionCookie: 会话Cookie
// baseURL: Dify 服务器地址，例如 "https://api.dify.ai"
func NewConsoleClientWithSession(sessionCookie, baseURL string) *console.Client {
	return console.NewClientWithSession(sessionCookie, baseURL)
}

// NewFilesClient 创建 Files API 客户端
// baseURL: Dify 服务器地址，例如 "https://api.dify.ai"
func NewFilesClient(baseURL string) *files.Client {
	return files.NewClient(baseURL)
}

// NewMCPClient 创建 MCP API 客户端 (Model Context Protocol)
// baseURL: Dify 服务器地址，例如 "https://api.dify.ai"
func NewMCPClient(baseURL string) *mcp.Client {
	return mcp.NewClient(baseURL)
}

// NewDatasetClient 创建Dataset API客户端 (面向数据集管理)
// datasetToken: 数据集API Token，可以从Dify控制台获取
func NewDatasetClient(datasetToken, baseURL string) *dataset.Client {
	return dataset.NewClient(datasetToken, baseURL)
}
