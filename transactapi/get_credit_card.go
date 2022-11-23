package transactapi

import (
	"net/http"
)

type getCreditCardPayload struct {
	ClientID        string `json:"clientID"`
	DeveloperAPIKey string `json:"developerAPIKey"`
	AccountID       string `json:"accountId"`
}

type creditCardDetails struct {
	AccountId        string `json:"accountId"`
	CreditCardName   string `json:"creditCardName"`
	CreditCardNumber string `json:"creditCardNumber"`
	ExpirationDate   string `json:"expirationDate"`
	CvvNumber        string `json:"cvvNumber"`
}

type getCreditCardResponse struct {
	StatusCode        string            `json:"statusCode"`
	StatusDesc        string            `json:"statusDesc"`
	CreditCardDetails creditCardDetails `json:"creditcardDetails"`
}

// # Get Credit Card
//
// This method is used to retrieve the credit card information that was
// previously saved to a specific account.
//
// ## Arguments
//
// - `accountId` - Account ID that is generated by the API once an account is
// created (createAccount).
func (c *client) GetCreditCard(accountId string) (*getCreditCardResponse, *errorResponse, error) {
	method := http.MethodPost
	endpoint := "/getCreditCard"
	payload := getCreditCardPayload{
		ClientID:        c.clientID,
		DeveloperAPIKey: c.developerApiKey,
		AccountID:       accountId,
	}
	var res getCreditCardResponse
	errRes, err := c.request(method, endpoint, payload, &res)

	if err != nil {
		return nil, nil, err
	}

	if errRes != nil {
		return nil, errRes, nil
	}

	return &res, nil, nil
}
