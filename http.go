package transactapi

import (
	"bytes"
	"encoding/json"
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
// standard error response given in the documentation, or an error related to creating the request.
func parseResponse[Res interface{}](resp *http.Response) (*Res, *ErrorResponse, error) {
	defer resp.Body.Close()

	var okResp Res
	var errResponse ErrorResponse

	if err := json.NewDecoder(resp.Body).Decode(&okResp); err != nil {
		if err := json.NewDecoder(resp.Body).Decode(&errResponse); err != nil {
			return nil, nil, err
		}
		return nil, &errResponse, nil
	}

	return &okResp, nil, nil
}

// request creates a generic http request that includes a payload
// (for our purposes here, we're mainly focused on GET, POST, and PATCH)
func request[Req, Res interface{}](client *http.Client, method, url string, payload *Req) (*Res, *ErrorResponse, error) {
	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, nil, err
	}

	req, err := makeRequestBody(method, url, rb)
	if err != nil {
		return nil, nil, err
	}

	resp, err := do(client, req)
	if err != nil {
		return nil, nil, err
	}

	return parseResponse[Res](resp)
}
