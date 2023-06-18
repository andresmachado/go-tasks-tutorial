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

	// Serve index page
	router.HandleFunc("/", serveIndexPage).Methods("GET")

	// Serve AddTask page
	router.HandleFunc("/task/add", serveAddTaskPage).Methods("GET")

	// Serve EditTask page
	router.HandleFunc("/task/edit", serveUpdateTaskPage).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func serveIndexPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func serveAddTaskPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/add-task.html")
}

func serveUpdateTaskPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/edit-task.html")
}
