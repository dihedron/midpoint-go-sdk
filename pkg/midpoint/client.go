package midpoint

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
)

// Client handles REST API interactions with Basic Auth.
type Client struct {
	baseURL  string
	username string
	password string
	client   *http.Client
	User     *UserService
	Self     *SelfService
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

	// new go 1.27 initialisation of nested anonymous structs
	c.User = &UserService{client: c}
	c.Self = &SelfService{client: c}
	return c
}

func (c *Client) Get[T any](ctx context.Context, path string) (*T, error) {
	result := new(T)
	if _, err := c.Do[T](ctx, http.MethodGet, path, nil, result); err != nil {
		slog.Error("failed performing GET request", "path", path, "error", err)
		return nil, err
	}
	return result, nil
}

// Do executes an HTTP request, attaching Basic Auth and standard headers.
func (c *Client) Do[T any, S any](ctx context.Context, method string, path string, payload *T, result *S) (*http.Response, error) {
	return c.DoAs(ctx, "", method, path, payload, result)
}

// DoAs executes an HTTP request, attaching Basic Auth and standard headers and impersonating
// the given principal in the request.
func (c *Client) DoAs[T any, S any](ctx context.Context, principal string, method string, path string, payload *T, result *S) (*http.Response, error) {
	url := c.baseURL + path

	var reader io.Reader
	if payload != nil {
		slog.Debug("request payload", "data", payload)
		data, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reader = bytes.NewBuffer(data)
	}

	request, err := http.NewRequestWithContext(ctx, method, url, reader)
	if err != nil {
		slog.Error("error creating request", "error", err)
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
		slog.Error("error performing request", "error", err)
		return nil, fmt.Errorf("request failed: %w", err)
	}

	// ensure the response body is drained to allow TCP connection reuse.
	// If the user wants to decode JSON into 'result', process it here.
	if result != nil && response.Body != nil {
		defer response.Body.Close()

		if response.StatusCode < 200 || response.StatusCode >= 300 {
			rawBody, err := io.ReadAll(response.Body)
			if err != nil {
				slog.Error("error reading response body", "error", err)
				return nil, err
			}
			return response, fmt.Errorf("API error [Status %d]: %s", response.StatusCode, string(rawBody))
		}

		data, _ := io.ReadAll(response.Body)
		slog.Debug("response body", "data", data)
		//fmt.Printf("%s\n", string(data))

		if err := json.NewDecoder(bytes.NewReader(data)).Decode(result); err != nil {
			return response, fmt.Errorf("failed to decode response body: %w", err)
		}
	}

	return response, nil
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
