package console

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/kingfs/godify/client"
	"github.com/kingfs/godify/types"
)

// InstalledAppAudioToText transcribes audio for an installed app.
func (c *client.Client) InstalledAppAudioToText(ctx context.Context, installedAppID, filename string, file io.Reader) (*types.AudioToTextResponse, error) {
	var result types.AudioToTextResponse
	path := fmt.Sprintf("/console/api/installed-apps/%s/audio-to-text", installedAppID)
	err := c.SendMultipartRequest(ctx, path, file, filename, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// InstalledAppTextToAudio converts text to audio for an installed app.
func (c *client.Client) InstalledAppTextToAudio(ctx context.Context, installedAppID string, req *types.TextToAudioRequest) ([]byte, http.Header, error) {
	path := fmt.Sprintf("/console/api/installed-apps/%s/text-to-audio", installedAppID)
	return c.SendRequestRaw(ctx, "POST", path, req)
}
