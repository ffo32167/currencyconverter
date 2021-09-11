package rate

import "time"

type Rate struct {
	Date     time.Time `db:"date"`
	CurrCode string    `db:"curr_code"`
	ExchRate float64   `db:"exchange_rate"`
}
