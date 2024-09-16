package customerio

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// EmailTestSuite
type EmailTestSuite struct {
	suite.Suite
	emailNoTemplate   *EmailRequest
	emailWithTemplate *EmailRequest
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *EmailTestSuite) SetupTest() {
	// Start the email request (no template)
	suite.emailNoTemplate = &EmailRequest{
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
	suite.emailWithTemplate = &EmailRequest{
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
}

// TestClient_SendEmail will test the method SendEmail()
func (suite *EmailTestSuite) TestClient_SendEmail() {
	// t.Parallel() (Cannot run in parallel - issues with overriding the mock client)

	suite.T().Run("successful response (no template)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		suite.SetupTest()
		mockSendEmail(http.StatusOK)

		var resp *EmailResponse
		resp, err = client.SendEmail(suite.emailNoTemplate)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})

	suite.T().Run("successful response (with template)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		suite.SetupTest()
		mockSendEmail(http.StatusOK)

		var resp *EmailResponse
		resp, err = client.SendEmail(suite.emailWithTemplate)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})

	suite.T().Run("missing email request", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		suite.SetupTest()
		mockSendEmail(http.StatusOK)

		var resp *EmailResponse
		resp, err = client.SendEmail(nil)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	suite.T().Run("template - missing to", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		suite.SetupTest()
		mockSendEmail(http.StatusOK)

		suite.emailWithTemplate.To = ""

		var resp *EmailResponse
		resp, err = client.SendEmail(suite.emailWithTemplate)
		assert.Error(t, err)
		assert.Nil(t, resp)
		checkParamError(t, err, "emailTo")
	})

	suite.T().Run("template - missing identifiers", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		suite.SetupTest()
		mockSendEmail(http.StatusOK)

		suite.emailWithTemplate.Identifiers = nil

		var resp *EmailResponse
		resp, err = client.SendEmail(suite.emailWithTemplate)
		assert.Error(t, err)
		assert.Nil(t, resp)
		checkParamError(t, err, "emailIdentifiers")
	})

	suite.T().Run("no template - missing to", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		suite.SetupTest()
		mockSendEmail(http.StatusOK)

		suite.emailNoTemplate.To = ""

		var resp *EmailResponse
		resp, err = client.SendEmail(suite.emailNoTemplate)
		assert.Error(t, err)
		assert.Nil(t, resp)
		checkParamError(t, err, "emailTo")
	})

	suite.T().Run("no template - missing identifiers", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		suite.SetupTest()
		mockSendEmail(http.StatusOK)

		suite.emailNoTemplate.Identifiers = nil

		var resp *EmailResponse
		resp, err = client.SendEmail(suite.emailNoTemplate)
		assert.Error(t, err)
		assert.Nil(t, resp)
		checkParamError(t, err, "emailIdentifiers")
	})

	suite.T().Run("no template - missing body", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		suite.SetupTest()
		mockSendEmail(http.StatusOK)

		suite.emailNoTemplate.Body = ""

		var resp *EmailResponse
		resp, err = client.SendEmail(suite.emailNoTemplate)
		assert.Error(t, err)
		assert.Nil(t, resp)
		checkParamError(t, err, "emailBody")
	})

	suite.T().Run("no template - missing subject", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		suite.SetupTest()
		mockSendEmail(http.StatusOK)

		suite.emailNoTemplate.Subject = ""

		var resp *EmailResponse
		resp, err = client.SendEmail(suite.emailNoTemplate)
		assert.Error(t, err)
		assert.Nil(t, resp)
		checkParamError(t, err, "emailSubject")
	})

	suite.T().Run("no template - missing from", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		suite.SetupTest()
		mockSendEmail(http.StatusOK)

		suite.emailNoTemplate.From = ""

		var resp *EmailResponse
		resp, err = client.SendEmail(suite.emailNoTemplate)
		assert.Error(t, err)
		assert.Nil(t, resp)
		checkParamError(t, err, "emailFrom")
	})

	suite.T().Run("email error", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		suite.SetupTest()
		mockSendEmail(http.StatusUnprocessableEntity)

		var resp *EmailResponse
		resp, err = client.SendEmail(suite.emailNoTemplate)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

// TestEmailTestSuite starts the suite testing
func TestEmailTestSuite(t *testing.T) {
	suite.Run(t, new(EmailTestSuite))
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
		fmt.Printf("error sending email: %s", err.Error())
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
