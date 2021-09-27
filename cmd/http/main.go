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
	"go.uber.org/zap"
)

func main() {
	log, err := zap.NewProduction()
	if err != nil {
		fmt.Println(fmt.Errorf("cant start logger: %w", err))
	}
	defer func() {
		err = log.Sync()
		if err != nil {
			fmt.Println(fmt.Errorf("cant sync logger: %w", err))
		}
	}()

	ctxTimeoutValue, err := strconv.ParseInt(os.Getenv("CTX_TIMEOUT"), 10, 64)

	ctxTimeout := time.Duration(ctxTimeoutValue)
	if err != nil {
		log.Error("CTX_TIMEOUT error", zap.Error(err))
		return
	}
	source := cbr.New(
		os.Getenv("CBR_CONN_STR"),
		os.Getenv("CURRENCIES"),
		ctxTimeout)

	rates, err := source.Rates()
	if err != nil {
		log.Error("cant call Rates", zap.Error(err))
	}
	fmt.Println(rates)

	storage, err := postgres.New(context.Background(), os.Getenv("PG_CONN_STR"))
	if err != nil {
		log.Error("pgCreate: ", zap.Error(err))
	}

	t := time.Now().Add(10 * time.Second)
	loc, _ := time.LoadLocation("Europe/Moscow")
	c := cron.New(t, *loc, internal.Sync, ctxTimeout, source, storage, log)
	go c.Action()

	apiServer := http.New(storage, os.Getenv("PORT"), ctxTimeout, log)
	apiServer.Run()

}
