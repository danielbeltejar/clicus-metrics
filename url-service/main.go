package main

import (
	"golang.org/x/net/context"
	"log"
	"net/http"
	"url-service/handlers"

	"github.com/gorilla/mux"
)

func main() {
	initMongoDB()
	defer mongoClient.Disconnect(context.Background())

	handlers.InitHandlers(urlCollection)

	r := mux.NewRouter()
	r.HandleFunc("/shorten", handlers.ShortenURL).Methods("POST")
	r.HandleFunc("/r", handlers.Redirect).Methods("GET")

	log.Println("URL service running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
