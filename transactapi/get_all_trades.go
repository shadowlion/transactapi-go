package transactapi

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type GetAllTradesPayload struct {
	ClientID        string `json:"clientID"`
	DeveloperAPIKey string `json:"developerAPIKey"`
}

type TradeFinancialDetail struct {
	Id              string `json:"id"`
	CreatedDate     string `json:"createdDate"`
	AccountId       string `json:"accountId"`
	IssueName       string `json:"issueName"`
	OrderId         string `json:"orderId"`
	AccountName     string `json:"accountName"`
	OfferingId      string `json:"offeringId"`
	TotalAmount     string `json:"totalAmount"`
	TransactionType string `json:"transactionType"`
	ApprovalStatus  string `json:"approvalStatus"`
	KycStatus       string `json:"kycStatus"`
	AmlStatus       string `json:"amlStatus"`
	OrderStatus     string `json:"orderStatus"`
	EsignStatus     string `json:"esignStatus"`
	DeveloperAPIKey string `json:"developerAPIKey"`
	ClientName      string `json:"clientName"`
	ClientId        string `json:"clientId"`
	Field1          string `json:"field1"`
}

type GetAllTradesResponse struct {
	StatusCode            string `json:"statusCode"`
	StatusDesc            string `json:"statusDesc"`
	TradeFinancialDetails []TradeFinancialDetail
}

func (c *client) GetAllTrades() (*GetAllTradesResponse, error) {
	endpoint := "/getAllTrades"

	payload := GetAllTradesPayload{
		ClientID:        c.clientID,
		DeveloperAPIKey: c.developerApiKey,
	}

	jsonData, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		c.baseURL+endpoint,
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return nil, err
	}

	var fullResponse GetAllTradesResponse

	if err := c.sendRequest(req, &fullResponse); err != nil {
		return nil, err
	}

	return &fullResponse, nil
}
