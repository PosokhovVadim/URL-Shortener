package main

import (
	"fmt"
	"os"
	"url-shortener/internal/config"
	"url-shortener/internal/logging"
	"url-shortener/internal/storage/mongodb"
)

func main() {

	if err != nil {
		os.Exit(1)
	}
}

func run() error {
	cfg := config.MustLoad()
	fmt.Printf("cfg: %v\n", cfg)

	logger := *logging.SetupLogger(cfg.Env)
	_ = logger

	//TODO: Implement storage: mongodb
	db, err := mongodb.ConnectStorage(cfg.StoragePath, logger)
	//mongodb.GetCollections(db)
	if err != nil {
		return fmt.Errorf(err.Error())

	}

	err = mongodb.InsertOneURL(db, "test", "T")
	if err != nil {
		///
	}
	err = mongodb.CloseStorage(db)
	if err != nil {

	}

	//TODO: Implement router: chi, "chi render"

	//TODO: To run server:
	return nil
}
