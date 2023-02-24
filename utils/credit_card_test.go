package utils

import (
	"testing"
)

// TestGetCreditCardTypeSuccess calls utils.GetCreditCardType with a card number, checking
// for a valid return value.
func TestGetCreditCardTypeSuccess(t *testing.T) {
	numbers := []string{
		"4111111111111111",
		"5555555555554444",
		"6011111111111117",
	}
	wanted := []string{
		"VI",
		"MC",
		"DI",
	}
	for i := 0; i < len(numbers); i++ {
		have := numbers[i]
		want := wanted[i]
		cardType, err := GetCreditCardType(have)
		if err != nil {
			t.Fatalf("GetCreditCardType('%s') error: %s", have, err)
		}
		if cardType != want {
			t.Fatalf("Have: %s, Want: %s", cardType, want)
		}
	}
}

// TestGetCreditCardTypeFailed calls utils.GetCreditCardType with a card number, checking
// for a valid return value.
func TestGetCreditCardTypeFailed(t *testing.T) {
	t.Fatalf("Unimplemented.")
}
