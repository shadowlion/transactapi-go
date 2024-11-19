package transactapi

import (
	"testing"
)

func TestEndpoints(t *testing.T) {
	type TestCase struct {
		input Client
		want  string
	}

	testCases := []TestCase{
		{
			input: Client{
				ClientID:        "",
				DeveloperAPIKey: "",
				Sandbox:         true,
			},
			want: "https://api-sandboxdash.norcapsecurities.com/tapiv3/index.php/v3/createAccount",
		},
		{
			input: Client{
				ClientID:        "",
				DeveloperAPIKey: "",
				Sandbox:         false,
			},
			want: "https://api.norcapsecurities.com/tapiv3/index.php/v3/createAccount",
		},
	}

	for _, tc := range testCases {
		have := tc.input.getURL(EndpointCreateAccount)
		if have != tc.want {
			t.Fatalf("Have: %s, Want: %s", have, tc.want)
		}
	}
}

func TestGetCreditCardTypeSuccess(t *testing.T) {
	type TestCase struct {
		input string
		want  string
	}

	testCases := []TestCase{
		{input: "4111111111111111", want: "VI"},
		{input: "5555555555554444", want: "MC"},
		{input: "6011111111111117", want: "DI"},
	}

	for _, tc := range testCases {
		have, err := GetCreditCardType(tc.input)
		if err != nil {
			t.Fatalf("GetCreditCardType('%s') error: %s", have, err)
		}

		if have != tc.want {
			t.Fatalf("Have: %s, Want: %s", have, tc.want)
		}
	}
}
