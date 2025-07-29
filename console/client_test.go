package console

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kingfs/godify/models"
)

func TestNewConsoleClient(t *testing.T) {
	client := NewClient("test-token", "https://api.example.com")
	if client == nil {
		t.Fatal("Expected client to be created")
	}
}

func TestGetApps(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/console/api/apps" {
			t.Errorf("Expected path /console/api/apps, got %s", r.URL.Path)
		}

		page := r.URL.Query().Get("page")
		if page != "1" {
			t.Errorf("Expected page '1', got %s", page)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{
			"data": [
				{
					"id": "app-1",
					"name": "Test App",
					"description": "Test Description",
					"mode": "chat",
					"enable_site": true,
					"enable_api": true,
					"created_at": "2023-01-01T00:00:00Z",
					"updated_at": "2023-01-01T00:00:00Z",
					"tenant_id": "tenant-1",
					"tags": []
				}
			],
			"has_more": false,
			"limit": 20,
			"total": 1,
			"page": 1
		}`))
		if err != nil {
			t.Errorf("Failed to write response: %v", err)
		}
	}))
	defer server.Close()

	client := NewClient("test-token", server.URL)
	apps, err := client.GetApps(context.Background(), 1, 20, "", "", nil, nil)
	if err != nil {
		t.Fatalf("GetApps failed: %v", err)
	}

	if len(apps.Data) != 1 {
		t.Errorf("Expected 1 app, got %d", len(apps.Data))
	}

	if apps.Data[0].Name != "Test App" {
		t.Errorf("Expected app name 'Test App', got %s", apps.Data[0].Name)
	}

	if apps.Data[0].Mode != models.AppModeChat {
		t.Errorf("Expected app mode 'chat', got %s", apps.Data[0].Mode)
	}
}

func TestCreateApp(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/console/api/apps" {
			t.Errorf("Expected path /console/api/apps, got %s", r.URL.Path)
		}

		if r.Method != "POST" {
			t.Errorf("Expected POST method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, err := w.Write([]byte(`{
			"id": "new-app-id",
			"name": "New Test App",
			"description": "New Test Description",
			"mode": "completion",
			"enable_site": false,
			"enable_api": false,
			"created_at": "2023-01-01T00:00:00Z",
			"updated_at": "2023-01-01T00:00:00Z",
			"tenant_id": "tenant-1",
			"tags": []
		}`))
		if err != nil {
			t.Errorf("Failed to write response: %v", err)
		}
	}))
	defer server.Close()

	client := NewClient("test-token", server.URL)
	req := &models.CreateAppRequest{
		Name:        "New Test App",
		Description: "New Test Description",
		Mode:        models.AppModeCompletion,
	}

	app, err := client.CreateApp(context.Background(), req)
	if err != nil {
		t.Fatalf("CreateApp failed: %v", err)
	}

	if app.Name != "New Test App" {
		t.Errorf("Expected app name 'New Test App', got %s", app.Name)
	}

	if app.Mode != models.AppModeCompletion {
		t.Errorf("Expected app mode 'completion', got %s", app.Mode)
	}
}

func TestGetDatasets(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/console/api/datasets" {
			t.Errorf("Expected path /console/api/datasets, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{
			"data": [
				{
					"id": "dataset-1",
					"name": "Test Dataset",
					"description": "Test Description",
					"permission": "only_me",
					"data_source_type": "upload_file",
					"indexing_technique": "high_quality",
					"created_by": "user-1",
					"created_at": "2023-01-01T00:00:00Z",
					"updated_at": "2023-01-01T00:00:00Z",
					"document_count": 5,
					"word_count": 1000,
					"app_count": 2,
					"embedding_model": "text-embedding-ada-002",
					"embedding_model_provider": "openai",
					"embedding_available": true,
					"tags": [],
					"partial_member_list": []
				}
			],
			"has_more": false,
			"limit": 20,
			"total": 1,
			"page": 1
		}`))
		if err != nil {
			t.Errorf("Failed to write response: %v", err)
		}
	}))
	defer server.Close()

	client := NewClient("test-token", server.URL)
	datasets, err := client.GetDatasets(context.Background(), 1, 20, "", nil, false)
	if err != nil {
		t.Fatalf("GetDatasets failed: %v", err)
	}

	if len(datasets.Data) != 1 {
		t.Errorf("Expected 1 dataset, got %d", len(datasets.Data))
	}

	if datasets.Data[0].Name != "Test Dataset" {
		t.Errorf("Expected dataset name 'Test Dataset', got %s", datasets.Data[0].Name)
	}

	if datasets.Data[0].DocumentCount != 5 {
		t.Errorf("Expected document count 5, got %d", datasets.Data[0].DocumentCount)
	}
}
