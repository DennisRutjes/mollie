package payment

type PaymentMethod int

const (
	Applepay PaymentMethod = iota
	Bancontact
	Banktransfer
	Belfius
	Creditcard
	Directdebit
	Eps
	Giftcard
	Giropay
	Ideal
	Inghomepay
	Kbc
	Mybank
	Paypal
	Paysafecard
	Przelewy24
	Sofort
)

type MethodMap map[PaymentMethod]string

var localeMap = MethodMap{
	Applepay:     "applepay",
	Bancontact:   "bancontact",
	Banktransfer: "banktransfer",
	Belfius:      "belfius",
	Creditcard:   "creditcard",
	Directdebit:  "directdebit",
	Eps:          "eps",
	Giftcard:     "giftcard",
	Giropay:      "giropay",
	Ideal:        "ideal",
	Inghomepay:   "Inghomepay",
	Kbc:          "kbc",
	Mybank:       "mybank",
	Paypal:       "paypal",
	Paysafecard:  "paysafecard",
	Przelewy24:   "przelewy24",
	Sofort:       "sofort",
}