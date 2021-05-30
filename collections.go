package customerio

import (
	"fmt"
	"net/http"
)

// UpdateCollection will create or update a collection with raw data
// See: https://customer.io/docs/api/#operation/addCollection
// See: https://customer.io/docs/api/#operation/updateCollection
//
// The name of the collection. This is how you'll reference your collection in message.
// Updating the data or url for your collection fully replaces the contents of the collection.
// Data example: {"data":[{"property1":null,"property2":null}]}}
func (c *Client) UpdateCollection(collectionID, collectionName string, items []map[string]interface{}) error {
	if collectionName == "" {
		return ParamError{Param: "collectionName"}
	}
	/*
		if len(data) > 56000 {   // todo: is there a limit?
			return errors.New("collection body size limited to 56000")
		}
	*/

	// Create or Update (if id is given)
	var err error
	if len(collectionID) == 0 {
		_, err = c.request(
			http.MethodPost,
			fmt.Sprintf("%s/v1/api/collections", c.options.betaURL),
			map[string]interface{}{
				"data": items,
				"name": collectionName,
			},
		)
	} else {
		_, err = c.request(
			http.MethodPut,
			fmt.Sprintf("%s/v1/api/collections/%s", c.options.betaURL, collectionID),
			map[string]interface{}{
				"data": items,
				"name": collectionName,
			},
		)
	}

	return err
}

// UpdateCollectionViaURL will create or update a collection using a URL to a JSON file
// See: https://customer.io/docs/api/#operation/addCollection
// See: https://customer.io/docs/api/#operation/updateCollection
//
// The name of the collection. This is how you'll reference your collection in message.
//
// The URL where your data is stored.
// If your URL does not include a Content-Type, Customer.io assumes your data is JSON.
// This URL can also be a google sheet that you've shared with cio_share@customer.io.
// Updating the data or url for your collection fully replaces the contents of the collection.
func (c *Client) UpdateCollectionViaURL(collectionID, collectionName string, jsonURL string) error {
	if collectionName == "" {
		return ParamError{Param: "collectionName"}
	}

	// Create or Update (if id is given)
	var err error
	if len(collectionID) == 0 {
		_, err = c.request(
			http.MethodPost,
			fmt.Sprintf("%s/v1/api/collections", c.options.betaURL),
			map[string]interface{}{
				"url":  jsonURL,
				"name": collectionName,
			},
		)
	} else {
		_, err = c.request(
			http.MethodPut,
			fmt.Sprintf("%s/v1/api/collections/%s", c.options.betaURL, collectionID),
			map[string]interface{}{
				"url":  jsonURL,
				"name": collectionName,
			},
		)
	}

	return err
}
