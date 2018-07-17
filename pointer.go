package main

import "fmt"

func main() {
	var i = 5
	fmt.Printf("An integer: %d, its location in memory: %p\n", i, &i)

	var intP *int
	intP = &i
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
}
