package test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kingfs/godify/models"
)

var apps_use_real_url = false

func SetupAppsMockServer() *httptest.Server {
	handler := http.NewServeMux()

	handler.HandleFunc("/console/api/apps/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		path := r.URL.Path

		if strings.HasSuffix(path, "/chat-messages") {
			resp := models.AppsChatMessageListApiResponse{
				Data: []models.AppsChatMessage{
					{
						ID:             "msg1",
						ConversationID: "conv1",
						FromAccountID:  "acc1",
						FromSource:     "user",
						CreatedAt:      1753153422,
						Status:         "success",
						Query:          "hello",
						Answer:         "world",
						Message: []models.AppsMessage{
							{Role: "user", Text: "hello"},
							{Role: "assistant", Text: "world"},
						},
					},
				},
				HasMore: false,
				Limit:   20,
			}
			json.NewEncoder(w).Encode(resp)
			return
		}
		if strings.Contains(path, "/messages/") {
			resp := models.AppsMessageApiResponse{
				ID:             "msg1",
				ConversationID: "conv1",
				FromAccountID:  "acc1",
				FromSource:     "user",
				CreatedAt:      1753153422,
				Status:         "success",
				Query:          "hello",
				Answer:         "world",
				Message: []models.AppsMessage{
					{Role: "user", Text: "hello"},
					{Role: "assistant", Text: "world"},
				},
			}
			json.NewEncoder(w).Encode(resp)
			return
		}
	})

	return httptest.NewServer(handler)
}

func TestGetAppsChatMessageList(t *testing.T) {
	mockServer := SetupAppsMockServer()
	defer mockServer.Close()

	client := TestNewClientWithBaseURL(mockServer.URL, apps_use_real_url)
	resp, err := client.GetAppsChatMessageList(context.Background(), "eb777f1d-77ac-4b67-b192-4b55371cef3d", "53662705-9c66-4b07-9690-aa6d597b4248", nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp: %+v", *resp)
}

func TestGetAppsMessage(t *testing.T) {
	mockServer := SetupAppsMockServer()
	defer mockServer.Close()

	client := TestNewClientWithBaseURL(mockServer.URL, apps_use_real_url)
	resp, err := client.GetAppsMessage(context.Background(), "eb777f1d-77ac-4b67-b192-4b55371cef3d", "5adcf08d-f535-4516-9bd6-eb71bc561b70")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp: %+v", *resp)
}
