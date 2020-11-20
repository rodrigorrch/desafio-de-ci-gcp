package main

import "fmt"

func Sum(x int, y int) int {
	return x + y
}

func main() {
	value_a := 5
	value_b := 5
	fmt.Printf("%d \n", Sum(value_a, value_b))
}
