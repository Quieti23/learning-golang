package main

import "fmt"

type Task struct {
    ID    int
    Title string
    Done  bool
}

func (task Task) Summary() string {
    return fmt.Sprintf("Task[%d]: %s, done=%v", task.ID, task.Title, task.Done)
}

func (task Task) IsDone() bool {
    return task.Done
}

func (task *Task) MarkDone() {
    task.Done = true
}

func (task *Task) Rename(newTitle string) {
    task.Title = newTitle
}

func main() {
    task := Task{
        ID:    1,
        Title: "learn go struct and method",
        Done:  false,
    }

    fmt.Println("before:", task.Summary())
    fmt.Println("is done?", task.IsDone())

    task.Rename("learn go receiver")
    task.MarkDone()

    fmt.Println("after:", task.Summary())
    fmt.Println("is done?", task.IsDone())
}