package main

import (
	"github.com/ffo32167/currencyconverter/pkg/server"
	"github.com/ffo32167/currencyconverter/pkg/storage"
)

type config struct {
	StorCfg storage.Config
	ServCfg server.Config
}

func newConfig() (*config, error) {
	return &config{
		storage.Config{},
		server.Config{},
	}, nil
}
