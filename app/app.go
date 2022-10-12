package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	router := mux.NewRouter()

	//defining routes
	router.HandleFunc("/hello", hello).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	//starting server
	log.Fatalf("error", http.ListenAndServe("localhost:8000", router))

}

func createCustomer(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Post request received")
}

func getCustomer(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprint(writer, vars["customer_id"])
}
