package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/types"
)

// InstalledAppGetMessages retrieves a list of messages for an installed app's conversation.
func (c *client.Client) InstalledAppGetMessages(ctx context.Context, installedAppID, conversationID, firstID string, limit int) (*MessageListResponse, error) {
	var result MessageListResponse
	path := fmt.Sprintf("/console/api/installed-apps/%s/messages?conversation_id=%s&limit=%d", installedAppID, conversationID, limit)
	if firstID != "" {
		path += fmt.Sprintf("&first_id=%s", firstID)
	}
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// InstalledAppCreateMessageFeedback creates feedback for a message in an installed app.
func (c *client.Client) InstalledAppCreateMessageFeedback(ctx context.Context, installedAppID, messageID string, req *CreateFeedbackRequest) (*types.StopResponse, error) {
	var result types.StopResponse
	path := fmt.Sprintf("/console/api/installed-apps/%s/messages/%s/feedbacks", installedAppID, messageID)
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// InstalledAppGetMoreLikeThis generates a "more like this" response for a message in an installed app.
func (c *client.Client) InstalledAppGetMoreLikeThis(ctx context.Context, installedAppID, messageID string) (*types.BlockingResponse, error) {
	var result types.BlockingResponse
	path := fmt.Sprintf("/console/api/installed-apps/%s/messages/%s/more-like-this?response_mode=blocking", installedAppID, messageID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// InstalledAppGetMoreLikeThisStream generates a "more like this" response in streaming mode.
func (c *client.Client) InstalledAppGetMoreLikeThisStream(ctx context.Context, installedAppID, messageID string) (<-chan *types.StreamEvent, error) {
	path := fmt.Sprintf("/console/api/installed-apps/%s/messages/%s/more-like-this?response_mode=streaming", installedAppID, messageID)
	return c.installedAppHandleStream(ctx, "GET", path, nil)
}

// InstalledAppGetSuggestedQuestions retrieves suggested questions for a message in an installed app.
func (c *client.Client) InstalledAppGetSuggestedQuestions(ctx context.Context, installedAppID, messageID string) (*SuggestedQuestionsResponse, error) {
	var result SuggestedQuestionsResponse
	path := fmt.Sprintf("/console/api/installed-apps/%s/messages/%s/suggested-questions", installedAppID, messageID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
