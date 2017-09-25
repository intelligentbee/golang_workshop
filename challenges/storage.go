package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var storage map[string]string

func Set(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// read key from url params
	key := ps.ByName("key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// read json from request body
	body := r.Body
	dec := json.NewDecoder(body)
	reqBody := map[string]string{}
	if err := dec.Decode(&reqBody); err != nil {
		if err.Error() == "EOF" {
			w.WriteHeader(http.StatusBadRequest)

			return
		}
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	value, ok := reqBody["value"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)

		return
	}
	//set key-value in storage
	storage[key] = value

	// respond back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// read key from url params
	key := ps.ByName("key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, ok := storage[key]
	if ok {
		delete(storage, key)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// read key from url params
	key := ps.ByName("key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	value, ok := storage[key]
	if !ok {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	result := map[string]string{key: value}

	enc := json.NewEncoder(w)
	if err := enc.Encode(result); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func main() {
	router := httprouter.New()

	router.DELETE("/storage/:key", Delete)
	router.POST("/storage/:key", Set)
	router.GET("/storage/:key", Get)

	storage = map[string]string{}

	log.Fatal(http.ListenAndServe(":8080", router))
}
