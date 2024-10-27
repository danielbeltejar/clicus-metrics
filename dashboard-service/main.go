// main.go
package main

import (
	"dashboard-service/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	InitConfig()
	db := InitMongo()
	handlers.InitHandlers(db)

	r := mux.NewRouter()
	r.HandleFunc("/dashboard", handlers.GetDashboardData).Methods("GET")

	log.Println("Dashboard service running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
