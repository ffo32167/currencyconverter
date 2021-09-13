package relation

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/ffo32167/currencyconverter/internal/postgres"
)

type Relation struct {
	pool *pgxpool.Pool
}

func NewRelation(pool *pgxpool.Pool) Relation {
	return Relation{pool: pool}
}

func (r Relation) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data, err := postgres.Rate(r.pool, mux.Vars(req)["date"])
	if err != nil {
		// check error
	}
	curr1 := mux.Vars(req)["curr1"]
	curr2 := mux.Vars(req)["curr2"]

	curr1index := sort.Search(len(data), func(i int) bool {
		return data[i].CurrCode >= curr1
	})

	curr2index := sort.Search(len(data), func(i int) bool {
		return data[i].CurrCode >= curr2
	})

	if data[curr2index].Rate == float64(0) {
		json.NewEncoder(res).Encode("error, rate of the " + curr2 + " is 0")
	}
	relation := data[curr2index].Rate / data[curr1index].Rate

	err = json.NewEncoder(res).Encode("exchange rate for " + curr1 + " to " + curr2 + " is " +
		strconv.FormatFloat(relation, 'f', 6, 64))
	if err != nil {
		// check error
	}
}
