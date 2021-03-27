package main

import (
	"fmt"
	"log"
	"net/http"

	"user-microservice/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/user/profile", handlers.UserHandler).Methods("GET")
	r.HandleFunc("/microservice/name", handlers.Microservicehandler).Methods("GET")

	fmt.Println("User microservice running...")
	log.Fatal(http.ListenAndServe("localhost:3001", r))
}
