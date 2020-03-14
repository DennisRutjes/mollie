package mollie

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

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
	sequenceType SequenceType
	locale       Locale
	testmode     bool
}

func (ms *MethodsService) WithLocale(locale Locale) *MethodsService {
	ms.locale = locale
	return ms
}

func (ms *MethodsService) WithSequenceType(sequenceType SequenceType) *MethodsService {
	ms.sequenceType = sequenceType
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

	if st, ok := sequenceTypMap[ms.sequenceType]; ok {
		params.Add("sequenceType", st)
	}

	if l, ok := localeMap[ms.locale]; ok {
		params.Add("locale", l)
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
	request, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
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
