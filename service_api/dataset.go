package service_api

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// RetrievalModel represents the retrieval model configuration for a dataset.
type RetrievalModel struct {
	SearchMethod         string      `json:"search_method"`
	RerankingEnable      bool        `json:"reranking_enable"`
	RerankingMode        string      `json:"reranking_mode"`
	RerankingModel       interface{} `json:"reranking_model"` // Can be complex
	Weights              interface{} `json:"weights"`         // Can be complex
	TopK                 int         `json:"top_k"`
	ScoreThresholdEnabled bool       `json:"score_threshold_enabled"`
	ScoreThreshold       float64     `json:"score_threshold"`
}

// Dataset represents the detailed information of a dataset.
type Dataset struct {
	ID                     string          `json:"id"`
	Name                   string          `json:"name"`
	Description            string          `json:"description"`
	Permission             string          `json:"permission"`
	DataSourceType         string          `json:"data_source_type"`
	IndexingTechnique      string          `json:"indexing_technique"`
	AppCount               int             `json:"app_count"`
	DocumentCount          int             `json:"document_count"`
	WordCount              int             `json:"word_count"`
	CreatedBy              string          `json:"created_by"`
	CreatedAt              int64           `json:"created_at"`
	EmbeddingModel         string          `json:"embedding_model"`
	EmbeddingModelProvider string          `json:"embedding_model_provider"`
	EmbeddingAvailable     bool            `json:"embedding_available"`
	RetrievalModel         *RetrievalModel `json:"retrieval_model_dict"`
}

// DatasetListResponse is the response for listing datasets.
type DatasetListResponse struct {
	Data    []Dataset `json:"data"`
	HasMore bool      `json:"has_more"`
	Limit   int       `json:"limit"`
	Total   int       `json:"total"`
	Page    int       `json:"page"`
}

// CreateDatasetRequest is the request body for creating a dataset.
type CreateDatasetRequest struct {
	Name              string          `json:"name"`
	Description       string          `json:"description,omitempty"`
	IndexingTechnique string          `json:"indexing_technique,omitempty"`
	RetrievalModel    *RetrievalModel `json:"retrieval_model,omitempty"`
}

// UpdateDatasetRequest is the request body for updating a dataset.
type UpdateDatasetRequest struct {
	Name           string          `json:"name,omitempty"`
	Description    string          `json:"description,omitempty"`
	Permission     string          `json:"permission,omitempty"`
	RetrievalModel *RetrievalModel `json:"retrieval_model,omitempty"`
}

// UpdateDocumentStatusRequest is the request body for updating document status.
type UpdateDocumentStatusRequest struct {
	DocumentIDs []string `json:"document_ids"`
}

// GetDatasets retrieves a list of datasets.
func (c *client.Client) GetDatasets(ctx context.Context, page, limit int) (*DatasetListResponse, error) {
	var result DatasetListResponse
	path := fmt.Sprintf("/v1/datasets?page=%d&limit=%d", page, limit)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateDataset creates a new dataset.
func (c *client.Client) CreateDataset(ctx context.Context, req *CreateDatasetRequest) (*Dataset, error) {
	var result Dataset
	err := c.sendRequest(ctx, "POST", "/v1/datasets", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDataset retrieves the details of a specific dataset.
func (c *client.Client) GetDataset(ctx context.Context, datasetID string) (*Dataset, error) {
	var result Dataset
	path := fmt.Sprintf("/v1/datasets/%s", datasetID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateDataset updates a dataset.
func (c *client.Client) UpdateDataset(ctx context.Context, datasetID string, req *UpdateDatasetRequest) (*Dataset, error) {
	var result Dataset
	path := fmt.Sprintf("/v1/datasets/%s", datasetID)
	err := c.sendRequest(ctx, "PATCH", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteDataset deletes a dataset.
func (c *client.Client) DeleteDataset(ctx context.Context, datasetID string) error {
	path := fmt.Sprintf("/v1/datasets/%s", datasetID)
	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}

// UpdateDocumentStatus updates the status of documents in a dataset.
func (c *client.Client) UpdateDocumentStatus(ctx context.Context, datasetID, action string, req *UpdateDocumentStatusRequest) (*StopResponse, error) {
	var result StopResponse
	path := fmt.Sprintf("/v1/datasets/%s/documents/status/%s", datasetID, action)
	err := c.sendRequest(ctx, "PATCH", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
