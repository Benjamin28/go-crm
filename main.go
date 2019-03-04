package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/benjamin28/go-crm/handlers"
)

func main() {
	router := mux.NewRouter()

	c := handlers.NewController()

	router.HandleFunc("/healthcheck", c.Healthcheck).Methods("GET")
	router.HandleFunc("/clients", c.GetClients).Methods("GET")
	router.HandleFunc("/client/{id}", c.GetClient).Methods("GET")
	router.HandleFunc("/client/{id}", c.CreateClient).Methods("POST")
	router.HandleFunc("/client/{id}", c.DeleteClient).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}