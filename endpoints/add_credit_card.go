package endpoints

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
