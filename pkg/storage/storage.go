package storage

import (
	"time"

	domain "github.com/ffo32167/currencyconverter"
	"github.com/ffo32167/currencyconverter/pkg/storage/postgres"
)

type Storage interface {
	Load(date time.Time) (domain.Rate, error)
	Save(domain.Rate) error
}

type Config struct {
	PgConnStr string
}

func New(c Config) (Storage, error) {
	return postgres.New(c.PgConnStr), nil
}
