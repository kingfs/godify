package web

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// Note: The following structs are duplicates from the service_api package.
// They should be refactored into a common types package.

// Message represents a single message in a conversation.
type Message struct {
	ID                 string                 `json:"id"`
	ConversationID     string                 `json:"conversation_id"`
	ParentMessageID    string                 `json:"parent_message_id"`
	Inputs             map[string]interface{} `json:"inputs"`
	Query              string                 `json:"query"`
	Answer             string                 `json:"answer"`
	MessageFiles       []interface{}          `json:"message_files"` // Simplified for now
	Feedback           *Feedback              `json:"feedback"`
	RetrieverResources []interface{}          `json:"retriever_resources"`
	CreatedAt          int64                  `json:"created_at"`
	AgentThoughts      []interface{}          `json:"agent_thoughts"` // Simplified for now
	Metadata           map[string]interface{} `json:"metadata"`
	Status             string                 `json:"status"`
	Error              string                 `json:"error"`
}

// Feedback represents feedback on a message.
type Feedback struct {
	Rating  string `json:"rating"`
	Content string `json:"content,omitempty"`
}

// MessageListResponse is the response for listing messages.
type MessageListResponse struct {
	Limit   int       `json:"limit"`
	HasMore bool      `json:"has_more"`
	Data    []Message `json:"data"`
}

// CreateFeedbackRequest is the request body for creating feedback.
type CreateFeedbackRequest struct {
	Rating  string `json:"rating"`
	Content string `json:"content,omitempty"`
	User    string `json:"user"`
}

// FeedbackResponse is the response for creating feedback.
type FeedbackResponse struct {
	Result string `json:"result"`
}

// SuggestedQuestionsResponse is the response for getting suggested questions.
type SuggestedQuestionsResponse struct {
	Data []string `json:"data"`
}

// WebGetMessages retrieves a list of messages in a conversation.
func (c *client.Client) WebGetMessages(ctx context.Context, user, conversationID, firstID string, limit int) (*MessageListResponse, error) {
	var result MessageListResponse
	path := fmt.Sprintf("/api/messages?user=%s&conversation_id=%s&limit=%d", user, conversationID, limit)
	if firstID != "" {
		path += fmt.Sprintf("&first_id=%s", firstID)
	}
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// WebCreateMessageFeedback creates feedback for a message.
func (c *client.Client) WebCreateMessageFeedback(ctx context.Context, messageID string, req *CreateFeedbackRequest) (*FeedbackResponse, error) {
	var result FeedbackResponse
	path := fmt.Sprintf("/api/messages/%s/feedbacks", messageID)
	err := c.sendRequest(ctx, "POST", path, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMoreLikeThis generates a "more like this" response for a message in blocking mode.
func (c *client.Client) GetMoreLikeThis(ctx context.Context, messageID, user string) (*BlockingResponse, error) {
	var result BlockingResponse
	path := fmt.Sprintf("/api/messages/%s/more-like-this?response_mode=blocking&user=%s", messageID, user)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMoreLikeThisStream generates a "more like this" response for a message in streaming mode.
func (c *client.Client) GetMoreLikeThisStream(ctx context.Context, messageID, user string) (<-chan *StreamEvent, error) {
	path := fmt.Sprintf("/api/messages/%s/more-like-this?response_mode=streaming&user=%s", messageID, user)
	return c.webHandleStream(ctx, "GET", path, nil)
}

// GetSuggestedQuestions retrieves suggested questions for a message.
func (c *client.Client) GetSuggestedQuestions(ctx context.Context, messageID, user string) (*SuggestedQuestionsResponse, error) {
	var result SuggestedQuestionsResponse
	path := fmt.Sprintf("/api/messages/%s/suggested-questions?user=%s", messageID, user)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
