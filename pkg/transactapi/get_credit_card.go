package transactapi

type GetCreditCardRequest struct {
	ClientID        string `json:"clientID"`
	DeveloperApiKey string `json:"developerAPIKey"`
	AccountID       string `json:"accountId"`
}

type GetCreditCardResponse struct {
	StatusCode        string            `json:"statusCode"`
	StatusDesc        string            `json:"statusDesc"`
	CreditcardDetails creditcardDetails `json:"creditcardDetails"`
}

type creditcardDetails struct {
	AccountID        string `json:"accountId"`
	CreditCardName   string `json:"creditCardName"`
	CreditCardNumber string `json:"creditCardNumber"`
	ExpirationDate   string `json:"expirationDate"`
	CvvNumber        string `json:"cvvNumber"`
}

// This method is used to retrieve the credit card information that was previously saved to a
// specific account.
//
// Reference: https://transactapi.readme.io/reference/getcreditcard
func (c *Client) GetCreditCard(req GetCreditCardRequest) (GetCreditCardResponse, error) {
	return PostRequest[GetCreditCardResponse](c.ctx, "/getCreditCard", &req)
}
