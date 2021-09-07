package server

import (
	"log"

	"github.com/ffo32167/currencyconverter/pkg/storage"
)

type Server struct {
	storage.Storage
	*log.Logger
	Config
}

type Config struct {
	port string
}

func Start(storage storage.Storage, log *log.Logger, cfg Config) error {
	return nil
}
