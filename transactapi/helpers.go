package transactapi

import (
	"fmt"
)

// Function getPrefix returns a string based on a conditional sandbox environment
func (c *Client) getPrefix() string {
	if c.Sandbox {
		return "api-sandboxdash"
	}
	return "api"
}

// Function baseUrl returns full api url, conditionally checking the prefix based on a possible
// sandbox environment.
func (c *Client) baseUrl() string {
	return fmt.Sprintf("https://%s.norcapsecurities.com/tapiv3/index.php/v3", c.getPrefix())
}
