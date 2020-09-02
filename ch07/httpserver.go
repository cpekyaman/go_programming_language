package main

import (
	"log"
	"net/http"

	"github.com/cpekyaman/go_programming_language/ch07/httpdb"
)

func main() {
	backend := httpdb.Database{"shoes": 50}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(backend.List))
	mux.Handle("/price", http.HandlerFunc(backend.Price))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
