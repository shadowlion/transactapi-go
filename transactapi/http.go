package transactapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
)

// Function basePostRequest will run a generic http post request.
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
