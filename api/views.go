package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var tasks []Task

type ResponseMessage struct {
	Message string `json:"message"`
}

// getTasks returns all tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if len(tasks) == 0 {
		json.NewEncoder(w).Encode([]Task{})
		return
	}

	// Fetch tasks from main module and return to the encoder
	json.NewEncoder(w).Encode(tasks)
}

// createTask create a new tasks with the params
func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Define a new task object
	var task Task

	// Generate a new ID
	task.ID = strconv.Itoa(len(tasks) + 1)

	_ = json.NewDecoder(r.Body).Decode(&task)

	// Create a new task by appending a new object to the list of tasks
	tasks = append(tasks, task)

	// Return the new created tasks
	json.NewEncoder(w).Encode(task)
}

// getTask fetch a new task with the given ID
func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get all request params
	params := mux.Vars(r)

	// Iterate over the list of tasks and check if the task exists, if not
	// return an err
	for _, task := range tasks {
		if task.ID == params["id"] {
			json.NewEncoder(w).Encode(task)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	errorResponse := ResponseMessage{Message: "Task not found"}
	json.NewEncoder(w).Encode(errorResponse)
}

// updateTask fetches and updates the task by the given id and body params
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var updatedTask Task
	_ = json.NewDecoder(r.Body).Decode(&updatedTask)

	for index, task := range tasks {
		if task.ID == params["id"] {
			updatedTask.ID = task.ID

			tasks[index] = updatedTask

			json.NewEncoder(w).Encode(updatedTask)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	errorResponse := ResponseMessage{Message: "Task not found"}
	json.NewEncoder(w).Encode(errorResponse)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, task := range tasks {
		if task.ID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			successResponse := ResponseMessage{Message: "Task deleted."}
			json.NewEncoder(w).Encode(successResponse)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	errorResponse := ResponseMessage{Message: "Task not found"}
	json.NewEncoder(w).Encode(errorResponse)
}
