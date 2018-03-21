package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.
		Path("/upload").
		Methods("POST").
		HandlerFunc(ReceiveFile)
	fmt.Println("Starting")
	log.Fatal(http.ListenAndServe(":8080", router))
}
