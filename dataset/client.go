package dataset

import (
	"context"
	"strconv"
	"time"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/models"
)

// Client Dataset API客户端
type Client struct {
	baseClient *client.BaseClient
}

// NewClient 创建Dataset API客户端
func NewClient(datasetToken, baseURL string) *Client {
	config := &client.ClientConfig{
		BaseURL:    baseURL + "/v1",
		AuthType:   client.AuthTypeBearer,
		Token:      datasetToken,
		Timeout:    30 * time.Second,
		MaxRetries: 3,
	}

	return &Client{
		baseClient: client.NewBaseClient(config),
	}
}

// ============ 数据集管理 ============

// GetDatasets 获取数据集列表
func (c *Client) GetDatasets(ctx context.Context, page, limit int, keyword string, tagIDs []string, includeAll bool) (*models.DatasetListForAPIResponse, error) {
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
	for _, tagID := range tagIDs {
		query["tag_ids"] = tagID
	}
	if includeAll {
		query["include_all"] = "true"
	}

	req := &client.Request{
		Method: "GET",
		Path:   "/datasets",
		Query:  query,
	}

	var result models.DatasetListForAPIResponse
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// CreateDataset 创建数据集
func (c *Client) CreateDataset(ctx context.Context, req *models.CreateDatasetForAPIRequest) (*models.DatasetForAPI, error) {
	httpReq := &client.Request{
		Method: "POST",
		Path:   "/datasets",
		Body:   req,
	}

	var result models.DatasetForAPI
	err := c.baseClient.DoJSON(ctx, httpReq, &result)
	return &result, err
}

// GetDataset 获取数据集详情
func (c *Client) GetDataset(ctx context.Context, datasetID string) (*models.DatasetForAPI, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/datasets/" + datasetID,
	}

	var result models.DatasetForAPI
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// UpdateDataset 更新数据集
func (c *Client) UpdateDataset(ctx context.Context, datasetID string, req *models.CreateDatasetForAPIRequest) (*models.DatasetForAPI, error) {
	httpReq := &client.Request{
		Method: "PATCH",
		Path:   "/datasets/" + datasetID,
		Body:   req,
	}

	var result models.DatasetForAPI
	err := c.baseClient.DoJSON(ctx, httpReq, &result)
	return &result, err
}

// DeleteDataset 删除数据集
func (c *Client) DeleteDataset(ctx context.Context, datasetID string) error {
	req := &client.Request{
		Method: "DELETE",
		Path:   "/datasets/" + datasetID,
	}

	var result map[string]string
	return c.baseClient.DoJSON(ctx, req, &result)
}

// ============ 文档管理 ============

// GetDatasetDocuments 获取数据集文档列表
func (c *Client) GetDatasetDocuments(ctx context.Context, datasetID string, page, limit int, keyword string) (*models.DocumentListResponse, error) {
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

	req := &client.Request{
		Method: "GET",
		Path:   "/datasets/" + datasetID + "/documents",
		Query:  query,
	}

	var result models.DocumentListResponse
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// CreateDocumentByText 通过文本创建文档
func (c *Client) CreateDocumentByText(ctx context.Context, datasetID string, req *models.CreateDocumentByTextRequest) (*models.DocumentForAPI, error) {
	httpReq := &client.Request{
		Method: "POST",
		Path:   "/datasets/" + datasetID + "/document/create-by-text",
		Body:   req,
	}

	var result models.DocumentForAPI
	err := c.baseClient.DoJSON(ctx, httpReq, &result)
	return &result, err
}

// CreateDocumentByFile 通过文件创建文档
func (c *Client) CreateDocumentByFile(ctx context.Context, datasetID string, filename string, fileData []byte, req *models.CreateDocumentByFileRequest) (*models.DocumentForAPI, error) {
	extraFields := make(map[string]string)
	if req != nil {
		// 将请求参数添加到表单字段中
		if req.DocForm != "" {
			extraFields["doc_form"] = req.DocForm
		}
		if req.DocLanguage != "" {
			extraFields["doc_language"] = req.DocLanguage
		}
		if req.IndexingTechnique != "" {
			extraFields["indexing_technique"] = req.IndexingTechnique
		}
	}

	_, err := c.baseClient.UploadFile(ctx, "/datasets/"+datasetID+"/document/create-by-file", "file", filename, fileData, extraFields)
	if err != nil {
		return nil, err
	}

	var result models.DocumentForAPI
	if err := c.baseClient.DoJSON(ctx, &client.Request{}, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateDocumentByText 通过文本更新文档
func (c *Client) UpdateDocumentByText(ctx context.Context, datasetID, documentID string, req *models.UpdateDocumentByTextRequest) (*models.DocumentForAPI, error) {
	httpReq := &client.Request{
		Method: "PUT",
		Path:   "/datasets/" + datasetID + "/documents/" + documentID + "/update-by-text",
		Body:   req,
	}

	var result models.DocumentForAPI
	err := c.baseClient.DoJSON(ctx, httpReq, &result)
	return &result, err
}

// UpdateDocumentByFile 通过文件更新文档
func (c *Client) UpdateDocumentByFile(ctx context.Context, datasetID, documentID string, filename string, fileData []byte, req *models.UpdateDocumentByFileRequest) (*models.DocumentForAPI, error) {
	extraFields := make(map[string]string)
	if req != nil {
		if req.DocForm != "" {
			extraFields["doc_form"] = req.DocForm
		}
		if req.DocLanguage != "" {
			extraFields["doc_language"] = req.DocLanguage
		}
		if req.IndexingTechnique != "" {
			extraFields["indexing_technique"] = req.IndexingTechnique
		}
	}

	_, err := c.baseClient.UploadFile(ctx, "/datasets/"+datasetID+"/documents/"+documentID+"/update-by-file", "file", filename, fileData, extraFields)
	if err != nil {
		return nil, err
	}

	var result models.DocumentForAPI
	if err := c.baseClient.DoJSON(ctx, &client.Request{}, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteDocument 删除文档
func (c *Client) DeleteDocument(ctx context.Context, datasetID, documentID string) error {
	req := &client.Request{
		Method: "DELETE",
		Path:   "/datasets/" + datasetID + "/documents/" + documentID,
	}

	var result map[string]string
	return c.baseClient.DoJSON(ctx, req, &result)
}

// GetDocument 获取文档详情
func (c *Client) GetDocument(ctx context.Context, datasetID, documentID string) (*models.DocumentForAPI, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/datasets/" + datasetID + "/documents/" + documentID,
	}

	var result models.DocumentForAPI
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// ============ 分段管理 ============

// GetSegments 获取文档分段列表
func (c *Client) GetSegments(ctx context.Context, datasetID, documentID string, page, limit int, status []string, keyword string) (*models.SegmentListResponse, error) {
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
	for _, s := range status {
		query["status"] = s
	}

	req := &client.Request{
		Method: "GET",
		Path:   "/datasets/" + datasetID + "/documents/" + documentID + "/segments",
		Query:  query,
	}

	var result models.SegmentListResponse
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// CreateSegments 创建分段
func (c *Client) CreateSegments(ctx context.Context, datasetID, documentID string, req *models.CreateSegmentsRequest) (*models.SegmentListResponse, error) {
	httpReq := &client.Request{
		Method: "POST",
		Path:   "/datasets/" + datasetID + "/documents/" + documentID + "/segments",
		Body:   req,
	}

	var result models.SegmentListResponse
	err := c.baseClient.DoJSON(ctx, httpReq, &result)
	return &result, err
}

// UpdateSegment 更新分段
func (c *Client) UpdateSegment(ctx context.Context, datasetID, documentID, segmentID string, req *models.UpdateSegmentRequest) (*models.SegmentForAPI, error) {
	httpReq := &client.Request{
		Method: "PUT",
		Path:   "/datasets/" + datasetID + "/documents/" + documentID + "/segments/" + segmentID,
		Body:   req,
	}

	var result models.SegmentForAPI
	err := c.baseClient.DoJSON(ctx, httpReq, &result)
	return &result, err
}

// DeleteSegment 删除分段
func (c *Client) DeleteSegment(ctx context.Context, datasetID, documentID, segmentID string) error {
	req := &client.Request{
		Method: "DELETE",
		Path:   "/datasets/" + datasetID + "/documents/" + documentID + "/segments/" + segmentID,
	}

	var result map[string]string
	return c.baseClient.DoJSON(ctx, req, &result)
}

// ============ 命中测试 ============

// HitTestDataset 数据集命中测试
func (c *Client) HitTestDataset(ctx context.Context, datasetID string, req *models.HitTestingRequest) (*models.HitTestingResponse, error) {
	httpReq := &client.Request{
		Method: "POST",
		Path:   "/datasets/" + datasetID + "/hit-testing",
		Body:   req,
	}

	var result models.HitTestingResponse
	err := c.baseClient.DoJSON(ctx, httpReq, &result)
	return &result, err
}

// ============ 元数据管理 ============

// GetMetadata 获取元数据列表
func (c *Client) GetMetadata(ctx context.Context, datasetID string) (*models.MetadataListResponse, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/datasets/" + datasetID + "/metadata",
	}

	var result models.MetadataListResponse
	err := c.baseClient.DoJSON(ctx, req, &result)
	return &result, err
}

// CreateMetadata 创建元数据
func (c *Client) CreateMetadata(ctx context.Context, datasetID string, req *models.CreateMetadataRequest) (*models.MetadataForAPI, error) {
	httpReq := &client.Request{
		Method: "POST",
		Path:   "/datasets/" + datasetID + "/metadata",
		Body:   req,
	}

	var result models.MetadataForAPI
	err := c.baseClient.DoJSON(ctx, httpReq, &result)
	return &result, err
}

// UpdateMetadata 更新元数据
func (c *Client) UpdateMetadata(ctx context.Context, datasetID, metadataID string, req *models.UpdateMetadataRequest) (*models.MetadataForAPI, error) {
	httpReq := &client.Request{
		Method: "PUT",
		Path:   "/datasets/" + datasetID + "/metadata/" + metadataID,
		Body:   req,
	}

	var result models.MetadataForAPI
	err := c.baseClient.DoJSON(ctx, httpReq, &result)
	return &result, err
}

// DeleteMetadata 删除元数据
func (c *Client) DeleteMetadata(ctx context.Context, datasetID, metadataID string) error {
	req := &client.Request{
		Method: "DELETE",
		Path:   "/datasets/" + datasetID + "/metadata/" + metadataID,
	}

	var result map[string]string
	return c.baseClient.DoJSON(ctx, req, &result)
}
