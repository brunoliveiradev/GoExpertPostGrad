package main

import (
	"fmt"
)

var (
	a = "Class about Anonymous Functions or Closures"
)

func main() {
	fmt.Println(a)

	total := func() int {
		return sum(50, 10, 15, 2, 3, 5, 8, 13, 21) * 2
	}()
	fmt.Println("Total Value:", total)

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
