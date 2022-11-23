package transactapi

import (
	"net/http"
)

type createPartyPayload struct {
	ClientID        string `json:"clientID"`
	DeveloperAPIKey string `json:"developerAPIKey"`
	DateOfBirth     string `json:"dob"`
	Domicile        string `json:"domicile"`
	EmailAddress    string `json:"emailAddress"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	PrimaryAddress1 string `json:"primAddress1"`
	PrimaryAddress2 string `json:"primAddress2,omitempty"`
	PrimaryCity     string `json:"primCity"`
	PrimaryCountry  string `json:"primCountry"`
	PrimaryState    string `json:"primState"`
	PrimaryZip      string `json:"primZip"`
}

type createPartyResponse struct {
	StatusCode        string                   `json:"statusCode"`
	StatusDescription string                   `json:"statusDesc"`
	PartyDetails      []createPartyPartyDetail `json:"accountDetails"`
}

type createPartyPartyDetail struct {
	PartyID   string `json:"partyId"`
	KycStatus string `json:"KYCstatus"`
	AmlStatus string `json:"AMLstatus"`
}

func (c *client) CreateParty(payload createPartyPayload) (*createPartyResponse, *errorResponse, error) {
	method := http.MethodPost
	endpoint := "/createParty"
	// payload := createPartyPayload{}
	var res createPartyResponse
	errRes, err := c.request(method, endpoint, payload, &res)

	if err != nil {
		return nil, nil, err
	}

	if errRes != nil {
		return nil, errRes, nil
	}

	return &res, nil, nil
}
