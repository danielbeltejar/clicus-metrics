package main

import (
	"log"
	"os"
)

var (
	MongoURI  string
	JwtSecret string
)

func initConfig() {
	MongoURI = os.Getenv("MONGO_URI")
	JwtSecret = os.Getenv("JWT_SECRET")

	if MongoURI == "" || JwtSecret == "" {
		log.Fatal("Required environment variables not set")
	}
}
