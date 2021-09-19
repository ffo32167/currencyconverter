package main

import (
	"fmt"
	"os"

	apiserver "github.com/ffo32167/currencyconverter/internal/http"
	"github.com/ffo32167/currencyconverter/internal/postgres"
)

func main() {
	storage, err := postgres.New()
	if err != nil {
		fmt.Println("pgCreate: ", err)
	}

	apiServer := apiserver.New(storage, os.Getenv("PORT"))
	apiServer.Run()
	/*
		cfr := currencyfreaks.New(
			os.Getenv("CURRENCYFREAKS_CONN_STR"),
			os.Getenv("CURRENCIES"))

		rates, err := cfr.Rates()
		if err != nil {
			fmt.Println("cfr.Rates: ", err)
		}*/
}
