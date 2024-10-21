package transactapi

import (
	"net/http"
)

type UpdateAccountRequest struct {
	ClientID              string `json:"clientID"`
	DeveloperAPIKey       string `json:"developerAPIKey"`
	AccountRegistration   string `json:"accountRegistration"`
	Type                  string `json:"type"`
	EntityType            string `json:"entityType,omitempty"`
	DomesticYN            string `json:"domesticYN"`
	StreetAddress1        string `json:"streetAddress1"`
	StreetAddress2        string `json:"streetAddress2,omitempty"`
	City                  string `json:"city"`
	State                 string `json:"state"`
	Zip                   string `json:"zip"`
	Country               string `json:"country"`
	Email                 string `json:"email,omitempty"`
	Phone                 int64  `json:"phone,omitempty"`
	TaxID                 int64  `json:"taxID,omitempty"`
	KYCstatus             string `json:"KYCstatus,omitempty"`
	AMLstatus             string `json:"AMLstatus,omitempty"`
	AMLdate               string `json:"AMLdate,omitempty"`
	SuitabilityScore      int64  `json:"suitabilityScore,omitempty"`
	SuitabilityDate       string `json:"suitabilityDate,omitempty"`
	SuitabilityApprover   string `json:"suitabilityApprover,omitempty"`
	AccreditedStatus      string `json:"AccreditedStatus,omitempty"`
	Allow                 string `json:"Allow,omitempty"`
	AIdate                string `json:"AIdate,omitempty"`
	The506CLimit          int64  `json:"506cLimit,omitempty"`
	AccountTotalLimit     int64  `json:"accountTotalLimit,omitempty"`
	SingleInvestmentLimit int64  `json:"singleInvestmentLimit,omitempty"`
	AssociatedAC          string `json:"associatedAC,omitempty"`
	Syndicate             string `json:"syndicate,omitempty"`
	Tags                  string `json:"tags,omitempty"`
	Notes                 string `json:"notes,omitempty"`
	ApprovalStatus        string `json:"approvalStatus"`
	ApprovalPrincipal     string `json:"approvalPrincipal,omitempty"`
	ApprovalLastReview    string `json:"approvalLastReview,omitempty"`
	Field1                string `json:"field1,omitempty"`
	Field2                string `json:"field2,omitempty"`
	Field3                string `json:"field3,omitempty"`
	UpdatedIPAddress      string `json:"updatedIpAddress"`
}

type UpdateAccountResponse struct {
	StatusCode     string                       `json:"statusCode"`
	StatusDesc     string                       `json:"statusDesc"`
	AccountDetails []updateAccountAccountDetail `json:"accountDetails"`
}

type updateAccountAccountDetail struct {
	AccountID        string `json:"accountId"`
	SuitabilityScore string `json:"suitabilityScore"`
	ApprovalStatus   string `json:"approvalStatus"`
}

// This method is used to update a specific account (updateAccount)
//
// Reference: https://transactapi.readme.io/reference/updateaccount
func (c *Client) UpdateAccount(req *UpdateAccountRequest) (*UpdateAccountResponse, error) {
	return request[UpdateAccountRequest, UpdateAccountResponse](c, http.MethodPut, EndpointUpdateAccount, req)
}
