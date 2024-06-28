package transactapi

import "net/http"

type Client struct {
	httpClient      *http.Client
	ClientID        string
	DeveloperAPIKey string
	Sandbox         bool
}

// NewClient returns an instance of a client used to run Transact API calls.
func NewClient(clientID, developerAPIKey string, sandbox bool) *Client {
	return &Client{
		httpClient:      &http.Client{},
		ClientID:        clientID,
		DeveloperAPIKey: developerAPIKey,
		Sandbox:         sandbox,
	}
}
