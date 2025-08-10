package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// GetSubscription retrieves the subscription payment link.
// The response is a dictionary with details about the subscription link, so we use map[string]interface{}.
func (c *client.Client) GetSubscription(ctx context.Context, plan, interval string) (map[string]interface{}, error) {
	var result map[string]interface{}
	path := fmt.Sprintf("/console/api/billing/subscription?plan=%s&interval=%s", plan, interval)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetInvoices retrieves the list of invoices for the current tenant.
// The response is a list of invoice objects, so we use []map[string]interface{}.
func (c *client.Client) GetInvoices(ctx context.Context) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := c.sendRequest(ctx, "GET", "/console/api/billing/invoices", nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}
