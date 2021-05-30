package customerio

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

// TestClient_UpdateCollection will test the method UpdateCollection()
func TestClient_UpdateCollection(t *testing.T) {
	// t.Parallel() (Cannot run in parallel - issues with overriding the mock client)

	t.Run("successful response (create)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockNewCollection(http.StatusOK)

		err = client.UpdateCollection(
			"",
			testCollectionName,
			[]map[string]interface{}{
				{
					"item_name":       "test_item_1",
					"id_field":        1,
					"timestamp_field": time.Now().UTC().Unix(),
				},
				{
					"item_name":       "test_item_2",
					"id_field":        2,
					"timestamp_field": time.Now().UTC().Unix(),
				},
			})
		assert.NoError(t, err)
	})

	t.Run("successful response (update)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockUpdateCollection(http.StatusOK, testCollectionID)

		err = client.UpdateCollection(
			testCollectionID,
			testCollectionName,
			[]map[string]interface{}{
				{
					"item_name":       "test_item_1",
					"id_field":        1,
					"timestamp_field": time.Now().UTC().Unix(),
				},
				{
					"item_name":       "test_item_2",
					"id_field":        2,
					"timestamp_field": time.Now().UTC().Unix(),
				},
			})
		assert.NoError(t, err)
	})

	t.Run("missing collection name", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockUpdateCollection(http.StatusOK, testCollectionID)

		err = client.UpdateCollection(
			testCollectionID,
			"",
			[]map[string]interface{}{
				{
					"item_name":       "test_item_1",
					"id_field":        1,
					"timestamp_field": time.Now().UTC().Unix(),
				},
				{
					"item_name":       "test_item_2",
					"id_field":        2,
					"timestamp_field": time.Now().UTC().Unix(),
				},
			})
		assert.Error(t, err)
		checkParamError(t, err, "collectionName")
	})

	t.Run("customerIo error", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockNewCollection(http.StatusUnprocessableEntity)

		err = client.UpdateCollection(
			"",
			testCollectionName,
			[]map[string]interface{}{
				{
					"item_name":       "test_item_1",
					"id_field":        1,
					"timestamp_field": time.Now().UTC().Unix(),
				},
				{
					"item_name":       "test_item_2",
					"id_field":        2,
					"timestamp_field": time.Now().UTC().Unix(),
				},
			})
		assert.Error(t, err)
	})
}

// ExampleClient_UpdateCollection example using UpdateCollection()
//
// See more examples in /examples/
func ExampleClient_UpdateCollection() {

	// Load the client
	client, err := newTestClient()
	if err != nil {
		fmt.Printf("error loading client: %s", err.Error())
		return
	}

	mockNewCollection(http.StatusOK)

	// New collection
	err = client.UpdateCollection(
		"",
		testCollectionName,
		[]map[string]interface{}{
			{
				"item_name":       "test_item_1",
				"id_field":        1,
				"timestamp_field": time.Now().UTC().Unix(),
			},
			{
				"item_name":       "test_item_2",
				"id_field":        2,
				"timestamp_field": time.Now().UTC().Unix(),
			},
		})
	if err != nil {
		fmt.Printf("error creating collection: " + err.Error())
		return
	}
	fmt.Printf("collection created: %s", testCollectionName)
	// Output:collection created: test_collection
}

// BenchmarkClient_UpdateCollection benchmarks the method UpdateCollection()
func BenchmarkClient_UpdateCollection(b *testing.B) {
	client, _ := newTestClient()
	mockUpdateCollection(http.StatusOK, testCollectionID)
	data := []map[string]interface{}{
		{
			"item_name":       "test_item_1",
			"id_field":        1,
			"timestamp_field": time.Now().UTC().Unix(),
		},
		{
			"item_name":       "test_item_2",
			"id_field":        2,
			"timestamp_field": time.Now().UTC().Unix(),
		},
	}
	for i := 0; i < b.N; i++ {
		_ = client.UpdateCollection(testCollectionID, testCollectionName, data)
	}
}

// TestClient_UpdateCollectionViaURL will test the method UpdateCollectionViaURL()
func TestClient_UpdateCollectionViaURL(t *testing.T) {
	// t.Parallel() (Cannot run in parallel - issues with overriding the mock client)

	t.Run("successful response (create)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockNewCollectionViaURL(http.StatusOK)

		err = client.UpdateCollectionViaURL(
			"",
			testCollectionName,
			testCollectionURL,
		)
		assert.NoError(t, err)
	})

	t.Run("successful response (update)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockUpdateCollectionViaURL(http.StatusOK, testCollectionID)

		err = client.UpdateCollectionViaURL(
			testCollectionID,
			testCollectionName,
			testCollectionURL,
		)
		assert.NoError(t, err)
	})

	t.Run("missing collection name", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockNewCollectionViaURL(http.StatusOK)

		err = client.UpdateCollectionViaURL(
			"",
			"",
			testCollectionURL,
		)
		assert.Error(t, err)
		checkParamError(t, err, "collectionName")
	})

	t.Run("customerIo error", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockNewCollectionViaURL(http.StatusUnprocessableEntity)

		err = client.UpdateCollectionViaURL(
			"",
			testCollectionName,
			testCollectionURL,
		)
		assert.Error(t, err)
	})
}

// ExampleClient_UpdateCollectionViaURL example using UpdateCollectionViaURL()
//
// See more examples in /examples/
func ExampleClient_UpdateCollectionViaURL() {

	// Load the client
	client, err := newTestClient()
	if err != nil {
		fmt.Printf("error loading client: %s", err.Error())
		return
	}

	mockNewCollectionViaURL(http.StatusOK)

	// New collection
	err = client.UpdateCollectionViaURL(
		"",
		testCollectionName,
		testCollectionURL,
	)
	if err != nil {
		fmt.Printf("error creating collection: " + err.Error())
		return
	}
	fmt.Printf("collection created: %s", testCollectionName)
	// Output:collection created: test_collection
}

// BenchmarkClient_UpdateCollectionViaURL benchmarks the method UpdateCollectionViaURL()
func BenchmarkClient_UpdateCollectionViaURL(b *testing.B) {
	client, _ := newTestClient()
	mockUpdateCollectionViaURL(http.StatusOK, testCollectionID)
	for i := 0; i < b.N; i++ {
		_ = client.UpdateCollectionViaURL(testCollectionID, testCollectionName, testCollectionURL)
	}
}

// mockNewCollection is used for mocking the response
func mockNewCollection(statusCode int) {
	httpmock.Reset()
	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%sv1/api/collections", testBetaAPIURL),
		httpmock.NewStringResponder(
			statusCode, "",
		),
	)
}

// mockUpdateCollection is used for mocking the response
func mockUpdateCollection(statusCode int, collectionID string) {
	httpmock.Reset()
	httpmock.RegisterResponder(http.MethodPut, fmt.Sprintf("%sv1/api/collections/%s", testBetaAPIURL, collectionID),
		httpmock.NewStringResponder(
			statusCode, "",
		),
	)
}

// mockNewCollectionViaURL is used for mocking the response
func mockNewCollectionViaURL(statusCode int) {
	httpmock.Reset()
	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%sv1/api/collections", testBetaAPIURL),
		httpmock.NewStringResponder(
			statusCode, "",
		),
	)
}

// mockUpdateCollectionViaURL is used for mocking the response
func mockUpdateCollectionViaURL(statusCode int, collectionID string) {
	httpmock.Reset()
	httpmock.RegisterResponder(http.MethodPut, fmt.Sprintf("%sv1/api/collections/%s", testBetaAPIURL, collectionID),
		httpmock.NewStringResponder(
			statusCode, "",
		),
	)
}
