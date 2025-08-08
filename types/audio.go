package types

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
	User      string `json:"user"`
}
