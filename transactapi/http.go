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

func BaseUrl(sandbox bool) string {
	var p string
	if sandbox {
		p = "api-sandboxdash"
	} else {
		p = "api"
	}
	return fmt.Sprintf("https://%s.norcapsecurities.com/tapiv3/index.php/v3", p)
}

func New(clientId string, developerApiKey string, sandbox bool) *client {
	return &client{
		baseURL:         BaseUrl(sandbox),
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
