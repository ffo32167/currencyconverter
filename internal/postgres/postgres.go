package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ffo32167/currencyconverter/internal"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PgDb struct {
	pool *pgxpool.Pool
}

type PgRates []pgCurr

type pgCurr struct {
	RateDate time.Time `json:"date"`
	CurrCode string    `json:"curr_code"`
	Rate     float64   `json:"rate"`
}

func New(ctx context.Context, connStr string) (PgDb, error) {
	db, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return PgDb{}, fmt.Errorf("Unable to connect to database: %w", err)
	}
	//	defer db.Close()
	return PgDb{pool: db}, nil
}

func (db PgDb) Rate(ctx context.Context, date time.Time) ([]internal.Rate, error) {
	rows, err := db.pool.Query(context.Background(),
		`SELECT rate_date,curr_code,rate FROM rates.rates r WHERE rate_date = $1 ORDER BY curr_code`,
		date)
	if err != nil {
		return nil, fmt.Errorf("unable to execute select query: %w ", err)
	}
	defer rows.Close()
	var rates PgRates
	var rate pgCurr
	for rows.Next() {
		err = rows.Scan(&rate.RateDate, &rate.CurrCode, &rate.Rate)
		if err != nil {
			return nil, fmt.Errorf("unable to scan query: %w ", err)
		}
		rates = append(rates, rate)
	}
	internalRates, err := rates.toDomain()
	if err != nil {
		return internalRates, fmt.Errorf("unable to convert PG Rates to internal Rates: %w ", err)
	}
	return internalRates, nil
}

func (db PgDb) Create(ctx context.Context, internalRates []internal.Rate) error {
	for _, v := range internalRates {
		ct, err := db.pool.Exec(ctx,
			"INSERT INTO rates.rates(rate_date,curr_code,rate) VALUES($1,$2,$3)",
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

func (rates PgRates) toDomain() ([]internal.Rate, error) {
	result := make([]internal.Rate, len(rates), len(rates))
	for i := range rates {
		result[i].CurrCode = rates[i].CurrCode
		result[i].Rate = rates[i].Rate
		result[i].RateDate = rates[i].RateDate
	}
	return result, nil
}
