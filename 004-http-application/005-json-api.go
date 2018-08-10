package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Route struct {
	uri     string
	method  string
	handler Handler
}

type Handler func(r *http.Request) (interface{}, error)

func main() {
	http.HandleFunc("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func router(w http.ResponseWriter, r *http.Request) {

	routes := []Route{
		{"/ping", "GET", handlePing},
		{"/health", "GET", handleHealth},
	}

	for _, route := range routes {
		if r.Method == route.method && r.RequestURI == route.uri {
			data, err := route.handler(r)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			response := struct {
				Status string
				Data   interface{}
			}{"SUCCESS", data}

			json.NewEncoder(w).Encode(response)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)

}

func handlePing(r *http.Request) (interface{}, error) {
	return "Pong", nil
}

func handleHealth(r *http.Request) (interface{}, error) {
	res := map[string]string{
		"MySQL Connectivity": "OK",
		"Redis Connectivity": "OK",
	}
	return res, nil
}
