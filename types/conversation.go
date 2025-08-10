package types

// SimpleConversation represents the basic details of a conversation.
type SimpleConversation struct {
	ID           string                 `json:"id"`
	Name         string                 `json:"name"`
	Inputs       map[string]interface{} `json:"inputs"`
	Status       string                 `json:"status"`
	Introduction string                 `json:"introduction"`
	CreatedAt    int64                  `json:"created_at"`
	UpdatedAt    int64                  `json:"updated_at"`
}

// ConversationListResponse is the response for listing conversations.
type ConversationListResponse struct {
	Limit   int                  `json:"limit"`
	HasMore bool                 `json:"has_more"`
	Data    []SimpleConversation `json:"data"`
}

// RenameConversationRequest is the request body for renaming a conversation.
type RenameConversationRequest struct {
	Name          string `json:"name,omitempty"`
	AutoGenerate  bool   `json:"auto_generate,omitempty"`
	User          string `json:"user"`
}

// ConversationDeleteResponse is the response for deleting a conversation.
type ConversationDeleteResponse struct {
	Result string `json:"result"`
}
