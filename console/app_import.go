package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// AppImportRequest is the request to import an app.
type AppImportRequest struct {
	Mode           string `json:"mode"`
	YAMLContent    string `json:"yaml_content,omitempty"`
	YAMLURL        string `json:"yaml_url,omitempty"`
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
	IconType       string `json:"icon_type,omitempty"`
	Icon           string `json:"icon,omitempty"`
	IconBackground string `json:"icon_background,omitempty"`
	AppID          string `json:"app_id,omitempty"`
}

// AppImportResponse is the response from an app import request.
type AppImportResponse struct {
	ID                 string `json:"id"`
	Status             string `json:"status"`
	AppID              string `json:"app_id"`
	AppMode            string `json:"app_mode"`
	CurrentDSLVersion  string `json:"current_dsl_version"`
	ImportedDSLVersion string `json:"imported_dsl_version"`
	Error              string `json:"error"`
}

// LeakedDependency represents a dependency that needs to be resolved.
type LeakedDependency struct {
	Type              string      `json:"type"`
	Value             interface{} `json:"value"`
	CurrentIdentifier string      `json:"current_identifier"`
}

// CheckAppDependenciesResponse is the response from checking app dependencies.
type CheckAppDependenciesResponse struct {
	LeakedDependencies []LeakedDependency `json:"leaked_dependencies"`
}

// ImportApp imports an application from a YAML definition.
func (c *client.Client) ImportApp(ctx context.Context, req *AppImportRequest) (*AppImportResponse, error) {
	var result AppImportResponse
	err := c.sendRequest(ctx, "POST", "/console/api/apps/imports", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ConfirmAppImport confirms a pending app import.
func (c *client.Client) ConfirmAppImport(ctx context.Context, importID string) (*AppImportResponse, error) {
	var result AppImportResponse
	path := fmt.Sprintf("/console/api/apps/imports/%s/confirm", importID)
	err := c.sendRequest(ctx, "POST", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CheckAppDependencies checks for leaked dependencies in an imported app.
func (c *client.Client) CheckAppDependencies(ctx context.Context, appID string) (*CheckAppDependenciesResponse, error) {
	var result CheckAppDependenciesResponse
	path := fmt.Sprintf("/console/api/apps/imports/%s/check-dependencies", appID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
