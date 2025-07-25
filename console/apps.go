package console

import (
	"context"
	"strconv"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/models"
)

// api.add_resource(CompletionMessageApi, "/apps/<uuid:app_id>/completion-messages")
// api.add_resource(CompletionMessageStopApi, "/apps/<uuid:app_id>/completion-messages/<string:task_id>/stop")
// api.add_resource(ChatMessageApi, "/apps/<uuid:app_id>/chat-messages")
// api.add_resource(ChatMessageStopApi, "/apps/<uuid:app_id>/chat-messages/<string:task_id>/stop")

// api.add_resource(MessageApi, "/apps/<uuid:app_id>/messages/<uuid:message_id>", endpoint="console_message")
// api.add_resource(ChatMessageListApi, "/apps/<uuid:app_id>/chat-messages", endpoint="console_chat_messages")

func (c *Client) GetAppsChatMessageList(ctx context.Context, appID string, conversationID string, firstID, limit *int) (*models.AppsChatMessageListApiResponse, error) {
	reqQuery := map[string]string{"conversation_id": conversationID}
	if firstID != nil {
		reqQuery["first_id"] = strconv.Itoa(*firstID)
	}
	if limit != nil {
		reqQuery["limit"] = strconv.Itoa(*limit)
	}
	req := &client.Request{
		Method: "GET",
		Path:   "/apps/" + appID + "/chat-messages",
		Query:  reqQuery,
	}
	var resp models.AppsChatMessageListApiResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

func (c *Client) GetAppsMessage(ctx context.Context, appID string, messageID string) (*models.AppsMessageApiResponse, error) {
	req := &client.Request{
		Method: "GET",
		Path:   "/apps/" + appID + "/messages/" + messageID,
	}
	var resp models.AppsMessageApiResponse
	err := c.baseClient.DoJSON(ctx, req, &resp)
	return &resp, err
}

