package transactapi

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type GetTradeStatusPayload struct {
	ClientID        string `json:"clientID"`
	DeveloperAPIKey string `json:"developerAPIKey"`
	TradeID         string `json:"tradeId"`
}

type GetTradeStatusResponse struct {
	StatusCode   string        `json:"statusCode"`
	StatusDesc   string        `json:"statusDesc"`
	TradeDetails []TradeDetail `json:"tradeDetails"`
}

type TradeDetail struct {
	ID               string `json:"id"`
	DeveloperAPIKey  string `json:"developerAPIKey"`
	OfferingID       string `json:"offeringId"`
	AccountID        string `json:"accountId"`
	PartyID          string `json:"partyId"`
	PartyType        string `json:"party_type"`
	EscrowID         string `json:"escrowId"`
	OrderID          string `json:"orderId"`
	TransactionType  string `json:"transactionType"`
	TotalAmount      string `json:"totalAmount"`
	TotalShares      string `json:"totalShares"`
	OrderStatus      string `json:"orderStatus"`
	CreatedDate      string `json:"createdDate"`
	CreatedIPAddress string `json:"createdIpAddress"`
	Errors           string `json:"errors"`
	DocumentKey      string `json:"documentKey"`
	EsignStatus      string `json:"esignStatus"`
	Users            string `json:"users"`
	ArchivedStatus   string `json:"archived_status"`
}

func (c *client) GetTradeStatus(tradeId string) (*GetTradeStatusResponse, error) {
	payload := GetTradeStatusPayload{
		ClientID:        c.clientID,
		DeveloperAPIKey: c.developerApiKey,
		TradeID:         tradeId,
	}

	jsonData, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		c.baseURL+"/getTradeStatus",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return nil, err
	}

	var fullResponse GetTradeStatusResponse

	if err := c.sendRequest(req, &fullResponse); err != nil {
		return nil, err
	}

	return &fullResponse, nil
}
