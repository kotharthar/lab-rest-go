package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

/* Get All Items */
func getItems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All Items")
}

func createItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create Item")
}

func getItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	item_id := vars["item_id"]
	fmt.Fprintf(w, "Get Item for %s", item_id)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	item_id := vars["item_id"]
	fmt.Fprintf(w, "Update Item for %s", item_id)
}

func destroyItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	item_id := vars["item_id"]
	fmt.Fprintf(w, "Destroy Item for %s", item_id)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/items", getItems).Methods("GET")
	r.HandleFunc("/items", createItem).Methods("POST")
	r.HandleFunc("/items/{item_id}", getItem).Methods("GET")
	r.HandleFunc("/items/{item_id}", updateItem).Methods("PUT")
	r.HandleFunc("/items/{item_id}", destroyItem).Methods("DELETE")
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}
