package transactapi

import (
	"fmt"
	"net/http"
)

type UpdateCreditCardRequest struct {
	ClientID         string `json:"clientID"`
	DeveloperAPIKey  string `json:"developerAPIKey"`
	AccountID        string `json:"accountId"`
	CreditCardName   string `json:"creditCardName"`
	CreditCardNumber string `json:"creditCardNumber"`
	ExpirationDate   string `json:"expirationDate"`
	CvvNumber        string `json:"cvvNumber"`
	CardType         string `json:"cardType"`
	UpdatedIPAddress string `json:"updatedIpAddress"`
}

type UpdateCreditCardResponse struct {
	StatusCode        string `json:"statusCode"`
	StatusDesc        string `json:"statusDesc"`
	CreditcardDetails string `json:"creditcardDetails"`
}

// This method is used to update the credit card information that is saved to a specific account (createExternalAccount).
//
// Reference: https://transactapi.readme.io/reference/updatecreditcard
func (c *Client) UpdateCreditCard(req *UpdateCreditCardRequest) (*UpdateCreditCardResponse, *ErrorResponse, error) {
	return request[UpdateCreditCardRequest, UpdateCreditCardResponse](c.httpClient, http.MethodPost, fmt.Sprintf("%s%s", c.baseURL(), EndpointUpdateCreditCard), req)
}
