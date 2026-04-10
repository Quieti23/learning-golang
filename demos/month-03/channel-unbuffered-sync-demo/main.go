package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		fmt.Println("receiver: waiting for value")
		message := <-ch
		fmt.Println("receiver: got value:", message)
	}()

	time.Sleep(20 * time.Millisecond)
	fmt.Println("sender: sending value")
	ch <- "ping"
	fmt.Println("sender: send completed after receiver was ready")

	time.Sleep(20 * time.Millisecond)
}
