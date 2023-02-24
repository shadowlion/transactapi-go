package transactapi

import (
	"net/http"

	"github.com/shadowlion/transactapi-go/endpoints"
)

type Client struct {
	HttpClient      *http.Client
	ClientID        string
	DeveloperAPIKey string
	Sandbox         bool
}

func New(clientId, developerApiKey string, sandbox bool) *Client {
	return &Client{
		HttpClient:      &http.Client{},
		ClientID:        clientId,
		DeveloperAPIKey: developerApiKey,
		Sandbox:         sandbox,
	}
}

func (c *Client) AddCreditCard(req endpoints.AddCreditCardRequest) endpoints.AddCreditCardResponse {
	return basePostRequest[endpoints.AddCreditCardRequest, endpoints.AddCreditCardResponse](c, "/addCreditCard", &req)
}

func (c *Client) ValidateABARoutingNumber(req endpoints.ValidateAbaRoutingNumberRequest) endpoints.ValidateAbaRoutingNumberResponse {
	return basePostRequest[endpoints.ValidateAbaRoutingNumberRequest, endpoints.ValidateAbaRoutingNumberResponse](c, "/validateABARoutingNumber", &req)
}

func (c *Client) CCFundMove(req endpoints.CCFundMoveRequest) endpoints.CCFundMoveResponse {
	return basePostRequest[endpoints.CCFundMoveRequest, endpoints.CCFundMoveResponse](c, "/ccFundMove", &req)
}
