package transactapi

import (
	"regexp"
)

func getCreditCardType(num string) string {
	visaRegex := regexp.MustCompile(`/^4\d{3}([-]?)\d{4}\1\d{4}\1\d{4}$/`)
	mastercardRegex := regexp.MustCompile(`/^5[1-5]\d{2}([-]?)\d{4}\1\d{4}\1\d{4}$/`)
	discoverRegex := regexp.MustCompile(`/^6(?:011|22(?:1(?=[-]?(?:2[6-9]|[3-9]))|[2-8]|9(?=[-]?(?:[01]|2[0-5])))|4[4-9]\d|5\d\d)([-]?)\d{4}\1\d{4}\1\d{4}$/`)
	if visaRegex.Find([]byte(num)) != nil {
		return "VI"
	} else if mastercardRegex.Find([]byte(num)) != nil {
		return "MC"
	} else if discoverRegex.Find([]byte(num)) != nil {
		return "DI"
	}
	return ""
}
