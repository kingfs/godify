package console_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

var mcp_use_real_url = true

// mock server，路径和返回内容严格参考workspaces.go接口实现
func SetupMCPMockServer() *httptest.Server {
	handler := http.NewServeMux()

	return httptest.NewServer(handler)
}

func TestCreateMCPProvider(t *testing.T) {
	mockServer := SetupMCPMockServer()
	defer mockServer.Close()

	client := TestNewClientWithBaseURL(mockServer.URL, mcp_use_real_url)

	resp, err := client.CreateMCPProvider(context.Background(), "http://host.docker.internal:8909/mcp", "mcp", "🧿", "emoji", "#EFF1F5", "mcp")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp: %+v", resp)
}

func TestUpdateMCPProvider(t *testing.T) {
	mockServer := SetupMCPMockServer()
	defer mockServer.Close()

	client := TestNewClientWithBaseURL(mockServer.URL, mcp_use_real_url)

	resp, err := client.UpdateMCPProvider(context.Background(), "b17ecd18-274c-49ec-8d3d-3122e2ecdf82", "http://host.docker.internal:8909/mcp", "mcp1", "🧿", "emoji", "#EFF1F5", "mcp1")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp: %+v", resp)
}

func TestGetMCPProviderDetail(t *testing.T) {
	mockServer := SetupMCPMockServer()
	defer mockServer.Close()

	client := TestNewClientWithBaseURL(mockServer.URL, mcp_use_real_url)

	resp, err := client.GetMCPProviderDetail(context.Background(), "b17ecd18-274c-49ec-8d3d-3122e2ecdf82")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp: %+v", resp)
}

func TestAuthMCPProvider(t *testing.T) {
	mockServer := SetupMCPMockServer()
	defer mockServer.Close()

	client := TestNewClientWithBaseURL(mockServer.URL, mcp_use_real_url)

	resp, err := client.AuthMCPProvider(context.Background(), "13e3a8f0-47cd-4284-a503-4154086dc05a", "123456")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp: %+v", resp)
}

func TestDeleteMCPProvider(t *testing.T) {
	mockServer := SetupMCPMockServer()
	defer mockServer.Close()

	client := TestNewClientWithBaseURL(mockServer.URL, mcp_use_real_url)

	resp, err := client.DeleteMCPProvider(context.Background(), "13e3a8f0-47cd-4284-a503-4154086dc05a")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp: %+v", resp)
}
