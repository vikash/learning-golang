package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	h := http.HandlerFunc(appHandler)
	log.Fatal(http.ListenAndServe(":8080", h))
}

func appHandler(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.RequestURI, "/allowed") {
		fmt.Fprint(w, "OK")
		return
	}

	fmt.Fprint(w, "NOK")
}
