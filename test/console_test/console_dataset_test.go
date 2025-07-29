package console_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

var dataset_use_real_url = true

func SetupDatasetMockServer() *httptest.Server {
	handler := http.NewServeMux()

	return httptest.NewServer(handler)
}

func TestGetDatasets(t *testing.T) {
	mockServer := SetupDatasetMockServer()
	defer mockServer.Close()

	client := TestNewClientWithBaseURL(mockServer.URL, dataset_use_real_url)
	resp, err := client.GetDatasets(context.Background(), 1, 10, "", []string{}, false)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp: %+v", *resp)
}
