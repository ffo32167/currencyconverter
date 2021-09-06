package storage

import (
	"time"

	domain "github.com/ffo32167/currencyconverter"
)

type Storage interface {
	Load(date time.Time) (domain.Rate, error)
	Save(domain.Rate) error
}

type Config struct {
	PgConnStr string
}
