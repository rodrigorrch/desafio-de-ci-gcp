package main

import "fmt"

func sum(value_a int, value_b int) int {
	result := value_a + value_b
	return result
}

func main() {
	value_a := 5
	value_b := 5
	fmt.Printf("%d \n", sum(value_a, value_b))
}
