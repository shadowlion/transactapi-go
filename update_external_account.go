package transactapi

import (
	"fmt"
	"net/http"
)

type UpdateExternalAccountRequest struct {
	ClientID           string `json:"clientID"`
	DeveloperAPIKey    string `json:"developerAPIKey"`
	ExtAccountfullname string `json:"ExtAccountfullname"`
	ExtNickname        string `json:"Extnickname"`
	Types              string `json:"types"`
	AccountID          string `json:"accountId"`
	ExtRoutingnumber   string `json:"ExtRoutingnumber"`
	ExtAccountnumber   string `json:"ExtAccountnumber"`
	UpdatedIPAddress   string `json:"updatedIpAddress"`
}

type UpdateExternalAccountResponse struct {
	StatusCode     string               `json:"statusCode"`
	StatusDesc     string               `json:"statusDesc"`
	AccountDetails []accountDetailUnion `json:"accountDetails"`
}

type accountDetailClass struct {
	AccountID          string `json:"accountId"`
	ExtAccountfullname string `json:"ExtAccountfullname"`
	Extnickname        string `json:"Extnickname"`
	ExtRoutingnumber   string `json:"ExtRoutingnumber"`
	ExtAccountnumber   string `json:"ExtAccountnumber"`
	AccountType        string `json:"accountType"`
	Types              string `json:"types"`
}

type accountDetailUnion struct {
	AccountDetailClassArray []accountDetailClass
	Bool                    *bool
}

// This method is used to update fields related to a particular external account for an Account (createAccount).
// The Account ID must be specified as a request parameter to update the record.
//
// Reference: https://transactapi.readme.io/reference/updateexternalaccount
func (c *Client) UpdateExternalAccount(req *UpdateExternalAccountRequest) (*UpdateExternalAccountResponse, *ErrorResponse, error) {
	return request[UpdateExternalAccountRequest, UpdateExternalAccountResponse](c.httpClient, http.MethodPost, fmt.Sprintf("%s%s", c.baseURL(), EndpointUpdateExternalAccount), req)
}
