package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
	"sync"
)

const (
	todosFile = "./todos.json"
	wrrPerm   = 0644
)

type Todo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Todos struct {
	sync.RWMutex
	NextId int     `json:"nextId"`
	Data   []*Todo `json:"data"`
}

var todos Todos

func LoadTodos() error {
	file, err := ioutil.ReadFile(todosFile)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &todos)
}

func saveTodos() error {
	b, err := json.Marshal(todos)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(todosFile, b, wrrPerm)
}

func findTodoIndex(id int) int {
	idx := sort.Search(len(todos.Data), func(i int) bool {
		return todos.Data[i].Id >= id
	})
	if idx < len(todos.Data) && todos.Data[idx].Id == id {
		return idx
	}
	return -1
}

func CreateTodo(name string) (*Todo, error) {
	todos.Lock()
	defer todos.Unlock()
	todo := &Todo{
		Id:   todos.NextId,
		Name: name,
	}
	todos.NextId++
	todos.Data = append(todos.Data, todo)
	if err := saveTodos(); err != nil {
		return nil, err
	}
	return todo, nil
}

func ListTodos() []*Todo {
	todos.RLock()
	defer todos.RUnlock()
	return todos.Data
}

func UpdateTodo(todo *Todo) error {
	todos.Lock()
	defer todos.Unlock()
	idx := findTodoIndex(todo.Id)
	if idx == -1 {
		return errors.New(fmt.Sprintf("Could not find todo with id: %d", todo.Id))
	}
	todos.Data[idx] = todo
	return saveTodos()
}

func DeleteTodo(id int) error {
	todos.Lock()
	defer todos.Unlock()
	idx := findTodoIndex(id)
	if idx == -1 {
		return errors.New(fmt.Sprintf("Could not find todo with id: %d", id))
	}
	todos.Data = append(todos.Data[:idx], todos.Data[idx+1:]...)
	return saveTodos()
}
