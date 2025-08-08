package service_api

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// Annotation represents a single annotation record.
type Annotation struct {
	ID        string `json:"id"`
	Question  string `json:"question"`
	Answer    string `json:"answer"`
	HitCount  int    `json:"hit_count"`
	CreatedAt int64  `json:"created_at"`
}

// EnableAnnotationRequest is the request body for enabling annotation reply.
type EnableAnnotationRequest struct {
	ScoreThreshold       float64 `json:"score_threshold"`
	EmbeddingProviderName string  `json:"embedding_provider_name"`
	EmbeddingModelName   string  `json:"embedding_model_name"`
}

// AnnotationActionResponse is the response for annotation actions.
type AnnotationActionResponse struct {
	Result string `json:"result"` // Assuming a simple success/failure message
}

// AnnotationStatusResponse is the response for checking an annotation job status.
type AnnotationStatusResponse struct {
	JobID     string `json:"job_id"`
	JobStatus string `json:"job_status"`
	ErrorMsg  string `json:"error_msg"`
}

// AnnotationListResponse is the response for listing annotations.
type AnnotationListResponse struct {
	Data    []Annotation `json:"data"`
	HasMore bool         `json:"has_more"`
	Limit   int          `json:"limit"`
	Total   int          `json:"total"`
	Page    int          `json:"page"`
}

// CreateAnnotationRequest is the request body for creating an annotation.
type CreateAnnotationRequest struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

// UpdateAnnotationRequest is the request body for updating an annotation.
type UpdateAnnotationRequest struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

// DeleteResponse is a generic response for delete operations.
type DeleteResponse struct {
	Result string `json:"result"`
}

// EnableAnnotationReply enables the annotation reply feature for an app.
func (c *client.Client) EnableAnnotationReply(ctx context.Context, req *EnableAnnotationRequest) (*AnnotationActionResponse, error) {
	var result AnnotationActionResponse
	err := c.sendRequest(ctx, "POST", "/v1/apps/annotation-reply/enable", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DisableAnnotationReply disables the annotation reply feature for an app.
func (c *client.Client) DisableAnnotationReply(ctx context.Context) (*AnnotationActionResponse, error) {
	var result AnnotationActionResponse
	err := c.sendRequest(ctx, "POST", "/v1/apps/annotation-reply/disable", nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAnnotationReplyStatus gets the status of an annotation reply job.
func (c *client.Client) GetAnnotationReplyStatus(ctx context.Context, action, jobID string) (*AnnotationStatusResponse, error) {
	var result AnnotationStatusResponse
	path := fmt.Sprintf("/v1/apps/annotation-reply/%s/status/%s", action, jobID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAnnotations retrieves a list of annotations for the app.
// Pass in optional query parameters for pagination and keyword search.
func (c *client.Client) GetAnnotations(ctx context.Context, page, limit int, keyword string) (*AnnotationListResponse, error) {
	var result AnnotationListResponse
	path := fmt.Sprintf("/v1/apps/annotations?page=%d&limit=%d&keyword=%s", page, limit, keyword)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateAnnotation creates a new annotation for the app.
func (c *client.Client) CreateAnnotation(ctx context.Context, req *CreateAnnotationRequest) (*Annotation, error) {
	var result Annotation
	err := c.sendRequest(ctx, "POST", "/v1/apps/annotations", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateAnnotation updates an existing annotation.
func (c *client.Client) UpdateAnnotation(ctx context.Context, annotationID string, req *UpdateAnnotationRequest) (*Annotation, error) {
	var result Annotation
	path := fmt.Sprintf("/v1/apps/annotations/%s", annotationID)
	err := c.sendRequest(ctx, "PUT", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteAnnotation deletes an annotation.
func (c *client.Client) DeleteAnnotation(ctx context.Context, annotationID string) (*DeleteResponse, error) {
	var result DeleteResponse
	path := fmt.Sprintf("/v1/apps/annotations/%s", annotationID)
	err := c.sendRequest(ctx, "DELETE", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
