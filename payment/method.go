package payment

type PaymentMethod int

const (
	Unknown                = -1
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

var PaymentMethodMap = MethodMap{
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

func (mm MethodMap) GetPaymentMethod(method string) PaymentMethod {
	for k, v := range mm {
		if v == method {
			return k
		}
	}
	return Unknown
}

