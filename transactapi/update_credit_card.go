package transactapi

import "net/http"

type updateCreditCardPayload struct {
	ClientID         string `json:"clientID"`
	DeveloperAPIKey  string `json:"developerAPIKey"`
	AccountID        string `json:"accountId"`
	CreditCardName   string `json:"creditCardName"`
	CreditCardNumber string `json:"creditCardNumber"`
	ExpirationDate   string `json:"expirationDate"`
	CvvNumber        string `json:"cvvNumber"`
	CardType         string `json:"cardType"`
	UpdatedIpAddress string `json:"updatedIpAddress"`
}

type updateCreditCardResponse struct {
	StatusCode        string `json:"statusCode"`
	StatusDesc        string `json:"statusDesc"`
	CreditCardDetails string `json:"creditcardDetails"`
}

func (c *client) updateCreditCard(accountId, name, number, month, year, cvv string) (*updateCreditCardResponse, *errorResponse, error) {
	method := http.MethodPost
	endpoint := "/updateCreditCard"
	payload := updateCreditCardPayload{
		ClientID:         c.clientID,
		DeveloperAPIKey:  c.developerApiKey,
		AccountID:        accountId,
		CreditCardName:   name,
		CreditCardNumber: number,
		ExpirationDate:   month + year,
		CvvNumber:        cvv,
		CardType:         getCreditCardType(number),
		UpdatedIpAddress: "69.61.177.110",
	}
	var res updateCreditCardResponse
	errRes, err := c.request(method, endpoint, payload, &res)

	if err != nil {
		return nil, nil, err
	}

	if errRes != nil {
		return nil, errRes, nil
	}

	return &res, nil, nil
}
