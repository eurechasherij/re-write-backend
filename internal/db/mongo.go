package db

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientInstance *mongo.Client
	clientOnce     sync.Once
)

func GetMongoClient() *mongo.Client {
	clientOnce.Do(func() {
		mongoURI := os.Getenv("MONGO_URI")
		if mongoURI == "" {
			log.Fatal("MONGO_URI environment variable not set")
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		clientOptions := options.Client().ApplyURI(mongoURI)

		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatalf("Failed to connect to MongoDB: %v", err)
		}

		if err = client.Ping(ctx, nil); err != nil {
			log.Fatalf("Failed to ping MongoDB: %v", err)
		}

		log.Println("Connected to MongoDB!")
		clientInstance = client
	})
	return clientInstance
}
