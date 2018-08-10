package main

import (
	"fmt"
	"log"
	"net/http"
)

type API struct {
	uri      string
	method   string
	response string
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	allowedAPIs := []API{
		{"/ping", "GET", "pong"},
		{"/health", "GET", "OK"},
		{"/yo", "GET", "yoyo"},
	}

	for _, api := range allowedAPIs {
		if r.Method == api.method && r.RequestURI == api.uri {
			fmt.Fprint(w, api.response)
			return
		}
	}

	w.WriteHeader(http.StatusMethodNotAllowed)

}
