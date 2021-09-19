package relation

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ffo32167/currencyconverter/internal"
)

type Relation struct {
	storage internal.Storage
}

func New(storage internal.Storage) Relation {
	return Relation{storage: storage}
}

func (r Relation) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	date := mux.Vars(req)["date"]
	curr1 := mux.Vars(req)["curr1"]
	curr2 := mux.Vars(req)["curr2"]

	result, err := internal.Relation(r.storage, date, curr1, curr2)

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
