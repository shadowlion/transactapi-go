package endpoints

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
