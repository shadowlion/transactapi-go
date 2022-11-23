package transactapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type client struct {
	HttpClient      *http.Client
	clientID        string
	developerApiKey string
	sandbox         bool
}

type errorResponse struct {
	StatusCode        string `json:"statusCode"`
	StatusDescription string `json:"statusDesc"`
}

func New(clientId string, developerApiKey string, sandbox bool) *client {
	return &client{
		HttpClient:      &http.Client{},
		clientID:        clientId,
		developerApiKey: developerApiKey,
	}
}

func (c *client) baseUrl() string {
	var p string
	if c.sandbox {
		p = "api-sandboxdash"
	} else {
		p = "api"
	}
	return fmt.Sprintf("https://%s.norcapsecurities.com/tapiv3/index.php/v3", p)
}

func (c *client) sendRequest(req *http.Request, v interface{}) (*errorResponse, error) {
	req.Header.Set("Content-Type", "application/json")

	res, err := c.HttpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var errRes errorResponse

		if err = json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return nil, errors.New(errRes.StatusDescription)
		}

		return &errRes, nil
	}

	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *client) request(method string, endpoint string, payload, res interface{}) (*errorResponse, error) {
	jsonData, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		method,
		c.baseUrl()+endpoint,
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return nil, err
	}

	errRes, err := c.sendRequest(req, &res)

	if err != nil {
		return nil, err
	}

	if errRes != nil {
		return errRes, nil
	}

	return nil, nil
}
