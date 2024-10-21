package transactapi

type ErrorResponse struct {
	StatusCode        string `json:"statusCode"`
	StatusDescription string `json:"statusDesc"`
}

// Returns the Status Description of an API error, in error interface format.
func (e *ErrorResponse) Error() string {
	return e.StatusDescription
}
