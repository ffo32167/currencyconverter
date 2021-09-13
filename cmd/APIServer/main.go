package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/ffo32167/currencyconverter/internal/currencyfreaks"
)

func main() {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("PG_CONN_STR"))
	if err != nil {
		fmt.Println("Unable to connect to database: ", err)
	}
	defer pool.Close()

	/*	rateHandler := rate.New(pool)
		relationHandler := relation.New(pool)

		router := mux.NewRouter()
		router.Handle("/rate/{date:[0-9]+}", rateHandler).Methods("GET")
		router.Handle("/relation/{date:[0-9]+}/{curr1:[A-Z]+}/{curr2:[A-Z]+}", relationHandler).Methods("GET")

		http.ListenAndServe(os.Getenv("PORT"), router)
	*/

	cfr := currencyfreaks.New(
		os.Getenv("CURRENCYFREAKS_CONN_STR"),
		os.Getenv("CURRENCIES"))
	cfr.Rates()
}
