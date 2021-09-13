package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/ffo32167/currencyconverter/internal/APIServer/handlers/rate"
	"github.com/ffo32167/currencyconverter/internal/APIServer/handlers/relation"
)

func main() {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("PG_CONN_STR"))
	if err != nil {
		fmt.Println("Unable to connect to database: ", err)
	}
	defer pool.Close()

	rateHandler := rate.NewRate(pool)
	relationHandler := relation.NewRelation(pool)

	router := mux.NewRouter()
	router.Handle("/rate/{date:[0-9]+}", rateHandler).Methods("GET")
	router.Handle("/relation/{date:[0-9]+}/{curr1:[A-Z]+}/{curr2:[A-Z]+}", relationHandler).Methods("GET")

	http.ListenAndServe(os.Getenv("PORT"), router)
}
