package main

import (
	"fmt"
	"url-shortener/internal/config"
	"url-shortener/internal/logging"
	"url-shortener/internal/storage/mongodb"
)

func main() {

	cfg := config.MustLoad()
	fmt.Printf("cfg: %v\n", cfg)

	logger := *logging.SetupLogger(cfg.Env)
	_ = logger

	//TODO: Implement storage: mongodb
	db, err := mongodb.ConnectStorage(cfg.StoragePath, logger)
	//mongodb.GetCollections(db)
	if err != nil {
		fmt.Println("Not succes connect into database", db.GetName())

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

}
