package transactapi

import (
	"testing"
)

func TestPrefix(t *testing.T) {
	type TestCase struct {
		have Client
		want string
	}

	testCases := []TestCase{
		{
			have: Client{
				ClientID:        "",
				DeveloperAPIKey: "",
				Sandbox:         true,
			},
			want: "api-Sandboxdash",
		},
		{
			have: Client{
				ClientID:        "",
				DeveloperAPIKey: "",
				Sandbox:         false,
			},
			want: "api",
		},
	}

	for _, tc := range testCases {
		have := tc.have.subdomain()
		if have != tc.want {
			t.Fatalf("Have: %s, Want: %s", have, tc.want)
		}
	}
}

func TestBaseUrl(t *testing.T) {
	type TestCase struct {
		have Client
		want string
	}

	testCases := []TestCase{
		{
			have: Client{
				ClientID:        "",
				DeveloperAPIKey: "",
				Sandbox:         true,
			},
			want: "https://api-Sandboxdash.norcapsecurities.com/tapiv3/index.php/v3",
		},
		{
			have: Client{
				ClientID:        "",
				DeveloperAPIKey: "",
				Sandbox:         false,
			},
			want: "https://api.norcapsecurities.com/tapiv3/index.php/v3",
		},
	}

	for _, tc := range testCases {
		have := tc.have.subdomain()
		if have != tc.want {
			t.Fatalf("Have: %s, Want: %s", have, tc.want)
		}
	}
}

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
