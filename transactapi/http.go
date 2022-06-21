package transactapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type client struct {
	HttpClient      *http.Client
	baseURL         string
	ClientID        string
	DeveloperApiKey string
}

type successResponse struct {
	StatusCode        string      `json:"statusCode"`
	StatusDescription string      `json:"statusDesc"`
	Data              interface{} `json:"data,omitempty"`
}

type errorResponse struct {
	StatusCode        string `json:"statusCode"`
	StatusDescription string `json:"statusDesc"`
}

func New(clientId string, developerApiKey string, sandbox bool) *client {
	var url string

	if sandbox {
		url = "https://api-sandboxdash.norcapsecurities.com/tapiv3/index.php/v3"
	} else {
		url = "https://api.norcapsecurities.com/tapiv3/index.php/v3"
	}

	return &client{
		baseURL:         url,
		ClientID:        clientId,
		DeveloperApiKey: developerApiKey,
		HttpClient:      &http.Client{},
	}
}

func (c *client) sendRequest(req *http.Request) (*http.Response, error) {
	res, err := c.HttpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return nil, errors.New(errRes.StatusDescription)
		}

		return nil, fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	return res, nil
}
