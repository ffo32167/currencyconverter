package server

import (
	"log"

	"github.com/ffo32167/currencyconverter/pkg/storage"
	"github.com/gorilla/mux"
)

type APIServer struct {
	storage *storage.Storage
	log     *log.Logger
	cfg     Config
	router  *mux.Router
}

type Config struct {
	port string
}

func NewConfig(port string) Config {
	return Config{port: port}
}

func (s *APIServer) NewRouter() {

}

func New(storage *storage.Storage, log *log.Logger, cfg Config) *APIServer {
	return &APIServer{storage: storage, log: log, cfg: cfg}
}

func (s APIServer) Start() error {

	log.Print("starting server")
	return nil
}
