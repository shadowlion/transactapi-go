package transactapi

import (
	"net/http"

	"github.com/shadowlion/transactapi-go/endpoints"
)

type Client struct {
	HttpClient      *http.Client
	ClientID        string
	DeveloperAPIKey string
	Sandbox         bool
}

// Create a new instance of the Transact API client struct
func New(clientId, developerApiKey string, sandbox bool) *Client {
	return &Client{
		HttpClient:      &http.Client{},
		ClientID:        clientId,
		DeveloperAPIKey: developerApiKey,
		Sandbox:         sandbox,
	}
}

// This method is used to save a credit card to an Account (createAccount). Only one credit card
// can be added for an account.
//
// Reference: https://transactapi.readme.io/reference/addcreditcard
func (c *Client) AddCreditCard(req endpoints.AddCreditCardRequest) endpoints.AddCreditCardResponse {
	return basePostRequest[endpoints.AddCreditCardResponse](c, "/addCreditCard", &req)
}

// This method is used to validate the routing number for an external account
// (createExternalAccount)
//
// Reference: https://transactapi.readme.io/reference/validateabaroutingnumber
func (c *Client) ValidateABARoutingNumber(req endpoints.ValidateAbaRoutingNumberRequest) endpoints.ValidateAbaRoutingNumberResponse {
	return basePostRequest[endpoints.ValidateAbaRoutingNumberResponse](c, "/validateABARoutingNumber", &req)
}

// Online Credit Card Transaction
//
// This Method has third party fees associated with it that will be charged for each use. This
// method is used to perform an online credit card transaction. Transactions are batch processed
// each business day at 6pm ET.
//
// The maximum amount per transaction is $5,000.00.
//
// Reference: https://transactapi.readme.io/reference/ccfundmove
func (c *Client) CCFundMove(req endpoints.CCFundMoveRequest) endpoints.CCFundMoveResponse {
	return basePostRequest[endpoints.CCFundMoveResponse](c, "/ccFundMove", &req)
}

// This method is used to create an account that can be linked (createLink) to an individual party
// (createParty) or an entity (createEntity).
//
// Reference: https://transactapi.readme.io/reference/createaccount
func (c *Client) CreateAccount(req endpoints.CreateAccountRequest) endpoints.CreateAccountResponse {
	return basePostRequest[endpoints.CreateAccountResponse](c, "/createAccount", &req)
}
