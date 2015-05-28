package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	// My backend data store.
	shareTodo = Todo{}

	PORT       = ":8888"
	SERVER_URL = "http://localhost" + PORT
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/todos", handleListTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", handleGetTodo).Methods("GET")
	router.HandleFunc("/todos", handleCreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", handlePatchTodo).Methods("PATCH")
	router.HandleFunc("/todos", handleDeleteTodo).Methods("DELETE")

	// Serve the assets from the app directory.
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./app")))
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
	sendJson(w, todo.All())
}

func handleGetTodo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	sendJson(w, todo.Find(id))
}

func handleCreateTodo(w http.ResponseWriter, r *http.Request) {
	newItem := &TodoItem{}
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func handlePatchTodo(w http.ResponseWriter, r *http.Request) {
}

func handleDeleteTodo(w http.ResponseWriter, r *http.Request) {
}
