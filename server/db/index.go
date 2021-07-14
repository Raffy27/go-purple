package db

import (
	"context"
	"log"
	"strings"

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

func C(path string) *mongo.Collection {
	i := strings.Index(path, "/")
	if i == -1 {
		// No database specified, use the main database
		return Main().Collection(path)
	}
	return client.Database(path[:i]).Collection(path[i+1:])
}

func Shutdown() {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Disconnected from MongoDB")
}
