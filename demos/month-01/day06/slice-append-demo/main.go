package main

import "fmt"

func main() {
	backing := make([]int, 2, 4)
	backing[0] = 1
	backing[1] = 2

	shared := backing
	fmt.Printf("before append: backing=%v len=%d cap=%d shared=%v\n", backing, len(backing), cap(backing), shared)

	shared = append(shared, 3)
	shared[0] = 99
	fmt.Printf("after append within capacity: backing=%v shared=%v\n", backing, shared)

	grown := append(shared, 4, 5)
	grown[1] = 777
	fmt.Printf("after append beyond capacity: shared=%v grown=%v\n", shared, grown)
}
