package transactapi

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type GetAccountPayload struct {
	ClientID        string `json:"clientID"`
	DeveloperAPIKey string `json:"developerAPIKey"`
	AccountID       string `json:"accountId"`
}

type GetAccountResponse struct {
	StatusCode     string          `json:"statusCode"`
	StatusDesc     string          `json:"statusDesc"`
	AccountDetails []AccountDetail `json:"accountDetails"`
}

type AccountDetail struct {
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

func (c *client) GetAccount(accountId string) (*GetAccountResponse, error) {
	payload := GetAccountPayload{
		ClientID:        c.clientID,
		DeveloperAPIKey: c.developerApiKey,
		AccountID:       accountId,
	}

	jsonData, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		c.baseURL+"/getAccount",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return nil, err
	}

	var fullResponse GetAccountResponse

	if err := c.sendRequest(req, &fullResponse); err != nil {
		return nil, err
	}

	return &fullResponse, nil
}
