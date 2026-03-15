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
	task := Task{
		ID:    1,
		Title: "learn json marshal",
		Done:  false,
	}

	data, err := json.Marshal(task)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(string(data))
}
