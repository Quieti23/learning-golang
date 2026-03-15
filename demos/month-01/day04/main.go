package main

import "fmt"
import . "Task"

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