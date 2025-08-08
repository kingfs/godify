package service_api

import (
	"context"
	"fmt"

	"github.com/kingfs/godify/client"
)

// MessageFile represents a file associated with a message.
type MessageFile struct {
	ID             string `json:"id"`
	Filename       string `json:"filename"`
	Type           string `json:"type"`
	URL            string `json:"url"`
	MimeType       string `json:"mime_type"`
	Size           int    `json:"size"`
	TransferMethod string `json:"transfer_method"`
	BelongsTo      string `json:"belongs_to"`
	UploadFileID   string `json:"upload_file_id"`
}

// Feedback represents feedback on a message.
type Feedback struct {
	Rating string `json:"rating"`
	Content string `json:"content,omitempty"`
}

// AgentThought represents an agent's thought process.
type AgentThought struct {
	ID          string      `json:"id"`
	ChainID     string      `json:"chain_id"`
	MessageID   string      `json:"message_id"`
	Position    int         `json:"position"`
	Thought     string      `json:"thought"`
	Tool        string      `json:"tool"`
	ToolLabels  interface{} `json:"tool_labels"`
	ToolInput   string      `json:"tool_input"`
	CreatedAt   int64       `json:"created_at"`
	Observation string      `json:"observation"`
	Files       []string    `json:"files"`
}

// Message represents a single message in a conversation.
type Message struct {
	ID                 string                 `json:"id"`
	ConversationID     string                 `json:"conversation_id"`
	ParentMessageID    string                 `json:"parent_message_id"`
	Inputs             map[string]interface{} `json:"inputs"`
	Query              string                 `json:"query"`
	Answer             string                 `json:"answer"`
	MessageFiles       []MessageFile          `json:"message_files"`
	Feedback           *Feedback              `json:"feedback"`
	RetrieverResources []interface{}          `json:"retriever_resources"`
	CreatedAt          int64                  `json:"created_at"`
	AgentThoughts      []AgentThought         `json:"agent_thoughts"`
	Status             string                 `json:"status"`
	Error              string                 `json:"error"`
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

// SuggestedMessagesResponse is the response for getting suggested messages.
type SuggestedMessagesResponse struct {
	Result string   `json:"result"`
	Data   []string `json:"data"`
}

// AppFeedback represents feedback with user info.
type AppFeedback struct {
	Rating        string `json:"rating"`
	Content       string `json:"content"`
	FromSource    string `json:"from_source"`
	FromEndUserID string `json:"from_end_user_id"`
}

// AppFeedbacksResponse is the response for getting all app feedbacks.
type AppFeedbacksResponse struct {
	Data []AppFeedback `json:"data"`
}


// GetMessages retrieves a list of messages in a conversation.
func (c *client.Client) GetMessages(ctx context.Context, user, conversationID, firstID string, limit int) (*MessageListResponse, error) {
	var result MessageListResponse
	path := fmt.Sprintf("/v1/messages?user=%s&conversation_id=%s&limit=%d", user, conversationID, limit)
	if firstID != "" {
		path += fmt.Sprintf("&first_id=%s", firstID)
	}
	err := c.sendRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateMessageFeedback creates feedback for a message.
func (c *client.Client) CreateMessageFeedback(ctx context.Context, messageID string, req *CreateFeedbackRequest) (*FeedbackResponse, error) {
	var result FeedbackResponse
	path := fmt.Sprintf("/v1/messages/%s/feedbacks", messageID)
	err := c.sendRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetSuggestedMessages retrieves suggested questions for a message.
func (c *client.Client) GetSuggestedMessages(ctx context.Context, messageID, user string) (*SuggestedMessagesResponse, error) {
	var result SuggestedMessagesResponse
	path := fmt.Sprintf("/v1/messages/%s/suggested?user=%s", messageID, user)
	err := c.sendRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAppFeedbacks retrieves all feedbacks for an app.
func (c *client.Client) GetAppFeedbacks(ctx context.Context, page, limit int) (*AppFeedbacksResponse, error) {
	var result AppFeedbacksResponse
	path := fmt.Sprintf("/v1/app/feedbacks?page=%d&limit=%d", page, limit)
	err := c.sendRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
