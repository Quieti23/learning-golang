package main

import (
	"encoding/json"
	"fmt"
)

type TaskWithoutTag struct {
	ID    int
	Title string
	Done  bool
}

type TaskWithTag struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func main() {
	withoutTag := TaskWithoutTag{ID: 1, Title: "without tag", Done: false}
	withTag := TaskWithTag{ID: 1, Title: "with tag", Done: false}

	withoutTagJSON, err := json.Marshal(withoutTag)
	if err != nil {
		fmt.Println("marshal without tag error:", err)
		return
	}

	withTagJSON, err := json.Marshal(withTag)
	if err != nil {
		fmt.Println("marshal with tag error:", err)
		return
	}

	fmt.Println("without tag:", string(withoutTagJSON))
	fmt.Println("with tag:", string(withTagJSON))

	badInput := []byte(`{"id":8,"title":"bad json","done":false}`)
	var task TaskWithTag
	if err := json.Unmarshal(badInput, &task); err != nil {
		fmt.Println("unmarshal error:", err)
	}
	fmt.Println("task:", task)
}
