package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", pingHandler)

	fmt.Println("server running at http://localhost:8080")
	fmt.Println("try GET /ping")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("server error:", err)
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request received:", r.Method, r.URL.Path)
	fmt.Fprintln(w, "pong")
}
