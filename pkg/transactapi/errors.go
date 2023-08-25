package transactapi

type TransactApiError struct {
	StatusCode        string `json:"statusCode"`
	StatusDescription string `json:"statusDesc"`
}
