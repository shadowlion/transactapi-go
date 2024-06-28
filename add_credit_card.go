package transactapi

import (
	"fmt"
	"net/http"
)

type AddCreditCardRequest struct {
	ClientID         string `json:"clientID"`
	DeveloperAPIKey  string `json:"developerAPIKey"`
	AccountID        string `json:"accountId"`
	CreditCardName   string `json:"creditCardName"`
	CreditCardNumber string `json:"creditCardNumber"`
	ExpirationDate   string `json:"expirationDate"`
	CvvNumber        string `json:"cvvNumber"`
	CardType         string `json:"cardType"`
	CreatedIpAddress string `json:"createdIpAddress"`
}

type AddCreditCardResponse struct {
	StatusCode        string `json:"statusCode"`
	StatusDesc        string `json:"statusDesc"`
	CreditCardDetails string `json:"creditcardDetails"`
}

// AddCreditCard is used to save a credit card to an Account (createAccount).
//
// Only one credit card can be added for each account.
//
// Reference: https://transactapi.readme.io/reference/addcreditcard
func (c *Client) AddCreditCard(req *AddCreditCardRequest) (*AddCreditCardResponse, *ErrorResponse, error) {
	return request[AddCreditCardRequest, AddCreditCardResponse](c.httpClient, http.MethodPost, fmt.Sprintf("%s%s", c.baseURL(), EndpointAddCreditCard), req)
}
