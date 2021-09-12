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

func NewRates(pool *pgxpool.Pool) Rate {
	return Rate{pool: pool}
}

func (h Rate) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data, err := postgres.Rates(h.pool, mux.Vars(req)["date"])
	if err != nil {

	}
	// check error
	json.NewEncoder(res).Encode(data)
}
