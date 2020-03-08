package mollie

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"mollie/payment"
)

// test_E4Afehj7hGDFG6uDVvT7yPVdC7hjHM

// Mollie define API client
type Mollie struct {
	ProfileId  string
	ApiKey     string
	Option     option
	BaseURL    string
	UserAgent  string
	HTTPClient *http.Client
	Debug      bool
	Logger     *log.Logger
	// do         doFunc
}

func NewMollieClient(profileId string, opts ...Option) Mollie {

	options := &option{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	return Mollie{
		ProfileId:  profileId,
		Option:     *options,
		BaseURL:    "https://api.mollie.com/v2",
		HTTPClient: options.client,
	}
}

// https://docs.mollie.com/reference/v2/payments-api/
func (c Mollie) NewPaymentsService() *PaymentsService {
	return &PaymentsService{
		version:      "v2",
		uri:          "payments",
		c:            c,
		SequenceType: Oneoff,
	}
}

type PaymentsService struct {
	version                         string
	uri                             string
	c                               Mollie
	Amount                          Amount
	Description                     string
	RedirectUrl                     *url.URL
	WebHookUrl                      *url.URL
	PaymentMethods                  []payment.PaymentMethod
	Metadata                        Metadata
	SequenceType                    SequenceType
	CustomerId                      string
	RestrictPaymentMethodsToCountry string
}

func (ps *PaymentsService) WithAmount(amount Amount) *PaymentsService {
	ps.Amount = amount
	return ps
}

func (ps *PaymentsService) WithDescription(description string) *PaymentsService {
	ps.Description = description
	return ps
}

func (ps *PaymentsService) WithRedirectUrl(url *url.URL) *PaymentsService {
	ps.RedirectUrl = url
	return ps
}

func (ps *PaymentsService) WithWebHookUrl(url *url.URL) *PaymentsService {
	ps.WebHookUrl = url
	return ps
}

func (ps *PaymentsService) WithPaymentMethods(paymentMethods ...payment.PaymentMethod) *PaymentsService {
	ps.PaymentMethods = paymentMethods
	return ps
}

func (ps *PaymentsService) WithMetadata(metadata Metadata) *PaymentsService {
	ps.Metadata = metadata
	return ps
}

func (ps *PaymentsService) WithSequenceType(sequenceType SequenceType) *PaymentsService {
	ps.SequenceType = sequenceType
	return ps
}

func (ps *PaymentsService) WithCustomerId(customerId string) *PaymentsService {
	ps.CustomerId = customerId
	return ps
}

func (ps *PaymentsService) WithRestrictPaymentMethodsToCountry(restrictPaymentMethodsToCountry string) *PaymentsService {
	ps.RestrictPaymentMethodsToCountry = restrictPaymentMethodsToCountry
	return ps
}

func (ps *PaymentsService) Do(ctx context.Context, ) (data []byte, err error) {

	return nil, nil
}

func (c Mollie) NewMethodsService() *MethodsService {
	return &MethodsService{
		version: "v2",
		uri:     "methods",
		c:       c,
	}
}

type MethodsService struct {
	version      string
	uri          string
	c            Mollie
	amount       *Amount
	sequenceType *SequenceType
	locale       *Locale
	testmode     bool
}

func (ms *MethodsService) WithLocale(locale Locale) *MethodsService {
	ms.locale = &locale
	return ms
}

func (ms *MethodsService) WithSequenceType(sequenceType SequenceType) *MethodsService {
	ms.sequenceType = &sequenceType
	return ms
}

func (ms *MethodsService) WithAmount(amount Amount) *MethodsService {
	ms.amount = &amount
	return ms
}

func (ms *MethodsService) WithTestMode(testmode bool) *MethodsService {
	ms.testmode = testmode
	return ms
}

func (ms *MethodsService) GetMethods(ctx context.Context) (status int, data []byte, err error) {

	params := url.Values{}
	if ms.amount != nil {
		params.Add("amount[value]", ms.amount.Value)
		params.Add("amount[currency]", ms.amount.Currency)
	}

	if ms.sequenceType != nil {
		params.Add("sequenceType", sequenceTypMap[*ms.sequenceType])
	}

	if ms.sequenceType != nil {
		params.Add("locale", localeMap[*ms.locale])
	}

	// TODO billingCountry
	// TODO includeWallets

	if ms.c.Option.authmode.ApiKey == "" {
		if ms.testmode {
			params.Add("testmode", strconv.FormatBool(ms.testmode))
		}

		params.Add("profileId", ms.c.ProfileId)
	}

	endpoint := fmt.Sprintf("%s/%s", ms.c.BaseURL, ms.uri)
	if len(params) > 0 {
		endpoint = fmt.Sprintf("%s/%s?%s", ms.c.BaseURL, ms.uri, params.Encode())
	}

	log.Print(endpoint)
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return
	}

	if ms.c.Option.authmode.ApiKey != "" {
		request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ms.c.Option.authmode.ApiKey))
	}

	response, err := ms.c.HTTPClient.Do(request)
	if err != nil {
		return
	}
	status = response.StatusCode
	data, err = ioutil.ReadAll(response.Body)

	return
}
