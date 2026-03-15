package main

import "fmt"

func main() {
	ages := map[string]int{
		"alice": 28,
		"bob":   32,
	}

	fmt.Printf("direct read missing key: ages[\"charlie\"] = %d\n", ages["charlie"])

	age, ok := ages["charlie"]
	fmt.Printf("with ok check: age=%d ok=%v\n", age, ok)

	age, ok = ages["alice"]
	fmt.Printf("existing key: age=%d ok=%v\n", age, ok)
}
