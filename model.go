package mollie

import (
	"time"
)

type AutoGenerated struct {
	Resource                        string      `json:"resource"`
	ID                              string      `json:"id"`
	Mode                            string      `json:"mode"`
	CreatedAt                       time.Time   `json:"createdAt"`
	Amount                          Amount      `json:"amount"`
	Description                     string      `json:"description"`
	Method                          interface{} `json:"method"`
	Metadata                        Metadata    `json:"metadata"`
	Status                          string      `json:"status"`
	IsCancelable                    bool        `json:"isCancelable"`
	Locale                          string      `json:"locale"`
	RestrictPaymentMethodsToCountry string      `json:"restrictPaymentMethodsToCountry"`
	ExpiresAt                       time.Time   `json:"expiresAt"`
	Details                         interface{} `json:"details"`
	ProfileID                       string      `json:"profileId"`
	SequenceType                    string      `json:"sequenceType"`
	RedirectURL                     string      `json:"redirectUrl"`
	WebhookURL                      string      `json:"webhookUrl"`
	Links                           Links       `json:"_links"`
}

type Metadata map[string]interface{}

type Amount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type Href struct {
	Href string `json:"href"`
	Type string `json:"type"`
}

type Links struct {
	Self          Href `json:"self"`
	Checkout      Href `json:"checkout"`
	Documentation Href `json:"documentation"`
}

type LocaleMap map[Locale]string

var localeMap = LocaleMap{
	EN_US: "en_US",
	NL_NL: "nl_NL",
	NL_BE: "nl_BE",
	FR_FR: "fr_FR",
	FR_BE: "fr_BE",
	DE_DE: "de_DE",
	DE_AT: "de_AT",
	DE_CH: "de_CH",
	ES_ES: "es_ES",
	CA_ES: "ca_ES",
	PT_PT: "pt_PT",
	IT_IT: "it_IT",
	NB_NO: "nb_NO",
	SV_SE: "sv_SE",
	FI_FI: "fi_FI",
	DA_DK: "da_DK",
	IS_IS: "is_IS",
	HU_HU: "hu_HU",
	PL_PL: "pl_PL",
	LV_LV: "lv_LV",
	LT_LT: "lt_LT",
}

type Locale int

const (
	EN_US Locale = iota
	NL_NL
	NL_BE
	FR_FR
	FR_BE
	DE_DE
	DE_AT
	DE_CH
	ES_ES
	CA_ES
	PT_PT
	IT_IT
	NB_NO
	SV_SE
	FI_FI
	DA_DK
	IS_IS
	HU_HU
	PL_PL
	LV_LV
	LT_LT
)

var listOfSequenTypes = []string{"oneoff", "first", "recurring"}

type SequenceType int

const (
	Oneoff SequenceType = iota
	First
	Recurring
)

type SequenceTypeMap map[SequenceType]string

var sequenceTypMap = SequenceTypeMap{
	Oneoff:    "oneoff",
	First:     "first",
	Recurring: "recurring",
}
