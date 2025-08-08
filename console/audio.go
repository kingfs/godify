package console

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/types"
)

// Voice represents a text-to-speech voice.
type Voice struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Gender string `json:"gender"`
	Language string `json:"language"`
}

// VoicesResponse is the response for getting available TTS voices.
type VoicesResponse struct {
	Voices []Voice `json:"voices"`
}

// ConsoleAudioToText transcribes an audio file to text within the console context.
func (c *client.Client) ConsoleAudioToText(ctx context.Context, appID, filename string, file io.Reader) (*types.AudioToTextResponse, error) {
	var result types.AudioToTextResponse
	path := fmt.Sprintf("/console/api/apps/%s/audio-to-text", appID)
	err := c.SendMultipartRequest(ctx, path, file, filename, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ConsoleTextToAudio converts text to an audio file within the console context.
func (c *client.Client) ConsoleTextToAudio(ctx context.Context, appID string, req *types.TextToAudioRequest) ([]byte, http.Header, error) {
	path := fmt.Sprintf("/console/api/apps/%s/text-to-audio", appID)
	return c.SendRequestRaw(ctx, "POST", path, req)
}

// GetTTSVoices retrieves the available text-to-speech voices for an app.
func (c *client.Client) GetTTSVoices(ctx context.Context, appID, language string) (*VoicesResponse, error) {
	var result VoicesResponse
	path := fmt.Sprintf("/console/api/apps/%s/text-to-audio/voices?language=%s", appID, language)
	err := c.sendRequest(ctx, "GET", path, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
