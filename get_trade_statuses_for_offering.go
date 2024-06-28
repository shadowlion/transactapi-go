package transactapi

import (
	"fmt"
	"net/http"
)

type GetTradeStatusesForOfferingRequest struct {
	ClientID        string `json:"clientID"`
	DeveloperApiKey string `json:"developerAPIKey"`
	OfferingID      string `json:"offeringId"`
	TradeStatus     string `json:"tradeStatus,omitempty"`
}

type GetTradeStatusesForOfferingResponse struct {
	StatusCode               string                    `json:"statusCode"`
	StatusDesc               string                    `json:"statusDesc"`
	OfferingPurchasedDetails []offeringPurchasedDetail `json:"Offering purchased details"`
}

type offeringPurchasedDetail struct {
	TradeID                 string `json:"tradeId"`
	AccountID               string `json:"accountId"`
	TransactionType         string `json:"transactionType"`
	TotalShares             string `json:"totalShares"`
	UnitPrice               string `json:"unitPrice"`
	TotalAmount             string `json:"totalAmount"`
	TradeStatus             string `json:"tradeStatus"`
	RRApprovalStatus        string `json:"RRApprovalStatus"`
	PrincipalApprovalStatus string `json:"PrincipalApprovalStatus"`
	Field1                  string `json:"field1"`
	Field2                  string `json:"field2"`
	Field3                  string `json:"field3"`
	EsignStatus             string `json:"esignStatus"`
}

// Returns an array of trade IDs and current trade statuses associated with a specific Offering ID.
//
// Reference: https://transactapi.readme.io/reference/gettradestatusesforoffering
func (c *Client) GetTradeStatusesForOffering(req *GetTradeStatusesForOfferingRequest) (*GetTradeStatusesForOfferingResponse, *ErrorResponse, error) {
	return request[GetTradeStatusesForOfferingRequest, GetTradeStatusesForOfferingResponse](c.httpClient, http.MethodPost, fmt.Sprintf("%s%s", c.baseURL(), EndpointGetTradeStatusesForOffering), req)
}
