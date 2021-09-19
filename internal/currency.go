package internal

import "time"

type Storage interface {
	Rate(date string) ([]Rate, error)
	Create(rates []Rate) error
}

type Rate struct {
	RateDate time.Time
	CurrCode string
	Rate     float64
}
