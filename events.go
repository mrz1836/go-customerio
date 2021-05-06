package customerio

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// NewEvent will create a new event for the supplied customer
// See: https://customer.io/docs/api/#tag/Track-Events
// AKA: Track()
// Only use "email" if the workspace is setup to use email instead of ID
// Use "timestamp" to send events in the past. If not set, it will use Now().UTC()
func (c *Client) NewEvent(customerIDOrEmail string, eventName string, timestamp time.Time,
	data map[string]interface{}) error {
	if customerIDOrEmail == "" {
		return ParamError{Param: "customerIDOrEmail"}
	}
	if eventName == "" {
		return ParamError{Param: "eventName"}
	}
	if timestamp.IsZero() {
		timestamp = time.Now().UTC()
	}
	_, err := c.request(
		http.MethodPost,
		fmt.Sprintf("%s/api/v1/customers/%s/events", c.options.trackURL, url.PathEscape(customerIDOrEmail)),
		map[string]interface{}{
			"data":      data,
			"name":      eventName,
			"timestamp": timestamp.Unix(),
			// "type":      "",  (set to Page for a page view) // todo: add support for this feature
		},
	)
	return err
}

// NewAnonymousEvent will create a new event for the anonymous visitor
// See: https://customer.io/docs/api/#operation/trackAnonymous
// AKA: TrackAnonymous()
// Use "timestamp" to send events in the past. If not set, it will use Now().UTC()
func (c *Client) NewAnonymousEvent(eventName string, timestamp time.Time, data map[string]interface{}) error {
	if eventName == "" {
		return ParamError{Param: "eventName"}
	}
	if timestamp.IsZero() {
		timestamp = time.Now().UTC()
	}
	_, err := c.request(
		http.MethodPost,
		fmt.Sprintf("%s/api/v1/events", c.options.trackURL),
		map[string]interface{}{
			"data":      data,
			"name":      eventName,
			"timestamp": timestamp.Unix(),
			// "type":      "",  (set to Page for a page view) // todo: add support for this feature
		},
	)
	return err
}
