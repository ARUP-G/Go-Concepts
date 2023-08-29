package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Mod in Golang")

	// Router for port connection
	router := mux.NewRouter()

	// Router handler
	router.HandleFunc("/", server).Methods("GET")

	// Error handler
	log.Fatal(http.ListenAndServe(":3000", router))

}

func greet() {
	fmt.Println("Welcome")
}

// Function for page output
func server(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> HELLO </h1>"))
}
