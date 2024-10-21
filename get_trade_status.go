package transactapi

import (
	"net/http"
)

type GetTradeStatusRequest struct {
	ClientID        string `json:"clientID"`
	DeveloperApiKey string `json:"developerAPIKey"`
	TradeID         string `json:"tradeId"`
}

type GetTradeStatusResponse struct {
	StatusCode   string        `json:"statusCode"`
	StatusDesc   string        `json:"statusDesc"`
	TradeDetails []tradeDetail `json:"tradeDetails"`
}

type tradeDetail struct {
	ID               string      `json:"id"`
	DeveloperAPIKey  string      `json:"developerAPIKey"`
	OfferingID       string      `json:"offeringId"`
	AccountID        string      `json:"accountId"`
	PartyID          string      `json:"partyId"`
	PartyType        string      `json:"party_type"`
	EscrowID         interface{} `json:"escrowId"`
	OrderID          string      `json:"orderId"`
	TransactionType  string      `json:"transactionType"`
	TotalAmount      string      `json:"totalAmount"`
	TotalShares      string      `json:"totalShares"`
	OrderStatus      string      `json:"orderStatus"`
	CreatedDate      string      `json:"createdDate"`
	CreatedIPAddress string      `json:"createdIpAddress"`
	Errors           string      `json:"errors"`
	DocumentKey      string      `json:"documentKey"`
	EsignStatus      string      `json:"esignStatus"`
	Users            string      `json:"users"`
	ArchivedStatus   string      `json:"archived_status"`
	Field1           string      `json:"field1"`
	Field2           string      `json:"field2"`
	Field3           string      `json:"field3"`
}

// This method is used to get the current trade details as an array with the current trade status.
//
// Reference: https://transactapi.readme.io/reference/gettradestatus
func (c *Client) GetTradeStatus(req *GetTradeStatusRequest) (*GetTradeStatusResponse, error) {
	return request[GetTradeStatusRequest, GetTradeStatusResponse](c, http.MethodPost, EndpointGetTradeStatus, req)
}
