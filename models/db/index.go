package db

import (
	"context"
	"log"

	"github.com/Raffy27/go-purple/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Init() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB")
}

func Client() *mongo.Client {
	return client
}

func Main() *mongo.Database {
	return client.Database(config.Get().GetString("database.main"))
}

func Shutdown() {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Disconnected from MongoDB")
}
