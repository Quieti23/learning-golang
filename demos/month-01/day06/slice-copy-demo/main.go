package main

import "fmt"

func main() {
	original := []string{"go", "java", "python"}
	shared := original[:]
	cloned := make([]string, len(original))
	copy(cloned, original)

	shared[0] = "golang"
	cloned[1] = "rust"

	fmt.Printf("original=%v\n", original)
	fmt.Printf("shared=%v\n", shared)
	fmt.Printf("cloned=%v\n", cloned)
}
