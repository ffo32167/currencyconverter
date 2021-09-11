package postgres

//	domain "github.com/ffo32167/currencyconverter"

type Postgres struct {
	connectionString string
}

func New(connectionString string) *Postgres {
	return &Postgres{connectionString: connectionString}
}

/*
func (p *Postgres) Load(date time.Time) (domain.Rate, error) {
	return domain.Rate{Base: "USD", Date: date, Rates: map[string]float64{"RUB": 75.00}}, nil
}
func (p *Postgres) Save(domain.Rate) error {
	return nil
}
*/
