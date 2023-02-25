package endpoints

type CreateTradeRequest struct {
	ClientID         string `json:"clientID"`
	DeveloperAPIKey  string `json:"developerAPIKey"`
	OfferingID       string `json:"offeringId"`
	AccountID        string `json:"accountId"`
	TransactionType  string `json:"transactionType"`
	TransactionUnits int64  `json:"transactionUnits"`
	CreatedIPAddress string `json:"createdIpAddress"`
}

type CreateTradeResponse struct {
	StatusCode      string                `json:"statusCode"`
	StatusDesc      string                `json:"statusDesc"`
	PurchaseDetails []purchaseDetailUnion `json:"purchaseDetails"`
}

type purchaseDetailClass struct {
	TradeID                 string `json:"tradeId"`
	TransactionID           string `json:"transactionId"`
	TransactionAmount       string `json:"transactionAmount"`
	TransactionDate         string `json:"transactionDate"`
	TransactionStatus       string `json:"transactionStatus"`
	RRApprovalStatus        string `json:"RRApprovalStatus"`
	RRName                  string `json:"RRName"`
	RRApprovalDate          string `json:"RRApprovalDate"`
	PrincipalApprovalStatus string `json:"PrincipalApprovalStatus"`
	PrincipalName           string `json:"PrincipalName"`
	PrincipalDate           string `json:"PrincipalDate"`
}

type purchaseDetailUnion struct {
	Bool                     *bool
	PurchaseDetailClassArray []purchaseDetailClass
}
