package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetupAppsMockServer() *httptest.Server {
	handler := http.NewServeMux()
	
	return httptest.NewServer(handler)
}

func TestGetAppsChatMessageList(t *testing.T) {
	// mockServer := SetupAppsMockServer()
	// defer mockServer.Close()

	test_url := "http://localhost"
	client := NewClientWithBaseURL(test_url)
	resp, err := client.GetAppsChatMessageList(context.Background(), "eb777f1d-77ac-4b67-b192-4b55371cef3d", "53662705-9c66-4b07-9690-aa6d597b4248", nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp: %+v", *resp)
}

func TestGetAppsMessage(t *testing.T) {
	// mockServer := SetupAppsMockServer()
	// defer mockServer.Close()

	test_url := "http://localhost"
	client := NewClientWithBaseURL(test_url)
	resp, err := client.GetAppsMessage(context.Background(), "eb777f1d-77ac-4b67-b192-4b55371cef3d", "5adcf08d-f535-4516-9bd6-eb71bc561b70")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp: %+v", *resp)
}