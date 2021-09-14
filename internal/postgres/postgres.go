package postgres

import (
	"context"
	"errors"
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
		return nil, fmt.Errorf("unable to execute select query: %w ", err)
	}
	defer rows.Close()
	var rates []StorageRate
	var rate StorageRate
	for rows.Next() {
		err = rows.Scan(&rate.RateDate, &rate.CurrCode, &rate.Rate)
		if err != nil {
			return nil, fmt.Errorf("unable to scan query: %w ", err)
		}
		rates = append(rates, rate)
	}
	return rates, nil
}

func Create(pool *pgxpool.Pool, rates []StorageRate) error {
	for _, v := range rates {
		ct, err := pool.Exec(context.Background(),
			"INSERT INTO employee_accounting.rates(rate_date,curr_code,rate) VALUES($1,$2,$3)",
			v.RateDate, v.CurrCode, v.Rate)
		if err != nil {
			return fmt.Errorf("unable to execute insert query: %w ", err)
		}
		if ct.RowsAffected() == 0 {
			return errors.New("execution of insert query affected 0 rows")
		}
	}
	return nil
}
