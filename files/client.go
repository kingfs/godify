package files

import (
	"context"
	"time"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/models"
)

// Client Files API客户端 (独立文件服务)
type Client struct {
	baseClient *client.BaseClient
}

// NewClient 创建Files API客户端
func NewClient(baseURL string) *Client {
	config := &client.ClientConfig{
		BaseURL:    baseURL + "/files",
		Timeout:    30 * time.Second,
		MaxRetries: 3,
	}

	return &Client{
		baseClient: client.NewBaseClient(config),
	}
}

// PluginUploadRequest 插件文件上传请求参数
type PluginUploadRequest struct {
	TenantID  string
	UserID    string
	Timestamp string
	Nonce     string
	Sign      string
}

// UploadForPlugin 为插件上传文件
func (c *Client) UploadForPlugin(ctx context.Context, filename string, fileData []byte, mimetype string, req *PluginUploadRequest) (*models.FileUpload, error) {
	query := map[string]string{
		"tenant_id": req.TenantID,
		"user_id":   req.UserID,
		"timestamp": req.Timestamp,
		"nonce":     req.Nonce,
		"sign":      req.Sign,
	}

	// 设置查询参数
	httpReq := &client.Request{
		Method: "POST",
		Path:   "/upload/for-plugin",
		Query:  query,
	}

	// 上传文件
	_, err := c.baseClient.UploadFile(ctx, httpReq.Path, "file", filename, fileData, nil)
	if err != nil {
		return nil, err
	}

	// 解析响应
	var result models.FileUpload
	if err := c.baseClient.DoJSON(ctx, &client.Request{}, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
