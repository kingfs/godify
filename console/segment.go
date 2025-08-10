package console

import (
	"context"
	"fmt"
	"io"

	"github.com/kingfs/godify/client"
)

// ConsoleSegment represents a document segment in the console.
type ConsoleSegment struct {
	ID            string                 `json:"id"`
	Position      int                    `json:"position"`
	DocumentID    string                 `json:"document_id"`
	Content       string                 `json:"content"`
	Answer        string                 `json:"answer"`
	WordCount     int                    `json:"word_count"`
	Tokens        int                    `json:"tokens"`
	Keywords      []string               `json:"keywords"`
	IndexNodeID   string                 `json:"index_node_id"`
	IndexNodeHash string                 `json:"index_node_hash"`
	HitCount      int                    `json:"hit_count"`
	Enabled       bool                   `json:"enabled"`
	Status        string                 `json:"status"`
	CreatedBy     string                 `json:"created_by"`
	CreatedAt     int64                  `json:"created_at"`
	Document      *ConsoleDocument       `json:"document"`
}

// ConsoleSegmentListResponse is the paginated response for listing segments.
type ConsoleSegmentListResponse struct {
	Data       []ConsoleSegment `json:"data"`
	Limit      int              `json:"limit"`
	Total      int              `json:"total"`
	TotalPages int              `json:"total_pages"`
	Page       int              `json:"page"`
}

// ConsoleChildChunk represents a child chunk of a segment.
type ConsoleChildChunk struct {
	ID       string  `json:"id"`
	Content  string  `json:"content"`
	Position int     `json:"position"`
	Score    float64 `json:"score"`
}

// ChildChunkListResponse is the paginated response for listing child chunks.
type ChildChunkListResponse struct {
	Data       []ConsoleChildChunk `json:"data"`
	Total      int                 `json:"total"`
	TotalPages int                 `json:"total_pages"`
	Page       int                 `json:"page"`
	Limit      int                 `json:"limit"`
}

// AddSegmentRequest is the request to add a new segment.
type AddSegmentRequest struct {
	Content  string   `json:"content"`
	Answer   string   `json:"answer,omitempty"`
	Keywords []string `json:"keywords,omitempty"`
}

// UpdateSegmentRequest is the request to update a segment.
type UpdateSegmentRequest struct {
	Content  string   `json:"content"`
	Answer   string   `json:"answer,omitempty"`
	Keywords []string `json:"keywords,omitempty"`
}

// BatchImportSegmentsResponse is the response from a batch import request.
type BatchImportSegmentsResponse struct {
	JobID     string `json:"job_id"`
	JobStatus string `json:"job_status"`
}


// GetDocumentSegments retrieves a list of segments for a document.
func (c *client.Client) GetDocumentSegments(ctx context.Context, datasetID, documentID string, page, limit int) (*ConsoleSegmentListResponse, error) {
	var result ConsoleSegmentListResponse
	path := fmt.Sprintf("/console/api/datasets/%s/documents/%s/segments?page=%d&limit=%d", datasetID, documentID, page, limit)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteDocumentSegments deletes multiple segments from a document.
func (c *client.Client) DeleteDocumentSegments(ctx context.Context, datasetID, documentID string, segmentIDs []string) error {
	path := fmt.Sprintf("/console/api/datasets/%s/documents/%s/segments", datasetID, documentID)
	for i, id := range segmentIDs {
		if i == 0 {
			path += fmt.Sprintf("?segment_id=%s", id)
		} else {
			path += fmt.Sprintf("&segment_id=%s", id)
		}
	}
	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}

// AddDocumentSegment adds a new segment to a document.
func (c *client.Client) AddDocumentSegment(ctx context.Context, datasetID, documentID string, req *AddSegmentRequest) (*ConsoleSegment, error) {
	var response struct {
		Data ConsoleSegment `json:"data"`
	}
	path := fmt.Sprintf("/console/api/datasets/%s/documents/%s/segment", datasetID, documentID)
	err := c.sendRequest(ctx, "POST", path, req, &response, nil)
	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

// UpdateDocumentSegment updates a segment in a document.
func (c *client.Client) UpdateDocumentSegment(ctx context.Context, datasetID, documentID, segmentID string, req *UpdateSegmentRequest) (*ConsoleSegment, error) {
	var response struct {
		Data ConsoleSegment `json:"data"`
	}
	path := fmt.Sprintf("/console/api/datasets/%s/documents/%s/segments/%s", datasetID, documentID, segmentID)
	err := c.sendRequest(ctx, "PATCH", path, req, &response, nil)
	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

// DeleteDocumentSegment deletes a segment from a document.
func (c *client.Client) DeleteDocumentSegment(ctx context.Context, datasetID, documentID, segmentID string) error {
	path := fmt.Sprintf("/console/api/datasets/%s/documents/%s/segments/%s", datasetID, documentID, segmentID)
	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}

// BatchImportSegments imports segments from a CSV file.
func (c *client.Client) BatchImportSegments(ctx context.Context, datasetID, documentID, uploadFileID string) (*BatchImportSegmentsResponse, error) {
	var result BatchImportSegmentsResponse
	path := fmt.Sprintf("/console/api/datasets/%s/documents/%s/segments/batch_import", datasetID, documentID)
	req := map[string]string{"upload_file_id": uploadFileID}
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetSegmentBatchImportStatus gets the status of a segment batch import job.
func (c *client.Client) GetSegmentBatchImportStatus(ctx context.Context, jobID string) (*BatchImportSegmentsResponse, error) {
	var result BatchImportSegmentsResponse
	path := fmt.Sprintf("/console/api/datasets/batch_import_status/%s", jobID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
// GetDocumentChildChunks retrieves a list of child chunks for a segment.
func (c *client.Client) GetDocumentChildChunks(ctx context.Context, datasetID, documentID, segmentID string, page, limit int, keyword string) (*ChildChunkListResponse, error) {
	var result ChildChunkListResponse
	path := fmt.Sprintf("/console/api/datasets/%s/documents/%s/segments/%s/child_chunks?page=%d&limit=%d&keyword=%s", datasetID, documentID, segmentID, page, limit, keyword)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// AddDocumentChildChunk adds a new child chunk to a segment.
func (c *client.Client) AddDocumentChildChunk(ctx context.Context, datasetID, documentID, segmentID, content string) (*ConsoleChildChunk, error) {
	var response struct {
		Data ConsoleChildChunk `json:"data"`
	}
	path := fmt.Sprintf("/console/api/datasets/%s/documents/%s/segments/%s/child_chunks", datasetID, documentID, segmentID)
	req := map[string]string{"content": content}
	err := c.sendRequest(ctx, "POST", path, req, &response, nil)
	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

// UpdateDocumentChildChunk updates a child chunk.
func (c *client.Client) UpdateDocumentChildChunk(ctx context.Context, datasetID, documentID, segmentID, chunkID, content string) (*ConsoleChildChunk, error) {
	var response struct {
		Data ConsoleChildChunk `json:"data"`
	}
	path := fmt.Sprintf("/console/api/datasets/%s/documents/%s/segments/%s/child_chunks/%s", datasetID, documentID, segmentID, chunkID)
	req := map[string]string{"content": content}
	err := c.sendRequest(ctx, "PATCH", path, req, &response, nil)
	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

// DeleteDocumentChildChunk deletes a child chunk.
func (c *client.Client) DeleteDocumentChildChunk(ctx context.Context, datasetID, documentID, segmentID, chunkID string) error {
	path := fmt.Sprintf("/console/api/datasets/%s/documents/%s/segments/%s/child_chunks/%s", datasetID, documentID, segmentID, chunkID)
	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}
