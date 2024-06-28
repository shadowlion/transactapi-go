package transactapi

import (
	"fmt"
	"net/http"
)

type GetExternalAccountRequest struct {
	ClientID        string `json:"clientID"`
	DeveloperApiKey string `json:"developerAPIKey"`
	Types           string `json:"types"`
	AccountID       string `json:"accountId"`
}

type GetExternalAccountResponse struct {
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

// Reference: https://transactapi.readme.io/reference/getexternalaccount
func (c *Client) GetExternalAccount(req *GetExternalAccountRequest) (*GetExternalAccountResponse, *ErrorResponse, error) {
	return request[GetExternalAccountRequest, GetExternalAccountResponse](c.httpClient, http.MethodPost, fmt.Sprintf("%s%s", c.baseURL(), EndpointGetExternalAccount), req)
}
