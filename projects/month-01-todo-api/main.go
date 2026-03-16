package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"month01todoapi/handler"
	"month01todoapi/service"
	"month01todoapi/store"
)

func main() {
	taskStore := store.NewTaskStore()
	taskService := service.NewTaskService(taskStore)
	taskHandler := handler.NewTaskHandler(taskService)

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
	})
	taskHandler.RegisterRoutes(mux)

	fmt.Println("todo api running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("server error:", err)
	}
}
