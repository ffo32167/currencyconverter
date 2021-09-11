package ratesondate

import (
	"time"

	domain "github.com/ffo32167/currencyconverter"
)

type ratesOnDate interface {
	RatesOnDate(date time.Time) (domain.Rate, error)
}
