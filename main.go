package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", staticTextHandler)

	fmt.Println("Starting server at port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalln(err)
	}
}

func staticTextHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Hello, world!")
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Printf("form parsing error %v", err)
		return
	}
	name := request.FormValue("name")
	address := request.FormValue("address")
	fmt.Fprintln(writer, "Name: ", name, "Address: ", address)
}
