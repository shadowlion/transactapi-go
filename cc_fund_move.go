package transactapi

import (
	"net/http"
)

type CCFundMoveRequest struct {
	ClientID         string `json:"clientID"`
	DeveloperAPIKey  string `json:"developerAPIKey"`
	AccountID        string `json:"accountId"`
	TradeID          string `json:"tradeId"`
	CreatedIPAddress string `json:"createdIpAddress"`
}

type CCFundMoveResponse struct {
	StatusCode         string             `json:"statusCode"`
	StatusDesc         string             `json:"statusDesc"`
	TransactionDetails transactionDetails `json:"transactionDetails"`
}

type transactionDetails struct {
	AccountID         string `json:"accountId"`
	TradeID           string `json:"tradeId"`
	OfferingID        string `json:"offeringId"`
	TotalAmount       string `json:"totalAmount"`
	Ccreferencenumber string `json:"ccreferencenumber"`
	FundStatus        string `json:"fundStatus"`
	Transactionstatus string `json:"transactionstatus"`
}

// Online Credit Card Transaction
//
// This Method has third party fees associated with it that will be charged for each use. This
// method is used to perform an online credit card transaction. Transactions are batch processed
// each business day at 6pm ET.
//
// The maximum amount per transaction is $5,000.00.
//
// Reference: https://transactapi.readme.io/reference/ccfundmove
func (c *Client) CCFundMove(req *CCFundMoveRequest) (*CCFundMoveResponse, error) {
	return request[CCFundMoveRequest, CCFundMoveResponse](c, http.MethodPost, EndpointCCFundMove, req)
}
