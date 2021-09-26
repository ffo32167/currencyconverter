package relation

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/ffo32167/currencyconverter/internal"
)

type Relation struct {
	storage    internal.Storage
	ctxTimeout time.Duration
}

func New(storage internal.Storage) Relation {
	return Relation{storage: storage}
}

func (r Relation) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ctxTimeout)
	defer cancel()

	date := mux.Vars(req)["date"]
	curr1 := mux.Vars(req)["curr1"]
	curr2 := mux.Vars(req)["curr2"]

	dt, err := time.Parse("2006-01-02", date)
	if err != nil {
		// log error
	}

	result, err := internal.Relation(ctx, r.storage, dt, curr1, curr2)

	if err != nil {
		// log error
		json.NewEncoder(res).Encode(err)
	}

	err = json.NewEncoder(res).Encode(result)

	if err != nil {
		// log error
		json.NewEncoder(res).Encode(err)
	}
}
