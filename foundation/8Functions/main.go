package main

import (
	"errors"
	"fmt"
)

var (
	a = "Class about Functions"
)

func main() {
	fmt.Println(a)

	value, err := sum(50, 10)

	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Value:", value)
}

func sum(a int, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("sum bigger than 50")
	}
	return a + b, nil
}
