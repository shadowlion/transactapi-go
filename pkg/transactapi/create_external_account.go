package transactapi

import (
	"fmt"
	"net/http"
)

type CreateExternalAccountRequest struct {
	ClientID           string `json:"clientID"`
	DeveloperAPIKey    string `json:"developerAPIKey"`
	Types              string `json:"types"`
	AccountID          string `json:"accountId"`
	ExtAccountfullname string `json:"ExtAccountfullname"`
	Extnickname        string `json:"Extnickname"`
	ExtRoutingnumber   string `json:"ExtRoutingnumber"`
	ExtAccountnumber   string `json:"ExtAccountnumber"`
	UpdatedIPAddress   string `json:"updatedIpAddress"`
}

type CreateExternalAccountResponse struct {
	StatusCode             string `json:"statusCode"`
	StatusDesc             string `json:"statusDesc"`
	ExternalAccountDetails []externalAccountDetailElement
}

type externalAccountDetailElement struct {
	Bool                       *bool
	ExternalAccountDetailClass *externalAccountDetailClass
}

type externalAccountDetailClass struct {
	AccountID          string `json:"accountId"`
	EXTAccountfullname string `json:"ExtAccountfullname"`
	Extnickname        string `json:"Extnickname"`
	ExtRoutingnumber   string `json:"ExtRoutingnumber"`
	ExtAccountnumber   string `json:"ExtAccountnumber"`
	AccountType        string `json:"accountType"`
	Types              string `json:"types"`
}

// This method is used to add information to an Account (createAccount) for an external bank
// account which an ACH transfer can be initiated from. Only one external account can be created
// for an account. External accounts can have funds debited from them (externalFundMove).
//
// Reference: https://transactapi.readme.io/reference/createexternalaccount
func (c *Client) CreateExternalAccount(req *CreateExternalAccountRequest) (*CreateExternalAccountResponse, *ErrorResponse, error) {
	return request[CreateExternalAccountRequest, CreateExternalAccountResponse](c.HttpClient, http.MethodPost, fmt.Sprintf("%s%s", c.baseURL(), EndpointCreateExternalAccount), req)
}
