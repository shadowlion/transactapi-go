package transactapi

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type CreatePartyPayload struct {
	basePayload
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

type CreatePartyResponse struct {
	baseResponse
	PartyDetails []CreatePartyPartyDetail `json:"accountDetails"`
}

type CreatePartyPartyDetail struct {
	PartyID   string `json:"partyId"`
	KycStatus string `json:"KYCstatus"`
	AmlStatus string `json:"AMLstatus"`
}

func (c *client) CreateParty(p CreatePartyPayload) (*CreatePartyResponse, error) {
	jsonData, err := json.Marshal(p)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		c.baseURL+"/createParty",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return nil, err
	}

	var fullResponse CreatePartyResponse

	if err := c.sendRequest(req, &fullResponse); err != nil {
		return nil, err
	}

	return &fullResponse, nil
}
