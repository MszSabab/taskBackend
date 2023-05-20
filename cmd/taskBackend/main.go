package main

import (
	"github.com/taskBackend/config"
	"github.com/taskBackend/httpserver"
	"github.com/taskBackend/repository"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := repository.NewPgRepository(cfg)
	if err != nil {
		log.Fatal(err)
	}

	rr := httpserver.NewServer(
		cfg,
		db,
	)
	rr.Serve()
}
