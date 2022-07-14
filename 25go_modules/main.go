package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello, world!")
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")

	// Listen and serve on port 2000
	//err := http.ListenAndServe(":2000", r)
	//if err != nil {
	//	return
	//}

	log.Fatal(http.ListenAndServe(":2000", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	write, err := w.Write([]byte("<h1>Hello, world!</h1>"))
	if err != nil {
		return
	}

	fmt.Println(write)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Connection", "close")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
}
