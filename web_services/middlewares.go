package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "index controller\n")
	fmt.Printf("index controller\n")
}

//func hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
func (c *Controller) hello(w http.ResponseWriter, r *http.Request) {
	params := c.getUrlParams(r)
	fmt.Fprintf(w, "Hello, %s!\n", params.ByName("name"))
}

func (c *Controller) getUrlParams(r *http.Request) httprouter.Params {
	_, params, _ := c.Router.Lookup("GET", r.URL.Path)
	return params

}

func Panic(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Panic("E panica man!\n")
}

func loginHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("LoginIn")
	fmt.Printf(">>>>>>>>ps: %+v\n", ps)
}

func auth(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// do some stuff before
	fmt.Printf("auth middleware -> before controller\n")
	next(rw, r)
	fmt.Printf("auth middleware -> after controller\n")
	// do some stuff after
}

type Controller struct {
	Router *httprouter.Router
}

func main() {
	router := httprouter.New()
	router.GET("/panic", Panic)
	router.POST("/login", loginHandler)

	// add middleware for a specific route
	nIndex := negroni.New()
	nIndex.Use(negroni.HandlerFunc(auth))
	nIndex.UseHandlerFunc(index)
	router.Handler("GET", "/", nIndex)

	controller := Controller{
		Router: router,
	}
	nHello := negroni.New()
	nHello.Use(negroni.HandlerFunc(auth))
	nHello.UseHandlerFunc(controller.hello)
	router.Handler("GET", "/hello/:name", nHello)

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(router)

	log.Fatal(http.ListenAndServe(":8080", n))
}
