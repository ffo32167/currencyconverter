package currencyfreak

import (
	"time"

	"github.com/ffo32167/currencyconverter/rate"
)

type CurrencyFreak struct {
	address string
}

func (c CurrencyFreak) RatesOnDate(date time.Time) (rate.Rate, error) {
	return rate.Rate{Base: "USD", Date: date, Rates: map[string]float64{"RUB": 78.00}}, nil
}
