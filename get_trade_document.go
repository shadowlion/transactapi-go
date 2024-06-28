package transactapi

import (
	"fmt"
	"net/http"
)

type GetTradeDocumentRequest struct {
	ClientID        string `json:"clientID"`
	DeveloperApiKey string `json:"developerAPIKey"`
	TradeID         string `json:"tradeId"`
}

type GetTradeDocumentResponse struct {
	StatusCode      string           `json:"statusCode"`
	StatusDesc      string           `json:"statusDesc"`
	DocumentDetails []documentDetail `json:"document_details"`
}

type documentDetail struct {
	TradeID                   string `json:"tradeId"`
	Documentid                string `json:"documentid"`
	DocumentTitle             string `json:"documentTitle"`
	DocumentFileName          string `json:"documentFileName"`
	DocumentFileReferenceCode string `json:"documentFileReferenceCode"`
	CreatedDate               string `json:"createdDate"`
	DocumentURL               string `json:"documentUrl"`
}

// Reference: https://transactapi.readme.io/reference/gettradedocument
func (c *Client) GetTradeDocument(req *GetTradeDocumentRequest) (*GetTradeDocumentResponse, *ErrorResponse, error) {
	return request[GetTradeDocumentRequest, GetTradeDocumentResponse](c.httpClient, http.MethodPost, fmt.Sprintf("%s%s", c.baseURL(), EndpointGetTradeDocument), req)
}
