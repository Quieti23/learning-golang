package main

import "fmt"

func main() {
	fmt.Println("main start")
	runDemo()
	fmt.Println("main end")
}

func runDemo() {
	fmt.Println("runDemo start")
	defer fmt.Println("defer 1: run before function returns")
	defer fmt.Println("defer 2: executed first because defer is LIFO")
	fmt.Println("runDemo working")
}
