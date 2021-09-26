package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/ffo32167/currencyconverter/internal"
	"github.com/ffo32167/currencyconverter/internal/cbr"
	"github.com/ffo32167/currencyconverter/internal/cron"
	"github.com/ffo32167/currencyconverter/internal/http"
	"github.com/ffo32167/currencyconverter/internal/postgres"
)

func main() {
	ctxTimeoutValue, err := strconv.ParseInt(os.Getenv("CTX_TIMEOUT"), 10, 64)

	ctxTimeout := time.Duration(ctxTimeoutValue)
	if err != nil {
		fmt.Println("CTX_TIMEOUT: ", err)
	}
	source := cbr.New(
		os.Getenv("CBR_CONN_STR"),
		os.Getenv("CURRENCIES"),
		ctxTimeout)
	fmt.Println(source.Rates())

	storage, err := postgres.New(context.Background(), os.Getenv("PG_CONN_STR"))

	if err != nil {
		fmt.Println("pgCreate: ", err)
	}

	t := time.Now().Add(10 * time.Second)
	loc, _ := time.LoadLocation("Europe/Moscow")
	c := cron.New(t, *loc, internal.Sync, ctxTimeout, source, storage)
	go c.Action()

	apiServer := http.New(storage, os.Getenv("PORT"), ctxTimeout)
	apiServer.Run()

}
