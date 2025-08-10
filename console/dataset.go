package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// ConsoleDataset represents the detailed information of a dataset in the console.
type ConsoleDataset struct {
	ID                     string                 `json:"id"`
	Name                   string                 `json:"name"`
	Description            string                 `json:"description"`
	Permission             string                 `json:"permission"`
	DataSourceType         string                 `json:"data_source_type"`
	IndexingTechnique      string                 `json:"indexing_technique"`
	AppCount               int                    `json:"app_count"`
	DocumentCount          int                    `json:"document_count"`
	WordCount              int                    `json:"word_count"`
	CreatedBy              string                 `json:"created_by"`
	CreatedAt              int64                  `json:"created_at"`
	EmbeddingModel         string                 `json:"embedding_model"`
	EmbeddingModelProvider string                 `json:"embedding_model_provider"`
	EmbeddingAvailable     bool                   `json:"embedding_available"`
	RetrievalModel         map[string]interface{} `json:"retrieval_model_dict"`
}

// ConsoleDatasetListResponse is the paginated response for listing datasets.
type ConsoleDatasetListResponse struct {
	Data    []ConsoleDataset `json:"data"`
	HasMore bool             `json:"has_more"`
	Limit   int              `json:"limit"`
	Total   int              `json:"total"`
	Page    int              `json:"page"`
}

// CreateConsoleDatasetRequest is the request to create a dataset from the console.
type CreateConsoleDatasetRequest struct {
	Name              string `json:"name"`
	Description       string `json:"description,omitempty"`
	IndexingTechnique string `json:"indexing_technique,omitempty"`
}

// UpdateConsoleDatasetRequest is the request to update a dataset from the console.
type UpdateConsoleDatasetRequest struct {
	Name              string                 `json:"name,omitempty"`
	Description       string                 `json:"description,omitempty"`
	IndexingTechnique string                 `json:"indexing_technique,omitempty"`
	Permission        string                 `json:"permission,omitempty"`
	RetrievalModel    map[string]interface{} `json:"retrieval_model,omitempty"`
}

// DatasetUseCheckResponse is the response for checking if a dataset is in use.
type DatasetUseCheckResponse struct {
	IsUsing bool `json:"is_using"`
}

// ... other structs for queries, error docs, related apps, etc. will be defined as needed.

// GetDatasets retrieves a list of datasets.
func (c *client.Client) ConsoleGetDatasets(ctx context.Context, page, limit int) (*ConsoleDatasetListResponse, error) {
	var result ConsoleDatasetListResponse
	path := fmt.Sprintf("/console/api/datasets?page=%d&limit=%d", page, limit)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateDataset creates a new dataset.
func (c *client.Client) ConsoleCreateDataset(ctx context.Context, req *CreateConsoleDatasetRequest) (*ConsoleDataset, error) {
	var result ConsoleDataset
	err := c.sendRequest(ctx, "POST", "/console/api/datasets", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDataset retrieves the details of a specific dataset.
func (c *client.Client) ConsoleGetDataset(ctx context.Context, datasetID string) (*ConsoleDataset, error) {
	var result ConsoleDataset
	path := fmt.Sprintf("/console/api/datasets/%s", datasetID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateDataset updates a dataset.
func (c *client.Client) ConsoleUpdateDataset(ctx context.Context, datasetID string, req *UpdateConsoleDatasetRequest) (*ConsoleDataset, error) {
	var result ConsoleDataset
	path := fmt.Sprintf("/console/api/datasets/%s", datasetID)
	err := c.sendRequest(ctx, "PATCH", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteDataset deletes a dataset.
func (c *client.Client) ConsoleDeleteDataset(ctx context.Context, datasetID string) error {
	path := fmt.Sprintf("/console/api/datasets/%s", datasetID)
	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}

// CheckDatasetInUse checks if a dataset is currently in use by any apps.
func (c *client.Client) CheckDatasetInUse(ctx context.Context, datasetID string) (*DatasetUseCheckResponse, error) {
	var result DatasetUseCheckResponse
	path := fmt.Sprintf("/console/api/datasets/%s/use-check", datasetID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDatasetAPIKeys retrieves the API keys for a dataset.
func (c *client.Client) GetDatasetAPIKeys(ctx context.Context) (*APIKeyListResponse, error) {
	var result APIKeyListResponse
	err := c.sendRequest(ctx, "GET", "/console/api/datasets/api-keys", nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateDatasetAPIKey creates an API key for all datasets.
func (c *client.Client) CreateDatasetAPIKey(ctx context.Context) (*APIKey, error) {
	var result APIKey
	err := c.sendRequest(ctx, "POST", "/console/api/datasets/api-keys", nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteDatasetAPIKey deletes an API key for all datasets.
func (c *client.Client) DeleteDatasetAPIKey(ctx context.Context, apiKeyID string) error {
	path := fmt.Sprintf("/console/api/datasets/api-keys/%s", apiKeyID)
	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}

// GetDatasetAPIBaseURL retrieves the base URL for the dataset API.
func (c *client.Client) GetDatasetAPIBaseURL(ctx context.Context) (string, error) {
	var result map[string]string
	err := c.sendRequest(ctx, "GET", "/console/api/datasets/api-base-info", nil, &result, nil)
	if err != nil {
		return "", err
	}
	return result["api_base_url"], nil
}

// GetDatasetRetrievalSetting retrieves the retrieval setting for the workspace.
func (c *client.Client) GetDatasetRetrievalSetting(ctx context.Context) (map[string][]string, error) {
	var result map[string][]string
	err := c.sendRequest(ctx, "GET", "/console/api/datasets/retrieval-setting", nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}
