# mollie


`
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
		WithPaymentMethods(payment.Ideal, payment.Inghomepay).
		WithWebHookUrl(MustParse("https://webshop.example.org/order/12345")).
		WithRedirectUrl(MustParse("https://webshop.example.org/payments/webhook/")).
		WithMetadata(Metadata{"orderId": "12345"}).
		Do(context.Background())

`