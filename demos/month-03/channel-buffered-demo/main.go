package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 2)

	fmt.Println("buffer capacity:", cap(ch))
	fmt.Println("buffer length:", len(ch))

	fmt.Println("main: send task-1")
	ch <- "task-1"
	fmt.Println("main: send task-2")
	ch <- "task-2"
	fmt.Println("main: buffer is full now, len =", len(ch))

	go func() {
		fmt.Println("sender: trying to send task-3 into full buffer")
		ch <- "task-3"
		fmt.Println("sender: task-3 send completed after a receive freed one slot")
	}()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("main: receive", <-ch)
	fmt.Println("main: buffer length right after one receive:", len(ch))
	fmt.Println("main: receive", <-ch)
	fmt.Println("main: receive", <-ch)
	fmt.Println("main: buffer is empty now, len =", len(ch))

	go func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("late sender: send task-4 to unblock empty receive")
		ch <- "task-4"
	}()

	fmt.Println("main: trying to receive from empty buffer")
	fmt.Println("main: receive", <-ch)
	fmt.Println("main: receive completed after late sender arrived")
}
