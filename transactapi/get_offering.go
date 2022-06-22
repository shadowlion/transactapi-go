package transactapi

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type GetOfferingPayload struct {
	ClientID        string `json:"clientID"`
	DeveloperAPIKey string `json:"developerAPIKey"`
	OfferingID      string `json:"offeringId"`
}

type GetOfferingResponse struct {
	StatusCode      string           `json:"statusCode"`
	StatusDesc      string           `json:"statusDesc"`
	OfferingDetails []OfferingDetail `json:"offeringDetails"`
}

type OfferingDetail struct {
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

func (c *client) GetOffering(offeringId string) (*GetOfferingResponse, error) {
	payload := GetOfferingPayload{
		ClientID:        c.clientID,
		DeveloperAPIKey: c.developerApiKey,
		OfferingID:      offeringId,
	}

	jsonData, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		c.baseURL+"/getOffering",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return nil, err
	}

	var fullResponse GetOfferingResponse

	if err := c.sendRequest(req, &fullResponse); err != nil {
		return nil, err
	}

	return &fullResponse, nil
}
