package main

import (
	"auth-service/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// @title Auth Service API
// @version 1.0
// @description API for managing user authentication
// @termsOfService http://clicus.danielbeltejar.es/terms/
// @contact.name API Support
// @contact.url http://danielbeltejar.es/
// @contact.email hi@danielbeltejar.es/
// @license.name MIT
// @license.url http://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /
func main() {
	initConfig()

	// Initialize auth service with mongoURI, jwtSecret, and allowUserCreation
	handlers.InitAuthService(mongoURI, jwtSecret, allowUserCreation)

	r := mux.NewRouter()
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	log.Println("Auth service running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
