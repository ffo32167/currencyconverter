package postgres

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/ffo32167/currencyconverter/internal"
	"github.com/jackc/pgx/v4/pgxpool"
)

type pgDb struct {
	pool *pgxpool.Pool
}

type PgRate struct {
	RateDate time.Time `json:"date"`
	CurrCode string    `json:"curr_code"`
	Rate     float64   `json:"rate"`
}

func New() (pgDb, error) {
	db, err := pgxpool.Connect(context.Background(), os.Getenv("PG_CONN_STR"))
	if err != nil {
		return pgDb{}, fmt.Errorf("Unable to connect to database: %w", err)
	}
	defer db.Close()
	return pgDb{pool: db}, nil
}

func (db pgDb) Rate(date string) ([]internal.Rate, error) {
	rows, err := db.pool.Query(context.Background(),
		"SELECT rate_date,curr_code,rate FROM employee_accounting.rates r WHERE rate_date = $1 ORDER BY curr_code",
		date)
	if err != nil {
		return nil, fmt.Errorf("unable to execute select query: %w ", err)
	}
	defer rows.Close()
	var rates []PgRate
	var rate PgRate
	for rows.Next() {
		err = rows.Scan(&rate.RateDate, &rate.CurrCode, &rate.Rate)
		if err != nil {
			return nil, fmt.Errorf("unable to scan query: %w ", err)
		}
		rates = append(rates, rate)
	}
	internalRates, err := toDomain(rates)
	if err != nil {
		return internalRates, fmt.Errorf("unable to convert PG Rates to internal Rates: %w ", err)
	}
	return internalRates, nil
}

func (db pgDb) Create(internalRates []internal.Rate) error {
	/*данный кусок кода не особо нужен*/
	pGrates, err := toPgRate(internalRates)
	if err != nil {
		return fmt.Errorf("unable to convert internal Rates to PG Rates: %w ", err)
	} /**/
	for _, v := range pGrates {
		ct, err := db.pool.Exec(context.Background(),
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

func toDomain(rates []PgRate) ([]internal.Rate, error) {
	result := make([]internal.Rate, len(rates))
	for i := range rates {
		result[i].CurrCode = rates[i].CurrCode
		result[i].Rate = rates[i].Rate
		result[i].RateDate = rates[i].RateDate
	}
	return result, nil
}

func toPgRate(rates []internal.Rate) ([]PgRate, error) {
	result := make([]PgRate, len(rates))
	for i := range rates {
		result[i].CurrCode = rates[i].CurrCode
		result[i].Rate = rates[i].Rate
		result[i].RateDate = rates[i].RateDate
	}
	return result, nil
}
