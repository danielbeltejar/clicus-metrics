// main.go
package main

import (
	"log"
	"net/http"
	"user-service/handlers"

	"github.com/gorilla/mux"
)

func main() {
	InitConfig()
	db := InitMongo()
	// Initialize handlers with the User collection
	handlers.InitHandlers(db)

	r := mux.NewRouter()
	r.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	r.HandleFunc("/user", handlers.GetUser).Methods("GET")

	log.Println("User service running on port 8083")
	log.Fatal(http.ListenAndServe(":8083", r))
}
