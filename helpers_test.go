package transactapi

import (
	"testing"
)

func TestPrefix(t *testing.T) {
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
			want: "api-sandboxdash",
		},
		{
			input: Client{
				ClientID:        "",
				DeveloperAPIKey: "",
				Sandbox:         false,
			},
			want: "api",
		},
	}

	for _, tc := range testCases {
		have := tc.input.subdomain()
		if have != tc.want {
			t.Fatalf("Have: %s, Want: %s", have, tc.want)
		}
	}
}

func TestBaseUrl(t *testing.T) {
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
			want: "https://api-sandboxdash.norcapsecurities.com/tapiv3/index.php/v3",
		},
		{
			input: Client{
				ClientID:        "",
				DeveloperAPIKey: "",
				Sandbox:         false,
			},
			want: "https://api.norcapsecurities.com/tapiv3/index.php/v3",
		},
	}

	for _, tc := range testCases {
		have := tc.input.baseURL()
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
