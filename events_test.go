package customerio

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

// TestClient_NewEvent will test the method NewEvent()
func TestClient_NewEvent(t *testing.T) {
	// t.Parallel() (Cannot run in parallel - issues with overriding the mock client)

	t.Run("successful response", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockNewEvent(http.StatusOK, testCustomerID)

		err = client.NewEvent(
			testCustomerID, testEventName, time.Now().UTC(),
			map[string]interface{}{
				"field_name":      "some_value",
				"int_field":       123,
				"timestamp_field": time.Now().UTC().Unix(),
			})
		assert.NoError(t, err)
	})

	t.Run("missing customer id", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockNewEvent(http.StatusOK, testCustomerID+"456")

		err = client.NewEvent(
			"", testEventName, time.Now().UTC(),
			map[string]interface{}{
				"field_name":      "some_value",
				"int_field":       123,
				"timestamp_field": time.Now().UTC().Unix(),
			})
		assert.Error(t, err)
		checkParamError(t, err, "customerIDOrEmail")
	})

	t.Run("missing event name", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockNewEvent(http.StatusOK, testCustomerID)

		err = client.NewEvent(
			testCustomerID, "", time.Now().UTC(),
			map[string]interface{}{
				"field_name":      "some_value",
				"int_field":       123,
				"timestamp_field": time.Now().UTC().Unix(),
			})
		assert.Error(t, err)
		checkParamError(t, err, "eventName")
	})

	t.Run("no time set", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockNewEvent(http.StatusOK, testCustomerID)

		err = client.NewEvent(
			testCustomerID, testEventName, time.Time{},
			map[string]interface{}{
				"field_name":      "some_value",
				"int_field":       123,
				"timestamp_field": time.Now().UTC().Unix(),
			})
		assert.NoError(t, err)
	})

	t.Run("customerIo error", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockNewEvent(http.StatusUnprocessableEntity, testCustomerID)

		err = client.NewEvent(
			testCustomerID, testEventName, time.Now().UTC(),
			map[string]interface{}{
				"field_name":      "some_value",
				"int_field":       123,
				"timestamp_field": time.Now().UTC().Unix(),
			})
		assert.Error(t, err)
	})
}

// ExampleClient_NewEvent example using NewEvent()
//
// See more examples in /examples/
func ExampleClient_NewEvent() {

	// Load the client
	client, err := newTestClient()
	if err != nil {
		fmt.Printf("error loading client: %s", err.Error())
		return
	}

	mockNewEvent(http.StatusOK, testCustomerID)

	// New event
	err = client.NewEvent(
		testCustomerID, testEventName, time.Now().UTC(),
		map[string]interface{}{
			"field_name":      "some_value",
			"int_field":       123,
			"timestamp_field": time.Now().UTC().Unix(),
		})
	if err != nil {
		fmt.Printf("error creating event: %s", err.Error())
		return
	}
	fmt.Printf("event created: %s", testEventName)
	// Output:event created: test_event
}

// BenchmarkClient_NewEvent benchmarks the method NewEvent()
func BenchmarkClient_NewEvent(b *testing.B) {
	client, _ := newTestClient()
	mockNewEvent(http.StatusOK, testCustomerID)
	data := map[string]interface{}{
		"field_name":      "some_value",
		"int_field":       123,
		"timestamp_field": time.Now().UTC().Unix(),
	}
	timestamp := time.Now().UTC()
	for i := 0; i < b.N; i++ {
		_ = client.NewEvent(testCustomerID, testEventName, timestamp, data)
	}
}

// TestClient_NewAnonymousEvent will test the method NewAnonymousEvent()
func TestClient_NewAnonymousEvent(t *testing.T) {
	// t.Parallel() (Cannot run in parallel - issues with overriding the mock client)

	t.Run("successful response", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockNewAnonymousEvent(http.StatusOK)

		err = client.NewAnonymousEvent(
			testEventName, time.Now().UTC(),
			map[string]interface{}{
				"field_name":      "some_value",
				"int_field":       123,
				"timestamp_field": time.Now().UTC().Unix(),
			})
		assert.NoError(t, err)
	})

	t.Run("missing event name", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockNewAnonymousEvent(http.StatusOK)

		err = client.NewAnonymousEvent(
			"", time.Now().UTC(),
			map[string]interface{}{
				"field_name":      "some_value",
				"int_field":       123,
				"timestamp_field": time.Now().UTC().Unix(),
			})
		assert.Error(t, err)
		checkParamError(t, err, "eventName")
	})

	t.Run("no time set", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockNewAnonymousEvent(http.StatusOK)

		err = client.NewAnonymousEvent(
			testEventName, time.Time{},
			map[string]interface{}{
				"field_name":      "some_value",
				"int_field":       123,
				"timestamp_field": time.Now().UTC().Unix(),
			})
		assert.NoError(t, err)
	})

	t.Run("customerIo error", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockNewAnonymousEvent(http.StatusUnprocessableEntity)

		err = client.NewEvent(
			testCustomerID, testEventName, time.Now().UTC(),
			map[string]interface{}{
				"field_name":      "some_value",
				"int_field":       123,
				"timestamp_field": time.Now().UTC().Unix(),
			})
		assert.Error(t, err)
	})
}

// ExampleClient_NewAnonymousEvent example using NewAnonymousEvent()
//
// See more examples in /examples/
func ExampleClient_NewAnonymousEvent() {

	// Load the client
	client, err := newTestClient()
	if err != nil {
		fmt.Printf("error loading client: %s", err.Error())
		return
	}

	mockNewAnonymousEvent(http.StatusOK)

	// New event
	err = client.NewAnonymousEvent(
		testEventName, time.Now().UTC(),
		map[string]interface{}{
			"field_name":      "some_value",
			"int_field":       123,
			"timestamp_field": time.Now().UTC().Unix(),
		})
	if err != nil {
		fmt.Printf("error creating event: %s", err.Error())
		return
	}
	fmt.Printf("event created: %s", testEventName)
	// Output:event created: test_event
}

// BenchmarkClient_NewAnonymousEvent benchmarks the method NewAnonymousEvent()
func BenchmarkClient_NewAnonymousEvent(b *testing.B) {
	client, _ := newTestClient()
	mockNewAnonymousEvent(http.StatusOK)
	data := map[string]interface{}{
		"field_name":      "some_value",
		"int_field":       123,
		"timestamp_field": time.Now().UTC().Unix(),
	}
	timestamp := time.Now().UTC()
	for i := 0; i < b.N; i++ {
		_ = client.NewAnonymousEvent(testEventName, timestamp, data)
	}
}

// TestClient_NewEventUsingInterface will test the method NewEventUsingInterface()
func TestClient_NewEventUsingInterface(t *testing.T) {
	// t.Parallel() (Cannot run in parallel - issues with overriding the mock client)

	t.Run("successful response", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockNewEvent(http.StatusOK, testCustomerID)

		someData := struct {
			FieldName      string `json:"field_name"`
			IntField       int    `json:"int_field"`
			TimestampField int64  `json:"timestamp_field"`
		}{
			FieldName:      "some_value",
			IntField:       123,
			TimestampField: time.Now().UTC().Unix(),
		}

		err = client.NewEventUsingInterface(
			testCustomerID, testEventName, time.Now().UTC(), someData,
		)
		assert.NoError(t, err)
	})

	t.Run("missing customer id", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockNewEvent(http.StatusOK, testCustomerID+"456")

		someData := struct {
			FieldName      string `json:"field_name"`
			IntField       int    `json:"int_field"`
			TimestampField int64  `json:"timestamp_field"`
		}{
			FieldName:      "some_value",
			IntField:       123,
			TimestampField: time.Now().UTC().Unix(),
		}

		err = client.NewEventUsingInterface(
			"", testEventName, time.Now().UTC(), someData,
		)
		assert.Error(t, err)
		checkParamError(t, err, "customerIDOrEmail")
	})

	t.Run("missing event name", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockNewEvent(http.StatusOK, testCustomerID)

		someData := struct {
			FieldName      string `json:"field_name"`
			IntField       int    `json:"int_field"`
			TimestampField int64  `json:"timestamp_field"`
		}{
			FieldName:      "some_value",
			IntField:       123,
			TimestampField: time.Now().UTC().Unix(),
		}

		err = client.NewEventUsingInterface(
			testCustomerID, "", time.Now().UTC(), someData,
		)
		assert.Error(t, err)
		checkParamError(t, err, "eventName")
	})

	t.Run("no time set", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockNewEvent(http.StatusOK, testCustomerID)

		someData := struct {
			FieldName      string `json:"field_name"`
			IntField       int    `json:"int_field"`
			TimestampField int64  `json:"timestamp_field"`
		}{
			FieldName:      "some_value",
			IntField:       123,
			TimestampField: time.Now().UTC().Unix(),
		}

		err = client.NewEventUsingInterface(
			testCustomerID, testEventName, time.Time{}, someData,
		)
		assert.NoError(t, err)
	})

	t.Run("customerIo error", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockNewEvent(http.StatusUnprocessableEntity, testCustomerID)

		someData := struct {
			FieldName      string `json:"field_name"`
			IntField       int    `json:"int_field"`
			TimestampField int64  `json:"timestamp_field"`
		}{
			FieldName:      "some_value",
			IntField:       123,
			TimestampField: time.Now().UTC().Unix(),
		}

		err = client.NewEventUsingInterface(
			testCustomerID, testEventName, time.Now().UTC(), someData,
		)
		assert.Error(t, err)
	})
}

// mockNewEvent is used for mocking the response
func mockNewEvent(statusCode int, customerID string) {
	httpmock.Reset()
	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%sapi/v1/customers/%s/events", testTrackingAPIURL, customerID),
		httpmock.NewStringResponder(
			statusCode, "",
		),
	)
}

// mockNewAnonymousEvent is used for mocking the response
func mockNewAnonymousEvent(statusCode int) {
	httpmock.Reset()
	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%sapi/v1/events", testTrackingAPIURL),
		httpmock.NewStringResponder(
			statusCode, "",
		),
	)
}
