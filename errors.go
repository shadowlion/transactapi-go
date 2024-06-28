package transactapi

type ErrorResponse struct {
	StatusCode        string `json:"statusCode"`
	StatusDescription string `json:"statusDesc"`
}
