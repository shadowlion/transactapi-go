package transactapi

import (
	"net/http"
)

type getAllTradesPayload struct {
	ClientID        string `json:"clientID"`
	DeveloperAPIKey string `json:"developerAPIKey"`
}

type tradeFinancialDetail struct {
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

type getAllTradesResponse struct {
	StatusCode            string `json:"statusCode"`
	StatusDesc            string `json:"statusDesc"`
	TradeFinancialDetails []tradeFinancialDetail
}

func (c *client) GetAllTrades() (*getAllTradesResponse, *errorResponse, error) {
	method := http.MethodPost
	endpoint := "/getAllTrades"
	payload := getAllTradesPayload{
		ClientID:        c.clientID,
		DeveloperAPIKey: c.developerApiKey,
	}
	var res getAllTradesResponse
	errRes, err := c.request(method, endpoint, payload, &res)

	if err != nil {
		return nil, nil, err
	}

	if errRes != nil {
		return nil, errRes, nil
	}

	return &res, nil, nil
}
