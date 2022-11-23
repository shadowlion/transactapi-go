package transactapi

import (
	"net/http"
)

type getExternalAccountPayload struct {
	ClientID        string `json:"clientID"`
	DeveloperApiKey string `json:"developerAPIKey"`
	Types           string `json:"types"`
	AccountID       string `json:"accountId"`
}

type getExternalAccountResponse struct {
	StatusCode string     `json:"statusCode"`
	AccountID  string     `json:"accountId"`
	StatusDesc statusDesc `json:"statusDesc"`
}

type statusDesc struct {
	AccountName          string `json:"AccountName"`
	AccountNickName      string `json:"AccountNickName"`
	AccountRoutingNumber string `json:"AccountRoutingNumber"`
	AccountNumber        string `json:"AccountNumber"`
	AccountType          string `json:"accountType"`
	UpdatedDate          string `json:"updatedDate"`
	CreatedDate          string `json:"createdDate"`
}

func (c *client) GetExternalAccount(accountId string) (*getExternalAccountResponse, *errorResponse, error) {
	method := http.MethodPost
	endpoint := "/getExternalAccount"
	payload := getExternalAccountPayload{
		ClientID:        c.clientID,
		DeveloperApiKey: c.developerApiKey,
		AccountID:       accountId,
	}
	var res getExternalAccountResponse
	errRes, err := c.request(method, endpoint, payload, &res)

	if err != nil {
		return nil, nil, err
	}

	if errRes != nil {
		return nil, errRes, nil
	}

	return &res, nil, nil
}
