package postgres

import (
	"time"

	domain "github.com/ffo32167/currencyconverter"
)

type postgres struct {
	connectionString string
}

func NewPostgres(connectionString string) *postgres {
	return &postgres{connectionString: connectionString}
}

func (p *postgres) Load(date time.Time) (domain.Rate, error) {
	return domain.Rate{Base: "USD", Date: time.Now(), Rates: map[string]float64{"RUB": 75.00}}, nil
}
func (p *postgres) Save(domain.Rate) error {
	return nil
}
