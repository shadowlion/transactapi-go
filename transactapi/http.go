package transactapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Function basePostRequest is a generic http POST request that takes an endpoint and request body
// parameters.
func basePostRequest[Response any](c *Client, endpoint string, req interface{}) Response {
	var res Response

	b, err := json.Marshal(req)
	if err != nil {
		log.Fatalf("Error marshalling json for %s: %s", endpoint, err)
	}

	url := fmt.Sprintf("%s%s", c.baseUrl(), endpoint)
	byteReader := bytes.NewReader(b)
	resp, err := c.HttpClient.Post(url, "application/json", byteReader)
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

// Function basePutRequest is a generic http PUT request that takes an endpoint and request body
// parameters.
func basePutRequest[Response any](c *Client, endpoint string, req interface{}) Response {
	var res Response

	b, err := json.Marshal(req)
	if err != nil {
		log.Fatalf("Error marshalling json for %s: %s", endpoint, err)
	}

	url := fmt.Sprintf("%s%s", c.baseUrl(), endpoint)
	byteReader := bytes.NewReader(b)

	// TODO: Figure out why the following doesn't work:
	request, err := http.NewRequest(http.MethodPut, url, byteReader)
	if err != nil {
		log.Fatalf("Error forming request: %s", err)
	}

	resp, err := c.HttpClient.Do(request)
	if err != nil {
		log.Fatalf("Error sending request to %s: %s", endpoint, err)
	}
	// End TODO

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
