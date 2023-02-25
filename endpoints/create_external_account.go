package endpoints

type CreateExternalAccountRequest struct {
	ClientID           string `json:"clientID"`
	DeveloperAPIKey    string `json:"developerAPIKey"`
	Types              string `json:"types"`
	AccountID          string `json:"accountId"`
	ExtAccountfullname string `json:"ExtAccountfullname"`
	Extnickname        string `json:"Extnickname"`
	ExtRoutingnumber   string `json:"ExtRoutingnumber"`
	ExtAccountnumber   string `json:"ExtAccountnumber"`
	UpdatedIPAddress   string `json:"updatedIpAddress"`
}

type CreateExternalAccountResponse struct {
	StatusCode             string `json:"statusCode"`
	StatusDesc             string `json:"statusDesc"`
	ExternalAccountDetails []externalAccountDetailElement
}

type externalAccountDetailElement struct {
	Bool                       *bool
	ExternalAccountDetailClass *externalAccountDetailClass
}

type externalAccountDetailClass struct {
	AccountID          string `json:"accountId"`
	EXTAccountfullname string `json:"ExtAccountfullname"`
	Extnickname        string `json:"Extnickname"`
	ExtRoutingnumber   string `json:"ExtRoutingnumber"`
	ExtAccountnumber   string `json:"ExtAccountnumber"`
	AccountType        string `json:"accountType"`
	Types              string `json:"types"`
}
