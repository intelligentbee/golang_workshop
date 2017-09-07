package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

// handler function for index endpoint
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the content of the index controller\n")
	log.Println("executing index controller")
}

// handler function for /login endpoint
func loginHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "This is the content of the login controller\n")
	log.Println("executing login controller")
}

// handler function for /hello/:name endpoint
func (c *Controller) helloHandler(w http.ResponseWriter, r *http.Request) {
	params := c.getUrlParams(r)
	fmt.Fprintf(w, "Hello, %s!\n", params.ByName("name"))
	log.Println("executing hello controller")
}

func (c *Controller) getUrlParams(r *http.Request) httprouter.Params {
	_, params, _ := c.Router.Lookup("GET", r.URL.Path)

	return params
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
	router.POST("/login", loginHandler)

	// add middleware for a specific route
	nIndex := negroni.New()
	nIndex.Use(negroni.HandlerFunc(auth))
	nIndex.UseHandlerFunc(indexHandler)
	router.Handler("GET", "/", nIndex)

	controller := Controller{
		Router: router,
	}
	nHello := negroni.New()
	nHello.Use(negroni.HandlerFunc(auth))
	nHello.UseHandlerFunc(controller.helloHandler)
	router.Handler("GET", "/hello/:name", nHello)

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(router)

	log.Fatal(http.ListenAndServe(":8080", n))
}
