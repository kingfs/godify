package web

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kingfs/godify/models"
)

func TestNewWebClient(t *testing.T) {
	client := NewClient("https://api.example.com")
	client.WithAppCode("test-app-code")
	if client == nil {
		t.Fatal("Expected client to be created")
	}
}

func TestGetPassport(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{"access_token": "test-token"}`))
		if err != nil {
			t.Errorf("Failed to write response: %v", err)
		}
	}))
	defer server.Close()

	client := NewClient(server.URL)
	client.WithAppCode("test-app-code")
	token, err := client.GetPassport(context.Background(), "")
	if err != nil {
		t.Fatalf("GetPassport failed: %v", err)
	}

	if token != "test-token" {
		t.Errorf("Expected token 'test-token', got %s", token)
	}
}

func TestGetWebAppAccessMode(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/webapp/access-mode" {
			t.Errorf("Expected path /api/webapp/access-mode, got %s", r.URL.Path)
		}

		appID := r.URL.Query().Get("appId")
		if appID != "test-app-id" {
			t.Errorf("Expected appId 'test-app-id', got %s", appID)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{"accessMode": "public"}`))
		if err != nil {
			t.Errorf("Failed to write response: %v", err)
		}
	}))
	defer server.Close()

	client := NewClient(server.URL)
	client.WithAppCode("test-app-code")
	accessMode, err := client.GetWebAppAccessMode(context.Background(), "test-app-id", "")

	if err != nil {
		t.Fatalf("GetWebAppAccessMode failed: %v", err)
	}

	if accessMode.AccessMode != models.AccessModePublic {
		t.Errorf("Expected access mode 'public', got %s", accessMode.AccessMode)
	}
}

func TestGetConversations(t *testing.T) {
	passportCalled := false
	conversationsCalled := false

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.URL.Path == "/api/passport" {
			passportCalled = true
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(`{"access_token": "test-token"}`))
			if err != nil {
				t.Errorf("Failed to write response: %v", err)
			}
			return
		}

		if r.URL.Path == "/api/conversations" {
			conversationsCalled = true
			limit := r.URL.Query().Get("limit")
			if limit != "20" {
				t.Errorf("Expected limit '20', got %s", limit)
			}

			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(`{
				"limit": 20,
				"has_more": false,
				"data": [
					{
						"id": "conv-1",
						"name": "Test Conversation",
						"inputs": {},
						"introduction": "Test intro",
						"created_at": "2023-01-01T00:00:00Z",
						"updated_at": "2023-01-01T00:00:00Z"
					}
				]
			}`))
			if err != nil {
				t.Errorf("Failed to write response: %v", err)
			}
			return
		}

		t.Errorf("Unexpected path: %s", r.URL.Path)
	}))
	defer server.Close()

	client := NewClient(server.URL)
	client.WithAppCode("test-app-code")
	resp, err := client.GetConversations(context.Background(), "", 20, nil, "")

	if err != nil {
		t.Fatalf("GetConversations failed: %v", err)
	}

	if !passportCalled {
		t.Error("Expected passport to be called")
	}

	if !conversationsCalled {
		t.Error("Expected conversations to be called")
	}

	if len(resp.Data) != 1 {
		t.Errorf("Expected 1 conversation, got %d", len(resp.Data))
	}

	if resp.Data[0].Name != "Test Conversation" {
		t.Errorf("Expected conversation name 'Test Conversation', got %s", resp.Data[0].Name)
	}
}

func TestWebChat(t *testing.T) {
	passportCalled := false
	chatCalled := false

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.URL.Path == "/api/passport" {
			passportCalled = true
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(`{"access_token": "test-token"}`))
			if err != nil {
				t.Errorf("Failed to write response: %v", err)
			}
			return
		}

		if r.URL.Path == "/api/chat-messages" {
			chatCalled = true
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(`{
				"event": "message",
				"message_id": "web-message-id",
				"conversation_id": "web-conversation-id",
				"answer": "Hello from web API!",
				"created_at": 1234567890
			}`))
			if err != nil {
				t.Errorf("Failed to write response: %v", err)
			}
			return
		}

		t.Errorf("Unexpected path: %s", r.URL.Path)
	}))
	defer server.Close()

	client := NewClient(server.URL)
	client.WithAppCode("test-app-code")
	req := &ChatRequest{
		Inputs: map[string]interface{}{},
		Query:  "Hello",
	}

	resp, err := client.Chat(context.Background(), req)
	if err != nil {
		t.Fatalf("Chat failed: %v", err)
	}

	if !passportCalled {
		t.Error("Expected passport to be called")
	}

	if !chatCalled {
		t.Error("Expected chat to be called")
	}

	if resp.Answer != "Hello from web API!" {
		t.Errorf("Expected answer 'Hello from web API!', got %s", resp.Answer)
	}
}
