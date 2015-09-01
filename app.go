package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func getItems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All Items")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/items", getItems)
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}
