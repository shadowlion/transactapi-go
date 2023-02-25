package endpoints

type ExternalFundMoveRequest struct {
	ClientID         string `json:"clientID"`
	DeveloperAPIKey  string `json:"developerAPIKey"`
	AccountID        string `json:"accountId"`
	OfferingID       string `json:"offeringId"`
	TradeID          string `json:"tradeId"`
	NickName         string `json:"NickName"`
	Amount           int64  `json:"amount"`
	Description      string `json:"description"`
	CheckNumber      string `json:"checkNumber"`
	CreatedIPAddress string `json:"createdIpAddress"`
}

type ExternalFundMoveResponse struct {
	StatusCode            string                 `json:"statusCode"`
	StatusDesc            string                 `json:"statusDesc"`
	TradeFinancialDetails []tradeFinancialDetail `json:"TradeFinancialDetails"`
}

type tradeFinancialDetail struct {
	AccountID   string `json:"accountId"`
	TradeID     string `json:"tradeId"`
	OfferingID  string `json:"offeringId"`
	TotalAmount string `json:"totalAmount"`
	RefNum      string `json:"RefNum"`
	FundStatus  string `json:"fundStatus"`
}
