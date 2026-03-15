package main

import "fmt"

func main() {
	scores := map[string]int{
		"go":   85,
		"java": 90,
	}

	fmt.Println("before update:", scores)

	scores["python"] = 80
	scores["go"] = 95

	fmt.Println("after update:", scores)
	fmt.Printf("java score = %d\n", scores["java"])
}
