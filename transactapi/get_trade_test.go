package transactapi

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestGetTrade(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error: %s", err)
	}

	clientId := os.Getenv("TRANSACT_API_CLIENT_ID")
	developerApiKey := os.Getenv("TRANSACT_API_DEVELOPER_API_KEY")
	accountId := os.Getenv("GET_TRADE_ACCOUNT_ID")
	tradeId := os.Getenv("GET_TRADE_TRADE_ID")

	client := New(clientId, developerApiKey, false)
	resp, err := client.GetTrade(accountId, tradeId)

	assert.NotEqual(t, clientId, "", "expecting client id to be set")
	assert.NotEqual(t, developerApiKey, "", "expecting api key to be set")
	assert.NotEqual(t, accountId, "", "expecting account id to be set")
	assert.NotEqual(t, tradeId, "", "expecting trade id to be set")
	assert.NotNil(t, resp, "expecting not-nil res")
	assert.Nil(t, err, "expecting nil err")
	assert.Equal(
		t,
		tradeId,
		resp.PartyDetails[0].OrderID,
		"expecting trade id to be the same, pre/post request",
	)
}
