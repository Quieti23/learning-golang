package main

import "fmt"

func main() {

	var a int = 10
	var b int = 20

	var c int = a + b

	var d = a + b

	const e = "hello world"

	string := "hello world1"

	int := 10
	bool := true

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(string)
	fmt.Println(int + a)
	fmt.Println(bool)

	for i := 0; i < 10; i++ {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}

	switch a {
	case 1:
		fmt.Println("a is 1")
	case 2:
		fmt.Println("a is 2")
	default:
		fmt.Println("a is not 1 or 2")
	}

}