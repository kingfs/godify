package service_api

import (
	"context"
	"io"
	"net/http"

	"github.com/kingfs/godify/client"
)

// AudioToTextResponse is the response from the audio-to-text endpoint.
type AudioToTextResponse struct {
	Text string `json:"text"`
}

// TextToAudioRequest is the request body for the text-to-audio endpoint.
type TextToAudioRequest struct {
	Text      string `json:"text"`
	Voice     string `json:"voice,omitempty"`
	Streaming bool   `json:"streaming,omitempty"`
	MessageID string `json:"message_id,omitempty"`
}

// AudioToText transcribes an audio file to text.
// The file should be an io.Reader, and filename should be provided.
func (c *client.Client) AudioToText(ctx context.Context, file io.Reader, filename string) (*AudioToTextResponse, error) {
	var result AudioToTextResponse
	err := c.SendMultipartRequest(ctx, "/v1/audio-to-text", file, filename, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// TextToAudio converts text to an audio file.
// It returns the raw audio data as a byte slice and the response headers.
func (c *client.Client) TextToAudio(ctx context.Context, req *TextToAudioRequest) ([]byte, http.Header, error) {
	return c.SendRequestRaw(ctx, "POST", "/v1/text-to-audio", req)
}
