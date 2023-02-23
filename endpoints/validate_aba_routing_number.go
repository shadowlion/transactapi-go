package endpoints

type ValidateAbaRoutingNumberRequest struct {
	ClientID        string `json:"clientID"`
	DeveloperApiKey string `json:"developerAPIKey"`
	RoutingNumber   string `json:"routingNumber"`
}

type ValidateAbaRoutingNumberResponse struct {
	StatusCode     string `json:"statusCode"`
	StatusDesc     string `json:"statusDesc"`
	AccountDetails string `json:"accountDetails"`
}
