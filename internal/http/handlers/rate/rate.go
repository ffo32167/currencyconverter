package rate

import (
	"encoding/json"
	"net/http"

	"github.com/ffo32167/currencyconverter/internal"
	"github.com/gorilla/mux"
)

type Rate struct {
	storage internal.Storage
}

func New(storage internal.Storage) Rate {
	return Rate{storage: storage}
}

func (r Rate) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data, err := r.storage.Rate(mux.Vars(req)["date"])
	if err != nil {
		// log error
	}

	err = json.NewEncoder(res).Encode(data)
	if err != nil {
		// log error
	}
}
