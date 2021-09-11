package ratesondate

import (
	"time"

	"github.com/ffo32167/currencyconverter/rate"
)

type ratesOnDate interface {
	RatesOnDate(date time.Time) (rate.Rate, error)
}
