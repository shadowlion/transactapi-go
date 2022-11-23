package transactapi

import "net/http"

type createExternalAccountPayload struct {
	ClientID           string `json:"clientID"`
	DeveloperAPIKey    string `json:"developerAPIKey"`
	Types              string `json:"types"`
	AccountID          string `json:"accountId"`
	ExtAccountFullName string `json:"ExtAccountfullname"`
	ExtRoutingNumber   string `json:"ExtRoutingnumber"`
	ExtAccountNumber   string `json:"ExtAccountnumber"`
	UpdatedIpAddress   string `json:"updatedIpAddress"`
}

type createExternalAccountResponse struct {
	StatusCode             string                  `json:"statusCode"`
	StatusDesc             string                  `json:"statusDesc"`
	ExternalAccountDetails []externalAccountDetail `json:"External Account Details"`
}

type externalAccountDetail struct {
	AccountID          string `json:"accountId"`
	ExtAccountfullname string `json:"ExtAccountfullname"`
	Extnickname        string `json:"Extnickname"`
	ExtRoutingnumber   string `json:"ExtRoutingnumber"`
	ExtAccountnumber   string `json:"ExtAccountnumber"`
	Types              string `json:"types"`
	AccountType        string `json:"accountType"`
}

func (c *client) createExternalAccount(accountId, fullName, routingNumber, accountNumber string) (*createExternalAccountResponse, *errorResponse, error) {
	method := http.MethodPost
	endpoint := "/createExternalAccount"
	payload := createExternalAccountPayload{
		ClientID:           c.clientID,
		DeveloperAPIKey:    c.developerApiKey,
		AccountID:          accountId,
		Types:              "Account",
		ExtAccountFullName: fullName,
		ExtRoutingNumber:   routingNumber,
		ExtAccountNumber:   accountNumber,
		UpdatedIpAddress:   "69.61.177.110",
	}
	var res createExternalAccountResponse
	errRes, err := c.request(method, endpoint, payload, &res)

	if err != nil {
		return nil, nil, err
	}

	if errRes != nil {
		return nil, errRes, nil
	}

	return &res, nil, nil
}
