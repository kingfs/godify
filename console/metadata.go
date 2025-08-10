package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// ConsoleDatasetMetadata represents a metadata field in a dataset.
type ConsoleDatasetMetadata struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

// CreateConsoleMetadataRequest is the request to create a new metadata field.
type CreateConsoleMetadataRequest struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

// UpdateConsoleMetadataRequest is the request to update a metadata field.
type UpdateConsoleMetadataRequest struct {
	Name string `json:"name"`
}

// UpdateConsoleDocumentsMetadataRequest is the request to update metadata for multiple documents.
type UpdateConsoleDocumentsMetadataRequest struct {
	OperationData []interface{} `json:"operation_data"`
}

// ConsoleBuiltInFieldsResponse is the response for getting built-in metadata fields.
type ConsoleBuiltInFieldsResponse struct {
	Fields []interface{} `json:"fields"`
}

// ConsoleCreateDatasetMetadata creates a new metadata field for a dataset.
func (c *client.Client) ConsoleCreateDatasetMetadata(ctx context.Context, datasetID string, req *CreateConsoleMetadataRequest) (*ConsoleDatasetMetadata, error) {
	var result ConsoleDatasetMetadata
	path := fmt.Sprintf("/console/api/datasets/%s/metadata", datasetID)
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ConsoleGetDatasetMetadata retrieves all metadata fields for a dataset.
func (c *client.Client) ConsoleGetDatasetMetadata(ctx context.Context, datasetID string) ([]ConsoleDatasetMetadata, error) {
	var result []ConsoleDatasetMetadata
	path := fmt.Sprintf("/console/api/datasets/%s/metadata", datasetID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ConsoleUpdateDatasetMetadata updates a metadata field.
func (c *client.Client) ConsoleUpdateDatasetMetadata(ctx context.Context, datasetID, metadataID, name string) (*ConsoleDatasetMetadata, error) {
	var result ConsoleDatasetMetadata
	path := fmt.Sprintf("/console/api/datasets/%s/metadata/%s", datasetID, metadataID)
	req := UpdateConsoleMetadataRequest{Name: name}
	err := c.sendRequest(ctx, "PATCH", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ConsoleDeleteDatasetMetadata deletes a metadata field.
func (c *client.Client) ConsoleDeleteDatasetMetadata(ctx context.Context, datasetID, metadataID string) error {
	path := fmt.Sprintf("/console/api/datasets/%s/metadata/%s", datasetID, metadataID)
	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}

// ConsoleGetBuiltInMetadataFields retrieves the list of built-in metadata fields.
func (c *client.Client) ConsoleGetBuiltInMetadataFields(ctx context.Context) (*ConsoleBuiltInFieldsResponse, error) {
	var result ConsoleBuiltInFieldsResponse
	err := c.sendRequest(ctx, "GET", "/console/api/datasets/metadata/built-in", nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ConsoleToggleBuiltInMetadata enables or disables the built-in metadata fields for a dataset.
func (c *client.Client) ConsoleToggleBuiltInMetadata(ctx context.Context, datasetID, action string) error {
	path := fmt.Sprintf("/console/api/datasets/%s/metadata/built-in/%s", datasetID, action)
	return c.sendRequest(ctx, "POST", path, nil, nil, nil)
}

// ConsoleUpdateDocumentsMetadata updates metadata for a batch of documents.
func (c *client.Client) ConsoleUpdateDocumentsMetadata(ctx context.Context, datasetID string, req *UpdateConsoleDocumentsMetadataRequest) error {
	path := fmt.Sprintf("/console/api/datasets/%s/documents/metadata", datasetID)
	return c.sendRequest(ctx, "POST", path, req, nil, nil)
}
