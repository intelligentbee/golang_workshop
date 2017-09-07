package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func respondJson(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	body := req.Body
	dec := json.NewDecoder(body)
	var v map[string]interface{}
	if err := dec.Decode(&v); err != nil {
		log.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)

	if err := enc.Encode(v); err != nil {
		log.Println(err.Error())
		return
	}
}

func respondJson1(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func main() {
	router := httprouter.New()
	router.POST("/respond", respondJson)
	router.POST("/respond1", respondJson1)

	log.Fatal(http.ListenAndServe(":8080", router))
}
