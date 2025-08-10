package client

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

const (
	defaultTimeout = 80 * time.Second
)

// Client is the main client for the Dify API.
type Client struct {
	host       string
	apiKey     string
	adminKey   string
	httpClient *http.Client
}

// NewClient creates a new Dify API client.
func NewClient(host, apiKey string) *Client {
	return &Client{
		host:   host,
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
	}
}

// WithAdminKey sets the admin API key for the client.
func (c *Client) WithAdminKey(adminKey string) *Client {
	c.adminKey = adminKey
	return c
}

// WithHTTPClient sets a custom http.Client for the Dify client.
func (c *Client) WithHTTPClient(httpClient *http.Client) *Client {
	c.httpClient = httpClient
	return c
}

// RedirectError is returned when the API returns a redirect status code.
type RedirectError struct {
	StatusCode int
	Location   string
}

func (e *RedirectError) Error() string {
	return fmt.Sprintf("api returned a redirect to %s (status: %d)", e.Location, e.StatusCode)
}

// sendRequest handles sending a JSON request and decoding a JSON response.
func (c *Client) sendRequest(ctx context.Context, method, path string, payload, result any, extraHeaders map[string]string) error {
	var body io.Reader
	if payload != nil {
		b, err := json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("failed to marshal payload: %w", err)
		}
		body = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.host+path, body)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Default to apiKey, but allow override from extraHeaders
	authHeader := "Bearer " + c.apiKey
	if extraHeaders != nil {
		if val, ok := extraHeaders["Authorization"]; ok {
			authHeader = val
			delete(extraHeaders, "Authorization") // remove it from extraHeaders to avoid setting it twice
		}
	}
	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Content-Type", "application/json")

	for key, value := range extraHeaders {
		req.Header.Set(key, value)
	}

	// Prevent auto-redirects
	c.httpClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 && resp.StatusCode < 400 {
		return &RedirectError{
			StatusCode: resp.StatusCode,
			Location:   resp.Header.Get("Location"),
		}
	}

	if resp.StatusCode >= http.StatusBadRequest {
		return c.decodeError(resp)
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

// SendMultipartRequestWithJSON sends a multipart/form-data request with a file and a JSON payload in a 'data' field.
func (c *Client) SendMultipartRequestWithJSON(ctx context.Context, path string, file io.Reader, filename string, jsonData any, result any) error {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	// Marshal JSON data and add it as a field
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		return fmt.Errorf("failed to marshal json data: %w", err)
	}
	if err := w.WriteField("data", string(jsonBytes)); err != nil {
		return fmt.Errorf("failed to write json data field: %w", err)
	}

	// Add file to the form
	if file != nil {
		fw, err := w.CreateFormFile("file", filename)
		if err != nil {
			return fmt.Errorf("failed to create form file: %w", err)
		}
		if _, err = io.Copy(fw, file); err != nil {
			return fmt.Errorf("failed to copy file to buffer: %w", err)
		}
	}

	w.Close()

	req, err := http.NewRequestWithContext(ctx, "POST", c.host+path, &b)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", w.FormDataContentType())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return c.decodeError(resp)
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

// SendMultipartRequest handles sending a multipart/form-data request, typically for file uploads.
func (c *Client) SendMultipartRequest(ctx context.Context, path string, file io.Reader, filename string, fields map[string]string, result any) error {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	// Add fields to the form
	for key, value := range fields {
		if err := w.WriteField(key, value); err != nil {
			return fmt.Errorf("failed to write field %s: %w", key, err)
		}
	}

	// Add file to the form
	if file != nil {
		fw, err := w.CreateFormFile("file", filename)
		if err != nil {
			return fmt.Errorf("failed to create form file: %w", err)
		}
		if _, err = io.Copy(fw, file); err != nil {
			return fmt.Errorf("failed to copy file to buffer: %w", err)
		}
	}

	w.Close()

	req, err := http.NewRequestWithContext(ctx, "POST", c.host+path, &b)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", w.FormDataContentType())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return c.decodeError(resp)
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

// SendRequestRaw handles requests where the response body is not JSON, but raw bytes.
func (c *Client) SendRequestRaw(ctx context.Context, method, path string, payload any) ([]byte, http.Header, error) {
	var body io.Reader
	if payload != nil {
		b, err := json.Marshal(payload)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to marshal payload: %w", err)
		}
		body = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.host+path, body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return nil, nil, c.decodeError(resp)
	}

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read raw response body: %w", err)
	}

	return rawBody, resp.Header, nil
}

// SendSSEQuest sends a request and returns a channel to read Server-Sent Events.
func (c *Client) SendSSEQuest(ctx context.Context, method, path string, payload any) (<-chan []byte, error) {
	var body io.Reader
	if payload != nil {
		b, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal payload: %w", err)
		}
		body = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.host+path, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	if resp.StatusCode >= http.StatusBadRequest {
		defer resp.Body.Close()
		return nil, c.decodeError(resp)
	}

	events := make(chan []byte)
	go func() {
		defer close(events)
		defer resp.Body.Close()

		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Bytes()
			// SSE data lines start with "data: "
			if bytes.HasPrefix(line, []byte("data: ")) {
				events <- bytes.TrimPrefix(line, []byte("data: "))
			}
		}
		if err := scanner.Err(); err != nil {
			// Handle error, maybe log it or send it over a separate error channel
			// For simplicity, we'll just log it here.
			// In a real-world scenario, a more robust error handling would be needed.
			fmt.Printf("error reading sse stream: %v\n", err)
		}
	}()

	return events, nil
}


func (c *Client) decodeError(resp *http.Response) error {
	// Try to decode the error response body
	var errResp struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Status  int    `json:"status"`
	}
	bodyBytes, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(bodyBytes, &errResp); err == nil && errResp.Message != "" {
		return fmt.Errorf("api error: %s (status: %d, code: %s)", errResp.Message, resp.StatusCode, errResp.Code)
	}

	// Fallback if error response is not in the expected format or body is empty
	return fmt.Errorf("api error: received status code %d, body: %s", resp.StatusCode, string(bodyBytes))
}
