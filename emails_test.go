package customerio

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

// TestClient_SendEmail will test the method SendEmail()
func TestClient_SendEmail(t *testing.T) {
	// t.Parallel() (Cannot run in parallel - issues with overriding the mock client)

	// Start the email request (no template)
	emailRequestNoTemplate := &EmailRequest{
		Body:        "This is an example body!",
		From:        "noreply@example.com",
		Identifiers: map[string]string{"id": "123"},
		MessageData: map[string]interface{}{
			"name": "Person",
			"items": map[string]interface{}{
				"name":  "shoes",
				"price": "59.99",
			},
			"products": []interface{}{},
		},
		PlaintextBody: "This is an example body!",
		ReplyTo:       "noreply@example.com",
		Subject:       "Customer io test email",
		To:            "bob@example.com",
	}

	// Start the email request (with template)
	emailRequestWithTemplate := &EmailRequest{
		TransactionalMessageID: "123",
		Identifiers:            map[string]string{"id": "123"},
		MessageData: map[string]interface{}{
			"name": "Person",
			"items": map[string]interface{}{
				"name":  "shoes",
				"price": "59.99",
			},
			"products": []interface{}{},
		},
		To: "bob@example.com",
	}

	t.Run("successful response (no template)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockSendEmail(http.StatusOK)

		var resp *EmailResponse
		resp, err = client.SendEmail(emailRequestNoTemplate)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("successful response (with template)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockSendEmail(http.StatusOK)

		var resp *EmailResponse
		resp, err = client.SendEmail(emailRequestWithTemplate)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("missing email request", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockSendEmail(http.StatusOK)

		var resp *EmailResponse
		resp, err = client.SendEmail(nil)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("template - missing to", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockSendEmail(http.StatusOK)

		emailRequestWithTemplate.To = ""

		var resp *EmailResponse
		resp, err = client.SendEmail(emailRequestWithTemplate)
		assert.Error(t, err)
		assert.Nil(t, resp)
		checkParamError(t, err, "emailTo")
	})

	t.Run("template - missing identifiers", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockSendEmail(http.StatusOK)

		emailRequestWithTemplate.Identifiers = nil

		var resp *EmailResponse
		resp, err = client.SendEmail(emailRequestWithTemplate)
		assert.Error(t, err)
		assert.Nil(t, resp)
		checkParamError(t, err, "emailIdentifiers")
	})

	t.Run("no template - missing to", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockSendEmail(http.StatusOK)

		emailRequestNoTemplate.To = ""

		var resp *EmailResponse
		resp, err = client.SendEmail(emailRequestNoTemplate)
		assert.Error(t, err)
		assert.Nil(t, resp)
		checkParamError(t, err, "emailTo")
	})

	t.Run("no template - missing identifiers", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockSendEmail(http.StatusOK)

		emailRequestNoTemplate.Identifiers = nil

		var resp *EmailResponse
		resp, err = client.SendEmail(emailRequestNoTemplate)
		assert.Error(t, err)
		assert.Nil(t, resp)
		checkParamError(t, err, "emailIdentifiers")
	})

	t.Run("no template - missing body", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockSendEmail(http.StatusOK)

		emailRequestNoTemplate.Body = ""

		var resp *EmailResponse
		resp, err = client.SendEmail(emailRequestNoTemplate)
		assert.Error(t, err)
		assert.Nil(t, resp)
		checkParamError(t, err, "emailBody")
	})

	t.Run("no template - missing subject", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockSendEmail(http.StatusOK)

		emailRequestNoTemplate.Subject = ""

		var resp *EmailResponse
		resp, err = client.SendEmail(emailRequestNoTemplate)
		assert.Error(t, err)
		assert.Nil(t, resp)
		checkParamError(t, err, "emailSubject")
	})

	t.Run("no template - missing from", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockSendEmail(http.StatusOK)

		emailRequestNoTemplate.From = ""

		var resp *EmailResponse
		resp, err = client.SendEmail(emailRequestNoTemplate)
		assert.Error(t, err)
		assert.Nil(t, resp)
		checkParamError(t, err, "emailFrom")
	})

	t.Run("email error", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		mockSendEmail(http.StatusUnprocessableEntity)

		var resp *EmailResponse
		resp, err = client.SendEmail(emailRequestWithTemplate)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

// ExampleClient_SendEmail example using SendEmail()
//
// See more examples in /examples/
func ExampleClient_SendEmail() {

	// Load the client
	client, err := newTestClient()
	if err != nil {
		fmt.Printf("error loading client: %s", err.Error())
		return
	}

	mockSendEmail(http.StatusOK)

	// Start the email request (with template)
	emailRequestWithTemplate := &EmailRequest{
		TransactionalMessageID: "123",
		Identifiers:            map[string]string{"id": "123"},
		MessageData: map[string]interface{}{
			"name": "Person",
			"items": map[string]interface{}{
				"name":  "shoes",
				"price": "59.99",
			},
			"products": []interface{}{},
		},
		To: "bob@example.com",
	}

	// Send email
	_, err = client.SendEmail(emailRequestWithTemplate)
	if err != nil {
		fmt.Printf("error sending email: " + err.Error())
		return
	}
	fmt.Printf("email sent to: %s", emailRequestWithTemplate.To)
	// Output:email sent to: bob@example.com
}

// BenchmarkClient_SendEmail benchmarks the method SendEmail()
func BenchmarkClient_SendEmail(b *testing.B) {
	client, _ := newTestClient()
	mockSendEmail(http.StatusOK)
	emailRequestWithTemplate := &EmailRequest{
		TransactionalMessageID: "123",
		Identifiers:            map[string]string{"id": "123"},
		MessageData: map[string]interface{}{
			"name": "Person",
			"items": map[string]interface{}{
				"name":  "shoes",
				"price": "59.99",
			},
			"products": []interface{}{},
		},
		To: "bob@example.com",
	}
	for i := 0; i < b.N; i++ {
		_, _ = client.SendEmail(emailRequestWithTemplate)
	}
}

// mockSendEmail is used for mocking the response
func mockSendEmail(statusCode int) {
	httpmock.Reset()
	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%sv1/send/email", testAppAPIURL),
		httpmock.NewStringResponder(
			statusCode, `{"delivery_id": "1234567890","queued_at": 1620313799}`,
		),
	)
}
