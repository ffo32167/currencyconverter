package ratesondate

import (
	"time"

	"github.com/ffo32167/currencyconverter/internal"
)

type ratesOnDate interface {
	RatesOnDate(date time.Time) (internal.Rate, error)
}
