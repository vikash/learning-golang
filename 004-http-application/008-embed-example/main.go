package main

import (
	"embed"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/static", Static)
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//go:embed static/index.html
var fs embed.FS

func Static(w http.ResponseWriter, r *http.Request) {
	data, err := fs.ReadFile("static/index.html")
	if err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something unexpected"))

		return
	}

	w.Write(data)
}
