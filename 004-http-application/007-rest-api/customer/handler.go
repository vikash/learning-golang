package customer

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Takes a customer object and returns the object with the JSON body if valid else error
// error on invalid JSON, type mismatch.
// NOTE: To show usage of json.Unmarshal
func Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid body"))

		return
	}

	var c Customer
	err = json.Unmarshal(body, &c)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid body"))

		return
	}

	w.Write(body)
}

// takes query parameter 'name' and prints hello +name
// NOTE:to show usage of r.URL.Query
func Greet(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("missing param: name"))

		return
	}

	resp := "hello " + name

	w.Write([]byte(resp))
}
