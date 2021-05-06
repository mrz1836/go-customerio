package customerio

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

// TestClient_UpdateCustomer will test the method UpdateCustomer()
func TestClient_UpdateCustomer(t *testing.T) {
	// t.Parallel() (Cannot run in parallel - issues with overriding the mock client)

	t.Run("successful response (ID)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockUpdateCustomer(http.StatusOK, testCustomerID)

		err = client.UpdateCustomer(testCustomerID, map[string]interface{}{
			"created_at": time.Now().Unix(),
			"email":      testCustomerEmail,
			"first_name": "Bob",
			"plan":       "basic",
		})
		assert.NoError(t, err)
	})

	t.Run("missing customer id", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockUpdateCustomer(http.StatusOK, testCustomerID+"456")

		err = client.UpdateCustomer("", map[string]interface{}{
			"created_at": time.Now().Unix(),
			"email":      testCustomerEmail,
			"first_name": "Bob",
			"plan":       "basic",
		})
		assert.Error(t, err)
		checkParamError(t, err, "customerIDOrEmail")
	})

	t.Run("customerIo error", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockUpdateCustomer(http.StatusUnprocessableEntity, testCustomerID)

		err = client.UpdateCustomer(testCustomerID, map[string]interface{}{
			"created_at": time.Now().Unix(),
			"email":      testCustomerEmail,
			"first_name": "Bob",
			"plan":       "basic",
		})
		assert.Error(t, err)
	})

	t.Run("successful response (Email)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockUpdateCustomerUsingEmail(http.StatusOK, testCustomerEmail)

		err = client.UpdateCustomer(testCustomerEmail, map[string]interface{}{
			"created_at": time.Now().Unix(),
			"email":      testCustomerEmail,
			"first_name": "Bob",
			"plan":       "basic",
		})
		assert.NoError(t, err)
	})
}

// ExampleClient_UpdateCustomer example using UpdateCustomer()
//
// See more examples in /examples/
func ExampleClient_UpdateCustomer() {

	// Load the client
	client, err := newTestClient()
	if err != nil {
		fmt.Printf("error loading client: %s", err.Error())
		return
	}

	mockUpdateCustomer(http.StatusOK, testCustomerID)

	// Update customer
	err = client.UpdateCustomer(testCustomerID, map[string]interface{}{
		"created_at": time.Now().Unix(),
		"email":      testCustomerEmail,
		"first_name": "Bob",
		"plan":       "basic",
	})
	if err != nil {
		fmt.Printf("error updating customer: " + err.Error())
		return
	}
	fmt.Printf("customer updated: %s", testCustomerID)
	// Output:customer updated: 123
}

// BenchmarkClient_UpdateCustomer benchmarks the method UpdateCustomer()
func BenchmarkClient_UpdateCustomer(b *testing.B) {
	client, _ := newTestClient()
	mockUpdateCustomer(http.StatusOK, testCustomerID)
	attributes := map[string]interface{}{
		"created_at": time.Now().Unix(),
		"email":      testCustomerEmail,
		"first_name": "Bob",
		"plan":       "basic",
	}
	for i := 0; i < b.N; i++ {
		_ = client.UpdateCustomer(testCustomerID, attributes)
	}
}

// TestClient_DeleteCustomer will test the method DeleteCustomer()
func TestClient_DeleteCustomer(t *testing.T) {
	// t.Parallel() (Cannot run in parallel - issues with overriding the mock client)

	t.Run("successful response (ID)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockDeleteCustomer(http.StatusOK, testCustomerID)

		err = client.DeleteCustomer(testCustomerID)
		assert.NoError(t, err)
	})

	t.Run("missing customer id", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockDeleteCustomer(http.StatusOK, testCustomerID+"456")

		err = client.DeleteCustomer("")
		assert.Error(t, err)
		checkParamError(t, err, "customerIDOrEmail")
	})

	t.Run("customerIo error", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockDeleteCustomer(http.StatusUnprocessableEntity, testCustomerID)

		err = client.DeleteCustomer(testCustomerID)
		assert.Error(t, err)
	})
}

// ExampleClient_DeleteCustomer example using DeleteCustomer()
//
// See more examples in /examples/
func ExampleClient_DeleteCustomer() {

	// Load the client
	client, err := newTestClient()
	if err != nil {
		fmt.Printf("error loading client: %s", err.Error())
		return
	}

	mockDeleteCustomer(http.StatusOK, testCustomerID)

	// Delete customer
	err = client.DeleteCustomer(testCustomerID)
	if err != nil {
		fmt.Printf("error deleting customer: " + err.Error())
		return
	}
	fmt.Printf("customer deleted: %s", testCustomerID)
	// Output:customer deleted: 123
}

// BenchmarkClient_DeleteCustomer benchmarks the method DeleteCustomer()
func BenchmarkClient_DeleteCustomer(b *testing.B) {
	client, _ := newTestClient()
	mockDeleteCustomer(http.StatusOK, testCustomerID)
	for i := 0; i < b.N; i++ {
		_ = client.DeleteCustomer(testCustomerID)
	}
}

// TestClient_UpdateDevice will test the method UpdateDevice()
func TestClient_UpdateDevice(t *testing.T) {
	// t.Parallel() (Cannot run in parallel - issues with overriding the mock client)

	t.Run("successful response (ID)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockUpdateDevice(http.StatusOK, testCustomerID)

		err = client.UpdateDevice(testCustomerID, &Device{
			ID:       testDeviceID,
			LastUsed: time.Now().UTC().Unix(),
			Platform: PlatformIOs,
		})
		assert.NoError(t, err)
	})

	t.Run("missing customer id", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockUpdateDevice(http.StatusOK, testCustomerID+"456")

		err = client.UpdateDevice("", &Device{
			ID:       testDeviceID,
			LastUsed: time.Now().UTC().Unix(),
			Platform: PlatformIOs,
		})
		assert.Error(t, err)
		checkParamError(t, err, "customerIDOrEmail")
	})

	t.Run("missing device", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockUpdateDevice(http.StatusOK, testCustomerID)

		err = client.UpdateDevice(testCustomerID, nil)
		assert.Error(t, err)
		checkParamError(t, err, "device")
	})

	t.Run("missing device id", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockUpdateDevice(http.StatusOK, testCustomerID)

		err = client.UpdateDevice(testCustomerID, &Device{
			ID:       "",
			LastUsed: time.Now().UTC().Unix(),
			Platform: PlatformIOs,
		})
		assert.Error(t, err)
		checkParamError(t, err, "deviceID")
	})

	t.Run("invalid platform", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockUpdateDevice(http.StatusOK, testCustomerID)

		err = client.UpdateDevice(testCustomerID, &Device{
			ID:       testDeviceID,
			LastUsed: time.Now().UTC().Unix(),
			Platform: "123",
		})
		assert.Error(t, err)
		checkParamError(t, err, "devicePlatform")
	})

	t.Run("customerIo error", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockUpdateDevice(http.StatusUnprocessableEntity, testCustomerID)

		err = client.UpdateDevice(testCustomerID, &Device{
			ID:       testDeviceID,
			LastUsed: time.Now().UTC().Unix(),
			Platform: PlatformIOs,
		})
		assert.Error(t, err)
	})
}

// ExampleClient_UpdateDevice example using UpdateDevice()
//
// See more examples in /examples/
func ExampleClient_UpdateDevice() {

	// Load the client
	client, err := newTestClient()
	if err != nil {
		fmt.Printf("error loading client: %s", err.Error())
		return
	}

	mockUpdateDevice(http.StatusOK, testCustomerID)

	// Delete customer
	err = client.UpdateDevice(testCustomerID, &Device{
		ID:       testDeviceID,
		LastUsed: time.Now().UTC().Unix(),
		Platform: PlatformIOs,
	})
	if err != nil {
		fmt.Printf("error updating device: " + err.Error())
		return
	}
	fmt.Printf("device updated: %s", testDeviceID)
	// Output:device updated: abcdefghijklmnopqrstuvwxyz
}

// BenchmarkClient_UpdateDevice benchmarks the method UpdateDevice()
func BenchmarkClient_UpdateDevice(b *testing.B) {
	client, _ := newTestClient()
	mockUpdateDevice(http.StatusOK, testCustomerID)
	timestamp := time.Now().UTC().Unix()
	device := &Device{
		ID:       testDeviceID,
		LastUsed: timestamp,
		Platform: PlatformIOs,
	}
	for i := 0; i < b.N; i++ {
		_ = client.UpdateDevice(testCustomerID, device)
	}
}

// TestClient_DeleteDevice will test the method DeleteDevice()
func TestClient_DeleteDevice(t *testing.T) {
	// t.Parallel() (Cannot run in parallel - issues with overriding the mock client)

	t.Run("successful response (ID)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockDeleteDevice(http.StatusOK, testCustomerID, testDeviceID)

		err = client.DeleteDevice(testCustomerID, testDeviceID)
		assert.NoError(t, err)
	})

	t.Run("missing customer id", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockDeleteDevice(http.StatusOK, testCustomerID+"456", testDeviceID+"456")

		err = client.DeleteDevice("", testDeviceID)
		assert.Error(t, err)
		checkParamError(t, err, "customerIDOrEmail")
	})

	t.Run("missing device id", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockDeleteDevice(http.StatusOK, testCustomerID, testDeviceID)

		err = client.DeleteDevice(testCustomerID, "")
		assert.Error(t, err)
		checkParamError(t, err, "deviceID")
	})

	t.Run("customerIo error", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockDeleteDevice(http.StatusUnprocessableEntity, testCustomerID, testDeviceID)

		err = client.DeleteDevice(testCustomerID, testDeviceID)
		assert.Error(t, err)
	})
}

// ExampleClient_DeleteDevice example using DeleteDevice()
//
// See more examples in /examples/
func ExampleClient_DeleteDevice() {

	// Load the client
	client, err := newTestClient()
	if err != nil {
		fmt.Printf("error loading client: %s", err.Error())
		return
	}

	mockDeleteDevice(http.StatusOK, testCustomerID, testDeviceID)

	// Delete device
	err = client.DeleteDevice(testCustomerID, testDeviceID)
	if err != nil {
		fmt.Printf("error deleting device: " + err.Error())
		return
	}
	fmt.Printf("device deleted: %s", testDeviceID)
	// Output:device deleted: abcdefghijklmnopqrstuvwxyz
}

// BenchmarkClient_DeleteDevice benchmarks the method DeleteDevice()
func BenchmarkClient_DeleteDevice(b *testing.B) {
	client, _ := newTestClient()
	mockDeleteDevice(http.StatusOK, testCustomerID, testDeviceID)
	for i := 0; i < b.N; i++ {
		_ = client.DeleteDevice(testCustomerID, testDeviceID)
	}
}

// mockUpdateCustomer is used for mocking the response
func mockUpdateCustomer(statusCode int, customerID string) {
	httpmock.Reset()
	httpmock.RegisterResponder(http.MethodPut, fmt.Sprintf("%sapi/v1/customers/%s", testTrackingAPIURL, customerID),
		httpmock.NewStringResponder(
			statusCode, "",
		),
	)
}

// mockUpdateCustomerUsingEmail is used for mocking the response
func mockUpdateCustomerUsingEmail(statusCode int, customerEmail string) {
	httpmock.Reset()
	httpmock.RegisterResponder(http.MethodPut, fmt.Sprintf("%sapi/v1/customers/%s", testTrackingAPIURL, customerEmail),
		httpmock.NewStringResponder(
			statusCode, "",
		),
	)
}

// mockDeleteCustomer is used for mocking the response
func mockDeleteCustomer(statusCode int, customerID string) {
	httpmock.Reset()
	httpmock.RegisterResponder(http.MethodDelete, fmt.Sprintf("%sapi/v1/customers/%s", testTrackingAPIURL, customerID),
		httpmock.NewStringResponder(
			statusCode, "",
		),
	)
}

// mockUpdateDevice is used for mocking the response
func mockUpdateDevice(statusCode int, customerID string) {
	httpmock.Reset()
	httpmock.RegisterResponder(http.MethodPut, fmt.Sprintf("%sapi/v1/customers/%s/devices", testTrackingAPIURL, customerID),
		httpmock.NewStringResponder(
			statusCode, "",
		),
	)
}

// mockDeleteDevice is used for mocking the response
func mockDeleteDevice(statusCode int, customerID, deviceID string) {
	httpmock.Reset()
	httpmock.RegisterResponder(
		http.MethodDelete,
		fmt.Sprintf(
			"%sapi/v1/customers/%s/devices/%s", testTrackingAPIURL, customerID, deviceID,
		),
		httpmock.NewStringResponder(
			statusCode, "",
		),
	)
}

// checkParamError will ensure the error is a Param error with the correct field
func checkParamError(t *testing.T, err error, param string) {
	if err == nil {
		t.Error("expected error")
		return
	}
	var pErr ParamError
	if !errors.As(err, &pErr) {
		t.Error("expected ParamError")
	} else {
		pe, ok := err.(ParamError)
		if !ok {
			t.Error("expected ParamError")
		}
		if pe.Param != param {
			t.Errorf("expected %s got %s", param, pe.Param)
		}
	}
}
