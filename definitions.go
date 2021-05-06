package customerio

import (
	"time"

	"github.com/go-resty/resty/v2"
)

// Defaults for all functions
const (
	defaultHTTPTimeout = 20 * time.Second            // Default timeout for all GET requests in seconds
	defaultRetryCount  = 2                           // Default retry count for HTTP requests
	defaultUserAgent   = "go-customerio: " + version // Default user agent
	version            = "v1.2.0"                    // CustomerIO version
)

// DevicePlatform is the platform for the customer device
type DevicePlatform string

// Allowed types of platforms
const (
	PlatformIOs     DevicePlatform = "ios"
	PlatformAndroid DevicePlatform = "android"
)

// acceptedPlatforms will return true if the platform is accepted
func acceptedPlatforms(platform DevicePlatform) bool {
	if platform == PlatformIOs || platform == PlatformAndroid {
		return true
	}
	return false
}

// Device is the customer device model
type Device struct {
	ID       string         `json:"id"`
	LastUsed int64          `json:"last_used"`
	Platform DevicePlatform `json:"platform"`
}

// StandardResponse is the standard fields returned on all responses
type StandardResponse struct {
	Body       []byte          `json:"-"` // Body of the response request
	StatusCode int             `json:"-"` // Status code returned on the request
	Tracing    resty.TraceInfo `json:"-"` // Trace information if enabled on the request
}
