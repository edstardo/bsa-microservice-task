package handlers

import (
	"fmt"
	"net/http"
)

func Microservicehandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "user-microservice")
}
