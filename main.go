package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	Address string `json:"address" xml:"address"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

func main() {

	//defining routes
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/customers", getAllCustomer)

	//starting server
	log.Fatalf("error", http.ListenAndServe("localhost:8000", nil))
}

func hello(writer http.ResponseWriter, request *http.Request) {
	println("request to hello")
	fmt.Fprintln(writer, "Hello World")
}

func getAllCustomer(writer http.ResponseWriter, request *http.Request) {
	customers := []Customer{
		{"Diogo Magliano", "Estrada da Vida 150, Casa 52", "22300000"},
		{"John Doe", "Rua de Baixo 455, Casa 101", "21500050"},
	}

	if request.Header.Get("Content-Type") == "application/xml" {
		writer.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(writer).Encode(customers)
	} else {
		writer.Header().Add("Content-Type", "application/json")
		//defining json, receive a writer and what to encode
		json.NewEncoder(writer).Encode(customers)
	}
}
