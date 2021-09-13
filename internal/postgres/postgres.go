package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type StorageRate struct {
	RateDate time.Time `json:"date"`
	CurrCode string    `json:"curr_code"`
	Rate     float64   `json:"rate"`
}

func Rate(pool *pgxpool.Pool, date string) ([]StorageRate, error) {
	rows, err := pool.Query(context.Background(),
		"SELECT rate_date,curr_code,rate FROM employee_accounting.rates r WHERE rate_date = $1 ORDER BY curr_code",
		date)
	if err != nil {
		return nil, fmt.Errorf("Unable to execute query: %w \n", err)
	}
	defer rows.Close()
	var rates []StorageRate
	var rate StorageRate
	for rows.Next() {
		err = rows.Scan(&rate.RateDate, &rate.CurrCode, &rate.Rate)
		if err != nil {
			return nil, fmt.Errorf("Unable to scan query: %w \n", err)
		}
		rates = append(rates, rate)
	}
	return rates, nil
}
