package endpoints

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
