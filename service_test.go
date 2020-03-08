package mollie

import (
	"context"
	"log"
	"testing"
)

func TestMethodsService_GetMethods(t *testing.T) {

	// curl -X GET "https://api.mollie.com/v2/methods?profileId=pfl_UCUhu7C7cq" -H "Authorization: Bearer test_2qhaUVDhEpvac95SN5gRdzTMcyqxNd"

	profileId := "fill in"
	apiKey := ApiKey("test_dHar4XY7LxsDOtmnkVtjNVWXLSlXsM")

	mollie := NewMollieClient(
		profileId,
		WithHTTPClient(NewHTTPClient(3000)),
		WithAuthMode(AuthMode{ApiKey: apiKey}))

	status, response, err := mollie.NewMethodsService().WithTestMode(true).GetMethods(context.Background())
	if err != nil && status != 200 {
		t.Fatalf("Dont want error, go %v", err)
	}

	log.Print(string(response))

}
