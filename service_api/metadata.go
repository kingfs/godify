package service_api

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// DatasetMetadata represents a metadata field in a dataset.
type DatasetMetadata struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

// CreateMetadataRequest is the request to create a new metadata field.
type CreateMetadataRequest struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

// UpdateMetadataRequest is the request to update a metadata field.
type UpdateMetadataRequest struct {
	Name string `json:"name"`
}

// MetadataOperationData represents a single operation in a batch metadata update.
type MetadataOperationData struct {
	DocumentID string `json:"document_id"`
	Metadata   map[string]interface{} `json:"metadata"`
}

// UpdateDocumentsMetadataRequest is the request to update metadata for multiple documents.
type UpdateDocumentsMetadataRequest struct {
	OperationData []MetadataOperationData `json:"operation_data"`
}

// BuiltInFieldsResponse is the response for getting built-in metadata fields.
type BuiltInFieldsResponse struct {
	Fields []interface{} `json:"fields"` // The structure is not defined, using interface{}
}

// CreateDatasetMetadata creates a new metadata field for a dataset.
func (c *client.Client) CreateDatasetMetadata(ctx context.Context, datasetID string, req *CreateMetadataRequest) (*DatasetMetadata, error) {
	var result DatasetMetadata
	path := fmt.Sprintf("/v1/datasets/%s/metadata", datasetID)
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDatasetMetadata retrieves all metadata fields for a dataset.
func (c *client.Client) GetDatasetMetadata(ctx context.Context, datasetID string) ([]DatasetMetadata, error) {
	var result []DatasetMetadata
	path := fmt.Sprintf("/v1/datasets/%s/metadata", datasetID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateDatasetMetadata updates a metadata field.
func (c *client.Client) UpdateDatasetMetadata(ctx context.Context, datasetID, metadataID, name string) (*DatasetMetadata, error) {
	var result DatasetMetadata
	path := fmt.Sprintf("/v1/datasets/%s/metadata/%s", datasetID, metadataID)
	req := UpdateMetadataRequest{Name: name}
	err := c.sendRequest(ctx, "PATCH", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteDatasetMetadata deletes a metadata field.
func (c *client.Client) DeleteDatasetMetadata(ctx context.Context, datasetID, metadataID string) error {
	path := fmt.Sprintf("/v1/datasets/%s/metadata/%s", datasetID, metadataID)
	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}

// GetBuiltInMetadataFields retrieves the list of built-in metadata fields.
func (c *client.Client) GetBuiltInMetadataFields(ctx context.Context) (*BuiltInFieldsResponse, error) {
	var result BuiltInFieldsResponse
	err := c.sendRequest(ctx, "GET", "/v1/datasets/metadata/built-in", nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ToggleBuiltInMetadata enables or disables the built-in metadata fields for a dataset.
func (c *client.Client) ToggleBuiltInMetadata(ctx context.Context, datasetID, action string) error {
	path := fmt.Sprintf("/v1/datasets/%s/metadata/built-in/%s", datasetID, action)
	return c.sendRequest(ctx, "POST", path, nil, nil, nil)
}

// UpdateDocumentsMetadata updates metadata for a batch of documents.
func (c *client.Client) UpdateDocumentsMetadata(ctx context.Context, datasetID string, req *UpdateDocumentsMetadataRequest) error {
	path := fmt.Sprintf("/v1/datasets/%s/documents/metadata", datasetID)
	return c.sendRequest(ctx, "POST", path, req, nil, nil)
}
