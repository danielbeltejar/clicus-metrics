package main

import (
	"log"
	"os"
)

var (
	mongoURI          string
	jwtSecret         string
	allowUserCreation bool
)

func initConfig() {
	mongoURI = os.Getenv("MONGO_URI")
	jwtSecret = os.Getenv("JWT_SECRET")
	allowUserCreation = os.Getenv("ALLOW_USER_CREATION") == "true"

	if mongoURI == "" || jwtSecret == "" {
		log.Fatal("Required environment variables not set")
	}
}
