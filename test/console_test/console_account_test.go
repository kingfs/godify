package console_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kingfs/godify/models"
)

var account_use_real_url = true

func SetupAccountMockServer() *httptest.Server {
	handler := http.NewServeMux()

	handler.HandleFunc("/console/api/setup", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "success"}`))
	})

	return httptest.NewServer(handler)
}

func TestSetupAccount(t *testing.T) {
	// 无需admin_api_key
	server := SetupAccountMockServer()
	defer server.Close()

	client := TestNewClientWithBaseURL(server.URL, account_use_real_url)

	resp, err := client.SetupAccount(context.Background(), "test@test.com", "test", "test123456")
	if err != nil {
		t.Fatalf("SetupAccount failed: %v", err)
	}

	t.Logf("SetupAccount response: %v", resp)
}

func TestLogin(t *testing.T) {
	server := SetupAccountMockServer()
	defer server.Close()

	client := TestNewClientWithBaseURL(server.URL, account_use_real_url)

	resp, err := client.Login(context.Background(), &models.LoginRequest{
		Email:    "test@test.com",
		Password: "test123456",
	})
	if err != nil {
		t.Fatalf("Login failed: %v", err)
	}
	t.Logf("Login response: %#v", resp)
}

func TestGetAccountProfile(t *testing.T) {
	server := SetupAccountMockServer()
	defer server.Close()

	client := TestNewClientWithBaseURL(server.URL, account_use_real_url)

	resp, err := client.GetAccountProfile(context.Background(), "test@test.com", "test123456")
	if err != nil {
		t.Fatalf("GetAccountProfile failed: %v", err)
	}
	t.Logf("GetAccountProfile response: %#v", resp)
}
