package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// DailyRunsStat represents a single day's workflow run statistics.
type DailyRunsStat struct {
	Date string `json:"date"`
	Runs int    `json:"runs"`
}

// DailyTerminalsStat represents a single day's workflow terminal statistics.
type DailyTerminalsStat struct {
	Date          string `json:"date"`
	TerminalCount int    `json:"terminal_count"`
}

// DailyWorkflowTokenCostStat represents a single day's workflow token cost statistics.
type DailyWorkflowTokenCostStat struct {
	Date       string `json:"date"`
	TokenCount int    `json:"token_count"`
}

// AverageAppInteractionStat represents a single day's average app interaction statistics.
type AverageAppInteractionStat struct {
	Date         string  `json:"date"`
	Interactions float64 `json:"interactions"`
}

func getWorkflowStatistic[T any](c *client.Client, ctx context.Context, appID, endpoint, start, end string) (*StatisticResponse[T], error) {
	var result StatisticResponse[T]
	path := fmt.Sprintf("/console/api/apps/%s/workflow/statistics/%s?start=%s&end=%s", appID, endpoint, start, end)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetWorkflowDailyRuns retrieves the daily run statistics for a workflow app.
func (c *client.Client) GetWorkflowDailyRuns(ctx context.Context, appID, start, end string) (*StatisticResponse[DailyRunsStat], error) {
	return getWorkflowStatistic[DailyRunsStat](c, ctx, appID, "daily-conversations", start, end)
}

// GetWorkflowDailyTerminals retrieves the daily terminal statistics for a workflow app.
func (c *client.Client) GetWorkflowDailyTerminals(ctx context.Context, appID, start, end string) (*StatisticResponse[DailyTerminalsStat], error) {
	return getWorkflowStatistic[DailyTerminalsStat](c, ctx, appID, "daily-terminals", start, end)
}

// GetWorkflowDailyTokenCosts retrieves the daily token cost statistics for a workflow app.
func (c *client.Client) GetWorkflowDailyTokenCosts(ctx context.Context, appID, start, end string) (*StatisticResponse[DailyWorkflowTokenCostStat], error) {
	return getWorkflowStatistic[DailyWorkflowTokenCostStat](c, ctx, appID, "token-costs", start, end)
}

// GetWorkflowAverageAppInteractions retrieves the average app interaction statistics for a workflow app.
func (c *client.Client) GetWorkflowAverageAppInteractions(ctx context.Context, appID, start, end string) (*StatisticResponse[AverageAppInteractionStat], error) {
	return getWorkflowStatistic[AverageAppInteractionStat](c, ctx, appID, "average-app-interactions", start, end)
}
