package transactapi

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// makeRequestBody simply checks if we're running a GET request. Otherwise, we
// return a request with a []byte body payload attached.
func makeRequestBody(method, url string, requestBody []byte) (*http.Request, error) {
	switch method {
	case http.MethodGet:
		return http.NewRequest(method, url, nil)
	default:
		return http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	}
}

// do attaches a `Content-Type` header to the request before sending.
func do(client *http.Client, req *http.Request) (*http.Response, error) {
	req.Header.Set("Content-Type", "application/json")
	return client.Do(req)
}

// parseResponse conforms the response body to one of three elements: an okay response, the
// standard error response given in the documentation, or an api error returned from Transact API
// (contained within the error handling code).
func parseResponse[OkResponse interface{}](resp *http.Response) (*OkResponse, error) {
	defer resp.Body.Close()

	// Since resp.Body gets used up if we conform it to a struct, we first need to
	// copy it to a different variable to be "re-used" (for the error struct handling)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		var okResp OkResponse
		if err := json.Unmarshal(body, &okResp); err != nil {
			return nil, err
		}
		return &okResp, nil
	} else {
		var errResp ErrorResponse
		if err := json.Unmarshal(body, &errResp); err != nil {
			return nil, err
		}
		return nil, &errResp
	}
}

// request creates a generic http request that includes a payload
// (for our purposes here, we're mainly focused on GET, POST, and PATCH)
func request[Request, Response interface{}](client *Client, method, endpoint string, payload *Request) (*Response, error) {
	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := makeRequestBody(method, client.formatURL(endpoint), rb)
	if err != nil {
		return nil, err
	}

	resp, err := do(client.httpClient, req)
	if err != nil {
		return nil, err
	}

	return parseResponse[Response](resp)
}
