package transactapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/shadowlion/transactapi-go/endpoints"
)

type Client struct {
	HttpClient      *http.Client
	Sandbox         bool
	ClientID        string
	DeveloperAPIKey string
}

func New(sandbox bool, clientId, developerApiKey string) *Client {
	return &Client{
		HttpClient:      &http.Client{},
		Sandbox:         sandbox,
		ClientID:        clientId,
		DeveloperAPIKey: developerApiKey,
	}
}

func (c *Client) baseUrl() string {
	var p string
	if c.Sandbox {
		p = "api-sandboxdash"
	} else {
		p = "api"
	}
	return fmt.Sprintf("https://%s.norcapsecurities.com/tapiv3/index.php/v3", p)
}

func BasePostRequest[Request, Response any](c *Client, endpoint string, req Request) Response {
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

func (c *Client) AddCreditCard(req endpoints.AddCreditCardRequest) endpoints.AddCreditCardResponse {
	return BasePostRequest[endpoints.AddCreditCardRequest, endpoints.AddCreditCardResponse](c, "/addCreditCard", req)
}

func (c *Client) ValidateABARoutingNumber(req endpoints.ValidateAbaRoutingNumberRequest) endpoints.ValidateAbaRoutingNumberResponse {
	return BasePostRequest[endpoints.ValidateAbaRoutingNumberRequest, endpoints.ValidateAbaRoutingNumberResponse](c, "/validateABARoutingNumber", req)
}
