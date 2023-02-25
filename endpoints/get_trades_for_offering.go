package endpoints

type GetTradesForOfferingRequest struct {
	ClientID        string `json:"clientID"`
	DeveloperApiKey string `json:"developerAPIKey"`
	OfferingID      string `json:"offeringId"`
}

type GetTradesForOfferingResponse struct {
	StatusCode               string                    `json:"statusCode"`
	StatusDesc               string                    `json:"statusDesc"`
	OfferingPurchasedDetails []offeringPurchaseHistory `json:"Offering purchased details"`
}

type offeringPurchaseHistory struct {
	TradeID             string `json:"tradeId"`
	AccountID           string `json:"accountId"`
	OfferingID          string `json:"offeringId"`
	TotalAmount         string `json:"totalAmount"`
	TotalShares         string `json:"totalShares"`
	PurchaseDate        string `json:"purchaseDate"`
	AccountName         string `json:"accountName"`
	OfferingTotalShares string `json:"offeringTotalShares"`
	TargetAmount        string `json:"targetAmount"`
	RemainingShares     string `json:"remainingShares"`
}