package rate

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/ffo32167/currencyconverter/internal/postgres"
)

type Rate struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) Rate {
	return Rate{pool: pool}
}

func (r Rate) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data, err := postgres.Rate(r.pool, mux.Vars(req)["date"])
	if err != nil {
		// log error
	}

	err = json.NewEncoder(res).Encode(data)
	if err != nil {
		// log error
	}
}
