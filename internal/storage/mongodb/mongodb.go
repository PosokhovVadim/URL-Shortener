package mongodb

import (
	"context"
	"fmt"
	"url-shortener/internal/lib/logger/sl"
	"url-shortener/internal/logging"
	"url-shortener/internal/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	db     *mongo.Database
	logger *logging.Logger
}

//TODO: remove logger from Storage, add to use const errors from storage.go

func isDuplicate(err error) error {
	if mongoErr, ok := err.(mongo.WriteException); ok {
		if mongoErr.WriteErrors[0].Code == 11000 {
			return storage.ErrURLExists

		}
	}
	return nil
}

func InsertOneURL(db *Storage, url, alias string) error {
	const fn = "storage.mongodb.InsertOneURL"

	coll := db.db.Collection("URL")

	_, err := coll.InsertOne(context.Background(), bson.D{
		{Key: "Url", Value: url},
		{Key: "Alias", Value: alias},
	})

	if err != nil {
		db.logger.Logger.Error(fmt.Sprintf("Insert error in func: %s", fn), sl.Err(err))
		if dupErr := isDuplicate(err); dupErr != nil {
			return fmt.Errorf("%w, %s", dupErr, fn)
		}

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
		db.logger.Logger.Error(fmt.Sprintf("Insert error in func: %s", fn), sl.Err(err))
		if dupErr := isDuplicate(err); dupErr != nil {
			return fmt.Errorf("%w, %s", dupErr, fn)
		}

		return fmt.Errorf("%w, %s", err, fn)
	}
	return nil

}

func SelectURL(db *Storage, alias string) (string, error) {
	const fn = "storage.mongodb.SelectURL"

	coll := db.db.Collection("URL")

	var url string
	err := coll.FindOne(context.Background(), bson.D{
		{Key: "Alias", Value: alias},
	}).Decode(&url)

	if err != nil {
		db.logger.Logger.Error(fmt.Sprintf("Select error in func: %s", fn), sl.Err(err))
		return "", fmt.Errorf("%w, %s", err, fn)
	}
	return url, nil
}

func DeleteOneURL(db *Storage, alias string) error {
	const fn = "storage.mongodb.DeleteURL"

	coll := db.db.Collection("URL")

	_, err := coll.DeleteOne(context.Background(), bson.D{
		{Key: "Alias", Value: alias},
	})

	if err != nil {
		db.logger.Logger.Error(fmt.Sprintf("Delete error in func: %s", fn), sl.Err(err))
		return fmt.Errorf("%w, %s", err, fn)
	}
	return nil
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
