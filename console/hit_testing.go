package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// ConsoleHitTestingRequest is the request for a hit testing query in the console.
type ConsoleHitTestingRequest struct {
	Query          string                 `json:"query"`
	RetrievalModel map[string]interface{} `json:"retrieval_model,omitempty"`
}

// ConsoleHitTestingResponse is the response from a hit testing query in the console.
type ConsoleHitTestingResponse struct {
	Query   string             `json:"query"`
	Records []HitTestingRecord `json:"records"`
}

// ConsoleHitTest performs a hit testing query on a dataset from the console.
func (c *client.Client) ConsoleHitTest(ctx context.Context, datasetID string, req *ConsoleHitTestingRequest) (*ConsoleHitTestingResponse, error) {
	var result ConsoleHitTestingResponse
	path := fmt.Sprintf("/console/api/datasets/%s/hit-testing", datasetID)
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
