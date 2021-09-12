package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/ffo32167/currencyconverter/internal/APIServer/handlers/rate"
)

func main() {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("PG_CONN_STR"))
	if err != nil {
		fmt.Println("Unable to connect to database: ", err)
	}
	defer pool.Close()

	ratesHandler := rate.NewRates(pool)

	router := mux.NewRouter()
	router.Handle("/rate/{date:[0-9]+}", ratesHandler).Methods("GET")

	http.ListenAndServe(os.Getenv("PORT"), router)
}
