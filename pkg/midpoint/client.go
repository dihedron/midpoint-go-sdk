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

func WithTraceRequest(enabled bool) Option {
	return func(c *resty.Client) {
		c.SetTrace(enabled)
	}
}

func WithSaveResponse(enabled bool, path string) Option {
	return func(c *resty.Client) {
		c.SetResponseSaveDirectory(path)
		c.SetResponseSaveToFile(enabled)
	}
}

// New initializes and returns a new REST API client.
func New(baseURL string, username string, password string, opts ...Option) *API {
	client := resty.
		New().
		SetBasicAuth(username, password).
		SetBaseURL(baseURL).
		SetHeader("Accept", "application/json").
		SetLoggerWarnLevel(true)
	for _, opt := range opts {
		opt(client)
	}

	// new go 1.27 initialisation of nested anonymous structs
	return &API{
		User: &UserService{client: client},
		Self: &SelfService{client: client},
	}
}

const (
	StatusHandledError = 240
	StatusPartialError = 250
)

type Error struct {
	Ns     *string `json:"@ns,omitempty"`
	Object *struct {
		Ns           *string    `json:"@ns,omitempty"`
		Type         *string    `json:"@type,omitempty"`
		Operation    *string    `json:"operation,omitempty"`
		Status       *string    `json:"status,omitempty"`
		Importance   *string    `json:"importance,omitempty"`
		Start        *time.Time `json:"start,omitempty"`
		End          *time.Time `json:"end,omitempty"`
		Microseconds int        `json:"microseconds,omitempty"`
		InvocationID int        `json:"invocationId,omitempty"`
		Params       *struct {
			Entry *[]struct {
				ParamValue struct {
					Type  string `json:"@type,omitempty"`
					Value string `json:"@value,omitempty"`
				} `json:"paramValue,omitempty"`
				Key *string `json:"key,omitempty"`
			} `json:"entry,omitempty"`
		} `json:"params,omitempty"`
		Token          int64   `json:"token,omitempty"`
		Message        *string `json:"message,omitempty"`
		Details        *string `json:"details,omitempty"`
		PartialResults *[]any  `json:"partialResults,omitempty"`
	} `json:"object,omitempty"`
}
