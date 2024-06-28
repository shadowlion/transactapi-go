package transactapi

import "net/http"

type Client struct {
	httpClient      *http.Client
	clientID        string
	developerAPIKey string
	sandbox         bool
}
