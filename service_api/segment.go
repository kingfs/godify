package service_api

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// CreateSegmentData represents the data for a single segment to be created.
type CreateSegmentData struct {
	Content  string   `json:"content"`
	Keywords []string `json:"keywords,omitempty"`
}

// CreateSegmentsRequest is the request to create multiple segments.
type CreateSegmentsRequest struct {
	Segments []CreateSegmentData `json:"segments"`
}

// SegmentListResponse is the paginated response for listing segments.
type SegmentListResponse struct {
	Data    []Segment `json:"data"`
	DocForm string    `json:"doc_form"`
	Total   int       `json:"total"`
	HasMore bool      `json:"has_more"`
	Limit   int       `json:"limit"`
	Page    int       `json:"page"`
}

// ChildChunkListResponse is the paginated response for listing child chunks.
type ChildChunkListResponse struct {
	Data       []ChildChunk `json:"data"`
	Total      int          `json:"total"`
	TotalPages int          `json:"total_pages"`
	Page       int          `json:"page"`
	Limit      int          `json:"limit"`
}

// CreateSegments creates multiple segments for a document.
func (c *client.Client) CreateSegments(ctx context.Context, datasetID, documentID string, req *CreateSegmentsRequest) ([]Segment, error) {
	var response struct {
		Data    []Segment `json:"data"`
		DocForm string    `json:"doc_form"`
	}
	path := fmt.Sprintf("/v1/datasets/%s/documents/%s/segments", datasetID, documentID)
	err := c.sendRequest(ctx, "POST", path, req, &response, nil)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// GetSegments retrieves a list of segments for a document.
func (c *client.Client) GetSegments(ctx context.Context, datasetID, documentID, keyword string, status []string, page, limit int) (*SegmentListResponse, error) {
	var result SegmentListResponse
	path := fmt.Sprintf("/v1/datasets/%s/documents/%s/segments?page=%d&limit=%d", datasetID, documentID, page, limit)
	if keyword != "" {
		path += fmt.Sprintf("&keyword=%s", keyword)
	}
	for _, s := range status {
		path += fmt.Sprintf("&status=%s", s)
	}
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetSegment retrieves a single segment.
func (c *client.Client) GetSegment(ctx context.Context, datasetID, documentID, segmentID string) (*Segment, error) {
	var response struct {
		Data    Segment `json:"data"`
		DocForm string  `json:"doc_form"`
	}
	path := fmt.Sprintf("/v1/datasets/%s/documents/%s/segments/%s", datasetID, documentID, segmentID)
	err := c.sendRequest(ctx, "GET", path, nil, &response, nil)
	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

// UpdateSegment updates a segment. The req should contain a "segment" key with the segment data.
func (c *client.Client) UpdateSegment(ctx context.Context, datasetID, documentID, segmentID string, req map[string]interface{}) (*Segment, error) {
	var response struct {
		Data    Segment `json:"data"`
		DocForm string  `json:"doc_form"`
	}
	path := fmt.Sprintf("/v1/datasets/%s/documents/%s/segments/%s", datasetID, documentID, segmentID)
	err := c.sendRequest(ctx, "POST", path, req, &response, nil)
	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

// DeleteSegment deletes a segment.
func (c *client.Client) DeleteSegment(ctx context.Context, datasetID, documentID, segmentID string) error {
	path := fmt.Sprintf("/v1/datasets/%s/documents/%s/segments/%s", datasetID, documentID, segmentID)
	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}

// GetChildChunks retrieves child chunks for a segment.
func (c *client.Client) GetChildChunks(ctx context.Context, datasetID, documentID, segmentID string, page, limit int, keyword string) (*ChildChunkListResponse, error) {
	var result ChildChunkListResponse
	path := fmt.Sprintf("/v1/datasets/%s/documents/%s/segments/%s/child_chunks?page=%d&limit=%d", datasetID, documentID, segmentID, page, limit)
	if keyword != "" {
		path += fmt.Sprintf("&keyword=%s", keyword)
	}
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateChildChunk creates a child chunk for a segment.
func (c *client.Client) CreateChildChunk(ctx context.Context, datasetID, documentID, segmentID, content string) (*ChildChunk, error) {
	var response struct {
		Data ChildChunk `json:"data"`
	}
	path := fmt.Sprintf("/v1/datasets/%s/documents/%s/segments/%s/child_chunks", datasetID, documentID, segmentID)
	req := map[string]string{"content": content}
	err := c.sendRequest(ctx, "POST", path, req, &response, nil)
	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

// UpdateChildChunk updates a child chunk.
func (c *client.Client) UpdateChildChunk(ctx context.Context, datasetID, documentID, segmentID, childChunkID, content string) (*ChildChunk, error) {
	var response struct {
		Data ChildChunk `json:"data"`
	}
	path := fmt.Sprintf("/v1/datasets/%s/documents/%s/segments/%s/child_chunks/%s", datasetID, documentID, segmentID, childChunkID)
	req := map[string]string{"content": content}
	err := c.sendRequest(ctx, "PATCH", path, req, &response, nil)
	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

// DeleteChildChunk deletes a child chunk.
func (c *client.Client) DeleteChildChunk(ctx context.Context, datasetID, documentID, segmentID, childChunkID string) error {
	path := fmt.Sprintf("/v1/datasets/%s/documents/%s/segments/%s/child_chunks/%s", datasetID, documentID, segmentID, childChunkID)
	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}
