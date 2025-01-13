package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", ProductsHandler)
	r.HandleFunc("/articles", ArticlesHandler)

	http.ListenAndServe(":8080", r)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home Page"))
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Products Page"))
}

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Articles Page"))
}
