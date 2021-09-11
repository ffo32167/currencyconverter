package currencyfreak

import (
	"time"

	domain "github.com/ffo32167/currencyconverter"
)

type CurrencyFreak struct {
	address string
}

func (c CurrencyFreak) RatesOnDate(date time.Time) (domain.Rate, error) {
	return domain.Rate{Base: "USD", Date: date, Rates: map[string]float64{"RUB": 78.00}}, nil
}
