package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ffo32167/currencyconverter/internal/apiserver"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	db, err := pgxpool.Connect(context.Background(), os.Getenv("PG_CONN_STR"))
	if err != nil {
		fmt.Println("Unable to connect to database: ", err)
	}
	defer db.Close()
	apiServer := apiserver.New(db, os.Getenv("PORT"))
	apiServer.Run()

	/*	cfr := currencyfreaks.New(
			os.Getenv("CURRENCYFREAKS_CONN_STR"),
			os.Getenv("CURRENCIES"))
		fmt.Println(cfr.Rates())
	*/
}
