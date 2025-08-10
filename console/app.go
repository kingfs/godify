package console

import (
	"context"
	"fmt"

	"github.comcom/kingfs/godify/client"
)

// ModelConfig represents the model configuration for an app.
type ModelConfig struct {
	OpeningStatement         string                 `json:"opening_statement"`
	SuggestedQuestions       []string               `json:"suggested_questions"`
	Model                    map[string]interface{} `json:"model"`
	UserInputForm            []interface{}          `json:"user_input_form"`
	PrePrompt                string                 `json:"pre_prompt"`
	AgentMode                map[string]interface{} `json:"agent_mode"`
}

// App represents the basic details of an application.
type App struct {
	ID              string       `json:"id"`
	Name            string       `json:"name"`
	Description     string       `json:"description"`
	Mode            string       `json:"mode"`
	Icon            string       `json:"icon"`
	IconBackground  string       `json:"icon_background"`
	ModelConfig     *ModelConfig `json:"model_config"`
	CreatedAt       int64        `json:"created_at"`
}

// AppListResponse is the paginated response for listing apps.
type AppListResponse struct {
	Data    []App `json:"data"`
	HasMore bool  `json:"has_more"`
	Limit   int   `json:"limit"`
	Total   int   `json:"total"`
	Page    int   `json:"page"`
}

// Site represents the site configuration for an app.
type Site struct {
	AccessToken string `json:"access_token"`
	Title       string `json:"title"`
	// ... other fields
}

// AppDetail represents the detailed information of an application.
type AppDetail struct {
	App
	EnableSite bool  `json:"enable_site"`
	EnableAPI  bool  `json:"enable_api"`
	Site       *Site `json:"site"`
}

// CreateAppRequest is the request to create a new app.
type CreateAppRequest struct {
	Name           string `json:"name"`
	Description    string `json:"description,omitempty"`
	Mode           string `json:"mode"`
	Icon           string `json:"icon,omitempty"`
	IconBackground string `json:"icon_background,omitempty"`
}

// UpdateAppRequest is the request to update an app.
type UpdateAppRequest struct {
	Name                string `json:"name"`
	Description         string `json:"description,omitempty"`
	Icon                string `json:"icon,omitempty"`
	IconBackground      string `json:"icon_background,omitempty"`
	UseIconAsAnswerIcon bool   `json:"use_icon_as_answer_icon,omitempty"`
}

// AppStatusRequest is the request to enable/disable site or api access.
type AppStatusRequest struct {
	Enable bool `json:"enable"`
}

// GetApps retrieves a list of applications.
func (c *client.Client) GetApps(ctx context.Context, page, limit int, mode, name string) (*AppListResponse, error) {
	var result AppListResponse
	path := fmt.Sprintf("/console/api/apps?page=%d&limit=%d&mode=%s&name=%s", page, limit, mode, name)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateApp creates a new application.
func (c *client.Client) CreateApp(ctx context.Context, req *CreateAppRequest) (*AppDetail, error) {
	var result AppDetail
	err := c.sendRequest(ctx, "POST", "/console/api/apps", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetApp retrieves the details of a specific application.
func (c *client.Client) GetApp(ctx context.Context, appID string) (*AppDetail, error) {
	var result AppDetail
	path := fmt.Sprintf("/console/api/apps/%s", appID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateApp updates an application.
func (c *client.Client) UpdateApp(ctx context.Context, appID string, req *UpdateAppRequest) (*AppDetail, error) {
	var result AppDetail
	path := fmt.Sprintf("/console/api/apps/%s", appID)
	err := c.sendRequest(ctx, "PUT", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteApp deletes an application.
func (c *client.Client) DeleteApp(ctx context.Context, appID string) error {
	path := fmt.Sprintf("/console/api/apps/%s", appID)
	return c.sendRequest(ctx, "DELETE", path, nil, nil, nil)
}

// CopyApp copies an application.
func (c *client.Client) CopyApp(ctx context.Context, appID string, req *CreateAppRequest) (*AppDetail, error) {
	var result AppDetail
	path := fmt.Sprintf("/console/api/apps/%s/copy", appID)
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ExportApp exports an application as a YAML string.
func (c *client.Client) ExportApp(ctx context.Context, appID string, includeSecret bool) (string, error) {
	var result map[string]string
	path := fmt.Sprintf("/console/api/apps/%s/export?include_secret=%t", appID, includeSecret)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return "", err
	}
	return result["data"], nil
}

// UpdateAppName updates the name of an application.
func (c *client.Client) UpdateAppName(ctx context.Context, appID, name string) (*AppDetail, error) {
	var result AppDetail
	path := fmt.Sprintf("/console/api/apps/%s/name", appID)
	req := map[string]string{"name": name}
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateAppIcon updates the icon of an application.
func (c *client.Client) UpdateAppIcon(ctx context.Context, appID, icon, iconBackground string) (*AppDetail, error) {
	var result AppDetail
	path := fmt.Sprintf("/console/api/apps/%s/icon", appID)
	req := map[string]string{"icon": icon, "icon_background": iconBackground}
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateAppSiteStatus enables or disables the site for an application.
func (c *client.Client) UpdateAppSiteStatus(ctx context.Context, appID string, enable bool) (*AppDetail, error) {
	var result AppDetail
	path := fmt.Sprintf("/console/api/apps/%s/site-enable", appID)
	req := AppStatusRequest{Enable: enable}
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateAppAPIStatus enables or disables the API for an application.
func (c *client.Client) UpdateAppAPIStatus(ctx context.Context, appID string, enable bool) (*AppDetail, error) {
	var result AppDetail
	path := fmt.Sprintf("/console/api/apps/%s/api-enable", appID)
	req := AppStatusRequest{Enable: enable}
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAppTraceConfig retrieves the trace configuration for an app.
func (c *client.Client) GetAppTraceConfig(ctx context.Context, appID string) (map[string]interface{}, error) {
	var result map[string]interface{}
	path := fmt.Sprintf("/console/api/apps/%s/trace", appID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateAppTraceConfig updates the trace configuration for an app.
func (c *client.Client) UpdateAppTraceConfig(ctx context.Context, appID string, enabled bool, provider string) error {
	path := fmt.Sprintf("/console/api/apps/%s/trace", appID)
	req := map[string]interface{}{"enabled": enabled, "tracing_provider": provider}
	return c.sendRequest(ctx, "POST", path, req, nil, nil)
}
