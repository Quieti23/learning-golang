package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("sender: trying to send 42")
		ch <- 42
		fmt.Println("sender: send finished after receiver arrived")
	}()

	fmt.Println("main: sleep 100ms, no receiver yet")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("main: ready to receive")
	value := <-ch
	fmt.Println("main: received", value)

	time.Sleep(20 * time.Millisecond)
}
