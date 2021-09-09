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

func NewPg(connectionString string) (*Storage, error) {
	return postgres.New(connectionString), nil
}
