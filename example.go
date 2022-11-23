package main

import (
	"fmt"

	"github.com/shadowlion/transactapi-go/transactapi"
)

func main() {
	clientID := ""
	developerAPIKey := ""
	sandbox := false

	client := transactapi.New(clientID, developerAPIKey, sandbox)

	tradeId := ""
	resp, _, err := client.GetTradeStatus(tradeId)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}
