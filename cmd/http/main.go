package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ffo32167/currencyconverter/internal"
	"github.com/ffo32167/currencyconverter/internal/cbr"
	"github.com/ffo32167/currencyconverter/internal/cron"
	"github.com/ffo32167/currencyconverter/internal/currencyfreaks"
	"github.com/ffo32167/currencyconverter/internal/http"
	"github.com/ffo32167/currencyconverter/internal/postgres"
	"github.com/ffo32167/currencyconverter/internal/redis"
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
	if err != nil {
		log.Error("cant parse env CTX_TIMEOUT:", zap.Error(err))
		return
	}
	ctxTimeout := time.Duration(ctxTimeoutValue)

	var source internal.Source
	switch os.Getenv("SOURCE") {
	case "cbr":
		source = cbr.New(
			os.Getenv("CBR_CONN_STR"),
			os.Getenv("CURRENCIES"),
			ctxTimeout)
	default:
		source = currencyfreaks.New(
			os.Getenv("CURRENCYFREAKS_CONN_STR"),
			os.Getenv("CURRENCIES"),
			ctxTimeout)
	}

	rates, err := source.Rates()
	if err != nil {
		log.Error("cant call Rates", zap.Error(err))
	}
	fmt.Println(rates)

	var storage internal.Storage
	switch os.Getenv("STORAGE") {
	case "redis":
		storage, err = redis.New(
			strings.Split(os.Getenv("REDIS_CONN_STR"), ","))
	default:
		storage, err = postgres.New(context.Background(), os.Getenv("PG_CONN_STR"))
	}
	if err != nil {
		log.Error("storage create error: ", zap.Error(err))
	}

	storage.Create(context.TODO(), rates)

	t := time.Now().Add(10 * time.Second)
	loc, _ := time.LoadLocation("Europe/Moscow")
	c := cron.New(t, *loc, internal.Sync, ctxTimeout, source, storage, log)
	go c.Action()

	apiServer := http.New(storage, os.Getenv("PORT"), ctxTimeout, log)
	err = apiServer.Run()
	if err != nil {
		log.Error("cant start api server:", zap.Error(err))
	}

}
