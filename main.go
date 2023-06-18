package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/andresmachado/go-tasks-tutorial/views"

)

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"Description"`
}

var tasks []Task

func main() {
	router := mux.NewRouter()

	// Endpoint to get all tasks
	router.HandleFunc("/tasks", GetTasks).Methods("GET")

	//Endpoint to create a new task
	router.HandleFunc("/tasks", CreateTask).Methods("POST")

	// Endpoint to retrieve a task
	router.HandleFunc("/tasks/{id}", GetTask).Methods("GET")

	// Endpoint to update a task
	router.HandleFunc("/tasks/{id}", UpdateTask).Methods([]string{"PUT", "PATCH"}...)

	// Endpoint to delete a task
	router.HandleFunc("/tasks/{id}", DeleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
