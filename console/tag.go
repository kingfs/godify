package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// ConsoleTag represents a tag in the console.
type ConsoleTag struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	BindingCount int    `json:"binding_count"`
}

// CreateTagRequest is the request to create a new tag.
type CreateTagRequest struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// UpdateTagRequest is the request to update a tag.
type UpdateTagRequest struct {
	Name string `json:"name"`
}

// CreateTagBindingRequest is the request to bind tags to a target.
type CreateTagBindingRequest struct {
	TagIDs   []string `json:"tag_ids"`
	TargetID string   `json:"target_id"`
	Type     string   `json:"type"`
}

// RemoveTagBindingRequest is the request to remove a tag binding.
type RemoveTagBindingRequest struct {
	TagID    string `json:"tag_id"`
	TargetID string `json:"target_id"`
	Type     string `json:"type"`
}

// GetTags retrieves a list of tags.
func (c *client.Client) GetTags(ctx context.Context, tagType, keyword string) ([]ConsoleTag, error) {
	var result []ConsoleTag
	path := fmt.Sprintf("/console/api/tags?type=%s&keyword=%s", tagType, keyword)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateTag creates a new tag.
func (c *client.Client) CreateTag(ctx context.Context, req *CreateTagRequest) (*ConsoleTag, error) {
	var result ConsoleTag
	err := c.sendRequest(ctx, "POST", "/console/api/tags", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateTag updates an existing tag.
func (c *client.Client) UpdateTag(ctx context.Context, tagID, name string) (*ConsoleTag, error) {
	var result ConsoleTag
	path := fmt.Sprintf("/console/api/tags/%s", tagID)
	req := UpdateTagRequest{Name: name}
	err := c.sendRequest(ctx, "PATCH", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteTag deletes a tag.
func (c *client.Client) DeleteTag(ctx context.Context, tagID string) error {
	path := fmt.Sprintf("/console/api/tags/%s", tagID)
	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}

// CreateTagBinding creates a new tag binding.
func (c *client.Client) CreateTagBinding(ctx context.Context, req *CreateTagBindingRequest) error {
	return c.sendRequest(ctx, "POST", "/console/api/tag-bindings/create", req, nil, nil)
}

// RemoveTagBinding removes a tag binding.
func (c *client.Client) RemoveTagBinding(ctx context.Context, req *RemoveTagBindingRequest) error {
	return c.sendRequest(ctx, "POST", "/console/api/tag-bindings/remove", req, nil, nil)
}
