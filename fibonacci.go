package main

import "fmt"

func fibonacci(num int) int {
	a, b := 0, 1
	for c := 0; c < num; c++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	c := fibonacci(100)
	fmt.Println(c)
}
