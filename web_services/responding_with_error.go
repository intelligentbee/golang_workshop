package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func respondWithError(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	err := errors.New("Error description") //generate an error

	enc := json.NewEncoder(w)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		enc.Encode(map[string]string{
			"error": err.Error(),
		})
		loggingErrors(err.Error())

		return
	}
}

func loggingErrors(errorMsg string) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic\n", r)
		}
	}()

	if errorMsg != "" {
		log.Panicf("Panic oops, something was too hard: %s\n", errorMsg)
	}

	log.Println("This message should not be logged")
}

func respondWithError1(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	log.Fatal("Fatal error message")
}

func main() {
	router := httprouter.New()
	router.GET("/respond-error", respondWithError)
	router.GET("/respond-error1", respondWithError1)

	log.Fatal(http.ListenAndServe(":8080", router))
}
