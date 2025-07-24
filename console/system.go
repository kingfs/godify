package console

import (
	"context"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/models"
)

const (
	DifyVersion = "1.6.0"
)

// Ping 检查Dify是否正常运行
func (c *Client) Ping(ctx context.Context) (*models.ResultResponse, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/ping",
	}

	var result models.ResultResponse
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// GetVersion 获取Dify版本信息
func (c *Client) GetVersion(ctx context.Context, currentVersion string) (*models.Version, error) {
	if currentVersion == "" {
		currentVersion = DifyVersion
	}

	req := &client.Request{
		Method: "GET",
		Path:   "/version",
		Query:  map[string]string{"current_version": currentVersion},
	}

	var result models.Version
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// GetSetupInfo 获取安装信息
func (c *Client) GetSetupInfo(ctx context.Context) (*models.StatusResponse, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/setup",
	}

	var result models.StatusResponse
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// GetInitInfo 获取Dify初始化信息
func (c *Client) GetInitInfo(ctx context.Context) (*models.StatusResponse, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/init",
	}

	var result models.StatusResponse
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// Setup 安装Dify
func (c *Client) Setup(ctx context.Context, setup *models.SetupRequest) (*models.ResultResponse, error) {
	req := &client.Request{
		Method: "POST",
		Path:   "/setup",
		Body:   setup,
	}

	var result models.ResultResponse
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// Init 初始化Dify
func (c *Client) Init(ctx context.Context, initPassword string) (*models.ResultResponse, error) {
	req := &client.Request{
		Method: "POST",
		Path:   "/init",
	}

	if initPassword != "" {
		req.Body = map[string]string{"password": initPassword}
	}

	var result models.ResultResponse
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}
