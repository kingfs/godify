package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// InsertExploreAppRequest is the request to add an app to the explore list.
type InsertExploreAppRequest struct {
	AppID             string `json:"app_id"`
	Description       string `json:"desc,omitempty"`
	Copyright         string `json:"copyright,omitempty"`
	PrivacyPolicy     string `json:"privacy_policy,omitempty"`
	CustomDisclaimer  string `json:"custom_disclaimer,omitempty"`
	Language          string `json:"language"`
	Category          string `json:"category"`
	Position          int    `json:"position"`
}

// AdminActionResponse is a generic response for admin actions.
type AdminActionResponse struct {
	Result string `json:"result"`
}

// AddAppToExplore adds an app to the explore list.
// This requires a special admin API key passed in the function call.
func (c *client.Client) AddAppToExplore(ctx context.Context, adminKey string, req *InsertExploreAppRequest) (*AdminActionResponse, error) {
	var result AdminActionResponse
	headers := map[string]string{
		"Authorization": "Bearer " + adminKey,
	}

	err := c.sendRequest(ctx, "POST", "/console/api/admin/insert-explore-apps", req, &result, headers)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// RemoveAppFromExplore removes an app from the explore list.
// This requires a special admin API key passed in the function call.
func (c *client.Client) RemoveAppFromExplore(ctx context.Context, adminKey, appID string) (*AdminActionResponse, error) {
	var result AdminActionResponse
	headers := map[string]string{
		"Authorization": "Bearer " + adminKey,
	}
	path := fmt.Sprintf("/console/api/admin/insert-explore-apps/%s", appID)
	err := c.sendRequest(ctx, "DELETE", path, nil, &result, headers)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
