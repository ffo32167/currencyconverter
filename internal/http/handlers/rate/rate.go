package rate

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/ffo32167/currencyconverter/internal"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Rate struct {
	currRepo   internal.CurrencyRepository
	ctxTimeout time.Duration
	log        *zap.Logger
}

func New(currRepo internal.CurrencyRepository, ctxTimeout time.Duration, log *zap.Logger) Rate {
	return Rate{currRepo: currRepo, ctxTimeout: ctxTimeout, log: log}
}

func (r Rate) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), r.ctxTimeout*time.Second)
	defer cancel()
	dt, err := time.Parse("20060102", mux.Vars(req)["date"])
	if err != nil {
		r.log.Error("rate handler time parse error:", zap.Error(err))
	}

	data, err := r.currRepo.Rates(ctx, dt)
	if err != nil {
		r.log.Error("rate handler rates error:", zap.Error(err))
	}

	err = json.NewEncoder(res).Encode(data)
	if err != nil {
		r.log.Error("rate handler encoder error:", zap.Error(err))
	}
}
