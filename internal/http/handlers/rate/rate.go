package rate

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ffo32167/currencyconverter/internal"
	"github.com/gorilla/mux"
)

type Rate struct {
	storage    internal.Storage
	ctxTimeout time.Duration
}

func New(storage internal.Storage, ctxTimeout time.Duration) Rate {
	return Rate{storage: storage, ctxTimeout: ctxTimeout}
}

func (r Rate) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println("http rate")
	ctx, cancel := context.WithTimeout(req.Context(), r.ctxTimeout*time.Second)
	defer cancel()
	dt, err := time.Parse("20060102", mux.Vars(req)["date"])
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
	fmt.Println("http rate")
}
