package transactapi

import (
	"net/http"
)

type getTradePayload struct {
	ClientID        string `json:"clientID"`
	DeveloperAPIKey string `json:"developerAPIKey"`
	AccountID       string `json:"accountId"`
	TradeID         string `json:"tradeId"`
}

type getTradeResponse struct {
	StatusCode   string        `json:"statusCode"`
	StatusDesc   string        `json:"statusDesc"`
	PartyDetails []partyDetail `json:"partyDetails"`
}

type partyDetail struct {
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

func (c *client) GetTrade(accountId string, tradeId string) (*getTradeResponse, *errorResponse, error) {
	method := http.MethodPost
	endpoint := "/getTrade"
	payload := getTradePayload{
		ClientID:        c.clientID,
		DeveloperAPIKey: c.developerApiKey,
		AccountID:       accountId,
		TradeID:         tradeId,
	}
	var res getTradeResponse
	errRes, err := c.request(method, endpoint, payload, &res)

	if err != nil {
		return nil, nil, err
	}

	if errRes != nil {
		return nil, errRes, nil
	}

	return &res, nil, nil
}
