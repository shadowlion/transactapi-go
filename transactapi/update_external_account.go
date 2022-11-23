package transactapi

import "net/http"

type updateExternalAccountPayload struct {
	ClientID           string `json:"clientID"`
	DeveloperAPIKey    string `json:"developerAPIKey"`
	Types              string `json:"types"`
	AccountID          string `json:"accountId"`
	ExtAccountFullName string `json:"ExtAccountfullname"`
	ExtRoutingNumber   string `json:"ExtRoutingnumber"`
	ExtAccountNumber   string `json:"ExtAccountnumber"`
	UpdatedIpAddress   string `json:"updatedIpAddress,omitempty"`
}

type updateExternalAccountResponse struct {
	StatusCode     string                               `json:"statusCode"`
	StatusDesc     string                               `json:"statusDesc"`
	AccountDetails []updateExternalAccountAccountDetail `json:"accountDetails"`
}

type updateExternalAccountAccountDetail struct {
	AccountID          string `json:"accountId"`
	ExtAccountfullname string `json:"ExtAccountfullname"`
	Extnickname        string `json:"Extnickname"`
	ExtRoutingnumber   string `json:"ExtRoutingnumber"`
	ExtAccountnumber   string `json:"ExtAccountnumber"`
	Types              string `json:"types"`
	AccountType        string `json:"accountType"`
}

func (c *client) updateExternalAccount(accountId, fullName, routingNumber, accountNumber string) (*updateExternalAccountResponse, *errorResponse, error) {
	method := http.MethodPost
	endpoint := "/updateExternalAccount"
	payload := updateExternalAccountPayload{
		ClientID:           c.clientID,
		DeveloperAPIKey:    c.developerApiKey,
		AccountID:          accountId,
		Types:              "Account",
		ExtAccountFullName: fullName,
		ExtRoutingNumber:   routingNumber,
		ExtAccountNumber:   accountNumber,
		UpdatedIpAddress:   "69.61.177.110",
	}
	var res updateExternalAccountResponse
	errRes, err := c.request(method, endpoint, payload, &res)

	if err != nil {
		return nil, nil, err
	}

	if errRes != nil {
		return nil, errRes, nil
	}

	return &res, nil, nil
}
