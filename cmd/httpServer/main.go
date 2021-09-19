package main

import (
	"fmt"
	"os"
	"strconv"

	apiserver "github.com/ffo32167/currencyconverter/internal/http"
	"github.com/ffo32167/currencyconverter/internal/postgres"
)

func main() {
	storage, err := postgres.New()
	if err != nil {
		fmt.Println("pgCreate: ", err)
	}

	ctxTimeout, err := strconv.ParseInt(os.Getenv("CTX_TIMEOUT"), 10, 64)
	if err != nil {
		fmt.Println("CTX_TIMEOUT: ", err)
	}

	apiServer := apiserver.New(storage, os.Getenv("PORT"), ctxTimeout)
	apiServer.Run()
	/*
		cfr := currencyfreaks.New(
			os.Getenv("CURRENCYFREAKS_CONN_STR"),
			os.Getenv("CURRENCIES"),
			ctxTimeout)

		rates, err := cfr.Rates()
		if err != nil {
			fmt.Println("cfr.Rates: ", err)
		}*/
}
