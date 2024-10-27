// config.go
package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var (
	MongoURI          string
	MongoUser         string
	MongoPass         string
	JwtSecret         string
	AllowUserCreation bool
)

func InitConfig() {
	MongoURI = os.Getenv("MONGO_URI")
	MongoUser = os.Getenv("MONGO_USER")
	MongoPass = os.Getenv("MONGO_PASSWORD")
	JwtSecret = os.Getenv("JWT_SECRET")
	AllowUserCreation = os.Getenv("ALLOW_USER_CREATION") == "true"

	if MongoURI == "" || JwtSecret == "" || MongoUser == "" || MongoPass == "" {
		log.Fatal("Required environment variables not set")
	}
}

func InitMongo() *mongo.Database {
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s", MongoUser, MongoPass, MongoURI)

	// Connect to MongoDB using MongoURI from config
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	return client.Database("clicus")
}
