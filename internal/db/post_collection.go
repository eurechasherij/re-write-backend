package db

import (
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetPostCollection() *mongo.Collection {
	client := GetMongoClient()
	dbName := os.Getenv("MONGO_DB_NAME")
	if dbName == "" {
		dbName = "rewrite"
	}

	return client.Database(dbName).Collection("posts")
}
