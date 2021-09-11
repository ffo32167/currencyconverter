package internal

import "time"

type Rate struct {
	Base  string
	Date  time.Time
	Rates map[string]float64
}

type Storage interface {
	Load(date time.Time) (Rate, error)
	Save(Rate) error
}
