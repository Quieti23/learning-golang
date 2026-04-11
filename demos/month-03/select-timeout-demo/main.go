package main

import (
	"fmt"
	"time"
)

func main() {
	runCase("fast job", 100*time.Millisecond, 300*time.Millisecond)
	fmt.Println()
	runCase("slow job", 400*time.Millisecond, 200*time.Millisecond)
}

func runCase(name string, workDuration time.Duration, timeout time.Duration) {
	resultCh := make(chan string, 1)

	go func() {
		fmt.Println(name + ": worker started")
		time.Sleep(workDuration)
		resultCh <- name + ": result ready"
	}()

	fmt.Println(name+": waiting with timeout", timeout)

	select {
	case result := <-resultCh:
		fmt.Println(name+": select received result ->", result)
	case <-time.After(timeout):
		fmt.Println(name + ": select hit timeout")
	}
}
