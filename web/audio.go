package web

import (
	"context"
	"io"
	"net/http"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/types"
)

// WebAudioToText transcribes an audio file to text.
// The file should be an io.Reader, and filename should be provided.
func (c *client.Client) WebAudioToText(ctx context.Context, user, filename string, file io.Reader) (*types.AudioToTextResponse, error) {
	var result types.AudioToTextResponse
	fields := map[string]string{"user": user}
	err := c.SendMultipartRequest(ctx, "/api/audio-to-text", file, filename, fields, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// WebTextToAudio converts text to an audio file.
// It returns the raw audio data as a byte slice and the response headers.
func (c *client.Client) WebTextToAudio(ctx context.Context, req *types.TextToAudioRequest) ([]byte, http.Header, error) {
	return c.SendRequestRaw(ctx, "POST", "/api/text-to-audio", req)
}
