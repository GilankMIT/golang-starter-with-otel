package money

type Money struct {
	Amount       int64  `json:"amount"`
	CurrencyCode uint16 `json:"currency_code"`
}
