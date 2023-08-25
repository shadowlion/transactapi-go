package transactapi

import (
	"context"
	"net/http"
)

type Client struct {
	ctx             *context.Context
	httpClient      *http.Client
	clientID        string
	developerAPIKey string
	sandbox         bool
}

// Create a new instance of the Transact API client struct
func New(clientId, developerApiKey string, sandbox bool) *Client {
	ctx := context.Background()

	return &Client{
		ctx:             &ctx,
		httpClient:      &http.Client{},
		clientID:        clientId,
		developerAPIKey: developerApiKey,
		sandbox:         sandbox,
	}
}
