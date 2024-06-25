package transactapi

import (
	"net/http"
)

type Client struct {
	HttpClient      *http.Client
	ClientID        string
	DeveloperAPIKey string
	Sandbox         bool
}

// Create a new instance of the Transact API client struct
func NewClient(clientId, developerApiKey string, sandbox bool) *Client {
	return &Client{
		HttpClient:      &http.Client{},
		ClientID:        clientId,
		DeveloperAPIKey: developerApiKey,
		Sandbox:         sandbox,
	}
}
