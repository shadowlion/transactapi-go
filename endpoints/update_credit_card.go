package endpoints

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
