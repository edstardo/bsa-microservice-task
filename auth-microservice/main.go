package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Username string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/auth", AuthHandler).Methods("GET")

	fmt.Println("Auth microservice running...")

	log.Fatal(http.ListenAndServe("localhost:3002", r))
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("Username")
	if username == "" || !findUser(username) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func getUsers() map[string]User {
	users := make(map[string]User)

	users["user1"] = User{"user1"}
	users["user2"] = User{"user2"}

	return users
}

func findUser(username string) bool {
	_, ok := getUsers()[username]
	return ok
}
