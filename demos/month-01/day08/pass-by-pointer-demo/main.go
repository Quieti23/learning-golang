package main

import "fmt"

type Task struct {
	Title string
	Done  bool
}

func increase(number *int) {
	*number = *number + 1
	fmt.Println("inside increase:", *number)
}

func markDone(task *Task) {
	task.Done = true
	task.Title = "learn pointers in go"
}

func main() {
	count := 10
	fmt.Println("before increase:", count)
	increase(&count)
	fmt.Println("after increase:", count)

	task := Task{Title: "learn go", Done: false}
	fmt.Printf("before markDone: %+v\n", task)
	markDone(&task)
	fmt.Printf("after markDone: %+v\n", task)
}
