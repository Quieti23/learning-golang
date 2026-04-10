package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("case 1: start goroutine and do not wait")
	launchWithoutWaiting()

	fmt.Println()
	fmt.Println("case 2: start goroutine and wait for it")
	launchAndWait()
}

func launchWithoutWaiting() {
	go func() {
		fmt.Println("worker: started")
		fmt.Println("worker: finished")
	}()

	fmt.Println("main: launched worker")
	fmt.Println("main: sleep 10ms to give scheduler a chance")
	time.Sleep(10 * time.Millisecond)
	fmt.Println("main: function returns, process may end before worker fully runs")
}

func launchAndWait() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("worker: started")
		time.Sleep(20 * time.Millisecond)
		fmt.Println("worker: finished")
	}()

	fmt.Println("main: launched worker")
	wg.Wait()
	fmt.Println("main: worker joined, function returns")
}
