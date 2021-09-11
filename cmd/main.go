package main

import (
	"fmt"
)

func main() {
	/*	log := log.New()

		cfg, err := newConfig()
		if err != nil {
			log.Fatal("cant get config: ", err)
		}
	*/
	/*	storage, err := pgxpool.Connect(context.Background(), dbURL)
		if err != nil {
			log.Fatalf("Unable to connection to database: %v\n", err)
		}
		defer storage.Close()

		server := APIServer.New(storage, log, APIServer.NewConfig(cfg.servPort))
		server.Start()*/
	fmt.Println("happy end!")
}
