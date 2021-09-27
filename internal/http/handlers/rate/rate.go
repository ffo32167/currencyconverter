package rate

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ffo32167/currencyconverter/internal"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Rate struct {
	storage    internal.Storage
	ctxTimeout time.Duration
	log        *zap.Logger
}

func New(storage internal.Storage, ctxTimeout time.Duration, log *zap.Logger) Rate {
	return Rate{storage: storage, ctxTimeout: ctxTimeout, log: log}
}

func (r Rate) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println("http rate")
	ctx, cancel := context.WithTimeout(req.Context(), r.ctxTimeout*time.Second)
	defer cancel()
	dt, err := time.Parse("20060102", mux.Vars(req)["date"])
	if err != nil {
		r.log.Error("rate handler time parse error:", zap.Error(err))
	}

	data, err := internal.Rates(ctx, r.storage, dt)
	if err != nil {
		r.log.Error("rate handler rates error:", zap.Error(err))
	}

	err = json.NewEncoder(res).Encode(data)
	if err != nil {
		r.log.Error("rate handler encoder error:", zap.Error(err))
	}
	fmt.Println("http rate")
}
