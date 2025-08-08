package service_api

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// Segment represents a segment of a document.
type Segment struct {
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
	Document      *Document              `json:"document"` // Using the Document struct from document.go
}

// ChildChunk represents a child chunk of a segment.
type ChildChunk struct {
	ID       string  `json:"id"`
	Content  string  `json:"content"`
	Position int     `json:"position"`
	Score    float64 `json:"score"`
}

// HitTestingRecord represents a single record from a hit testing query.
type HitTestingRecord struct {
	Segment     *Segment     `json:"segment"`
	ChildChunks []ChildChunk `json:"child_chunks"`
	Score       float64      `json:"score"`
}

// HitTestingResponse is the response from a hit testing query.
type HitTestingResponse struct {
	Query   string             `json:"query"`
	Records []HitTestingRecord `json:"records"`
}

// HitTestingRequest is the request body for a hit testing query.
type HitTestingRequest struct {
	Query          string          `json:"query"`
	RetrievalModel *RetrievalModel `json:"retrieval_model,omitempty"`
}

// HitTest performs a hit testing query on a dataset.
func (c *client.Client) HitTest(ctx context.Context, datasetID string, req *HitTestingRequest) (*HitTestingResponse, error) {
	var result HitTestingResponse
	path := fmt.Sprintf("/v1/datasets/%s/hit-testing", datasetID)
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
