package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const baseURL string = "https://api.norcapsecurities.com/tapiv3/index.php/v3"

type BasePayload struct {
	ClientID        string `json:"clientID"`
	DeveloperAPIKey string `json:"developerAPIKey"`
}

type GetTradeStatusPayload struct {
	ClientID        string `json:"clientID"`
	DeveloperAPIKey string `json:"developerAPIKey"`
	TradeID         string `json:"tradeId"`
}

type GetTradeStatusResponse struct {
	StatusCode   string        `json:"statusCode"`
	StatusDesc   string        `json:"statusDesc"`
	TradeDetails []TradeDetail `json:"tradeDetails"`
}

type TradeDetail struct {
	ID               string      `json:"id"`
	DeveloperAPIKey  string      `json:"developerAPIKey"`
	OfferingID       string      `json:"offeringId"`
	AccountID        string      `json:"accountId"`
	PartyID          string      `json:"partyId"`
	PartyType        string      `json:"party_type"`
	EscrowID         interface{} `json:"escrowId"`
	OrderID          string      `json:"orderId"`
	TransactionType  string      `json:"transactionType"`
	TotalAmount      string      `json:"totalAmount"`
	TotalShares      string      `json:"totalShares"`
	OrderStatus      string      `json:"orderStatus"`
	CreatedDate      string      `json:"createdDate"`
	CreatedIPAddress string      `json:"createdIpAddress"`
	Errors           string      `json:"errors"`
	DocumentKey      string      `json:"documentKey"`
	EsignStatus      string      `json:"esignStatus"`
	Users            string      `json:"users"`
	ArchivedStatus   string      `json:"archived_status"`
}

func TransactApiClient(clientId string, developerApiKey string) *BasePayload {
	return &BasePayload{ClientID: clientId, DeveloperAPIKey: developerApiKey}
}

func postRequest(url string, payload io.Reader) ([]byte, error) {
	res, err := http.Post(url, "application/json", payload)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if 200 != res.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}

	return body, err
}

func GetTradeStatus(payload *GetTradeStatusPayload) (*GetTradeStatusResponse, error) {
	url := baseURL + "/getTradeStatus"
	j, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	res, err := postRequest(url, bytes.NewBuffer(j))

	if err != nil {
		return nil, err
	}

	var data GetTradeStatusResponse
	err = json.Unmarshal(res, &data)

	if err != nil {
		return nil, err
	}

	return &data, err
}

func main() {
	payload := GetTradeStatusPayload{
		ClientID:        "",
		DeveloperAPIKey: "",
		TradeID:         "",
	}

	res, err := GetTradeStatus(&payload)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
