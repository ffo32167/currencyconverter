package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ffo32167/currencyconverter/internal/currencyfreaks"
	"github.com/ffo32167/currencyconverter/internal/postgres"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	db, err := pgxpool.Connect(context.Background(), os.Getenv("PG_CONN_STR"))
	if err != nil {
		fmt.Println("Unable to connect to database: ", err)
	}
	defer db.Close()
	/*
		apiServer := apiserver.New(db, os.Getenv("PORT"))
		apiServer.Run()
	*/
	cfr := currencyfreaks.New(
		os.Getenv("CURRENCYFREAKS_CONN_STR"),
		os.Getenv("CURRENCIES"))
	rates, err := cfr.Rates()
	if err != nil {
		fmt.Println("cfr.Rates: ", err)
	}
	err = postgres.Create(db, rates)
	if err != nil {
		fmt.Println("pgCreate: ", err)
	}
}
