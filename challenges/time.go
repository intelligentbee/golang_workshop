package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func Time(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t := time.Now().Format(time.RFC3339Nano)
	fmt.Println(t)
	response := map[string]string{
		"now": string(t),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)

	if err := enc.Encode(response); err != nil {
		log.Println(err.Error())
		return
	}
}

func main() {
	router := httprouter.New()

	router.GET("/time", Time)

	log.Fatal(http.ListenAndServe(":8080", router))
}
