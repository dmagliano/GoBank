package app

import (
	"log"
	"net/http"
)

func Start() {

	//defining routes
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/customers", getAllCustomer)

	//starting server
	log.Fatalf("error", http.ListenAndServe("localhost:8000", nil))

}
