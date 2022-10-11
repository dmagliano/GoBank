package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		println("request to hello")
		fmt.Fprintln(writer, "Hello World")
	})

	http.ListenAndServe("localhost:8000", nil)

}
