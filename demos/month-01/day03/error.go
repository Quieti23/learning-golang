package main

import "fmt"

func main() { 
 var a = 10
 var b = 0

 result, err := divide(a, b)
 if err != nil {
	 fmt.Println("-----:", err.Error())
	 return
 }
 fmt.Println(result)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}