package rate

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/ffo32167/currencyconverter/internal"
	"github.com/gorilla/mux"
)

type Rate struct {
	storage    internal.Storage
	ctxTimeout int64
}

func New(storage internal.Storage) Rate {
	return Rate{storage: storage}
}

func (r Rate) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.ctxTimeout)*time.Millisecond)
	defer cancel()
	dt, err := time.Parse("2006-01-02", mux.Vars(req)["date"])
	if err != nil {
		// log error
	}
	data, err := internal.Rates(ctx, r.storage, dt)
	if err != nil {
		// log error
	}

	err = json.NewEncoder(res).Encode(data)
	if err != nil {
		// log error
	}
}
