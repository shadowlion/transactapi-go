package transactapi

import "github.com/shadowlion/transactapi-go/endpoints"

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

// This method is used to create a trade/investment for an offering. This requires Account ID and
// the total number of units/shares to be purchased by the account. Creating a trade represents the
// intention to invest and does NOT initiate any sort of fund move. To initiate an ACH transfer for
// a trade, you will need to use the externalFundMove method.
//
// Reference: https://transactapi.readme.io/reference/createtrade
func (c *Client) CreateTrade(req endpoints.CreateTradeRequest) endpoints.CreateTradeResponse {
	return basePostRequest[endpoints.CreateTradeResponse](c, "/createTrade", &req)
}
