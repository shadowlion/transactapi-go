package endpoints

type GetExternalAccountRequest struct {
	ClientID        string `json:"clientID"`
	DeveloperApiKey string `json:"developerAPIKey"`
	Types           string `json:"types"`
	AccountID       string `json:"accountId"`
}

type GetExternalAccountResponse struct {
	StatusCode string     `json:"statusCode"`
	AccountID  string     `json:"accountId"`
	StatusDesc statusDesc `json:"statusDesc"`
}

type statusDesc struct {
	AccountName          string `json:"AccountName"`
	AccountNickName      string `json:"AccountNickName"`
	AccountRoutingNumber string `json:"AccountRoutingNumber"`
	AccountNumber        string `json:"AccountNumber"`
	AccountType          string `json:"accountType"`
	UpdatedDate          string `json:"updatedDate"`
	CreatedDate          string `json:"createdDate"`
}
