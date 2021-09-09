package main

import (
	"github.com/ffo32167/currencyconverter/log"
	APIServer "github.com/ffo32167/currencyconverter/pkg/APIServer"
	"github.com/ffo32167/currencyconverter/pkg/storage"
)

func main() {
	log := log.New()

	cfg, err := newConfig()
	if err != nil {
		log.Fatal("cant get config: ", err)
	}

	storage, err := storage.NewPg(cfg.pgConnStr)
	if err != nil {
		log.Fatal("cant get storage: ", err)
	}
	server := APIServer.New(storage, log, APIServer.NewConfig(cfg.servPort))
	server.Start()
}
