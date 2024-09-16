package customerio

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const (
	testSiteID         = "TestSiteID1234567"
	testAppAPIKey      = "TestAPIKey1234567"
	testTrackingAPIKey = "TestTrackingAPIKey1234567"
)

// newTestClient will return a client for testing purposes
func newTestClient() (*Client, error) {
	// Create a Resty Client
	client := resty.New()

	// Get the underlying HTTP Client and set it to Mock
	httpmock.ActivateNonDefault(client.GetClient())

	// Create a new client
	newClient, err := NewClient(WithRequestTracing(), WithTrackingKey(testSiteID, testTrackingAPIKey))
	if err != nil {
		return nil, err
	}
	newClient.WithCustomHTTPClient(client)

	// Return the mocking client
	return newClient, nil
}

// TestNewClient will test the method NewClient()
func TestNewClient(t *testing.T) {
	t.Parallel()

	t.Run("default client", func(t *testing.T) {
		client, err := NewClient(WithTrackingKey(testSiteID, testTrackingAPIKey))
		assert.NoError(t, err)
		assert.NotNil(t, client)
		assert.Equal(t, defaultHTTPTimeout, client.options.httpTimeout)
		assert.Equal(t, defaultRetryCount, client.options.retryCount)
		assert.Equal(t, defaultUserAgent, client.options.userAgent)
		assert.Equal(t, false, client.options.requestTracing)
		assert.Equal(t, RegionUS.apiURL, client.options.apiURL)
		assert.Equal(t, RegionUS.betaURL, client.options.betaURL)
		assert.Equal(t, RegionUS.trackURL, client.options.trackURL)
	})

	t.Run("custom http client", func(t *testing.T) {
		customHTTPClient := resty.New()
		customHTTPClient.SetTimeout(defaultHTTPTimeout)
		client, err := NewClient(WithTrackingKey(testSiteID, testTrackingAPIKey))
		assert.NoError(t, err)
		assert.NotNil(t, client)
		client.WithCustomHTTPClient(customHTTPClient)
	})

	t.Run("custom http timeout", func(t *testing.T) {
		client, err := NewClient(WithTrackingKey(testSiteID, testTrackingAPIKey), WithHTTPTimeout(10*time.Second))
		assert.NoError(t, err)
		assert.NotNil(t, client)
		assert.Equal(t, 10*time.Second, client.options.httpTimeout)
	})

	t.Run("custom retry count", func(t *testing.T) {
		client, err := NewClient(WithTrackingKey(testSiteID, testTrackingAPIKey), WithRetryCount(3))
		assert.NoError(t, err)
		assert.NotNil(t, client)
		assert.Equal(t, 3, client.options.retryCount)
	})

	t.Run("custom options", func(t *testing.T) {
		client, err := NewClient(WithTrackingKey(testSiteID, testTrackingAPIKey), WithUserAgent("custom user agent"))
		assert.NotNil(t, client)
		assert.NoError(t, err)
		assert.Equal(t, "custom user agent", client.GetUserAgent())
	})

	t.Run("custom region (EU)", func(t *testing.T) {
		client, err := NewClient(WithTrackingKey(testSiteID, testTrackingAPIKey), WithRegion(RegionEU))
		assert.NotNil(t, client)
		assert.NoError(t, err)
		assert.Equal(t, client.options.apiURL, "https://api-eu.customer.io")
		assert.Equal(t, client.options.betaURL, "https://beta-api.customer.io")
		assert.Equal(t, client.options.trackURL, "https://track-eu.customer.io")
	})

	t.Run("custom region (US)", func(t *testing.T) {
		client, err := NewClient(WithTrackingKey(testSiteID, testTrackingAPIKey), WithRegion(RegionUS))
		assert.NotNil(t, client)
		assert.NoError(t, err)
		assert.Equal(t, client.options.apiURL, "https://api.customer.io")
		assert.Equal(t, client.options.betaURL, "https://beta-api.customer.io")
		assert.Equal(t, client.options.trackURL, "https://track.customer.io")
	})

	t.Run("test auth (tracking)", func(t *testing.T) {
		client, err := NewClient(WithTrackingKey(testSiteID, testTrackingAPIKey), WithRegion(RegionUS))
		assert.NotNil(t, client)
		assert.NoError(t, err)
		assert.Equal(t, "VGVzdFNpdGVJRDEyMzQ1Njc6VGVzdFRyYWNraW5nQVBJS2V5MTIzNDU2Nw==", client.auth())
		assert.Equal(t, testSiteID, client.options.siteID)
		assert.Equal(t, testTrackingAPIKey, client.options.trackingAPIKey)
	})

	t.Run("test app api key", func(t *testing.T) {
		client, err := NewClient(WithAppKey(testAppAPIKey))
		assert.NotNil(t, client)
		assert.NoError(t, err)
		assert.Equal(t, testAppAPIKey, client.options.appAPIKey)
	})
}

// TestClient_GetUserAgent will test the method GetUserAgent()
func TestClient_GetUserAgent(t *testing.T) {
	t.Parallel()

	t.Run("get user agent", func(t *testing.T) {
		client, err := NewClient(WithTrackingKey(testSiteID, testTrackingAPIKey))
		assert.NoError(t, err)
		assert.NotNil(t, client)
		userAgent := client.GetUserAgent()
		assert.Equal(t, defaultUserAgent, userAgent)
	})
}

// ExampleNewClient example using NewClient()
//
// See more examples in /examples/
func ExampleNewClient() {
	client, err := NewClient(WithTrackingKey(testSiteID, testTrackingAPIKey))
	if err != nil {
		fmt.Printf("error loading client: %s", err.Error())
		return
	}
	fmt.Printf("loaded client: %s", client.options.userAgent)
	// Output:loaded client: go-customerio: v1.5.0
}

// BenchmarkNewClient benchmarks the method NewClient()
func BenchmarkNewClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = NewClient(WithTrackingKey(testSiteID, testTrackingAPIKey))
	}
}

// TestDefaultClientOptions will test the method defaultClientOptions()
func TestDefaultClientOptions(t *testing.T) {
	t.Parallel()

	options := defaultClientOptions()
	assert.NotNil(t, options)

	assert.Equal(t, defaultUserAgent, options.userAgent)
	assert.Equal(t, defaultHTTPTimeout, options.httpTimeout)
	assert.Equal(t, defaultRetryCount, options.retryCount)
	assert.Equal(t, false, options.requestTracing)
}

// BenchmarkDefaultClientOptions benchmarks the method defaultClientOptions()
func BenchmarkDefaultClientOptions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = defaultClientOptions()
	}
}
