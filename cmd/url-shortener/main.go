package main

import (
	"fmt"
	"url-shortener/internal/config"
	"url-shortener/internal/logging"
	"url-shortener/internal/storage/mongodb"
)

func main() {
	//DONE: Implement config: cleanenv

	cfg := config.MustLoad()
	fmt.Printf("cfg: %v\n", cfg)
	//TODO: Implement logger: slog

	logger := *logging.SetupLogger(cfg.Env)
	_ = logger
	//TODO: Implement storage: mongodb
	db, err := mongodb.ConnectStorage(cfg.StoragePath, logger)
	mongodb.GetCollections(db)
	if err != nil {
		fmt.Println("Not succes connect into database", db.GetName())

	}
	//mongodb.GetCollections(db)

	err = mongodb.CloseStorage(db)
	if err != nil {
	}

	//TODO: Implement router: chi, "chi render"

	//TODO: To run server:

}
