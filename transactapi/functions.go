package transactapi

import "github.com/shadowlion/transactapi-go/endpoints"

// This method is used to save a credit card to an Account (createAccount). Only one credit card
// can be added for an account.
//
// Reference: https://transactapi.readme.io/reference/addcreditcard
func (c *Client) AddCreditCard(req endpoints.AddCreditCardRequest) endpoints.AddCreditCardResponse {
	return basePostRequest[endpoints.AddCreditCardResponse](c, "/addCreditCard", &req)
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

// Reference: https://transactapi.readme.io/reference/externalfundmove
func (c *Client) ExternalFundMove(req endpoints.ExternalFundMoveRequest) endpoints.ExternalFundMoveResponse {
	return basePostRequest[endpoints.ExternalFundMoveResponse](c, "/externalFundMove", &req)
}

// This method is used to get all information for an account (createAccount).
//
// Reference: https://transactapi.readme.io/reference/getaccount
func (c *Client) GetAccount(req endpoints.GetAccountRequest) endpoints.GetAccountResponse {
	return basePostRequest[endpoints.GetAccountResponse](c, "/getAccount", &req)
}

// This method is used to retrieve the credit card information that was previously saved to a
// specific account.
//
// Reference: https://transactapi.readme.io/reference/getcreditcard
func (c *Client) GetCreditCard(req endpoints.GetCreditCardRequest) endpoints.GetCreditCardResponse {
	return basePostRequest[endpoints.GetCreditCardResponse](c, "/getCreditCard", &req)
}

// Reference: https://transactapi.readme.io/reference/getexternalaccount
func (c *Client) GetExternalAccount(req endpoints.GetExternalAccountRequest) endpoints.GetExternalAccountResponse {
	return basePostRequest[endpoints.GetExternalAccountResponse](c, "/getExternalAccount", &req)
}

// This method is used to get all the details of an offering. The Offering ID is required to get
// the information.
//
// Reference: https://transactapi.readme.io/reference/getoffering
func (c *Client) GetOffering(req endpoints.GetOfferingRequest) endpoints.GetOfferingResponse {
	return basePostRequest[endpoints.GetOfferingResponse](c, "/getOffering", &req)
}

// This method is used to get all information about an individual Party. The Party ID must be
// specified as a request parameter to get the party information.
//
// Reference: https://transactapi.readme.io/reference/getparty
func (c *Client) GetParty(req endpoints.GetPartyRequest) endpoints.GetPartyResponse {
	return basePostRequest[endpoints.GetPartyResponse](c, "/getParty", &req)
}

// Reference: https://transactapi.readme.io/reference/gettradedocument
func (c *Client) GetTradeDocument(req endpoints.GetTradeDocumentRequest) endpoints.GetTradeDocumentResponse {
	return basePostRequest[endpoints.GetTradeDocumentResponse](c, "/getTradeDocument", &req)
}

// This method is used to get the current trade details as an array with the current trade status.
//
// Reference: https://transactapi.readme.io/reference/gettradestatus
func (c *Client) GetTradeStatus(req endpoints.GetTradeStatusRequest) endpoints.GetTradeStatusResponse {
	return basePostRequest[endpoints.GetTradeStatusResponse](c, "/getTradeStatus", &req)
}

// Returns an array of trade IDs and current trade statuses associated with a specific Offering ID.
//
// Reference: https://transactapi.readme.io/reference/gettradestatusesforoffering
func (c *Client) GetTradeStatusesForOffering(req endpoints.GetTradeStatusesForOfferingRequest) endpoints.GetTradeStatusesForOfferingResponse {
	return basePostRequest[endpoints.GetTradeStatusesForOfferingResponse](c, "/getTradeStatusesForOffering", &req)
}

// This method is used to retrieve the history of all trades (and details of the trades) created
// for an offering. The Offering ID is required as a request parameter to fetch the purchase
// history.
//
// Reference: https://transactapi.readme.io/reference/gettradesforoffering
func (c *Client) GetTradesForOffering(req endpoints.GetTradesForOfferingRequest) endpoints.GetTradesForOfferingResponse {
	return basePostRequest[endpoints.GetTradesForOfferingResponse](c, "/getTradesForOffering", &req)
}

// This method is used to update a specific account (updateAccount)
//
// Reference: https://transactapi.readme.io/reference/updateaccount
func (c *Client) UpdateAccount(req endpoints.UpdateAccountRequest) endpoints.UpdateAccountResponse {
	return basePutRequest[endpoints.UpdateAccountResponse](c, "/updateAccount", &req)
}

// This method is used to update the credit card information that is saved to a specific account
// (createExternalAccount).
//
// Reference: https://transactapi.readme.io/reference/updatecreditcard
func (c *Client) UpdateCreditCard(req endpoints.UpdateCreditCardRequest) endpoints.UpdateCreditCardResponse {
	return basePostRequest[endpoints.UpdateCreditCardResponse](c, "/updateCreditCard", &req)
}

// This method is used to update fields related to a particular external account for an Account
// (createAccount). The Account ID must be specified as a request parameter to update the record.
//
// Reference: https://transactapi.readme.io/reference/updateexternalaccount
func (c *Client) UpdateExternalAccount(req endpoints.UpdateExternalAccountRequest) endpoints.UpdateExternalAccountResponse {
	return basePostRequest[endpoints.UpdateExternalAccountResponse](c, "/updateExternalAccount", &req)
}

// Update trade transaction type
//
// Reference: https://transactapi.readme.io/reference/updatetradetransactiontype
func (c *Client) UpdateTradeTransactionType(req endpoints.UpdateTradeTransactionTypeRequest) endpoints.UpdateTradeTransactionTypeResponse {
	return basePostRequest[endpoints.UpdateTradeTransactionTypeResponse](c, "/updateTradeTransactionType", &req)
}

// This method is used to validate the routing number for an external account
// (createExternalAccount)
//
// Reference: https://transactapi.readme.io/reference/validateabaroutingnumber
func (c *Client) ValidateABARoutingNumber(req endpoints.ValidateAbaRoutingNumberRequest) endpoints.ValidateAbaRoutingNumberResponse {
	return basePostRequest[endpoints.ValidateAbaRoutingNumberResponse](c, "/validateABARoutingNumber", &req)
}
