package customerio

import (
	"fmt"
	"net/http"
	"net/url"
)

// UpdateCustomer will add/update a customer and sets their attributes
// If not found, a customer will be created. If found, the attributes will be updated
// See: https://customer.io/docs/api/#operation/identify
// AKA: Identify()
// Only use "email" if the workspace is setup to use email instead of ID
func (c *Client) UpdateCustomer(customerIDOrEmail string, attributes map[string]interface{}) error {
	if customerIDOrEmail == "" {
		return ParamError{Param: "customerIDOrEmail"}
	}
	_, err := c.request(
		http.MethodPut,
		fmt.Sprintf("%s/api/v1/customers/%s", c.options.trackURL, url.PathEscape(customerIDOrEmail)),
		attributes,
	)
	return err
}

// DeleteCustomer will remove a customer given their id or email
// If not found, a customer will be created. If found, the attributes will be updated
// See: https://customer.io/docs/api/#operation/delete
// Only use "email" if the workspace is setup to use email instead of ID
func (c *Client) DeleteCustomer(customerIDOrEmail string) error {
	if customerIDOrEmail == "" {
		return ParamError{Param: "customerIDOrEmail"}
	}
	_, err := c.request(
		http.MethodDelete,
		fmt.Sprintf("%s/api/v1/customers/%s", c.options.trackURL, url.PathEscape(customerIDOrEmail)),
		nil,
	)
	return err
}

// UpdateDevice will add/update a customer's device
// If not found, a device will be created. If found, the attributes will be updated
// See: https://customer.io/docs/api/#operation/add_device
// Only use "email" if the workspace is setup to use email instead of ID
func (c *Client) UpdateDevice(customerIDOrEmail string, device *Device) error {
	if customerIDOrEmail == "" {
		return ParamError{Param: "customerIDOrEmail"}
	}
	if device == nil {
		return ParamError{Param: "device"}
	} else if device.ID == "" {
		return ParamError{Param: "deviceID"}
	} else if !acceptedPlatforms(device.Platform) {
		return ParamError{Param: "devicePlatform"}
	}
	_, err := c.request(
		http.MethodPut,
		fmt.Sprintf("%s/api/v1/customers/%s/devices", c.options.trackURL, url.PathEscape(customerIDOrEmail)),
		map[string]interface{}{
			"device": device,
		},
	)
	return err
}

// DeleteDevice will remove a customer's device
// See: https://customer.io/docs/api/#operation/delete_device
// Only use "email" if the workspace is setup to use email instead of ID
func (c *Client) DeleteDevice(customerIDOrEmail, deviceID string) error {
	if customerIDOrEmail == "" {
		return ParamError{Param: "customerIDOrEmail"}
	}
	if deviceID == "" {
		return ParamError{Param: "deviceID"}
	}
	_, err := c.request(
		http.MethodDelete,
		fmt.Sprintf(
			"%s/api/v1/customers/%s/devices/%s",
			c.options.trackURL,
			url.PathEscape(customerIDOrEmail),
			url.PathEscape(deviceID),
		),
		nil,
	)
	return err
}
