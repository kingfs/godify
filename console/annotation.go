package console

import (
	"context"
	"fmt"
	"io"

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

// AnnotationSetting is the response for getting annotation settings.
type AnnotationSetting struct {
	ID                  string  `json:"id"`
	CollectionMethod    string  `json:"collection_method"`
	ScoreThreshold      float64 `json:"score_threshold"`
	EmbeddingModel      string  `json:"embedding_model"`
	EmbeddingModelProvider string `json:"embedding_model_provider"`
}

// UpdateAnnotationSettingRequest is the request to update annotation settings.
type UpdateAnnotationSettingRequest struct {
	ScoreThreshold float64 `json:"score_threshold"`
}

// EnableAnnotationRequest is the request to enable annotation reply.
type EnableAnnotationRequest struct {
	ScoreThreshold       float64 `json:"score_threshold"`
	EmbeddingProviderName string  `json:"embedding_provider_name"`
	EmbeddingModelName   string  `json:"embedding_model_name"`
}

// AnnotationActionResponse is a generic response for annotation actions.
type AnnotationActionResponse struct {
	Result string `json:"result"`
}

// AnnotationStatusResponse is the response for checking an annotation job status.
type AnnotationStatusResponse struct {
	JobID     string `json:"job_id"`
	JobStatus string `json:"job_status"`
	ErrorMsg  string `json:"error_msg"`
}

// AnnotationListResponse is the paginated response for listing annotations.
type AnnotationListResponse struct {
	Data    []Annotation `json:"data"`
	HasMore bool         `json:"has_more"`
	Limit   int          `json:"limit"`
	Total   int          `json:"total"`
	Page    int          `json:"page"`
}

// CreateAnnotationRequest is the request to create an annotation.
type CreateAnnotationRequest struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

// AnnotationHitHistory represents a single hit history record.
type AnnotationHitHistory struct {
	ID        string  `json:"id"`
	Source    string  `json:"source"`
	Score     float64 `json:"score"`
	Question  string  `json:"question"`
	CreatedAt int64   `json:"created_at"`
	Match     string  `json:"match"`
	Response  string  `json:"response"`
}

// AnnotationHitHistoryListResponse is the paginated response for hit histories.
type AnnotationHitHistoryListResponse struct {
	Data    []AnnotationHitHistory `json:"data"`
	HasMore bool                   `json:"has_more"`
	Limit   int                    `json:"limit"`
	Total   int                    `json:"total"`
	Page    int                    `json:"page"`
}

// EnableAnnotationReply enables the annotation reply feature for an app.
func (c *client.Client) EnableAnnotationReply(ctx context.Context, appID string, req *EnableAnnotationRequest) (*AnnotationActionResponse, error) {
	var result AnnotationActionResponse
	path := fmt.Sprintf("/console/api/apps/%s/annotation-reply/enable", appID)
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DisableAnnotationReply disables the annotation reply feature for an app.
func (c *client.Client) DisableAnnotationReply(ctx context.Context, appID string, req *EnableAnnotationRequest) (*AnnotationActionResponse, error) {
	var result AnnotationActionResponse
	path := fmt.Sprintf("/console/api/apps/%s/annotation-reply/disable", appID)
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAnnotationReplyStatus gets the status of an annotation reply job.
func (c *client.Client) GetAnnotationReplyStatus(ctx context.Context, appID, action, jobID string) (*AnnotationStatusResponse, error) {
	var result AnnotationStatusResponse
	path := fmt.Sprintf("/console/api/apps/%s/annotation-reply/%s/status/%s", appID, action, jobID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAnnotations retrieves a list of annotations for an app.
func (c *client.Client) GetAnnotations(ctx context.Context, appID, keyword string, page, limit int) (*AnnotationListResponse, error) {
	var result AnnotationListResponse
	path := fmt.Sprintf("/console/api/apps/%s/annotations?page=%d&limit=%d&keyword=%s", appID, page, limit, keyword)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteAnnotationsInBatch deletes multiple annotations for an app.
func (c *client.Client) DeleteAnnotationsInBatch(ctx context.Context, appID string, annotationIDs []string) error {
	path := fmt.Sprintf("/console/api/apps/%s/annotations", appID)
	// The API expects the IDs as repeated query parameters, which is not well supported by the current client.
	// I will construct the URL manually.
	for _, id := range annotationIDs {
		path += fmt.Sprintf("&annotation_id=%s", id)
	}
	path =- "/console/api/apps/" + appID + "/annotations?" + path[1:]

	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}

// ClearAllAnnotations deletes all annotations for an app.
func (c *client.Client) ClearAllAnnotations(ctx context.Context, appID string) error {
	path := fmt.Sprintf("/console/api/apps/%s/annotations", appID)
	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}


// ExportAnnotations exports all annotations for an app.
func (c *client.Client) ExportAnnotations(ctx context.Context, appID string) (*AnnotationListResponse, error) {
	var result AnnotationListResponse
	path := fmt.Sprintf("/console/api/apps/%s/annotations/export", appID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateAnnotation updates a single annotation.
func (c *client.Client) UpdateAnnotation(ctx context.Context, appID, annotationID string, req *CreateAnnotationRequest) (*Annotation, error) {
	var result Annotation
	path := fmt.Sprintf("/console/api/apps/%s/annotations/%s", appID, annotationID)
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteAnnotation deletes a single annotation.
func (c *client.Client) DeleteAnnotation(ctx context.Context, appID, annotationID string) error {
	path := fmt.Sprintf("/console/api/apps/%s/annotations/%s", appID, annotationID)
	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}

// BatchImportAnnotations imports annotations from a CSV file.
func (c *client.Client) BatchImportAnnotations(ctx context.Context, appID string, file io.Reader) (*AnnotationStatusResponse, error) {
	var result AnnotationStatusResponse
	path := fmt.Sprintf("/console/api/apps/%s/annotations/batch-import", appID)
	err := c.SendMultipartRequest(ctx, path, file, "annotations.csv", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetBatchImportStatus gets the status of a batch import job.
func (c *client.Client) GetBatchImportStatus(ctx context.Context, appID, jobID string) (*AnnotationStatusResponse, error) {
	var result AnnotationStatusResponse
	path := fmt.Sprintf("/console/api/apps/%s/annotations/batch-import-status/%s", appID, jobID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAnnotationHitHistories retrieves the hit history for an annotation.
func (c *client.Client) GetAnnotationHitHistories(ctx context.Context, appID, annotationID string, page, limit int) (*AnnotationHitHistoryListResponse, error) {
	var result AnnotationHitHistoryListResponse
	path := fmt.Sprintf("/console/api/apps/%s/annotations/%s/hit-histories?page=%d&limit=%d", appID, annotationID, page, limit)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAnnotationSetting retrieves the annotation settings for an app.
func (c *client.Client) GetAnnotationSetting(ctx context.Context, appID string) (*AnnotationSetting, error) {
	var result AnnotationSetting
	path := fmt.Sprintf("/console/api/apps/%s/annotation-setting", appID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateAnnotationSetting updates the annotation settings for an app.
func (c *client.Client) UpdateAnnotationSetting(ctx context.Context, appID, settingID string, req *UpdateAnnotationSettingRequest) (*AnnotationSetting, error) {
	var result AnnotationSetting
	path := fmt.Sprintf("/console/api/apps/%s/annotation-settings/%s", appID, settingID)
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
