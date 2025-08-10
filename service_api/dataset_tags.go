package service_api

import (
	"context"
	"fmt"

	"github.comcom/kingfs/godify/client"
)

// Tag represents a tag that can be applied to datasets.
type Tag struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	BindingCount int    `json:"binding_count"`
}

// CreateTagRequest is the request body for creating a tag.
type CreateTagRequest struct {
	Name string `json:"name"`
}

// UpdateTagRequest is the request body for updating a tag.
type UpdateTagRequest struct {
	Name  string `json:"name"`
	TagID string `json:"tag_id"`
}

// BindTagsRequest is the request body for binding tags to a target.
type BindTagsRequest struct {
	TargetID string   `json:"target_id"`
	TagIDs   []string `json:"tag_ids"`
}

// UnbindTagRequest is the request body for unbinding a tag from a target.
type UnbindTagRequest struct {
	TargetID string `json:"target_id"`
	TagID    string `json:"tag_id"`
}

// DatasetTagsResponse is the response for getting tags for a specific dataset.
type DatasetTagsResponse struct {
	Data  []Tag `json:"data"`
	Total int   `json:"total"`
}

// GetTags retrieves all knowledge type tags.
func (c *client.Client) GetTags(ctx context.Context, datasetID string) ([]Tag, error) {
	var result []Tag
	// The python code indicates this endpoint is /datasets/tags, but it is associated with a dataset_id in the class.
	// This seems to be a generic get for all tags of type 'knowledge'.
	// The `dataset_id` is in the URL for the class, but not used in the `get` method path.
	// I'll use the path from the `api.add_resource` call.
	err := c.sendRequest(ctx, "GET", "/v1/datasets/tags", nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateTag creates a new knowledge type tag.
func (c *client.Client) CreateTag(ctx context.Context, name string) (*Tag, error) {
	var result Tag
	req := CreateTagRequest{Name: name}
	err := c.sendRequest(ctx, "POST", "/v1/datasets/tags", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateTag updates an existing tag.
func (c *client.Client) UpdateTag(ctx context.Context, tagID, name string) (*Tag, error) {
	var result Tag
	req := UpdateTagRequest{TagID: tagID, Name: name}
	err := c.sendRequest(ctx, "PATCH", "/v1/datasets/tags", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteTag deletes a tag.
func (c *client.Client) DeleteTag(ctx context.Context, tagID string) error {
	req := map[string]string{"tag_id": tagID}
	return c.sendRequest(ctx, "DELETE", "/v1/datasets/tags", req, nil, nil)
}

// BindTags binds one or more tags to a target dataset.
func (c *client.Client) BindTags(ctx context.Context, targetID string, tagIDs []string) error {
	req := BindTagsRequest{TargetID: targetID, TagIDs: tagIDs}
	return c.sendRequest(ctx, "POST", "/v1/datasets/tags/binding", req, nil, nil)
}

// UnbindTag unbinds a tag from a target dataset.
func (c *client.Client) UnbindTag(ctx context.Context, targetID, tagID string) error {
	req := UnbindTagRequest{TargetID: targetID, TagID: tagID}
	return c.sendRequest(ctx, "POST", "/v1/datasets/tags/unbinding", req, nil, nil)
}

// GetDatasetTags retrieves the tags for a specific dataset.
func (c *client.Client) GetDatasetTags(ctx context.Context, datasetID string) (*DatasetTagsResponse, error) {
	var result DatasetTagsResponse
	path := fmt.Sprintf("/v1/datasets/%s/tags", datasetID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
