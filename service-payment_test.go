package mollie

import (
	"context"
	"log"
	"testing"
)

func TestMethodsService_GetMethods(t *testing.T) {

	// curl -X GET "https://api.mollie.com/v2/methods?profileId=pfl_UCUhu7C7cq" -H "Authorization: Bearer test_2qhaUVDhEpvac95SN5gRdzTMcyqxNd"

	profileId := "fill in"
	apiKey := ApiKey("test_......")

	mollie := NewMollieClient(
		profileId,
		WithHTTPClient(NewHTTPClient(3000)),
		WithAuthMode(AuthMode{ApiKey: apiKey}))

	status, response, err := mollie.NewMethodsService().Do(context.Background())
	if err != nil && status != 200 {
		t.Fatalf("Dont want error, go %v", err)
	}

	log.Print(string(response))

}

func TestMethodsService_CreatePayment(t *testing.T) {

	// curl -X GET "https://api.mollie.com/v2/methods?profileId=pfl_UCUhu7C7cq" -H "Authorization: Bearer test_2qhaUVDhEpvac95SN5gRdzTMcyqxNd"

	profileId := "fill in"
	apiKey := ApiKey("test_......")

	mollie := NewMollieClient(
		profileId,
		WithHTTPClient(NewHTTPClient(3000)),
		WithAuthMode(AuthMode{ApiKey: apiKey}))

	status, response, err := mollie.NewPaymentsService().
		WithAmount(Amount{
			Value:    "10.00",
			Currency: "EUR",
		}).
		WithDescription("Payment of product").
		WithWebHookUrl(MustParse("https://webshop.example.org/order/12345")).
		WithRedirectUrl(MustParse("https://webshop.example.org/payments/webhook/")).
		WithMetadata(Metadata{"orderId": "12345"}).
		Do(context.Background())
	if err != nil && status != 200 {
		t.Fatalf("Dont want error, go %v", err)
	}

	log.Print(string(response))

}
