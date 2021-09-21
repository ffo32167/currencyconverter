package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ffo32167/currencyconverter/internal/cron"
	"github.com/go-redis/redis/v8"
)

func hello() {
	fmt.Println("hello from cron,", time.Now())
}

func main() {
	/*	ctxTimeout, err := strconv.ParseInt(os.Getenv("CTX_TIMEOUT"), 10, 64)
		if err != nil {
			fmt.Println("CTX_TIMEOUT: ", err)
		}
		curr := cbr.New(
			os.Getenv("CBR_CONN_STR"),
			os.Getenv("CURRENCIES"),
			ctxTimeout)

		fmt.Println(curr.Rates())
	*/

	/*	storage, err := postgres.New()
		if err != nil {
			fmt.Println("pgCreate: ", err)
		}

		ctxTimeout, err := strconv.ParseInt(os.Getenv("CTX_TIMEOUT"), 10, 64)
		if err != nil {
			fmt.Println("CTX_TIMEOUT: ", err)
		}

		apiServer := apiserver.New(storage, os.Getenv("PORT"), ctxTimeout)
		apiServer.Run()*/
	/*
		cfr := currencyfreaks.New(
			os.Getenv("CURRENCYFREAKS_CONN_STR"),
			os.Getenv("CURRENCIES"),
			ctxTimeout)

		rates, err := cfr.Rates()
		if err != nil {
			fmt.Println("cfr.Rates: ", err)
		}*/

	fmt.Println("Go Redis Tutorial")
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.TODO()

	//	err = client.Set(ctx, "name", "Elliot", 0).Err()

	val, err := client.Get(ctx, "name").Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val, err)

	t := time.Now().Add(10 * time.Second)
	loc, _ := time.LoadLocation("Europe/Moscow")
	c := cron.New(t, *loc, hello)

	c.Action()
	time.Sleep(30)
}
