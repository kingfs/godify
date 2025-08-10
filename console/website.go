package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// WebsiteCrawlRequest is the request to crawl a website.
type WebsiteCrawlRequest struct {
	Provider string                 `json:"provider"`
	URL      string                 `json:"url"`
	Options  map[string]interface{} `json:"options"`
}

// WebsiteCrawlResponse is the response from a website crawl request.
type WebsiteCrawlResponse struct {
	JobID string `json:"job_id"`
}

// WebsiteCrawlStatusResponse is the response for getting the status of a crawl job.
type WebsiteCrawlStatusResponse struct {
	Status   string      `json:"status"`
	Data     interface{} `json:"data,omitempty"`
	Error    string      `json:"error,omitempty"`
	Progress float64     `json:"progress,omitempty"`
}

// CrawlWebsite starts a website crawl job.
func (c *client.Client) CrawlWebsite(ctx context.Context, req *WebsiteCrawlRequest) (*WebsiteCrawlResponse, error) {
	var result WebsiteCrawlResponse
	err := c.sendRequest(ctx, "POST", "/console/api/website/crawl", req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetWebsiteCrawlStatus gets the status of a website crawl job.
func (c *client.Client) GetWebsiteCrawlStatus(ctx context.Context, jobID, provider string) (*WebsiteCrawlStatusResponse, error) {
	var result WebsiteCrawlStatusResponse
	path := fmt.Sprintf("/console/api/website/crawl/status/%s?provider=%s", jobID, provider)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
