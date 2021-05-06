package customerio

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// EmailRequest is the request structure for sending an email
type EmailRequest struct {
	AMPBody                 string                 `json:"amp_body,omitempty"`
	Attachments             map[string]string      `json:"attachments,omitempty"`
	BCC                     string                 `json:"bcc,omitempty"`
	Body                    string                 `json:"body,omitempty"`
	DisableMessageRetention *bool                  `json:"disable_message_retention,omitempty"`
	EnableTracking          *bool                  `json:"tracked,omitempty"`
	FakeBCC                 *bool                  `json:"fake_bcc,omitempty"`
	From                    string                 `json:"from,omitempty"`
	Headers                 map[string]string      `json:"headers,omitempty"`
	Identifiers             map[string]string      `json:"identifiers"`
	MessageData             map[string]interface{} `json:"message_data,omitempty"`
	PlaintextBody           string                 `json:"plaintext_body,omitempty"`
	Preheader               string                 `json:"preheader,omitempty"`
	QueueDraft              *bool                  `json:"queue_draft,omitempty"`
	ReplyTo                 string                 `json:"reply_to,omitempty"`
	SendToUnsubscribed      *bool                  `json:"send_to_unsubscribed,omitempty"`
	Subject                 string                 `json:"subject,omitempty"`
	To                      string                 `json:"to,omitempty"`
	TransactionalMessageID  string                 `json:"transactional_message_id,omitempty"`
}

// ErrAttachmentExists is the error message if the attachment already exists
var ErrAttachmentExists = errors.New("attachment with this name already exists")

// Attach will add a new file to the email
func (e *EmailRequest) Attach(name string, value io.Reader) error {
	if e.Attachments == nil {
		e.Attachments = map[string]string{}
	}
	if _, ok := e.Attachments[name]; ok {
		return ErrAttachmentExists
	}

	var buf bytes.Buffer
	enc := base64.NewEncoder(base64.StdEncoding, &buf)
	if _, err := io.Copy(enc, value); err != nil {
		return err
	}

	e.Attachments[name] = buf.String()
	return nil
}

// EmailResponse is the response from sending the email
type EmailResponse struct {
	TransactionalResponse
}

// TransactionalResponse  is a response to the send of a transactional message.
type TransactionalResponse struct {
	DeliveryID string    `json:"delivery_id"` // DeliveryID is a unique id for the given message.
	QueuedAt   time.Time `json:"queued_at"`   // QueuedAt is when the message was queued.
}

// UnmarshalJSON will unmarshall the response
func (t *TransactionalResponse) UnmarshalJSON(b []byte) (err error) {
	var r struct {
		DeliveryID string `json:"delivery_id"`
		QueuedAt   int64  `json:"queued_at"`
	}
	if err = json.Unmarshal(b, &r); err != nil {
		return
	}
	t.DeliveryID = r.DeliveryID
	t.QueuedAt = time.Unix(r.QueuedAt, 0)
	return
}

// TransactionalError is returned if a transactional message fails to send.
type TransactionalError struct {
	Err        string // Err is a more specific error message.
	StatusCode int    // StatusCode is the http status code for the error.
}

// Error with display the string error message
func (e *TransactionalError) Error() string {
	return e.Err
}

// SendEmail sends a single transactional email using the Customer.io transactional API
// See: https://customer.io/docs/api/#tag/Transactional
func (c *Client) SendEmail(emailRequest *EmailRequest) (*EmailResponse, error) {

	// Request cannot be nil, don't panic dude!
	if emailRequest == nil {
		return nil, ParamError{Param: "emailRequest"}
	}

	// If a template is set (advanced error checking)
	if len(emailRequest.TransactionalMessageID) > 0 {
		if emailRequest.To == "" {
			return nil, ParamError{Param: "emailTo"}
		} else if len(emailRequest.Identifiers) == 0 {
			return nil, ParamError{Param: "emailIdentifiers"}
		}
	} else { // NOT using a template
		if emailRequest.Body == "" {
			return nil, ParamError{Param: "emailBody"}
		} else if emailRequest.Subject == "" {
			return nil, ParamError{Param: "emailSubject"}
		} else if emailRequest.To == "" {
			return nil, ParamError{Param: "emailTo"}
		} else if emailRequest.From == "" {
			return nil, ParamError{Param: "emailFrom"}
		} else if len(emailRequest.Identifiers) == 0 {
			return nil, ParamError{Param: "emailIdentifiers"}
		}
	}

	// Attempt to send the email
	response, err := c.request(
		http.MethodPost,
		fmt.Sprintf("%s/v1/send/email", c.options.apiURL),
		emailRequest,
	)
	if err != nil {
		return nil, err
	}

	// Unmarshal the response
	var r EmailResponse
	if err = json.Unmarshal(response.Body, &r); err != nil {
		return nil, err
	}

	return &r, nil
}
