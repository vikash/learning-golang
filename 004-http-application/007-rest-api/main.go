package main

import (
	"log"
	"net/http"

	"github.com/vikash/learning-golang/004-http-application/007-rest-api/customer"
)

func main() {
	h := http.HandlerFunc(appHandler)
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", h))
}

func appHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		customer.Greet(w, r)

	case http.MethodPost:
		customer.Create(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("invalid method"))
	}
}
