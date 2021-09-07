package main

import (
	"github.com/ffo32167/currencyconverter/log"
	"github.com/ffo32167/currencyconverter/pkg/server"
	"github.com/ffo32167/currencyconverter/pkg/storage"
)

func main() {
	log := log.New()

	cfg, err := newConfig()
	if err != nil {
		log.Fatal("cant get config: ", err)
	}

	storage, err := storage.New(cfg.StorCfg)
	if err != nil {
		log.Fatal("cant get storage: ", err)
	}

	server.Start(storage, log, cfg.ServCfg)
}
