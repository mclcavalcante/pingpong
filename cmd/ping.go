package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/pong", pongFunc)
	log.Fatal(http.ListenAndServe(":8099", myRouter))
}
