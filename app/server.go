package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harrisonzhao/simple-golang-webapp/app/controller"
	"github.com/harrisonzhao/simple-golang-webapp/app/model"
)

func main() {
	if err := model.LoadTodos(); err != nil {
		log.Fatal("Could not load todos from file ", err)
	}
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/todos", controller.TodoCreate).Methods("POST")
	r.HandleFunc("/todos", controller.TodosList).Methods("GET")
	r.HandleFunc("/todos", controller.TodoUpdate).Methods("PUT")
	r.HandleFunc("/todos/{id}", controller.TodoDelete).Methods("DELETE")
	r.HandleFunc("/", controller.Index).Methods("GET")
	r.PathPrefix("/static").Handler(http.FileServer(http.Dir("./")))
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
