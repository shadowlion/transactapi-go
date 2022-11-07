package transactapi

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type GetExternalAccountPayload struct {
	basePayload
	Types     string `json:"types"`
	AccountID string `json:"accountId"`
}

type GetExternalAccountResponse struct {
	StatusCode string     `json:"statusCode"`
	AccountID  string     `json:"accountId"`
	StatusDesc StatusDesc `json:"statusDesc"`
}

type StatusDesc struct {
	AccountName          string `json:"AccountName"`
	AccountNickName      string `json:"AccountNickName"`
	AccountRoutingNumber string `json:"AccountRoutingNumber"`
	AccountNumber        string `json:"AccountNumber"`
	AccountType          string `json:"accountType"`
	UpdatedDate          string `json:"updatedDate"`
	CreatedDate          string `json:"createdDate"`
}

func (c *client) GetExternalAccount(p GetExternalAccountPayload) (*GetExternalAccountResponse, error) {
	jsonData, err := json.Marshal(p)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		c.baseURL+"/getExternalAccount",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return nil, err
	}

	var fullResponse GetExternalAccountResponse

	if err := c.sendRequest(req, &fullResponse); err != nil {
		return nil, err
	}

	return &fullResponse, nil
}
