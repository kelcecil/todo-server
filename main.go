package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

func main() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: MakeTodoHandler(),
	}
	log.Println("Starting server.")
	server.ListenAndServe()
}

// TodoHandler is a handler intended to add, list, and remove
// TODO items via an API.
type TodoHandler struct {
	storage *sync.Map
}

// MakeTodoHandler is a convenience functin to create a
// new TodoHandler.
func MakeTodoHandler() TodoHandler {
	return TodoHandler{
		storage: &sync.Map{},
	}
}

// Todo is the primary structure for working with Todos.
type Todo struct {
	Key         string `json:"key"`
	Description string `json:"description,omitempty"`
}

// ListTodo is the expected format for retrieving TODOs.
type ListTodo struct {
	Items []Todo `json:"items"`
}

func (h TodoHandler) ServeHTTP(response http.ResponseWriter, req *http.Request) {
	parts := strings.Split(strings.TrimPrefix(req.URL.Path, "/"), "/")
	action := parts[0]
	switch req.Method {
	case http.MethodPost:
		switch action {
		case "add":
			h.AddTodo(response, req)
		case "delete":
			h.DeleteTodo(response, req)
		}
	case http.MethodGet:
		switch action {
		case "list":
			h.ListTodos(response)
		}
	}
}

// AddTodo is responsible for writing a TODO.
func (h *TodoHandler) AddTodo(response http.ResponseWriter, req *http.Request) {
	newTodo := parseTodoFromBytes(req)
	h.AddTodoToStore(newTodo)
	response.WriteHeader(200)
}

func parseTodoFromBytes(req *http.Request) Todo {
	body, _ := ioutil.ReadAll(req.Body)
	log.Println(string(body))
	var newTodo Todo
	err := json.Unmarshal(body, &newTodo)
	if err != nil {
		log.Println(err)
	}
	return newTodo
}

func (h *TodoHandler) AddTodoToStore(item Todo) {
	log.Printf("Adding %v to storage", item.Key)
	h.storage.Store(item.Key, item)
}

func (h *TodoHandler) DeleteTodo(response http.ResponseWriter, req *http.Request) {
	candidateTodo := parseTodoFromBytes(req)
	h.storage.Delete(candidateTodo.Key)
	response.WriteHeader(200)
}

func (h *TodoHandler) ListTodos(response http.ResponseWriter) {
	todos := h.ListTodosFromStore()
	jsonResponse, err := json.Marshal(todos)
	if err != nil {
		log.Println("Failed to write JSON response")
	}
	response.WriteHeader(200)
	response.Write(jsonResponse)
}

func (h *TodoHandler) ListTodosFromStore() ListTodo {
	todos := make([]Todo, 0)
	h.storage.Range(func(key, value interface{}) bool {
		todos = append(todos, value.(Todo))
		return true
	})

	response := ListTodo{}
	response.Items = todos
	return response
}

func formatErrorsResponse(errors ...string) interface{} {
	return struct {
		items []string
	}{
		items: errors,
	}
}
