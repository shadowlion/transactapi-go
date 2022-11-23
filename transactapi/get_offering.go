package transactapi

import (
	"net/http"
)

type getOfferingPayload struct {
	ClientID        string `json:"clientID"`
	DeveloperAPIKey string `json:"developerAPIKey"`
	OfferingID      string `json:"offeringId"`
}

type getOfferingResponse struct {
	StatusCode      string           `json:"statusCode"`
	StatusDesc      string           `json:"statusDesc"`
	OfferingDetails []offeringDetail `json:"offeringDetails"`
}

type offeringDetail struct {
	IssuerID        string `json:"issuerId"`
	OfferingID      string `json:"offeringId"`
	IssueName       string `json:"issueName"`
	IssueType       string `json:"issueType"`
	TargetAmount    string `json:"targetAmount"`
	MinAmount       string `json:"minAmount"`
	MaxAmount       string `json:"maxAmount"`
	UnitPrice       string `json:"unitPrice"`
	TotalShares     string `json:"totalShares"`
	RemainingShares string `json:"remainingShares"`
	StartDate       string `json:"startDate"`
	EndDate         string `json:"endDate"`
	OfferingStatus  string `json:"offeringStatus"`
	OfferingText    string `json:"offeringText"`
	StampingText    string `json:"stampingText"`
}

func (c *client) GetOffering(offeringId string) (*getOfferingResponse, *errorResponse, error) {
	method := http.MethodPost
	endpoint := "/getOffering"
	payload := getOfferingPayload{
		ClientID:        c.clientID,
		DeveloperAPIKey: c.developerApiKey,
		OfferingID:      offeringId,
	}
	var res getOfferingResponse
	errRes, err := c.request(method, endpoint, payload, &res)

	if err != nil {
		return nil, nil, err
	}

	if errRes != nil {
		return nil, errRes, nil
	}

	return &res, nil, nil
}
