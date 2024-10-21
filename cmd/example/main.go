package main

import (
	"log"

	"github.com/shadowlion/transactapi-go"
)

func main() {
	// Initialize the client
	client := transactapi.NewClient("...", "...", false)

	// Form the data necessary to run an API call
	payload := &transactapi.ValidateAbaRoutingNumberRequest{
		ClientID:        client.ClientID,
		DeveloperApiKey: client.DeveloperAPIKey,
		RoutingNumber:   "",
	}

	// Run the API call, check for error response format
	okResp, err := client.ValidateABARoutingNumber(payload)
	if err != nil {
		// Using type assertion here, to see if it returns an API Error-formatted response first:
		if apiErr, ok := err.(*transactapi.ErrorResponse); ok {
			log.Fatalf("[%s] %s", apiErr.StatusCode, apiErr.StatusDescription)
		} else {
			log.Fatal(err.Error())
		}
		return
	}

	// Else, do something with the data
	log.Println(okResp.AccountDetails)
}
