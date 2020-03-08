package mollie

import (
	"net/http"
)

type AuthMode struct {
	ApiKey ApiKey
}

// options
type option struct {
	locale   string
	testmode bool
	client   *http.Client
	authmode AuthMode
}

type ApiKey string

type Option func(*option)

// todo throw error if wrong locale or default to en_US
func WithLocale(locale Locale) Option {
	return func(o *option) {
		o.locale = localeMap[locale]
	}
}

func WithTestmode(testmode bool) Option {
	return func(o *option) {
		o.testmode = testmode
	}
}

func WithHTTPClient(client *http.Client) Option {
	return func(o *option) {
		o.client = client
	}
}

func WithAuthMode(authmode AuthMode) Option {
	return func(o *option) {
		o.authmode = authmode
	}
}
