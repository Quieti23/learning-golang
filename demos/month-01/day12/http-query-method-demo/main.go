package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/tasks/", taskDetailHandler)

	fmt.Println("server running at http://localhost:8080")
	fmt.Println("try GET /ping")
	fmt.Println("try GET /greet?name=eson")
	fmt.Println("try GET /tasks/123")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("server error:", err)
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("request received:", r.Method, r.URL.Path)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "pong")
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	fmt.Println("request received:", r.Method, r.URL.Path, "name=", name)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello, %s\n", name)
}

func taskDetailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/tasks/")
	if path == "" || strings.Contains(path, "/") {
		http.Error(w, "invalid task path", http.StatusBadRequest)
		return
	}

	fmt.Println("request received:", r.Method, r.URL.Path)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "task id: %s\n", path)
}
