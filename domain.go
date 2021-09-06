package domain

import "time"

type Rate struct {
	Base  string
	Date  time.Time
	Rates map[string]float64
}
