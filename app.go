package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

const templ = `<h1>New Item</h1>
{{.Aok}}
<form action="items" method="post">
    <input type="" id="name" value="" name="" />
    <input type="Submit" class="button" value="" name="Create" />
</form>
`

type Data struct {
	Aok string
}

/* Get All Items */
func getItems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All Items")
}

func newItem(w http.ResponseWriter, r *http.Request) {
	data := Data{
		Aok: "aok",
	}

	t := template.New("newItemForm")
	//t, _ = t.ParseFiles(templ)
	t, _ = t.ParseFiles("newItem.html")
	t.Execute(w, data)
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
	r.HandleFunc("/items/new", newItem).Methods("GET")
	r.HandleFunc("/items/{item_id}", getItem).Methods("GET")
	r.HandleFunc("/items/{item_id}", updateItem).Methods("PUT")
	r.HandleFunc("/items/{item_id}", destroyItem).Methods("DELETE")
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}
