package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// ExternalAPITemplate represents an external knowledge API template.
type ExternalAPITemplate struct {
	ID        string                 `json:"id"`
	Name      string                 `json:"name"`
	Settings  map[string]interface{} `json:"settings"`
	CreatedAt int64                  `json:"created_at"`
}

// ExternalAPITemplateListResponse is the paginated response for listing external API templates.
type ExternalAPITemplateListResponse struct {
	Data    []ExternalAPITemplate `json:"data"`
	HasMore bool                  `json:"has_more"`
	Limit   int                   `json:"limit"`
	Total   int                   `json:"total"`
	Page    int                   `json:"page"`
}

// CreateExternalAPITemplateRequest is the request to create an external API template.
type CreateExternalAPITemplateRequest struct {
	Name     string                 `json:"name"`
	Settings map[string]interface{} `json:"settings"`
}

// UpdateExternalAPITemplateRequest is the request to update an external API template.
type UpdateExternalAPITemplateRequest struct {
	Name     string                 `json:"name"`
	Settings map[string]interface{} `json:"settings"`
}

// ExternalAPIUseCheckResponse is the response for checking if an external API is in use.
type ExternalAPIUseCheckResponse struct {
	IsUsing bool `json:"is_using"`
	Count   int  `json:"count"`
}

// CreateExternalDatasetRequest is the request to create an external dataset.
type CreateExternalDatasetRequest struct {
	ExternalKnowledgeAPIID string                 `json:"external_knowledge_api_id"`
	ExternalKnowledgeID    string                 `json:"external_knowledge_id"`
	Name                   string                 `json:"name"`
	Description            string                 `json:"description,omitempty"`
	ExternalRetrievalModel map[string]interface{} `json:"external_retrieval_model,omitempty"`
}

// ExternalHitTestingRequest is the request for external hit testing.
type ExternalHitTestingRequest struct {
	Query                       string                 `json:"query"`
	ExternalRetrievalModel      map[string]interface{} `json:"external_retrieval_model,omitempty"`
	MetadataFilteringConditions map[string]interface{} `json:"metadata_filtering_conditions,omitempty"`
}

// GetExternalAPITemplates retrieves a list of external knowledge API templates.
func (c *client.Client) GetExternalAPITemplates(ctx context.Context, page, limit int, keyword string) (*ExternalAPITemplateListResponse, error) {
	var result ExternalAPITemplateListResponse
	path := fmt.Sprintf("/console/api/datasets/external-knowledge-api?page=%d&limit=%d&keyword=%s", page, limit, keyword)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateExternalAPITemplate creates a new external knowledge API template.
func (c *client.Client) CreateExternalAPITemplate(ctx context.Context, req *CreateExternalAPITemplateRequest) (*ExternalAPITemplate, error) {
	var result ExternalAPITemplate
	err := c.sendRequest(ctx, "POST", "/console/api/datasets/external-knowledge-api", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetExternalAPITemplate retrieves a single external knowledge API template.
func (c *client.Client) GetExternalAPITemplate(ctx context.Context, templateID string) (*ExternalAPITemplate, error) {
	var result ExternalAPITemplate
	path := fmt.Sprintf("/console/api/datasets/external-knowledge-api/%s", templateID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateExternalAPITemplate updates an external knowledge API template.
func (c *client.Client) UpdateExternalAPITemplate(ctx context.Context, templateID string, req *UpdateExternalAPITemplateRequest) (*ExternalAPITemplate, error) {
	var result ExternalAPITemplate
	path := fmt.Sprintf("/console/api/datasets/external-knowledge-api/%s", templateID)
	err := c.sendRequest(ctx, "PATCH", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteExternalAPITemplate deletes an external knowledge API template.
func (c *client.Client) DeleteExternalAPITemplate(ctx context.Context, templateID string) error {
	path := fmt.Sprintf("/console/api/datasets/external-knowledge-api/%s", templateID)
	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}

// CheckExternalAPIInUse checks if an external knowledge API is in use.
func (c *client.Client) CheckExternalAPIInUse(ctx context.Context, templateID string) (*ExternalAPIUseCheckResponse, error) {
	var result ExternalAPIUseCheckResponse
	path := fmt.Sprintf("/console/api/datasets/external-knowledge-api/%s/use-check", templateID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateExternalDataset creates a new external dataset.
func (c *client.Client) CreateExternalDataset(ctx context.Context, req *CreateExternalDatasetRequest) (*ConsoleDataset, error) {
	var result ConsoleDataset
	err := c.sendRequest(ctx, "POST", "/console/api/datasets/external", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ExternalHitTesting performs a hit testing query on an external dataset.
func (c *client.Client) ExternalHitTesting(ctx context.Context, datasetID string, req *ExternalHitTestingRequest) (map[string]interface{}, error) {
	var result map[string]interface{}
	path := fmt.Sprintf("/console/api/datasets/%s/external-hit-testing", datasetID)
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}
