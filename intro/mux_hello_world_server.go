package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Index1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world 1")
}

func Index2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world 2")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/index1", Index1)
	router.HandleFunc("/index2", Index2)
	log.Fatal(http.ListenAndServe(":8000", router))
}
