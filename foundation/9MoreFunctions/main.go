package main

import (
	"fmt"
)

var (
	a = "Class about Functions"
)

func main() {
	fmt.Println(a)

	totalSum := sum(50, 10, 15, 2, 3, 5, 8, 13, 21)

	fmt.Println("Sum:", totalSum)
}

func sum(numbers ...int) int {
	total := 0

	for _, num := range numbers {
		total += num
	}
	return total
}
