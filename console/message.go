package console

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/types"
)

// MessageDetail represents the detailed information of a message.
type MessageDetail struct {
	ID                   string                 `json:"id"`
	ConversationID       string                 `json:"conversation_id"`
	Inputs               map[string]interface{} `json:"inputs"`
	Query                string                 `json:"query"`
	Answer               string                 `json:"answer"`
	MessageTokens        int                    `json:"message_tokens"`
	AnswerTokens         int                    `json:"answer_tokens"`
	CreatedAt            int64                  `json:"created_at"`
	// ... other fields
}

// MessageListResponse is the paginated response for listing messages.
type MessageListResponse struct {
	Data    []MessageDetail `json:"data"`
	HasMore bool            `json:"has_more"`
	Limit   int             `json:"limit"`
}

// CreateFeedbackRequest is the request to create feedback for a message.
type CreateFeedbackRequest struct {
	MessageID string `json:"message_id"`
	Rating    string `json:"rating"`
}

// MessageAnnotation represents an annotation created from a message.
type MessageAnnotation struct {
	ID        string `json:"id"`
	Question  string `json:"question"`
	Answer    string `json:"answer"`
	HitCount  int    `json:"hit_count"`
	CreatedAt int64  `json:"created_at"`
}

// CreateAnnotationRequest is the request to create an annotation for a message.
type CreateAnnotationRequest struct {
	MessageID string `json:"message_id,omitempty"`
	Question  string `json:"question"`
	Answer    string `json:"answer"`
}

// AnnotationCountResponse is the response for getting the annotation count.
type AnnotationCountResponse struct {
	Count int `json:"count"`
}

// SuggestedQuestionsResponse is the response for getting suggested questions.
type SuggestedQuestionsResponse struct {
	Data []string `json:"data"`
}


// GetChatMessages retrieves a list of chat messages for a conversation.
func (c *client.Client) GetChatMessages(ctx context.Context, appID, conversationID, firstID string, limit int) (*MessageListResponse, error) {
	var result MessageListResponse
	path := fmt.Sprintf("/console/api/apps/%s/chat-messages?conversation_id=%s&limit=%d", appID, conversationID, limit)
	if firstID != "" {
		path += fmt.Sprintf("&first_id=%s", firstID)
	}
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateMessageFeedback creates feedback for a message.
func (c *client.Client) CreateMessageFeedback(ctx context.Context, appID string, req *CreateFeedbackRequest) (*types.StopResponse, error) {
	var result types.StopResponse
	path := fmt.Sprintf("/console/api/apps/%s/feedbacks", appID)
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateMessageAnnotation creates an annotation from a message.
func (c *client.Client) CreateMessageAnnotation(ctx context.Context, appID string, req *CreateAnnotationRequest) (*MessageAnnotation, error) {
	var result MessageAnnotation
	path := fmt.Sprintf("/console/api/apps/%s/annotations", appID)
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAnnotationCount gets the total count of annotations for an app.
func (c *client.Client) GetAnnotationCount(ctx context.Context, appID string) (*AnnotationCountResponse, error) {
	var result AnnotationCountResponse
	path := fmt.Sprintf("/console/api/apps/%s/annotations/count", appID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMessageSuggestedQuestions retrieves suggested questions for a message.
func (c *client.Client) GetMessageSuggestedQuestions(ctx context.Context, appID, messageID string) (*SuggestedQuestionsResponse, error) {
	var result SuggestedQuestionsResponse
	path := fmt.Sprintf("/console/api/apps/%s/chat-messages/%s/suggested-questions", appID, messageID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMessage retrieves the details of a single message.
func (c *client.Client) GetMessage(ctx context.Context, appID, messageID string) (*MessageDetail, error) {
	var result MessageDetail
	path := fmt.Sprintf("/console/api/apps/%s/messages/%s", appID, messageID)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
