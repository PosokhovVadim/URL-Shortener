package mongodb

import (
	"context"
	"fmt"
	"log"
	"url-shortener/internal/logging"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	for _, collection := range collections {
		fmt.Println(collection)
	}

}


func InsertOneURL(db *Storage, url, alias string) error {
	const fn = "storage.mongodb.InsertOneURL"

	coll := db.db.Collection("URL")

	_, err := coll.InsertOne(context.Background(), bson.D{
		{Key: "Url", Value: url},
		{Key: "Alias", Value: alias},
	})

	if err != nil {
		db.logger.Logger.Info(fmt.Sprintf("Insert error in func: %s", fn))
		return fmt.Errorf("%w, %s", err, fn)
	}
	return nil
}


func InsertManyURL(db *Storage, values map[string]string) error {
	const fn = "storage.mongodb.InsertManyURL"

	coll := db.db.Collection("URL")
	docs := []interface{}{}
	for key, val := range values {
		docs = append(docs, bson.D{
			{Key: "Url", Value: key},
			{Key: "Alias", Value: val},
		})
	}

	_, err := coll.InsertMany(context.Background(), docs)
	if err != nil {
		db.logger.Logger.Info(fmt.Sprintf("Insert error in func: %s", fn))
		return fmt.Errorf("%w, %s", err, fn)
	}
	return nil

}

func SelectURL() {

}

func DeleteUrl() {
	
}

func ConnectStorage(storagePath string, log logging.Logger) (*Storage, error) {
	const fn = "storage.mongodb.ConnectStorage"
	clientOptions := options.Client().ApplyURI(storagePath)

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Logger.Info(fmt.Sprintf("Connection error in func: %s", fn))
		return nil, fmt.Errorf("%w, %s", err, fn)
	}

	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Logger.Info(fmt.Sprintf("Incorrect connection in func: %s", fn))
		return nil, fmt.Errorf("%w, %s", err, fn)
	}

	log.Logger.Info("Connection established")

	return &Storage{
		db:     client.Database("url-shortener"),
		logger: &log,
	}, nil

}

func CloseStorage(db *Storage) error {
	const fn = "storage.mongodb.CloseStorage"

	err := db.db.Client().Disconnect(context.Background())

	if err != nil {
		db.logger.Logger.Info(fmt.Sprintf("Disconnect error in func: %s", fn))
		return fmt.Errorf("%w, %s", err, fn)
	}

	db.logger.Logger.Info("Database Disconnected")

	return nil

}
