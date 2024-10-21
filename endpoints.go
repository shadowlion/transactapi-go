package transactapi

type Endpoint string

const (
	EndpointAddCreditCard               Endpoint = "addCreditCard"
	EndpointCCFundMove                  Endpoint = "ccFundMove"
	EndpointCreateAccount               Endpoint = "createAccount"
	EndpointCreateExternalAccount       Endpoint = "createExternalAccount"
	EndpointExternalFundMove            Endpoint = "externalFundMove"
	EndpointGetAccount                  Endpoint = "getAccount"
	EndpointGetCreditCard               Endpoint = "getCreditCard"
	EndpointGetExternalAccount          Endpoint = "getExternalAccount"
	EndpointGetOffering                 Endpoint = "getOffering"
	EndpointGetParty                    Endpoint = "getParty"
	EndpointGetTradeDocument            Endpoint = "getTradeDocument"
	EndpointGetTradeStatus              Endpoint = "getTradeStatus"
	EndpointGetTradeStatusesForOffering Endpoint = "getTradeStatusesForOffering"
	EndpointGetTradesForOffering        Endpoint = "getTradesForOffering"
	EndpointUpdateAccount               Endpoint = "updateAccount"
	EndpointUpdateCreditCard            Endpoint = "updateCreditCard"
	EndpointUpdateExternalAccount       Endpoint = "updateExternalAccount"
	EndpointUpdateTradeTransactionType  Endpoint = "updateTradeTransactionType"
	EndpointValidateABARoutingNumber    Endpoint = "validateABARoutingNumber"
)
