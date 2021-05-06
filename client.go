package customerio

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

// Client is the CustomerIO client/configuration
type Client struct {
	httpClient *resty.Client
	options    *clientOptions // Options are all the default settings / configuration
}

// ClientOptions holds all the configuration for client requests and default resources
// See: https://fly.customer.io/settings/api_credentials
type clientOptions struct {
	apiURL         string        // Regional API endpoint (URL)
	appAPIKey      string        // App or Beta API key
	httpTimeout    time.Duration // Default timeout in seconds for GET requests
	requestTracing bool          // If enabled, it will trace the request timing
	retryCount     int           // Default retry count for HTTP requests
	siteID         string        // Used in conjunction with the Tracking API key
	trackingAPIKey string        // Tracking API key (Only tracking API requests)
	trackURL       string        // Regional Tracking API endpoint (URL)
	userAgent      string        // User agent for all outgoing requests
}

// region is used for changing the location of the API endpoints
type region struct {
	apiURL   string
	trackURL string
}

// Current regions available
var (
	RegionUS = region{
		apiURL:   "https://api.customer.io",
		trackURL: "https://track.customer.io",
	}
	RegionEU = region{
		apiURL:   "https://api-eu.customer.io",
		trackURL: "https://track-eu.customer.io",
	}
)

// ClientOps allow functional options to be supplied
// that overwrite default client options.
type ClientOps func(c *clientOptions)

// WithRegion will change the region API endpoints
func WithRegion(r region) ClientOps {
	return func(c *clientOptions) {
		c.apiURL = r.apiURL
		c.trackURL = r.trackURL
	}
}

// WithHTTPTimeout can be supplied to adjust the default http client timeouts.
// The http client is used when creating requests
// Default timeout is 20 seconds.
func WithHTTPTimeout(timeout time.Duration) ClientOps {
	return func(c *clientOptions) {
		c.httpTimeout = timeout
	}
}

// WithRequestTracing will enable tracing.
// Tracing is disabled by default.
func WithRequestTracing() ClientOps {
	return func(c *clientOptions) {
		c.requestTracing = true
	}
}

// WithRetryCount will overwrite the default retry count for http requests.
// Default retries is 2.
func WithRetryCount(retries int) ClientOps {
	return func(c *clientOptions) {
		c.retryCount = retries
	}
}

// WithUserAgent will overwrite the default useragent.
// Default is package name + version.
func WithUserAgent(userAgent string) ClientOps {
	return func(c *clientOptions) {
		c.userAgent = userAgent
	}
}

// WithTrackingKey will provide the SiteID and Tracking API key
// See: https://fly.customer.io/settings/api_credentials
func WithTrackingKey(siteID, trackingAPIKey string) ClientOps {
	return func(c *clientOptions) {
		c.siteID = siteID
		c.trackingAPIKey = trackingAPIKey
	}
}

// WithAppKey will provide the App or Beta API key
// See: https://fly.customer.io/settings/api_credentials?keyType=app
func WithAppKey(appAPIKey string) ClientOps {
	return func(c *clientOptions) {
		c.appAPIKey = appAPIKey
	}
}

// WithCustomHTTPClient will overwrite the default client with a custom client.
func (c *Client) WithCustomHTTPClient(client *resty.Client) *Client {
	c.httpClient = client
	return c
}

// GetUserAgent will return the user agent string of the client
func (c *Client) GetUserAgent() string {
	return c.options.userAgent
}

// defaultClientOptions will return an Options struct with the default settings
//
// Useful for starting with the default and then modifying as needed
func defaultClientOptions() (opts *clientOptions, err error) {
	// Set the default options
	opts = &clientOptions{
		apiURL:         RegionUS.apiURL,
		httpTimeout:    defaultHTTPTimeout,
		requestTracing: false,
		retryCount:     defaultRetryCount,
		trackURL:       RegionUS.trackURL,
		userAgent:      defaultUserAgent,
	}
	return
}

// NewClient creates a new client for all CustomerIO requests (tracking, app, beta)
//
// If no options are given, it will use the DefaultClientOptions()
// If no client is supplied it will use a default Resty HTTP client
func NewClient(opts ...ClientOps) (*Client, error) {
	defaults, err := defaultClientOptions()
	if err != nil {
		return nil, err
	}
	// Create a new client
	client := &Client{
		options: defaults,
	}
	// overwrite defaults with any set by user
	for _, opt := range opts {
		opt(client.options)
	}
	// Check for at least one type of API key
	if client.options.trackingAPIKey == "" && client.options.appAPIKey == "" {
		return nil, errors.New("missing an API Key (Tracking or App)")
	}
	// Set the Resty HTTP client
	if client.httpClient == nil {
		client.httpClient = resty.New()
		// Set defaults (for GET requests)
		client.httpClient.SetTimeout(client.options.httpTimeout)
		client.httpClient.SetRetryCount(client.options.retryCount)
	}
	return client, nil
}

// auth creates the Basic Auth string using the SiteID and API Key
func (c *Client) auth() string {
	return base64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("%v:%v", c.options.siteID, c.options.trackingAPIKey)))
}

// request is a standard GET / POST / PUT / DELETE request for all outgoing HTTP requests
// Omit the data attribute if using a GET request
func (c *Client) request(httpMethod string, requestURL string, expectedStatusCode int,
	data interface{}) (response StandardResponse, err error) {

	// Set the user agent
	req := c.httpClient.R().SetHeader("User-Agent", c.options.userAgent)

	// Set the body if (PUT || POST)
	if httpMethod != http.MethodGet && httpMethod != http.MethodDelete {
		var j []byte
		j, err = json.Marshal(data)
		if err != nil {
			return
		}
		req = req.SetBody(string(j))
		req.Header.Add("Content-Length", strconv.Itoa(len(j)))
		req.Header.Set("Content-Type", "application/json")
	}

	// Enable tracing
	if c.options.requestTracing {
		req.EnableTrace()
	}

	// Set the authorization and content type
	if strings.Contains(requestURL, c.options.trackURL) {
		req.Header.Add("Authorization", "Basic "+c.auth())
	} else { // App or Beta
		req.Header.Set("Authorization", "Bearer "+c.options.appAPIKey)
	}

	// Fire the request
	var resp *resty.Response
	switch httpMethod {
	case http.MethodPost:
		if resp, err = req.Post(requestURL); err != nil {
			return
		}
	case http.MethodPut:
		if resp, err = req.Put(requestURL); err != nil {
			return
		}
	case http.MethodDelete:
		if resp, err = req.Delete(requestURL); err != nil {
			return
		}
	case http.MethodGet:
		if resp, err = req.Get(requestURL); err != nil {
			return
		}
	}

	// Tracing enabled?
	if c.options.requestTracing {
		response.Tracing = resp.Request.TraceInfo()
	}

	// Set the status code & body
	response.StatusCode = resp.StatusCode()
	response.Body = resp.Body()

	// Process if error (different error formats for different API endpoint/urls)
	if expectedStatusCode != response.StatusCode {
		if strings.Contains(requestURL, "/v1/send/email") {
			var meta struct {
				Meta struct {
					Err string `json:"error"`
				} `json:"meta"`
			}
			if err = json.Unmarshal(response.Body, &meta); err != nil {
				err = &TransactionalError{
					StatusCode: response.StatusCode,
					Err:        string(response.Body),
				}
			} else {
				err = &TransactionalError{
					StatusCode: response.StatusCode,
					Err:        meta.Meta.Err,
				}
			}
			return
		} else {
			err = &APIError{
				status: response.StatusCode,
				url:    requestURL,
				body:   response.Body,
			}
		}
	}
	return
}
