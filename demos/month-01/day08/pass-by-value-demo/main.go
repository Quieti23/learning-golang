package main

import "fmt"

func increase(number int) {
	number = number + 1
	fmt.Println("inside increase:", number)
}

func main() {
	count := 10
	fmt.Println("before increase:", count)
	increase(count)
	fmt.Println("after increase:", count)
}
