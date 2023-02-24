package endpoints

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
