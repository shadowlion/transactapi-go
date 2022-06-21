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
	clientID        string
	developerApiKey string
}

type successResponse struct {
	StatusCode        string `json:"statusCode"`
	StatusDescription string `json:"statusDesc"`
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
		clientID:        clientId,
		developerApiKey: developerApiKey,
		HttpClient:      &http.Client{},
	}
}

func (c *client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json")

	res, err := c.HttpClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var errRes errorResponse

		if err = json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return errors.New(errRes.StatusDescription)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return err
}
