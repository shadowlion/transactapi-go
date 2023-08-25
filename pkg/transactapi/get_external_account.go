package transactapi

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

// Reference: https://transactapi.readme.io/reference/getexternalaccount
func (c *Client) GetExternalAccount(req GetExternalAccountRequest) (GetExternalAccountResponse, error) {
	return PostRequest[GetExternalAccountResponse](c.ctx, "/getExternalAccount", &req)
}
