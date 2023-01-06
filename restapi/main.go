package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	handler "github.com/siddhantprateek/ohmyback-end/restapi/handler"
)

func main() {

	router := mux.NewRouter()
	// the Methods function adds a matcher for HTTP methods
	// It accept a sequence of one or more methods to be matched
	router.HandleFunc("/hello", handler.HelloHandler).Methods("HEAD")
	router.HandleFunc("/query", handler.QueryParamHandler)

	log.Println("Listening...")
	http.ListenAndServe(":8000", router)

}
