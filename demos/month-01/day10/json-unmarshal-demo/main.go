package main

import (
	"encoding/json"
	"fmt"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func main() {
	input := []byte(`{"id":2,"title":"learn json unmarshal","done":true}`)

	var task Task
	if err := json.Unmarshal(input, &task); err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("decoded task: %+v\n", task)
}
