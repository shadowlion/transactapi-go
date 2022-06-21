package transactapi

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestGetTradeStatus(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error: %s", err)
	}

	clientId := os.Getenv("TRANSACT_API_CLIENT_ID")
	developerApiKey := os.Getenv("TRANSACT_API_DEVELOPER_API_KEY")
	tradeId := os.Getenv("GET_TRADE_STATUS_TRADE_ID")

	client := New(clientId, developerApiKey, false)
	resp, err := client.GetTradeStatus(tradeId)

	assert.NotEqual(t, clientId, "", "expecting client id to be set")
	assert.NotEqual(t, developerApiKey, "", "expecting api key to be set")
	assert.NotEqual(t, tradeId, "", "expecting trade id to be set")
	assert.NotNil(t, resp, "expecting not-nil res")
	assert.Nil(t, err, "expecting nil err")
	assert.Equal(
		t,
		tradeId,
		resp.TradeDetails[0].OrderID,
		"expecting trade id to be the same, pre/post request",
	)
}
