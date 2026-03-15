package main

import "fmt"

func main() {
	count := 10
	pointer := &count

	fmt.Println("count:", count)
	fmt.Println("address of count:", &count)
	fmt.Println("pointer value:", pointer)
	fmt.Println("value through pointer:", *pointer)

	*pointer = 20
	fmt.Println("count after *pointer = 20:", count)
}
