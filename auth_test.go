package customerio

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

// TestClient_FindRegion will test the method FindRegion()
func TestClient_FindRegion(t *testing.T) {
	// t.Parallel() (Cannot run in parallel - issues with overriding the mock client)

	t.Run("successful response", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockFindRegion(http.StatusOK)

		var regionInfo *RegionInfo
		regionInfo, err = client.FindRegion()
		assert.NoError(t, err)
		assert.NotNil(t, regionInfo)
		assert.Equal(t, "https://track.customer.io", regionInfo.URL)
		assert.Equal(t, "us", regionInfo.DataCenter)
		assert.Equal(t, uint64(3), regionInfo.EnvironmentID)
	})

	t.Run("customerIo error", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockFindRegion(http.StatusUnprocessableEntity)

		var regionInfo *RegionInfo
		regionInfo, err = client.FindRegion()
		assert.Error(t, err)
		assert.Nil(t, regionInfo)
	})
}

// TestClient_TestAuth will test the method TestAuth()
func TestClient_TestAuth(t *testing.T) {
	// t.Parallel() (Cannot run in parallel - issues with overriding the mock client)

	t.Run("successful response", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockTestAuth(http.StatusOK)

		err = client.TestAuth()
		assert.NoError(t, err)
	})

	t.Run("customerIo error", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockTestAuth(http.StatusUnauthorized)

		err = client.TestAuth()
		assert.Error(t, err)
	})
}

// mockFindRegion is used for mocking the response
func mockFindRegion(statusCode int) {
	httpmock.Reset()
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%sapi/v1/accounts/region", testTrackingAPIURL),
		httpmock.NewStringResponder(
			statusCode, `{"url": "https://track.customer.io","data_center": "us","environment_id": 3}`,
		),
	)
}

// mockTestAuth is used for mocking the response
func mockTestAuth(statusCode int) {
	httpmock.Reset()
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%sauth", testTrackingAPIURL),
		httpmock.NewStringResponder(
			statusCode, `{"meta": {"message": "Nice credentials."}}`,
		),
	)
}
