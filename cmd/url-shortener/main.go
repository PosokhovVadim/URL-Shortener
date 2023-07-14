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
	db, err := mongodb.ConnectStorage(cfg.StoragePath)

	if err != nil {
		//some to do
	}
	fmt.Println("Succes connect into database", db)
	//TODO: Implement router: chi, "chi render"

	//TODO: To run server:

}
