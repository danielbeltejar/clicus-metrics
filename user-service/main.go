// main.go
package main

import (
	"context"
	"log"
	"net/http"
	"user-service/handlers"

	"github.com/gorilla/mux"
)

func main() {
	initMongoDB()
	defer mongoClient.Disconnect(context.Background())

	// Initialize handlers with the User collection
	handlers.InitHandlers(userCollection)

	r := mux.NewRouter()
	r.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	r.HandleFunc("/user", handlers.GetUser).Methods("GET")

	log.Println("User service running on port 8083")
	log.Fatal(http.ListenAndServe(":8083", r))
}
