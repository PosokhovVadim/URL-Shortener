package main

import (
	"fmt"
	"os"
	"url-shortener/internal/config"
	"url-shortener/internal/logging"
	"url-shortener/internal/storage/mongodb"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	if  run() != nil {
		os.Exit(1)
	}
}

func run() error {
	cfg := config.MustLoad()
	fmt.Printf("cfg: %v\n", cfg)

	logger := *logging.SetupLogger(cfg.Env)

	db, err := mongodb.ConnectStorage(cfg.StoragePath, logger)
	
 
	if err != nil {
		return fmt.Errorf(err.Error())

	}
	defer db.CloseStorage() 

	router := chi.NewRouter()
	router.Use(middleware.RequestID)

	//TODO: check thunder-tests (dont remeber pull it on gitignore) 
	//middleware 


	//TODO: To run server:
	return nil
}
