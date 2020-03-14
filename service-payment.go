package mollie

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/DennisRutjes/mollie/payment"
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
		sequenceType: Oneoff,
	}
}

type PaymentsService struct {
	version                         string
	uri                             string
	c                               Mollie
	amount                          *Amount
	description                     string
	redirectUrl                     *url.URL
	webHookUrl                      *url.URL
	paymentMethods                  []payment.PaymentMethod
	metadata                        Metadata
	sequenceType                    SequenceType
	customerId                      string
	restrictPaymentMethodsToCountry string
	locale                          Locale
	testmode                        bool
}

type Payment struct {
	Amount      Amount   `json:"amount"`
	Description string   `json:"description"`
	WebhookURL  string   `json:"webhookUrl"`
	RedirectURL string   `json:"redirectUrl"`
	Metadata    Metadata `json:"metadata"`
	Locale      string   `json:"locale"`
}

func (ps *PaymentsService) WithAmount(amount Amount) *PaymentsService {
	ps.amount = &amount
	return ps
}

func (ps *PaymentsService) WithDescription(description string) *PaymentsService {
	ps.description = description
	return ps
}

func (ps *PaymentsService) WithLocale(locale Locale) *PaymentsService {
	ps.locale = locale
	return ps
}

func (ps *PaymentsService) WithRedirectUrl(url *url.URL) *PaymentsService {
	ps.redirectUrl = url
	return ps
}

func (ps *PaymentsService) WithWebHookUrl(url *url.URL) *PaymentsService {
	ps.webHookUrl = url
	return ps
}

func (ps *PaymentsService) WithPaymentMethods(paymentMethods ...payment.PaymentMethod) *PaymentsService {
	ps.paymentMethods = paymentMethods
	return ps
}

func (ps *PaymentsService) WithMetadata(metadata Metadata) *PaymentsService {
	ps.metadata = metadata
	return ps
}

func (ps *PaymentsService) WithSequenceType(sequenceType SequenceType) *PaymentsService {
	ps.sequenceType = sequenceType
	return ps
}

func (ps *PaymentsService) WithCustomerId(customerId string) *PaymentsService {
	ps.customerId = customerId
	return ps
}

func (ps *PaymentsService) WithRestrictPaymentMethodsToCountry(restrictPaymentMethodsToCountry string) *PaymentsService {
	ps.restrictPaymentMethodsToCountry = restrictPaymentMethodsToCountry
	return ps
}

func (ps *PaymentsService) WithTestMode(testmode bool) *PaymentsService {
	ps.testmode = testmode
	return ps
}

func (ps *PaymentsService) Do(ctx context.Context) (status int, data []byte, err error) {

	// params := url.Values{}
	// if ps.amount != nil {
	// 	params.Add("amount[value]", ps.amount.Value)
	// 	params.Add("amount[currency]", ps.amount.Currency)
	// }
	//
	// if st, ok := sequenceTypMap[ps.sequenceType]; ok {
	// 	params.Add("sequenceType", st)
	// }
	//
	// if l, ok := localeMap[ps.locale]; ok {
	// 	params.Add("locale", l)
	// }
	//
	// if ps.c.Option.authmode.ApiKey == "" {
	// 	if ps.testmode {
	// 		params.Add("testmode", strconv.FormatBool(ps.testmode))
	// 	}
	//
	// 	params.Add("profileId", ps.c.ProfileId)
	// }

	endpoint := fmt.Sprintf("%s/%s", ps.c.BaseURL, ps.uri)

	pay := &Payment{
		Amount:      *ps.amount,
		Description: ps.description,
		WebhookURL:  ps.webHookUrl.String(),
		RedirectURL: ps.redirectUrl.String(),
		Metadata:    ps.metadata,
	}

	if locale, ok := localeMap[ps.locale]; ok {
		pay.Locale = locale
	}

	pd, err := json.Marshal(pay)
	if err != nil {
		log.Printf("marshal payment error : %v", err)
		return
	}

	request, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewBuffer(pd))
	if err != nil {
		return
	}

	if ps.c.Option.authmode.ApiKey != "" {
		request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ps.c.Option.authmode.ApiKey))
	}

	response, err := ps.c.HTTPClient.Do(request)
	if err != nil {
		return
	}
	status = response.StatusCode
	data, err = ioutil.ReadAll(response.Body)

	return
}
