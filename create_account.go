package transactapi

import (
	"net/http"
)

type CreateAccountRequest struct {
	ClientID              string `json:"clientID"`
	DeveloperAPIKey       string `json:"developerAPIKey"`
	AccountRegistration   string `json:"accountRegistration"`
	Type                  string `json:"type"`
	EntityType            string `json:"entityType"`
	DomesticYN            string `json:"domesticYN"`
	StreetAddress1        string `json:"streetAddress1"`
	StreetAddress2        string `json:"streetAddress2"`
	City                  string `json:"city"`
	State                 string `json:"state"`
	Zip                   string `json:"zip"`
	Country               string `json:"country"`
	Email                 string `json:"email"`
	Phone                 int64  `json:"phone"`
	TaxID                 int64  `json:"taxID"`
	KYCstatus             string `json:"KYCstatus"`
	AMLstatus             string `json:"AMLstatus"`
	AMLdate               string `json:"AMLdate"`
	SuitabilityScore      int64  `json:"suitabilityScore"`
	SuitabilityDate       string `json:"suitabilityDate"`
	SuitabilityApprover   string `json:"suitabilityApprover"`
	AccreditedStatus      string `json:"AccreditedStatus"`
	Allow                 string `json:"Allow"`
	AIdate                string `json:"AIdate"`
	The506CLimit          int64  `json:"506cLimit"`
	AccountTotalLimit     int64  `json:"accountTotalLimit"`
	SingleInvestmentLimit int64  `json:"singleInvestmentLimit"`
	AssociatedAC          string `json:"associatedAC"`
	Syndicate             string `json:"syndicate"`
	Tags                  string `json:"tags"`
	Notes                 string `json:"notes"`
	ApprovalStatus        string `json:"approvalStatus"`
	ApprovalPrincipal     string `json:"approvalPrincipal"`
	ApprovalLastReview    string `json:"approvalLastReview"`
	Field1                string `json:"field1"`
	Field2                string `json:"field2"`
	Field3                string `json:"field3"`
}

type CreateAccountResponse struct {
	StatusCode     string                       `json:"statusCode"`
	StatusDesc     string                       `json:"statusDesc"`
	AccountDetails []createAccountAccountDetail `json:"accountDetails"`
}

type createAccountAccountDetail struct {
	AccountID string `json:"accountId"`
}

// This method is used to create an account that can be linked (createLink) to an individual party
// (createParty) or an entity (createEntity).
//
// Reference: https://transactapi.readme.io/reference/createaccount
func (c *Client) CreateAccount(req *CreateAccountRequest) (*CreateAccountResponse, error) {
	return request[CreateAccountRequest, CreateAccountResponse](c, http.MethodPost, EndpointCreateAccount, req)
}
