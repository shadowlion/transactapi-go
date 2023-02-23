package transactapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
// func (c *Client) baseUrl() string {
// 	return fmt.Sprintf("https://%s.norcapsecurities.com/tapiv3/index.php/v3", c.getPrefix())
// }

// Function basePostRequest will run a generic http post request.
func basePostRequest[Request, Response any](c *Client, endpoint string, req *Request) Response {
	var res Response

	b, err := json.Marshal(req)
	if err != nil {
		log.Fatalf("Error marshalling json for %s: %s", endpoint, err)
	}

	url := fmt.Sprintf("%s%s", c.baseUrl(), endpoint)
	byteReader := bytes.NewReader(b)
	httpRequest, err := http.NewRequest(http.MethodPost, url, byteReader)
	if err != nil {
		log.Fatalf("Error forming request for %s: %s", endpoint, err)
	}

	resp, err := c.HttpClient.Do(httpRequest)
	if err != nil {
		log.Fatalf("Error sending request to %s: %s", endpoint, err)
	}

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("Error reading response body for %s: %s", endpoint, err)
	}

	if err := json.Unmarshal(body, &res); err != nil {
		log.Fatalf("Error unmarshalling response body for %s: %s", endpoint, err)
	}

	return res
}
