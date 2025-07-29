package console

import (
	"context"
	"strconv"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/models"
)

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
