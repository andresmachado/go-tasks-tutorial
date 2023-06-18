package main

import (
	"log"
	"net/http"

	"example.com/go-tasks-tutorial/api"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Endpoint to get all tasks
	router.HandleFunc("/tasks", api.GetTasks).Methods("GET")

	//Endpoint to create a new task
	router.HandleFunc("/tasks", api.CreateTask).Methods("POST")

	// Endpoint to retrieve a task
	router.HandleFunc("/tasks/{id}", api.GetTask).Methods("GET")

	// Endpoint to update a task
	router.HandleFunc("/tasks/{id}", api.UpdateTask).Methods([]string{"PUT", "PATCH"}...)

	// Endpoint to delete a task
	router.HandleFunc("/tasks/{id}", api.DeleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
