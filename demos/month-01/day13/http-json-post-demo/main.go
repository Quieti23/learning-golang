package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Task struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/tasks", createTaskHandler)

	fmt.Println("server running at http://localhost:8080")
	fmt.Println("try GET /ping")
	fmt.Println("try POST /tasks with JSON body")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("server error:", err)
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
}

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, ErrorResponse{Message: "method not allowed"})
		return
	}

	defer r.Body.Close()

	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Message: "invalid JSON body"})
		return
	}

	if task.Title == "" {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Message: "title is required"})
		return
	}

	writeJSON(w, http.StatusCreated, task)
}

func writeJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
