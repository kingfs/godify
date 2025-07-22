package service

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kingfs/godify/models"
)

func TestNewClient(t *testing.T) {
	client := NewClient("test-token", "https://api.example.com")
	if client == nil {
		t.Fatal("Expected client to be created")
	}
}

func TestGetAppInfo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/info" {
			t.Errorf("Expected path /v1/info, got %s", r.URL.Path)
		}

		if r.Header.Get("Authorization") != "Bearer test-token" {
			t.Errorf("Expected Authorization header 'Bearer test-token', got %s", r.Header.Get("Authorization"))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{
			"id": "test-app-id",
			"name": "Test App",
			"description": "Test Description", 
			"tags": ["tag1", "tag2"],
			"mode": "chat",
			"author_name": "Test Author"
		}`))
		if err != nil {
			t.Errorf("Failed to write response: %v", err)
		}
	}))
	defer server.Close()

	client := NewClient("test-token", server.URL)
	appInfo, err := client.GetAppInfo(context.Background())

	if err != nil {
		t.Fatalf("GetAppInfo failed: %v", err)
	}

	if appInfo.Name != "Test App" {
		t.Errorf("Expected app name 'Test App', got %s", appInfo.Name)
	}

	if appInfo.Mode != models.AppModeChat {
		t.Errorf("Expected app mode 'chat', got %s", appInfo.Mode)
	}
}

func TestChat(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/chat-messages" {
			t.Errorf("Expected path /v1/chat-messages, got %s", r.URL.Path)
		}

		if r.Method != "POST" {
			t.Errorf("Expected POST method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{
			"event": "message",
			"message_id": "test-message-id",
			"conversation_id": "test-conversation-id",
			"answer": "Hello! How can I help you?",
			"created_at": 1234567890
		}`))
		if err != nil {
			t.Errorf("Failed to write response: %v", err)
		}
	}))
	defer server.Close()

	client := NewClient("test-token", server.URL)
	req := &ChatRequest{
		Inputs: map[string]interface{}{},
		Query:  "Hello",
		User:   "test-user",
	}

	resp, err := client.Chat(context.Background(), req)
	if err != nil {
		t.Fatalf("Chat failed: %v", err)
	}

	if resp.Answer != "Hello! How can I help you?" {
		t.Errorf("Expected answer 'Hello! How can I help you?', got %s", resp.Answer)
	}

	if resp.MessageID != "test-message-id" {
		t.Errorf("Expected message ID 'test-message-id', got %s", resp.MessageID)
	}
}

func TestCompletion(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/completion-messages" {
			t.Errorf("Expected path /v1/completion-messages, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{
			"event": "message",
			"message_id": "test-message-id",
			"answer": "This is a completion response.",
			"created_at": 1234567890
		}`))
		if err != nil {
			t.Errorf("Failed to write response: %v", err)
		}
	}))
	defer server.Close()

	client := NewClient("test-token", server.URL)
	req := &CompletionRequest{
		Inputs: map[string]interface{}{
			"query": "Complete this text",
		},
		Query: "Complete this text",
		User:  "test-user",
	}

	resp, err := client.Completion(context.Background(), req)
	if err != nil {
		t.Fatalf("Completion failed: %v", err)
	}

	if resp.Answer != "This is a completion response." {
		t.Errorf("Expected answer 'This is a completion response.', got %s", resp.Answer)
	}
}
