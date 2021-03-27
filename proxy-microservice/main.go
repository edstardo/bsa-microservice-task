package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

func AuthHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		authMSURL, _ := url.Parse("http://127.0.0.1:3002/auth")

		authRequest := r.Clone(context.Background())
		authRequest.Host = authMSURL.Host
		authRequest.URL.Host = authMSURL.Host
		authRequest.URL.Scheme = authMSURL.Scheme
		authRequest.URL.Path = authMSURL.Path
		authRequest.RequestURI = ""
		_, err := http.DefaultClient.Do(authRequest)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		handler.ServeHTTP(rw, r)
	}
}

func UserMicroserviceHandler(rw http.ResponseWriter, r *http.Request) {
	userMSURL, _ := url.Parse("http://127.0.0.1:3001")

	userRequest := r.Clone(context.Background())
	userRequest.Host = userMSURL.Host
	userRequest.URL.Host = userMSURL.Host
	userRequest.URL.Scheme = userMSURL.Scheme
	userRequest.RequestURI = ""

	res, err := http.DefaultClient.Do(userRequest)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(rw, err)
		return
	}
	for key, values := range res.Header {
		for _, value := range values {
			rw.Header().Set(key, value)
		}
	}
	rw.WriteHeader(res.StatusCode)
	io.Copy(rw, res.Body)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/user/profile", AuthHandler(UserMicroserviceHandler))
	r.HandleFunc("/microservice/name", UserMicroserviceHandler)

	fmt.Println("Proxy microservice running...")

	log.Fatal(http.ListenAndServe("localhost:3000", r))
}
