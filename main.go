package main

import (
	"fmt"
	"net/http"

	"task/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	events := make(chan string)

	r.HandleFunc("/task", handlers.Create).Methods("POST")
	r.HandleFunc("/list", handlers.GetList).Methods("GET")
	r.HandleFunc("/task", handlers.Update).Methods("PATCH")

	fmt.Printf("Starting server on 8000")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		fmt.Printf("Failed to start server; err: %v", err)
	}
}
