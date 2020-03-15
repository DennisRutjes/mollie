package mollie

import (
	"context"
	"log"
	"testing"
)

var test_key = "test_..."

/*
 * curl -X GET "https://api.mollie.com/v2/methods?profileId=pfl_UCUhu7C7cq" -H "Authorization: Bearer test_2qhaUVDhEpvac95SN5gRdzTMcyqxNd"
 */
func TestMethodsService_GetMethods(t *testing.T) {

	profileId := "fill in"
	apiKey := ApiKey(test_key)

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

/*
 *	curl -X POST https://api.mollie.com/v2/payments \
 *	   -H "Authorization: Bearer test_dHar4XY7LxsDOtmnkVtjNVWXLSlXsM" \
 *	   -d "amount[currency]=EUR" \
 *	   -d "amount[value]=10.00" \
 *	   -d "description=Order #12345" \
 *	   -d "redirectUrl=https://webshop.example.org/order/12345/" \
 *	   -d "webhookUrl=https://webshop.example.org/payments/webhook/" \
 *	   -d "metadata={\"order_id\": \"12345\"}"
 */
func TestMethodsService_Payment(t *testing.T) {

	profileId := "fill in"
	apiKey := ApiKey(test_key)

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
		WithLocale(NL_NL).
		// todo add payment specific data, for now leave this blank mollie wil take care of it
		// WithPaymentMethods(payment.Ideal, payment.Inghomepay).
		WithWebHookUrl(MustParse("https://webshop.example.org/order/12345")).
		WithRedirectUrl(MustParse("https://webshop.example.org/payments/webhook/")).
		WithMetadata(Metadata{"orderId": "12345"}).
		Do(context.Background())
	if err != nil && status != 200 {
		t.Fatalf("Dont want error, go %v", err)
	}

	log.Print(string(response))

}
