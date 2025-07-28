package console_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

var plugin_use_real_url = true

// mock server，路径和返回内容严格参考workspaces.go接口实现
func SetupPluginMockServer() *httptest.Server {
	handler := http.NewServeMux()

	return httptest.NewServer(handler)
}

func TestGetPluginList(t *testing.T) {
	mockServer := SetupPluginMockServer()
	defer mockServer.Close()

	client := TestNewClientWithBaseURL(mockServer.URL, plugin_use_real_url)

	resp, err := client.GetPluginList(context.Background(), 1, 10)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp: %+v", resp)
}

func TestInstallPluginFromPkg(t *testing.T) {
	mockServer := SetupPluginMockServer()
	defer mockServer.Close()

	client := TestNewClientWithBaseURL(mockServer.URL, plugin_use_real_url)

	resp, err := client.InstallPluginFromPkg(context.Background(), "./vuln_info_query.difypkg")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp: %+v", resp)
}
