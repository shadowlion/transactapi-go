package transactapi

import (
	"net/http"
)

type getAccountPayload struct {
	ClientID        string `json:"clientID"`
	DeveloperAPIKey string `json:"developerAPIKey"`
	AccountID       string `json:"accountId"`
}

type getAccountResponse struct {
	StatusCode     string          `json:"statusCode"`
	StatusDesc     string          `json:"statusDesc"`
	AccountDetails []accountDetail `json:"accountDetails"`
}

type accountDetail struct {
	AccountID              string `json:"accountId"`
	AccountName            string `json:"accountName"`
	Type                   string `json:"type"`
	EntityType             string `json:"entityType"`
	ResidentType           string `json:"residentType"`
	Address1               string `json:"address1"`
	Address2               string `json:"address2"`
	City                   string `json:"city"`
	State                  string `json:"state"`
	Zip                    string `json:"zip"`
	Country                string `json:"country"`
	Phone                  string `json:"phone"`
	TaxId                  string `json:"taxID"`
	KycStatus              string `json:"kycStatus"`
	KycDate                string `json:"kycDate"`
	AmlStatus              string `json:"amlStatus"`
	AmlDate                string `json:"amlDate"`
	SuitabilityScore       string `json:"suitabilityScore"`
	SuitabilityDate        string `json:"suitabilityDate"`
	SuitabilityApprover    string `json:"suitabilityApprover"`
	AccreditedStatus       string `json:"accreditedStatus"`
	AccreditedInvestor     string `json:"accreditedInvestor"`
	AccreditedInvestorDate string `json:"accreditedInvestorDate"`
	Limit506c              string `json:"506cLimit"`
	AccountTotalLimit      string `json:"accountTotalLimit"`
	SingleInvestmentLimit  string `json:"singleInvestmentLimit"`
	AssociatedAC           string `json:"associatedAC"`
	Syndicate              string `json:"syndicate"`
	Tags                   string `json:"tags"`
	Notes                  string `json:"notes"`
	ApprovalStatus         string `json:"approvalStatus"`
	ApprovalPrincipal      string `json:"approvalPrincipal"`
	ApprovalLastReview     string `json:"approvalLastReview"`
	Field1                 string `json:"field1"`
	Field2                 string `json:"field2"`
	Field3                 string `json:"field3"`
}

func (c *client) GetAccount(accountId string) (*getAccountResponse, *errorResponse, error) {
	method := http.MethodPost
	endpoint := "/getAccount"
	payload := getAccountPayload{
		ClientID:        c.clientID,
		DeveloperAPIKey: c.developerApiKey,
		AccountID:       accountId,
	}
	var res getAccountResponse
	errRes, err := c.request(method, endpoint, payload, &res)

	if err != nil {
		return nil, nil, err
	}

	if errRes != nil {
		return nil, errRes, nil
	}

	return &res, nil, nil
}
