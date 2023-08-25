package transactapi

type GetPartyRequest struct {
	ClientID        string `json:"clientID"`
	DeveloperApiKey string `json:"developerAPIKey"`
	PartyID         string `json:"partyId"`
}

type GetPartyResponse struct {
	StatusCode   string        `json:"statusCode"`
	StatusDesc   string        `json:"statusDesc"`
	PartyDetails []partyDetail `json:"partyDetails"`
}

type partyDetail struct {
	PartyID                string `json:"partyId"`
	FirstName              string `json:"firstName"`
	MiddleInitial          string `json:"middleInitial"`
	LastName               string `json:"lastName"`
	Domicile               string `json:"domicile"`
	SocialSecurityNumber   string `json:"socialSecurityNumber"`
	Dob                    string `json:"dob"`
	PrimAddress1           string `json:"primAddress1"`
	PrimAddress2           string `json:"primAddress2"`
	PrimCity               string `json:"primCity"`
	PrimState              string `json:"primState"`
	PrimZip                string `json:"primZip"`
	PrimCountry            string `json:"primCountry"`
	EmailAddress           string `json:"emailAddress"`
	EmailAddress2          string `json:"emailAddress2"`
	Phone                  string `json:"phone"`
	Phone2                 string `json:"phone2"`
	Occupation             string `json:"occupation"`
	AssociatedPerson       string `json:"associatedPerson"`
	EmpCountry             string `json:"empCountry"`
	EmpAddress1            string `json:"empAddress1"`
	EmpAddress2            string `json:"empAddress2"`
	EmpCity                string `json:"empCity"`
	EmpState               string `json:"empState"`
	EmpZip                 string `json:"empZip"`
	CurrentANNIncome       string `json:"currentAnnIncome"`
	AvgANNIncome           string `json:"avgAnnIncome"`
	CurrentHouseholdIncome string `json:"currentHouseholdIncome"`
	AvgHouseholdIncome     string `json:"avgHouseholdIncome"`
	HouseholdNetworth      string `json:"householdNetworth"`
	KycStatus              string `json:"kycStatus"`
	AmlStatus              string `json:"amlStatus"`
	AmlDate                string `json:"amlDate"`
	Tags                   string `json:"tags"`
	Notes                  string `json:"notes"`
	Field1                 string `json:"field1"`
	Field2                 string `json:"field2"`
	Field3                 string `json:"field3"`
}

// This method is used to get all information about an individual Party. The Party ID must be
// specified as a request parameter to get the party information.
//
// Reference: https://transactapi.readme.io/reference/getparty
func (c *Client) GetParty(req GetPartyRequest) (GetPartyResponse, error) {
	return PostRequest[GetPartyResponse](c.ctx, "/getParty", &req)
}
