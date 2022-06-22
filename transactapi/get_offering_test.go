package transactapi

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestGetOffering(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error: %s", err)
	}

	clientId := os.Getenv("TRANSACT_API_CLIENT_ID")
	developerApiKey := os.Getenv("TRANSACT_API_DEVELOPER_API_KEY")
	offeringId := os.Getenv("GET_OFFERING_OFFERING_ID")

	client := New(clientId, developerApiKey, false)
	resp, err := client.GetOffering(offeringId)

	assert.NotEqual(t, clientId, "", "expecting client id to be set")
	assert.NotEqual(t, developerApiKey, "", "expecting api key to be set")
	assert.NotEqual(t, offeringId, "", "expecting offering id to be set")
	assert.NotNil(t, resp, "expecting not-nil res")
	assert.Nil(t, err, "expecting nil err")
	assert.Equal(
		t,
		offeringId,
		resp.OfferingDetails[0].OfferingID,
		"expecting offering id to be the same, pre/post request",
	)
}
