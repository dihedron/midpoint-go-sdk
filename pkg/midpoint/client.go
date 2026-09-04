package midpoint

import (
	"crypto/tls"
	"net/http"
	"time"

	"resty.dev/v3"
)

// API handles REST API interactions with Basic Auth.
type API struct {
	User *UserService
	Self *SelfService
}

// option defines functional options for configuring the API.
type Option func(*resty.Client)

// WithTimeout configures a custom timeout for the HTTP client.
func WithTimeout(timeout time.Duration) Option {
	return func(c *resty.Client) {
		c.SetTimeout(timeout)
	}
}

// WithClientTrustAnchors configures the root certificates for the HTTP client.
func WithClientTrustAnchors(paths ...string) Option {
	return func(c *resty.Client) {
		c.SetClientRootCertificates(paths...)
	}
}

// WithClientTrustAnchorsWatcher configures the root certificates for the HTTP client.
func WithClientTrustAnchorsWatcher(duration time.Duration, paths ...string) Option {
	return func(c *resty.Client) {
		c.SetClientRootCertificatesWatcher(&resty.CertWatcherOptions{PoolInterval: duration}, paths...)
	}
}

// WithClientCertificates configures the client certificates for the HTTP client;
// load the keypair from file (e.g. tlx.LoadX509KeyPair()) or use an in-memory certificate.
func WithClientCertificates(certs ...tls.Certificate) Option {
	return func(c *resty.Client) {
		c.SetCertificates(certs...)
	}
}

// WithClientCertificateFromFiles configures the client certificates for the HTTP client;
// it loads the keypair from the provided files.
func WithClientCertificateFromFiles(certPath string, keyPath string) Option {
	return func(c *resty.Client) {
		c.SetCertificateFromFile(certPath, keyPath)
	}
}

// WithClientCertificateFromString configures the client certificates for the HTTP client;
// it uses the provided certificate and key as strings.
func WithClientCertificateFromString(cert string, key string) Option {
	return func(c *resty.Client) {
		c.SetCertificateFromString(cert, key)
	}
}

// WithClientCertificatesWatcher configures the client certificates for the HTTP client.
func WithClientCertificatesWatcher(duration time.Duration, paths ...string) Option {
	return func(c *resty.Client) {
		c.SetClientRootCertificatesWatcher(&resty.CertWatcherOptions{PoolInterval: duration}, paths...)
	}
}

// WithTransport configures the transport for the HTTP client.
func WithTransport(transport http.RoundTripper) Option {
	return func(c *resty.Client) {
		c.SetTransport(transport)
	}
}

// WithRetry configures the retry count for the HTTP client.
func WithRetry(count int) Option {
	return func(c *resty.Client) {
		c.SetRetryCount(count)
	}
}

// WithDebug configures the debug mode for the HTTP client.
func WithDebug(enable bool) Option {
	return func(c *resty.Client) {
		c.SetDebug(enable)
	}
}

// WithImpersonation configures the principal to use for impersonation in API calls.
func WithImpersonation(principal string) Option {
	return func(c *resty.Client) {
		c.SetHeader("Switch-To-Principal", principal)
	}
}

// New initializes and returns a new REST API client.
func New(baseURL string, username string, password string, opts ...Option) *API {

	client := resty.
		New().
		SetBasicAuth(username, password).
		SetBaseURL(baseURL).
		SetHeader("Accept", "application/json")

	for _, opt := range opts {
		opt(client)
	}

	// new go 1.27 initialisation of nested anonymous structs
	return &API{
		User: &UserService{client: client},
		Self: &SelfService{client: client},
	}
}

/*
func (c *Client) Get[T any](ctx context.Context, path string) (*T, *Result, error) {
	entity := new(T)
	result, err := c.Do[T](ctx, http.MethodGet, path, nil, entity)
	if err != nil {
		slog.Error("failed to perform GET request", "path", path, "error", err)
		return nil, result, err
	}
	return entity, result, nil
}

func (c *Client) Post[T any](ctx context.Context, path string, object *T) (*Result, error) {
	entity := new(T)
	result, err := c.Do(ctx, http.MethodPost, path, object, entity)
	if err != nil {
		slog.Error("failed to perform POST request", "path", path, "object", object, "error", err)
		return result, err
	}
	return result, nil
}

type Result struct {
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Location   *url.URL
	Headers    http.Header
	Body       io.Reader
}

// Do executes an HTTP request, attaching Basic Auth and standard headers.
func (c *Client) Do[T any, S any](ctx context.Context, method string, path string, payload *T, entity *S) (*Result, error) {
	return c.DoAs(ctx, "", method, path, payload, entity)
}

// DoAs executes an HTTP request, attaching Basic Auth and standard headers and impersonating
// the given principal in the request.
func (c *Client) DoAs[T any, S any](ctx context.Context, principal string, method string, path string, payload *T, entity *S) (*Result, error) {
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

	response, err := c.Do(request)
	if err != nil {
		slog.Error("error performing request", "error", err)
		return nil, fmt.Errorf("request failed: %w", err)
	}

	r := &Result{
		Status:     response.Status,
		StatusCode: response.StatusCode,
		Headers:    response.Header,
	}

	location, err := response.Location()
	if err != nil && err != http.ErrNoLocation {
		slog.Warn("failed to retrieve Location header")
	} else if location != nil {
		slog.Debug("location header", "location", location.String())
		r.Location = location
	}

	if response.Body != nil {
		// ensure the response body is drained to allow TCP connection reuse;
		// use separate function to allow immediate response body close
		if err := func() error {
			defer response.Body.Close()
			data, err := io.ReadAll(response.Body)
			if err != nil {
				slog.Error("error reading response body", "error", err)
				return err
			}
			slog.Debug("response body", "data", data)
			r.Body = bytes.NewReader(data)
			return nil
		}(); err != nil {
			return nil, err
		}
	}

	// if the user wants to decode JSON into 'entity', process it here
	if entity != nil && r.Body != nil {
		if response.StatusCode < 200 || response.StatusCode >= 300 {
			data, err := io.ReadAll(r.Body)
			if err != nil {
				slog.Error("error reading response body", "error", err)
				return r, err
			}
			r.Body = bytes.NewReader(data)
			if err := json.NewDecoder(r.Body).Decode(entity); err != nil {
				slog.Error("failed to decode response", "error", err)
				return nil, fmt.Errorf("failed to decode response body: %w", err)
			}
		}
	}

	return r, nil
}
*/

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
