package main

// Adapted from TODO backend at
// https://github.com/savaki/todo-backend-gin

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	// My backend data store.
	sharedTodo = Todo{}

	PORT       = ":8888"
	SERVER_URL = "http://localhost" + PORT
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/todos", handleListTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", handleGetTodo).Methods("GET")
	router.HandleFunc("/todos", handleCreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", handlePatchTodo).Methods("PATCH")
	router.HandleFunc("/todos/{id}", handleDeleteTodo).Methods("DELETE")
	router.HandleFunc("/todos", handleDeleteAllTodos).Methods("DELETE")
	http.Handle("/", router)

	fmt.Println("Serving on " + SERVER_URL)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func sendJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleListTodos(w http.ResponseWriter, r *http.Request) {
	sendJson(w, sharedTodo.All())
}

func handleGetTodo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	sendJson(w, sharedTodo.Find(id))
}

func handleCreateTodo(w http.ResponseWriter, r *http.Request) {
	var newItem TodoItem
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sendJson(w, sharedTodo.Create(newItem, func(path string) string {
		return SERVER_URL + path
	}))
}

func handlePatchTodo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if item := sharedTodo.Find(id); item != nil {
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		sendJson(w, item)
	}
}

func handleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	sendJson(w, sharedTodo.Delete(id))
}

func handleDeleteAllTodos(w http.ResponseWriter, r *http.Request) {
	sendJson(w, sharedTodo.DeleteAll())
}
