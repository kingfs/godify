package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// RecommendedApp represents an app in the explore section.
type RecommendedApp struct {
	App            App    `json:"app"`
	AppID          string `json:"app_id"`
	Description    string `json:"description"`
	Copyright      string `json:"copyright"`
	PrivacyPolicy  string `json:"privacy_policy"`
	CustomDisclaimer string `json:"custom_disclaimer"`
	Category       string `json:"category"`
	Position       int    `json:"position"`
	IsListed       bool   `json:"is_listed"`
}

// RecommendedAppListResponse is the response for listing recommended apps.
type RecommendedAppListResponse struct {
	RecommendedApps []RecommendedApp `json:"recommended_apps"`
	Categories      []string         `json:"categories"`
}

// GetRecommendedApps retrieves the list of recommended apps from the explore section.
func (c *client.Client) GetRecommendedApps(ctx context.Context, language string) (*RecommendedAppListResponse, error) {
	var result RecommendedAppListResponse
	path := "/console/api/explore/apps"
	if language != "" {
		path += fmt.Sprintf("?language=%s", language)
	}
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetRecommendedAppDetail retrieves the details of a single recommended app.
// The response for this is not marshaled, so it's likely a raw dictionary.
func (c *client.Client) GetRecommendedAppDetail(ctx context.Context, appID string) (map[string]interface{}, error) {
	var result map[string]interface{}
	path := fmt.Sprintf("/console/api/explore/apps/%s", appID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}
