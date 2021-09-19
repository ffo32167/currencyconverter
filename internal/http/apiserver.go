package apiserver

import (
	"net/http"

	"github.com/ffo32167/currencyconverter/internal"
	"github.com/ffo32167/currencyconverter/internal/http/handlers/rate"
	"github.com/ffo32167/currencyconverter/internal/http/handlers/relation"
	"github.com/gorilla/mux"
)

type ApiServer struct {
	storage internal.Storage
	port    string
	timeout int64
}

func New(storage internal.Storage, port string, timeout int64) ApiServer {
	return ApiServer{storage: storage, port: port, timeout: timeout}
}

func (as ApiServer) Run() error {
	rateHandler := rate.New(as.storage)
	relationHandler := relation.New(as.storage)

	router := mux.NewRouter()
	router.Handle("/rate/{date:[0-9]+}", rateHandler).Methods("GET")
	router.Handle("/relation/{date:[0-9]+}/{curr1:[A-Z]+}/{curr2:[A-Z]+}", relationHandler).Methods("GET")

	err := http.ListenAndServe(as.port, router)
	if err != nil {
		return err
	}
	return nil
}
