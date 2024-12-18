package transactapi

import (
	"errors"
	"fmt"
	"regexp"
)

// getURL returns the full url for any particular endpoint.
func (c *Client) getURL(endpoint Endpoint) string {
	var prefix string
	if c.Sandbox {
		prefix = "api-sandboxdash"
	} else {
		prefix = "api"
	}

	return fmt.Sprintf("https://%s.norcapsecurities.com/tapiv3/index.php/v3/%s", prefix, endpoint)
}

// Function GetCreditCardType uses regex to determine if the input number is either a
// VISA (VI), MASTERCARD (MC), or DISCOVER (DI) card.
// All other cards return an error, as they are non-compliant with North Capital's transactions.
func GetCreditCardType(cardNumber string) (string, error) {
	visaRegex := regexp.MustCompile(`4\d{3}([-]?)\d{4}\d{4}\d{4}`)
	mastercardRegex := regexp.MustCompile(`5[1-5]\d{2}([-]?)\d{4}\d{4}\d{4}`)
	discoverRegex := regexp.MustCompile(`6(?:011|22(?:1([-]?(?:2[6-9]|[3-9]))|[2-8]|9([-]?(?:[01]|2[0-5])))|4[4-9]\d|5\d\d)([-]?)\d{4}\d{4}\d{4}`)

	if visaRegex.Find([]byte(cardNumber)) != nil {
		return "VI", nil
	} else if mastercardRegex.Find([]byte(cardNumber)) != nil {
		return "MC", nil
	} else if discoverRegex.Find([]byte(cardNumber)) != nil {
		return "DI", nil
	}
	return "", errors.New("Invalid card number.")
}
