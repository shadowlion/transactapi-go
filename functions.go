package transactapi

import "github.com/shadowlion/transactapi-go/endpoints"

func (c *Client) CreatePayloadValidateAbaRoutingNumber(routingNumber string) endpoints.ValidateAbaRoutingNumberRequest {
	return endpoints.ValidateAbaRoutingNumberRequest{
		ClientID:        c.ClientID,
		DeveloperApiKey: c.DeveloperAPIKey,
		RoutingNumber:   routingNumber,
	}
}
