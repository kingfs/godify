package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// DailyMessagesStat represents a single day's message statistics.
type DailyMessagesStat struct {
	Date         string `json:"date"`
	MessageCount int    `json:"message_count"`
}

// DailyConversationsStat represents a single day's conversation statistics.
type DailyConversationsStat struct {
	Date             string `json:"date"`
	ConversationCount int    `json:"conversation_count"`
}

// DailyEndUsersStat represents a single day's end user statistics.
type DailyEndUsersStat struct {
	Date          string `json:"date"`
	TerminalCount int    `json:"terminal_count"`
}

// DailyTokenCostStat represents a single day's token cost statistics.
type DailyTokenCostStat struct {
	Date       string  `json:"date"`
	TokenCount int     `json:"token_count"`
	TotalPrice float64 `json:"total_price"`
	Currency   string  `json:"currency"`
}

// AverageSessionInteractionStat represents a single day's session interaction statistics.
type AverageSessionInteractionStat struct {
	Date         string  `json:"date"`
	Interactions float64 `json:"interactions"`
}

// UserSatisfactionRateStat represents a single day's user satisfaction rate.
type UserSatisfactionRateStat struct {
	Date string  `json:"date"`
	Rate float64 `json:"rate"`
}

// AverageResponseTimeStat represents a single day's average response time.
type AverageResponseTimeStat struct {
	Date    string  `json:"date"`
	Latency float64 `json:"latency"`
}

// TokensPerSecondStat represents a single day's tokens per second.
type TokensPerSecondStat struct {
	Date string  `json:"date"`
	TPS  float64 `json:"tps"`
}

// StatisticResponse is a generic wrapper for all statistic responses.
type StatisticResponse[T any] struct {
	Data []T `json:"data"`
}

func getStatistic[T any](c *client.Client, ctx context.Context, appID, endpoint, start, end string) (*StatisticResponse[T], error) {
	var result StatisticResponse[T]
	path := fmt.Sprintf("/console/api/apps/%s/statistics/%s?start=%s&end=%s", appID, endpoint, start, end)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDailyMessages retrieves the daily message statistics for an app.
func (c *client.Client) GetDailyMessages(ctx context.Context, appID, start, end string) (*StatisticResponse[DailyMessagesStat], error) {
	return getStatistic[DailyMessagesStat](c, ctx, appID, "daily-messages", start, end)
}

// GetDailyConversations retrieves the daily conversation statistics for an app.
func (c *client.Client) GetDailyConversations(ctx context.Context, appID, start, end string) (*StatisticResponse[DailyConversationsStat], error) {
	return getStatistic[DailyConversationsStat](c, ctx, appID, "daily-conversations", start, end)
}

// GetDailyEndUsers retrieves the daily end user statistics for an app.
func (c *client.Client) GetDailyEndUsers(ctx context.Context, appID, start, end string) (*StatisticResponse[DailyEndUsersStat], error) {
	return getStatistic[DailyEndUsersStat](c, ctx, appID, "daily-end-users", start, end)
}

// GetDailyTokenCosts retrieves the daily token cost statistics for an app.
func (c *client.Client) GetDailyTokenCosts(ctx context.Context, appID, start, end string) (*StatisticResponse[DailyTokenCostStat], error) {
	return getStatistic[DailyTokenCostStat](c, ctx, appID, "token-costs", start, end)
}

// GetAverageSessionInteractions retrieves the average session interaction statistics for an app.
func (c *client.Client) GetAverageSessionInteractions(ctx context.Context, appID, start, end string) (*StatisticResponse[AverageSessionInteractionStat], error) {
	return getStatistic[AverageSessionInteractionStat](c, ctx, appID, "average-session-interactions", start, end)
}

// GetUserSatisfactionRate retrieves the user satisfaction rate statistics for an app.
func (c *client.Client) GetUserSatisfactionRate(ctx context.Context, appID, start, end string) (*StatisticResponse[UserSatisfactionRateStat], error) {
	return getStatistic[UserSatisfactionRateStat](c, ctx, appID, "user-satisfaction-rate", start, end)
}

// GetAverageResponseTime retrieves the average response time statistics for an app.
func (c *client.Client) GetAverageResponseTime(ctx context.Context, appID, start, end string) (*StatisticResponse[AverageResponseTimeStat], error) {
	return getStatistic[AverageResponseTimeStat](c, ctx, appID, "average-response-time", start, end)
}

// GetTokensPerSecond retrieves the tokens per second statistics for an app.
func (c *client.Client) GetTokensPerSecond(ctx context.Context, appID, start, end string) (*StatisticResponse[TokensPerSecondStat], error) {
	return getStatistic[TokensPerSecondStat](c, ctx, appID, "tokens-per-second", start, end)
}
