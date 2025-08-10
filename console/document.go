package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// ConsoleDocument represents a document in a dataset.
type ConsoleDocument struct {
	ID                   string                 `json:"id"`
	Position             int                    `json:"position"`
	Name                 string                 `json:"name"`
	DataSourceType       string                 `json:"data_source_type"`
	DataSourceInfo       map[string]interface{} `json:"data_source_info"`
	DatasetProcessRuleID string                 `json:"dataset_process_rule_id"`
	Batch                string                 `json:"batch"`
	IndexingStatus       string                 `json:"indexing_status"`
	// ... other fields
}

// ConsoleDocumentListResponse is the paginated response for listing documents.
type ConsoleDocumentListResponse struct {
	Data    []ConsoleDocument `json:"data"`
	HasMore bool              `json:"has_more"`
	Limit   int               `json:"limit"`
	Total   int               `json:"total"`
	Page    int               `json:"page"`
}

// CreateDocumentsRequest is the request to create documents in a dataset.
type CreateDocumentsRequest struct {
	IndexingTechnique string                 `json:"indexing_technique"`
	DataSource        map[string]interface{} `json:"data_source,omitempty"`
	ProcessRule       map[string]interface{} `json:"process_rule,omitempty"`
}

// CreateDocumentsResponse is the response after creating documents.
type CreateDocumentsResponse struct {
	Dataset   ConsoleDataset    `json:"dataset"`
	Documents []ConsoleDocument `json:"documents"`
	Batch     string            `json:"batch"`
}

// GetProcessRuleResponse is the response for getting process rules.
type GetProcessRuleResponse struct {
	Mode  string                 `json:"mode"`
	Rules map[string]interface{} `json:"rules"`
}

// GetDocuments retrieves a list of documents in a dataset.
func (c *client.Client) ConsoleGetDocuments(ctx context.Context, datasetID, keyword, sort string, page, limit int, fetch bool) (*ConsoleDocumentListResponse, error) {
	var result ConsoleDocumentListResponse
	path := fmt.Sprintf("/console/api/datasets/%s/documents?page=%d&limit=%d&sort=%s&fetch=%t", datasetID, page, limit, sort, fetch)
	if keyword != "" {
		path += fmt.Sprintf("&keyword=%s", keyword)
	}
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateDocuments creates one or more documents in a dataset.
func (c *client.Client) ConsoleCreateDocuments(ctx context.Context, datasetID string, req *CreateDocumentsRequest) (*CreateDocumentsResponse, error) {
	var result CreateDocumentsResponse
	path := fmt.Sprintf("/console/api/datasets/%s/documents", datasetID)
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteDocuments deletes one or more documents from a dataset.
func (c *client.Client) ConsoleDeleteDocuments(ctx context.Context, datasetID string, documentIDs []string) error {
	path := fmt.Sprintf("/console/api/datasets/%s/documents", datasetID)
	for _, id := range documentIDs {
		path += fmt.Sprintf("&document_id=%s", id)
	}
	path = "/console/api/datasets/" + datasetID + "/documents?" + path[1:]
	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}

// GetProcessRule retrieves the processing rule for a document or dataset.
func (c *client.Client) GetProcessRule(ctx context.Context, documentID string) (*GetProcessRuleResponse, error) {
	var result GetProcessRuleResponse
	path := "/console/api/datasets/process-rule"
	if documentID != "" {
		path += fmt.Sprintf("?document_id=%s", documentID)
	}
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// InitDatasetWithDocument creates a new dataset and document in one call.
func (c *client.Client) InitDatasetWithDocument(ctx context.Context, req *CreateDocumentsRequest) (*CreateDocumentsResponse, error) {
	var result CreateDocumentsResponse
	err := c.sendRequest(ctx, "POST", "/console/api/datasets/init", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDocumentDetail retrieves the details of a single document.
func (c *client.Client) GetDocumentDetail(ctx context.Context, datasetID, documentID string) (*ConsoleDocument, error) {
	var result ConsoleDocument
	path := fmt.Sprintf("/console/api/datasets/%s/documents/%s", datasetID, documentID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ConsoleDeleteDocument deletes a single document.
func (c *client.Client) ConsoleDeleteDocument(ctx context.Context, datasetID, documentID string) error {
	path := fmt.Sprintf("/console/api/datasets/%s/documents/%s", datasetID, documentID)
	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}

// UpdateDocumentMetadata updates the metadata for a document.
func (c *client.Client) UpdateDocumentMetadata(ctx context.Context, datasetID, documentID string, docType string, metadata map[string]interface{}) error {
	path := fmt.Sprintf("/console/api/datasets/%s/documents/%s/metadata", datasetID, documentID)
	req := map[string]interface{}{
		"doc_type":     docType,
		"doc_metadata": metadata,
	}
	return c.sendRequest(ctx, "PUT", path, req, nil, nil)
}

// RenameDocument renames a document.
func (c *client.Client) RenameDocument(ctx context.Context, datasetID, documentID, name string) (*ConsoleDocument, error) {
	var result ConsoleDocument
	path := fmt.Sprintf("/console/api/datasets/%s/documents/%s/rename", datasetID, documentID)
	req := map[string]string{"name": name}
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// PauseDocumentIndexing pauses the indexing of a document.
func (c *client.Client) PauseDocumentIndexing(ctx context.Context, datasetID, documentID string) error {
	path := fmt.Sprintf("/console/api/datasets/%s/documents/%s/processing/pause", datasetID, documentID)
	return c.sendRequest(ctx, "PATCH", path, nil, nil, nil)
}

// ResumeDocumentIndexing resumes the indexing of a document.
func (c *client.Client) ResumeDocumentIndexing(ctx context.Context, datasetID, documentID string) error {
	path := fmt.Sprintf("/console/api/datasets/%s/documents/%s/processing/resume", datasetID, documentID)
	return c.sendRequest(ctx, "PATCH", path, nil, nil, nil)
}
