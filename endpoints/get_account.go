package endpoints

type GetAccountRequest struct {
	ClientID        string `json:"clientID"`
	DeveloperApiKey string `json:"developerAPIKey"`
	AccountID       string `json:"accountId"`
}

type GetAccountResponse struct {
	StatusCode     string         `json:"statusCode"`
	StatusDesc     string         `json:"statusDesc"`
	AccountDetails accountDetails `json:"accountDetails"`
}

type accountDetails struct {
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
	TaxID                  string `json:"taxID"`
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
	The506CLimit           string `json:"506cLimit"`
	AccountTotalLimit      string `json:"accountTotalLimit"`
	SingleInvestmentLimit  string `json:"singleInvestmentLimit"`
	AssociatedAC           string `json:"associatedAC"`
	Syndicate              string `json:"syndicate"`
	Tags                   string `json:"tags"`
	Notes                  string `json:"notes"`
	ApprovalStatus         string `json:"approvalStatus"`
	ApprovalPrincipal      string `json:"approvalPrincipal"`
	ApprovalLastReview     string `json:"approvalLastReview"`
	ArchivedStatus         string `json:"archived_status"`
	Field1                 string `json:"field1"`
	Field2                 string `json:"field2"`
	Field3                 string `json:"field3"`
}
