package midpoint

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client handles REST API interactions with Basic Auth.
type Client struct {
	baseURL  string
	username string
	password string
	client   *http.Client
}

// option defines functional options for configuring the Client.
type Option func(*Client)

// WithTimeout configures a custom timeout for the HTTP client.
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.client.Timeout = timeout
	}
}

// WithHTTPClient allows passing a custom net/http Client.
func WithHTTPClient(custom *http.Client) Option {
	return func(c *Client) {
		c.client = custom
	}
}

// New initializes and returns a new REST API client.
func New(baseURL string, username string, password string, opts ...Option) *Client {
	c := &Client{
		baseURL:  baseURL,
		username: username,
		password: password,
		client: &http.Client{
			Timeout: 15 * time.Second,
		},
	}

	for _, opt := range opts {
		opt(c)
	}
	return c
}

// DoAs executes an HTTP request, attaching Basic Auth and standard headers and impersonating
// the given principal in the request.
func (c *Client) DoAs(ctx context.Context, principal string, method, path string, payload any, result any) (*http.Response, error) {
	url := c.baseURL + path

	var reader io.Reader
	if payload != nil {
		jsonBytes, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reader = bytes.NewBuffer(jsonBytes)
	}

	request, err := http.NewRequestWithContext(ctx, method, url, reader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// attach Basic Authentication header
	request.SetBasicAuth(c.username, c.password)

	// do impersonation if so instructed
	if principal != "" {
		request.Header.Set("Switch-To-Principal", principal)
	}

	// set standard API headers
	request.Header.Set("Accept", "application/json")
	if payload != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	response, err := c.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	// ensure the response body is drained to allow TCP connection reuse.
	// If the user wants to decode JSON into 'result', process it here.
	if result != nil && response.Body != nil {
		defer response.Body.Close()

		if response.StatusCode < 200 || response.StatusCode >= 300 {
			rawBody, _ := io.ReadAll(response.Body)
			return response, fmt.Errorf("API error [Status %d]: %s", response.StatusCode, string(rawBody))
		}

		if err := json.NewDecoder(response.Body).Decode(result); err != nil {
			return response, fmt.Errorf("failed to decode response body: %w", err)
		}
	}

	return response, nil
}

// Do executes an HTTP request, attaching Basic Auth and standard headers.
func (c *Client) Do(ctx context.Context, method string, path string, body any, result any) (*http.Response, error) {
	return c.DoAs(ctx, "", method, path, body, result)
}

type Error struct {
	Namespace string `json:"@ns"`
	Object    struct {
		Type           string    `json:"@type"`
		Operation      string    `json:"operation"`
		Status         string    `json:"status"`
		Importance     string    `json:"importance"`
		Start          time.Time `json:"start"`
		End            time.Time `json:"end"`
		Microseconds   int       `json:"microseconds"`
		InvocationID   int       `json:"invocationId"`
		Token          int64     `json:"token"`
		Message        string    `json:"message"`
		PartialResults []any     `json:"partialResults"`
	} `json:"object"`
}
