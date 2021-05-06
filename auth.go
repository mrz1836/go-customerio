package customerio

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// RegionInfo is the response from FindRegion
type RegionInfo struct {
	DataCenter    string `json:"data_center"`
	EnvironmentID uint64 `json:"environment_id"`
	URL           string `json:"url"`
}

/*
// Example Response
{
	"url": "https://track.customer.io",
	"data_center": "us",
	"environment_id": 3
}
*/

// FindRegion will return the url, data center and environment id
// See: https://customer.io/docs/api/#tag/trackAuth
func (c *Client) FindRegion() (region *RegionInfo, err error) {
	var resp StandardResponse
	if resp, err = c.request(
		http.MethodGet,
		fmt.Sprintf("%s/api/v1/accounts/region", c.options.trackURL),
		nil,
	); err != nil {
		return
	}
	err = json.Unmarshal(resp.Body, &region)
	return
}

// TestAuth will test your current Tracking API credentials
// See: https://customer.io/docs/api/#tag/trackAuth
func (c *Client) TestAuth() error {
	_, err := c.request(
		http.MethodGet,
		fmt.Sprintf("%s/auth", c.options.trackURL),
		nil,
	)
	return err
}
