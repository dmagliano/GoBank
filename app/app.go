package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	router := mux.NewRouter()

	//defining routes
	router.HandleFunc("/hello", hello)
	router.HandleFunc("/customers", getAllCustomer)

	//starting server
	log.Fatalf("error", http.ListenAndServe("localhost:8000", router))

}
