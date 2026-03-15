package main

import "fmt"

func main() {
	status := map[string]string{
		"task-1": "todo",
		"task-2": "done",
	}

	fmt.Println("before delete:", status)

	delete(status, "task-2")
	delete(status, "task-3")

	fmt.Println("after delete:", status)

	value, ok := status["task-2"]
	fmt.Printf("task-2 => value=%q ok=%v\n", value, ok)
}
