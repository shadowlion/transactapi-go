package transactapi

import (
	"net/http"
)

type UpdateTradeTransactionTypeRequest struct {
	ClientID        string `json:"clientID"`
	DeveloperAPIKey string `json:"developerAPIKey"`
	TradeID         string `json:"tradeId"`
	TransactionType string `json:"transactionType"`
}

type UpdateTradeTransactionTypeResponse struct {
	StatusCode   string                                  `json:"statusCode"`
	StatusDesc   string                                  `json:"statusDesc"`
	TradeDetails []updateTradeTransactionTypetradeDetail `json:"TradeDetails"`
}

type updateTradeTransactionTypetradeDetail struct {
	OrderID         string `json:"orderId"`
	ID              string `json:"id"`
	TransactionType string `json:"transactionType"`
	DeveloperAPIKey string `json:"developerAPIKey"`
	OfferingID      string `json:"offeringId"`
	AccountID       string `json:"accountId"`
	PartyID         string `json:"partyId"`
	PartyType       string `json:"party_type"`
	TotalAmount     string `json:"totalAmount"`
	TotalShares     string `json:"totalShares"`
	OrderStatus     string `json:"orderStatus"`
}

// Update trade transaction type
//
// Reference: https://transactapi.readme.io/reference/updatetradetransactiontype
func (c *Client) UpdateTradeTransactionType(req *UpdateTradeTransactionTypeRequest) (*UpdateTradeTransactionTypeResponse, error) {
	return request[UpdateTradeTransactionTypeRequest, UpdateTradeTransactionTypeResponse](c, http.MethodPost, EndpointUpdateTradeTransactionType, req)
}
