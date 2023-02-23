package utils

import (
	"errors"
	"regexp"
)

// Function GetCreditCardType uses regex to determine if the input number is either a VISA,
// MASTERCARD, or DISCOVER. All other cards return an error, as they are non-compliant with North
// Capital's transactions.
func GetCreditCardType(num string) (string, error) {
	visaRegex := regexp.MustCompile(`/^4\d{3}([-]?)\d{4}\1\d{4}\1\d{4}$/`)
	mastercardRegex := regexp.MustCompile(`/^5[1-5]\d{2}([-]?)\d{4}\1\d{4}\1\d{4}$/`)
	discoverRegex := regexp.MustCompile(`/^6(?:011|22(?:1(?=[-]?(?:2[6-9]|[3-9]))|[2-8]|9(?=[-]?(?:[01]|2[0-5])))|4[4-9]\d|5\d\d)([-]?)\d{4}\1\d{4}\1\d{4}$/`)

	if visaRegex.Find([]byte(num)) != nil {
		return "VI", nil
	} else if mastercardRegex.Find([]byte(num)) != nil {
		return "MC", nil
	} else if discoverRegex.Find([]byte(num)) != nil {
		return "DI", nil
	}
	return "", errors.New("Invalid card number.")
}
