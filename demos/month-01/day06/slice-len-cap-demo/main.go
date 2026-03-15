package main

import "fmt"

func main() {
	empty := []int{}
	fixed := make([]int, 3, 5)
	arr := [5]int{10, 20, 30, 40, 50}
	sliced := arr[1:4]

	fmt.Printf("empty: value=%v len=%d cap=%d\n", empty, len(empty), cap(empty))
	fmt.Printf("fixed: value=%v len=%d cap=%d\n", fixed, len(fixed), cap(fixed))
	fmt.Printf("sliced from array: value=%v len=%d cap=%d\n", sliced, len(sliced), cap(sliced))
}
