package transactapi

import (
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

func (c *Client) AddCreditCard(req endpoints.AddCreditCardRequest) endpoints.AddCreditCardResponse {
	return basePostRequest[endpoints.AddCreditCardRequest, endpoints.AddCreditCardResponse](c, "/addCreditCard", &req)
}

func (c *Client) ValidateABARoutingNumber(req endpoints.ValidateAbaRoutingNumberRequest) endpoints.ValidateAbaRoutingNumberResponse {
	return basePostRequest[endpoints.ValidateAbaRoutingNumberRequest, endpoints.ValidateAbaRoutingNumberResponse](c, "/validateABARoutingNumber", &req)
}
