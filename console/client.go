package console

import (
	"context"
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
		BaseURL:  baseURL + "/console/api",
		AuthType: client.AuthTypeBearer,
		// Token:      accessToken,
		Timeout:    30 * time.Second,
		SkipTLS:    true,
		MaxRetries: 3,
		Cookies: map[string]string{
			"access_token": accessToken,
		},
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

func (c *Client) WithCookies(cookies map[string]string) *Client {
	c.baseClient.WithCookies(cookies)
	return c
}

func (c *Client) WithToken(token string) *Client {
	c.baseClient.WithToken(token)
	return c
}

// ============ 认证相关 ============

// Login 用户登录
func (c *Client) Login(ctx context.Context, req *models.LoginRequest) (*models.LoginResponse, error) {
	httpReq := &client.Request{
		Method: "POST",
		Path:   "/login",
		Body:   req,
	}

	var result models.LoginResponse
	resp, err := c.baseClient.Do(ctx, httpReq)
	if err != nil {
		return nil, err
	}
	// 将resp.Cookies中的csrf_token、access_token、refresh_token赋值给result.Data.CSRFToken、result.Data.AccessToken、result.Data.RefreshToken
	if resp.Cookies != nil {
		result.Data.CSRFToken = resp.Cookies["csrf_token"]
		result.Data.AccessToken = resp.Cookies["access_token"]
		result.Data.RefreshToken = resp.Cookies["refresh_token"]
	}
	return &result, err
}

// refresh token
func (c *Client) RefreshToken(ctx context.Context) (*models.LoginResponse, error) {
	httpReq := &client.Request{
		Method: "GET",
		Path:   "/refresh-token",
	}

	var result models.LoginResponse
	resp, err := c.baseClient.Do(ctx, httpReq)
	if err != nil {
		return nil, err
	}
	if resp.Cookies != nil {
		result.Data.CSRFToken = resp.Cookies["csrf_token"]
		result.Data.AccessToken = resp.Cookies["access_token"]
		result.Data.RefreshToken = resp.Cookies["refresh_token"]
	}
	return &result, err
}
