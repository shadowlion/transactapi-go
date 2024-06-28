package transactapi

import (
	"fmt"
	"net/http"
)

type ValidateAbaRoutingNumberRequest struct {
	ClientID        string `json:"clientID"`
	DeveloperApiKey string `json:"developerAPIKey"`
	RoutingNumber   string `json:"routingNumber"`
}

type ValidateAbaRoutingNumberResponse struct {
	StatusCode     string `json:"statusCode"`
	StatusDesc     string `json:"statusDesc"`
	AccountDetails string `json:"accountDetails"`
}

// This method is used to validate the routing number for an external account (createExternalAccount)
//
// Reference: https://transactapi.readme.io/reference/validateabaroutingnumber
func (c *Client) ValidateABARoutingNumber(req *ValidateAbaRoutingNumberRequest) (*ValidateAbaRoutingNumberResponse, *ErrorResponse, error) {
	return request[ValidateAbaRoutingNumberRequest, ValidateAbaRoutingNumberResponse](c.httpClient, http.MethodPost, fmt.Sprintf("%s%s", c.baseURL(), EndpointValidateABARoutingNumber), req)
}
