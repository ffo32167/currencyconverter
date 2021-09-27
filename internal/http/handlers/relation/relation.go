package relation

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/ffo32167/currencyconverter/internal"
)

type Relation struct {
	storage    internal.Storage
	ctxTimeout time.Duration
	log        *zap.Logger
}

func New(storage internal.Storage, log *zap.Logger) Relation {
	return Relation{storage: storage, log: log}
}

func (r Relation) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ctxTimeout)
	defer cancel()

	date := mux.Vars(req)["date"]
	curr1 := mux.Vars(req)["curr1"]
	curr2 := mux.Vars(req)["curr2"]

	dt, err := time.Parse("2006-01-02", date)
	if err != nil {
		r.log.Error("rate handler time parse error:", zap.Error(err))
	}

	result, err := internal.Relation(ctx, r.storage, dt, curr1, curr2)

	if err != nil {
		r.log.Error("rate handler relation error:", zap.Error(err))
		json.NewEncoder(res).Encode(err)
	}

	err = json.NewEncoder(res).Encode(result)
	if err != nil {
		r.log.Error("rate handler encoder error:", zap.Error(err))
		json.NewEncoder(res).Encode(err)
	}
}
