package transactapi

import (
	"net/http"
)

type getTradeStatusPayload struct {
	ClientID        string `json:"clientID"`
	DeveloperAPIKey string `json:"developerAPIKey"`
	TradeID         string `json:"tradeId"`
}

type getTradeStatusResponse struct {
	StatusCode   string        `json:"statusCode"`
	StatusDesc   string        `json:"statusDesc"`
	TradeDetails []tradeDetail `json:"tradeDetails"`
}

type tradeDetail struct {
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

func (c *client) GetTradeStatus(tradeId string) (*getTradeStatusResponse, *errorResponse, error) {
	method := http.MethodPost
	endpoint := "/getTradeStatus"
	payload := getTradePayload{
		ClientID:        c.clientID,
		DeveloperAPIKey: c.developerApiKey,
		TradeID:         tradeId,
	}
	var res getTradeStatusResponse
	errRes, err := c.request(method, endpoint, payload, &res)

	if err != nil {
		return nil, nil, err
	}

	if errRes != nil {
		return nil, errRes, nil
	}

	return &res, nil, nil
}
