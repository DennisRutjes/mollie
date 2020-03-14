package mollie

import (
	"log"
	"net/url"
)

func MustParse(rawurl string) *url.URL {
	correctURL, err := url.Parse(rawurl)
	if err != nil {
		log.Fatalf("cannot parse url : %s, err : %v", rawurl, err)
	}

	return correctURL
}
