package httpdb

import (
	"fmt"
	"net/http"
)

type dollars float32

func (p dollars) String() string { return fmt.Sprintf("$%.2f", p) }

type Database map[string]dollars

func (db Database) List(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s", item, price)
	}
}

func (db Database) Price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "price of item %s : %s", item, price)
}
