package transactapi

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// ErrorResponse is the standard shape of the json object that gets returned to us when
// we make an API call, but comes back with a (although status code 200) error shape.
type ErrorResponse struct {
	StatusCode        string `json:"statusCode"`
	StatusDescription string `json:"statusDesc"`
}

// request parses a payload struct to json, sends the http request, then conforms the response body
// to either an ok response shape that's determined at runtime, an error shape if the error is from
// North Capital's server side, or an error related to the conforming to either shape.
func request[Req, Res interface{}](client *http.Client, method, url string, payload *Req) (*Res, *ErrorResponse, error) {
	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(rb))
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}

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
