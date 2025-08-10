package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// DataSourceIntegrate represents a data source integration.
type DataSourceIntegrate struct {
	ID         string                 `json:"id"`
	Provider   string                 `json:"provider"`
	CreatedAt  int64                  `json:"created_at"`
	IsBound    bool                   `json:"is_bound"`
	Disabled   bool                   `json:"disabled"`
	Link       string                 `json:"link"`
	SourceInfo map[string]interface{} `json:"source_info"`
}

// DataSourceIntegrateListResponse is the response for listing data source integrations.
type DataSourceIntegrateListResponse struct {
	Data []DataSourceIntegrate `json:"data"`
}

// NotionPage represents a Notion page for import.
type NotionPage struct {
	PageName string      `json:"page_name"`
	PageID   string      `json:"page_id"`
	PageIcon interface{} `json:"page_icon"`
	IsBound  bool        `json:"is_bound"`
	ParentID string      `json:"parent_id"`
	Type     string      `json:"type"`
}

// NotionWorkspace represents a Notion workspace with its pages.
type NotionWorkspace struct {
	WorkspaceName string       `json:"workspace_name"`
	WorkspaceID   string       `json:"workspace_id"`
	WorkspaceIcon string       `json:"workspace_icon"`
	Pages         []NotionPage `json:"pages"`
}

// NotionInfoListResponse is the response for listing Notion pages for import.
type NotionInfoListResponse struct {
	NotionInfo []NotionWorkspace `json:"notion_info"`
}

// NotionIndexingEstimateRequest is the request to estimate indexing for Notion pages.
type NotionIndexingEstimateRequest struct {
	NotionInfoList []interface{} `json:"notion_info_list"`
	ProcessRule    interface{}   `json:"process_rule"`
	DocForm        string        `json:"doc_form,omitempty"`
	DocLanguage    string        `json:"doc_language,omitempty"`
}

// GetDataSources retrieves the list of data source integrations.
func (c *client.Client) GetDataSources(ctx context.Context) (*DataSourceIntegrateListResponse, error) {
	var result DataSourceIntegrateListResponse
	err := c.sendRequest(ctx, "GET", "/console/api/data-source/integrates", nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ToggleDataSource enables or disables a data source integration.
func (c *client.Client) ToggleDataSource(ctx context.Context, bindingID, action string) error {
	path := fmt.Sprintf("/console/api/data-source/integrates/%s/%s", bindingID, action)
	return c.sendRequest(ctx, "PATCH", path, nil, nil, nil)
}

// GetNotionPreImportPages retrieves a list of Notion pages available for import.
func (c *client.Client) GetNotionPreImportPages(ctx context.Context, datasetID string) (*NotionInfoListResponse, error) {
	var result NotionInfoListResponse
	path := "/console/api/notion/pre-import/pages"
	if datasetID != "" {
		path += fmt.Sprintf("?dataset_id=%s", datasetID)
	}
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetNotionPagePreview retrieves a preview of a Notion page.
func (c *client.Client) GetNotionPagePreview(ctx context.Context, workspaceID, pageID, pageType string) (string, error) {
	var result struct {
		Content string `json:"content"`
	}
	path := fmt.Sprintf("/console/api/notion/workspaces/%s/pages/%s/%s/preview", workspaceID, pageID, pageType)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return "", err
	}
	return result.Content, nil
}

// EstimateNotionIndexing estimates the indexing cost for a set of Notion pages.
func (c *client.Client) EstimateNotionIndexing(ctx context.Context, req *NotionIndexingEstimateRequest) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := c.sendRequest(ctx, "POST", "/console/api/datasets/notion-indexing-estimate", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SyncNotionDataset triggers a sync for all documents in a Notion-based dataset.
func (c *client.Client) SyncNotionDataset(ctx context.Context, datasetID string) error {
	path := fmt.Sprintf("/console/api/datasets/%s/notion/sync", datasetID)
	return c.sendRequest(ctx, "GET", path, nil, nil, nil)
}

// SyncNotionDocument triggers a sync for a single document in a Notion-based dataset.
func (c *client.Client) SyncNotionDocument(ctx context.Context, datasetID, documentID string) error {
	path := fmt.Sprintf("/console/api/datasets/%s/documents/%s/notion/sync", datasetID, documentID)
	return c.sendRequest(ctx, "GET", path, nil, nil, nil)
}
