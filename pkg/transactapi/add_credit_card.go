package transactapi

type AddCreditCardRequest struct {
	BaseRequest
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

// This method is used to save a credit card to an Account (createAccount). Only one credit card
// can be added for an account.
//
// Reference: https://transactapi.readme.io/reference/addcreditcard
func (c *Client) AddCreditCard(req *AddCreditCardRequest) (AddCreditCardResponse, error) {
	return PostRequest[AddCreditCardResponse](c.ctx, "/addCreditCard", &req)
}
