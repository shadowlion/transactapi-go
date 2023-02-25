package endpoints

type GetOfferingRequest struct {
	ClientID        string `json:"clientID"`
	DeveloperApiKey string `json:"developerAPIKey"`
	OfferingID      string `json:"offeringId"`
}

type GetOfferingResponse struct {
	StatusCode      string           `json:"statusCode"`
	StatusDesc      string           `json:"statusDesc"`
	OfferingDetails []offeringDetail `json:"offeringDetails"`
}

type offeringDetail struct {
	IssuerID            string `json:"issuerId"`
	OfferingID          string `json:"offeringId"`
	IssueName           string `json:"issueName"`
	IssueType           string `json:"issueType"`
	TargetAmount        string `json:"targetAmount"`
	MinAmount           string `json:"minAmount"`
	MaxAmount           string `json:"maxAmount"`
	UnitPrice           string `json:"unitPrice"`
	TotalShares         string `json:"totalShares"`
	RemainingShares     string `json:"remainingShares"`
	StartDate           string `json:"startDate"`
	EndDate             string `json:"endDate"`
	OfferingStatus      string `json:"offeringStatus"`
	OfferingText        string `json:"offeringText"`
	StampingText        string `json:"stampingText"`
	EscrowAccountNumber string `json:"escrowAccountNumber"`
	Field1              string `json:"field1"`
	Field2              string `json:"field2"`
	Field3              string `json:"field3"`
}
