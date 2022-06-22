package main

import "fmt"

func main() {
	clientID := ""
	developerAPIKey := ""
	sandbox := false

	client := New(clientID, developerAPIKey, sandbox)

	tradeId := ""
	resp, err := client.GetTradeStatus(tradeId)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}
