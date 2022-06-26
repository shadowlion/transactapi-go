package transactapi

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type GetTradePayload struct {
	ClientID        string `json:"clientID"`
	DeveloperAPIKey string `json:"developerAPIKey"`
	AccountID       string `json:"accountId"`
	TradeID         string `json:"tradeId"`
}

type GetTradeResponse struct {
	StatusCode   string        `json:"statusCode"`
	StatusDesc   string        `json:"statusDesc"`
	PartyDetails []PartyDetail `json:"partyDetails"`
}

type PartyDetail struct {
	ID                      string `json:"id"`
	DeveloperAPIKey         string `json:"developerAPIKey"`
	OfferingID              string `json:"offeringId"`
	AccountID               string `json:"accountId"`
	PartyID                 string `json:"partyId"`
	PartyType               string `json:"party_type"`
	EscrowID                string `json:"escrowId"`
	OrderID                 string `json:"orderId"`
	TransactionType         string `json:"transactionType"`
	TotalAmount             string `json:"totalAmount"`
	TotalShares             string `json:"totalShares"`
	OrderStatus             string `json:"orderStatus"`
	CreatedDate             string `json:"createdDate"`
	CreatedIPAddress        string `json:"createdIpAddress"`
	Errors                  string `json:"errors"`
	DocumentKey             string `json:"documentKey"`
	EsignStatus             string `json:"esignStatus"`
	Users                   string `json:"users"`
	ArchivedStatus          string `json:"archived_status"`
	RRAprovalStatus         string `json:"RRAprovalStatus"`
	RRName                  string `json:"RRName"`
	RRApprovalDate          string `json:"RRApprovalDate"`
	PrincipalApprovalStatus string `json:"PrincipalApprovalStatus"`
	PrincipalName           string `json:"PrincipalName"`
	PrincipalDate           string `json:"PrincipalDate"`
}

func (c *client) GetTrade(accountId string, tradeId string) (*GetTradeResponse, error) {
	payload := GetTradePayload{
		ClientID:        c.clientID,
		DeveloperAPIKey: c.developerApiKey,
		AccountID:       accountId,
		TradeID:         tradeId,
	}

	jsonData, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		c.baseURL+"/getTrade",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return nil, err
	}

	var fullResponse GetTradeResponse

	if err := c.sendRequest(req, &fullResponse); err != nil {
		return nil, err
	}

	return &fullResponse, nil
}
