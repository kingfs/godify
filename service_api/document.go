package service_api

import (
	"context"
	"fmt"
	"io"

	"github.com/kingfs/godify/client"
)

// Document represents a document in a dataset.
type Document struct {
	ID                    string                 `json:"id"`
	DocType               string                 `json:"doc_type"`
	DocMetadata           map[string]interface{} `json:"doc_metadata"`
	DatasetID             string                 `json:"dataset_id"`
	Position              int                    `json:"position"`
	Name                  string                 `json:"name"`
	DataSourceType        string                 `json:"data_source_type"`
	DataSourceInfo        map[string]interface{} `json:"data_source_info"`
	DatasetProcessRuleID  string                 `json:"dataset_process_rule_id"`
	DatasetProcessRule    map[string]interface{} `json:"dataset_process_rule"`
	Batch                 string                 `json:"batch"`
	CreatedFrom           string                 `json:"created_from"`
	CreatedBy             string                 `json:"created_by"`
	CreatedAt             int64                  `json:"created_at"`
	Tokens                int                    `json:"tokens"`
	WordCount             int                    `json:"word_count"`
	IndexingStatus        string                 `json:"indexing_status"`
	Error                 string                 `json:"error"`
	Enabled               bool                   `json:"enabled"`
	DisabledAt            int64                  `json:"disabled_at"`
	DisabledBy            string                 `json:"disabled_by"`
	Archived              bool                   `json:"archived"`
	DisplayStatus         string                 `json:"display_status"`
	DocForm               string                 `json:"doc_form"`
	DocLanguage           string                 `json:"doc_language"`
}

// DocumentIndexingStatus represents the indexing status of a document.
type DocumentIndexingStatus struct {
	ID                   string `json:"id"`
	IndexingStatus       string `json:"indexing_status"`
	ProcessingStartedAt  int64  `json:"processing_started_at"`
	ParsingCompletedAt   int64  `json:"parsing_completed_at"`
	CleaningCompletedAt  int64  `json:"cleaning_completed_at"`
	SplittingCompletedAt int64  `json:"splitting_completed_at"`
	CompletedAt          int64  `json:"completed_at"`
	PausedAt             int64  `json:"paused_at"`
	Error                string `json:"error"`
	StoppedAt            int64  `json:"stopped_at"`
	CompletedSegments    int    `json:"completed_segments"`
	TotalSegments        int    `json:"total_segments"`
}

// DocumentResponse is the response for creating or updating a document.
type DocumentResponse struct {
	Document Document `json:"document"`
	Batch    string   `json:"batch"`
}

// DocumentListResponse is the paginated response for listing documents.
type DocumentListResponse struct {
	Data    []Document `json:"data"`
	HasMore bool       `json:"has_more"`
	Limit   int        `json:"limit"`
	Total   int        `json:"total"`
	Page    int        `json:"page"`
}

// IndexingStatusResponse contains a list of document indexing statuses.
type IndexingStatusResponse struct {
	Data []DocumentIndexingStatus `json:"data"`
}


// CreateDocumentRequest is the base request for creating/updating a document.
type CreateDocumentRequest struct {
	Name             string                 `json:"name,omitempty"`
	Text             string                 `json:"text,omitempty"`
	ProcessRule      map[string]interface{} `json:"process_rule,omitempty"`
	DocForm          string                 `json:"doc_form,omitempty"`
	DocLanguage      string                 `json:"doc_language,omitempty"`
	IndexingTechnique string                `json:"indexing_technique,omitempty"`
}


// CreateDocumentByText creates a document in a dataset from raw text.
func (c *client.Client) CreateDocumentByText(ctx context.Context, datasetID string, req *CreateDocumentRequest) (*DocumentResponse, error) {
	var result DocumentResponse
	path := fmt.Sprintf("/v1/datasets/%s/document/create_by_text", datasetID)
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateDocumentByFile creates a document in a dataset from a file upload.
func (c *client.Client) CreateDocumentByFile(ctx context.Context, datasetID string, file io.Reader, filename string, jsonData *CreateDocumentRequest) (*DocumentResponse, error) {
	var result DocumentResponse
	path := fmt.Sprintf("/v1/datasets/%s/document/create_by_file", datasetID)
	err := c.SendMultipartRequestWithJSON(ctx, path, file, filename, jsonData, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateDocumentByText updates a document with raw text.
func (c *client.Client) UpdateDocumentByText(ctx context.Context, datasetID, documentID string, req *CreateDocumentRequest) (*DocumentResponse, error) {
	var result DocumentResponse
	path := fmt.Sprintf("/v1/datasets/%s/documents/%s/update_by_text", datasetID, documentID)
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateDocumentByFile updates a document with a file upload.
func (c *client.Client) UpdateDocumentByFile(ctx context.Context, datasetID, documentID string, file io.Reader, filename string, jsonData *CreateDocumentRequest) (*DocumentResponse, error) {
	var result DocumentResponse
	path := fmt.Sprintf("/v1/datasets/%s/documents/%s/update_by_file", datasetID, documentID)
	err := c.SendMultipartRequestWithJSON(ctx, path, file, filename, jsonData, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDocuments retrieves a list of documents in a dataset.
func (c *client.Client) GetDocuments(ctx context.Context, datasetID string, page, limit int, keyword string) (*DocumentListResponse, error) {
	var result DocumentListResponse
	path := fmt.Sprintf("/v1/datasets/%s/documents?page=%d&limit=%d", datasetID, page, limit)
	if keyword != "" {
		path += fmt.Sprintf("&keyword=%s", keyword)
	}
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteDocument deletes a document from a dataset.
func (c *client.Client) DeleteDocument(ctx context.Context, datasetID, documentID string) error {
	path := fmt.Sprintf("/v1/datasets/%s/documents/%s", datasetID, documentID)
	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}

// GetDocumentIndexingStatus retrieves the indexing status for a batch of documents.
func (c *client.Client) GetDocumentIndexingStatus(ctx context.Context, datasetID, batch string) (*IndexingStatusResponse, error) {
	var result IndexingStatusResponse
	path := fmt.Sprintf("/v1/datasets/%s/documents/%s/indexing-status", datasetID, batch)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDocument retrieves the details of a single document.
func (c *client.Client) GetDocument(ctx context.Context, datasetID, documentID string) (*Document, error) {
	var result Document
	path := fmt.Sprintf("/v1/datasets/%s/documents/%s", datasetID, documentID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
