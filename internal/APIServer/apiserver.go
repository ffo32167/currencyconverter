package apiserver

import (
	"net/http"

	"github.com/ffo32167/currencyconverter/internal/apiserver/handlers/rate"
	"github.com/ffo32167/currencyconverter/internal/apiserver/handlers/relation"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ApiServer struct {
	db   *pgxpool.Pool
	port string
}

func New(db *pgxpool.Pool, port string) ApiServer {
	return ApiServer{db: db, port: port}
}

func (as ApiServer) Run() error {
	rateHandler := rate.New(as.db)
	relationHandler := relation.New(as.db)

	router := mux.NewRouter()
	router.Handle("/rate/{date:[0-9]+}", rateHandler).Methods("GET")
	router.Handle("/relation/{date:[0-9]+}/{curr1:[A-Z]+}/{curr2:[A-Z]+}", relationHandler).Methods("GET")

	err := http.ListenAndServe(as.port, router)
	return err
}
