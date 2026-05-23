package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("WELCOME TO MY GO LANG MODULES")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))

}

func greeter() {
	fmt.Println("hello my mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome To Golang Series  on YT<h1>"))
}

// COMMANDS FOR GO
// go mod why github.com/gorilla/mux
// go mod graph
// go list -m all
// go mod verify
// go list
// go mod tidy
