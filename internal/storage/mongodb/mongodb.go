package mongodb

import (
	"context"
	"fmt"
	"log"
	"url-shortener/internal/logging"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/exp/slog"
)

type Storage struct {
	db     *mongo.Database
	logger *logging.Logger
}

func (s *Storage) GetName() string {
	return string(s.db.Name())
}

func GetCollections(db *Storage) {

	collections, err := db.db.ListCollectionNames(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	// Вывод списка коллекций
	for _, collection := range collections {
		fmt.Println(collection)
	}

}
func ConnectStorage(storagePath string, log logging.Logger) (*Storage, error) {
	const fn = "storage.mongodb.ConnectStorage"

	clientOptions := options.Client().ApplyURI(storagePath)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, fmt.Errorf("%w, %s", err, fn)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, fmt.Errorf("%w, %s", err, fn)
	}

	slog.Info("Connection established")

	return &Storage{
		db:     client.Database("url-shortener"),
		logger: &log,
	}, nil

}

func CloseStorage(db *Storage) error {
	const fn = "storage.mongodb.CloseStorage"

	err := db.db.Client().Disconnect(context.TODO())

	if err != nil {
		return fmt.Errorf("%w, %s", err, fn)
	}

	slog.Info("Database Disconnected")
	return nil

}
