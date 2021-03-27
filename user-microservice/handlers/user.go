package handlers

import (
	"encoding/json"
	"net/http"

	"user-microservice/data"
)

func UserHandler(rw http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("Username")
	if username == "" {
		http.Error(rw, "Unauthorized", http.StatusUnauthorized)
		return
	}
	user := data.GetUsers()[username]
	if user == nil {
		http.Error(rw, "Unauthorized", http.StatusUnauthorized)
		return
	}
	data, err := json.Marshal(user)
	if err != nil {
		http.Error(rw, "An error occured.", http.StatusInternalServerError)
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write(data)
}
