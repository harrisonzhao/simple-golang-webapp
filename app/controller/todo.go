package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/harrisonzhao/simple-golang-webapp/app/model"
)

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Name string `json:"name"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if checkError(w, err) {
		return
	}
	todo, err := model.CreateTodo(data.Name)
	if checkError(w, err) {
		return
	}
	jsonResponse(w, todo)
}

func TodosList(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Data []*model.Todo `json:"data"`
	}{
		Data: model.ListTodos(),
	}
	jsonResponse(w, data)
}

func TodoUpdate(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if checkError(w, err) {
		return
	}
	err = model.UpdateTodo(&todo)
	checkError(w, err)
}

func TodoDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.Atoi(idString)
	if checkError(w, err) {
		return
	}
	if checkError(w, model.DeleteTodo(id)) {
		return
	}
	w.WriteHeader(http.StatusOK)
}

func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	payload, err := json.Marshal(data)
	if checkError(w, err) {
		return
	}

	fmt.Fprintf(w, string(payload))
}

func checkError(w http.ResponseWriter, err error) bool {
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}
