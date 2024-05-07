package main

import (
	"log"
	"net/http"
)

// home handler
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the appy."))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a particular snippet."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display form for creating snippet."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("starting server on: 4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
