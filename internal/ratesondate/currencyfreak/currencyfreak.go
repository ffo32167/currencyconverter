package currencyfreak

import (
	"time"

	"github.com/ffo32167/currencyconverter/internal"
)

type CurrencyFreak struct {
	address string
}

func (c CurrencyFreak) RatesOnDate(date time.Time) (internal.Rate, error) {
	return internal.Rate{Base: "USD", Date: date, Rates: map[string]float64{"RUB": 78.00}}, nil
}
