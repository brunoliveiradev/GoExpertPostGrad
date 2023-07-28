package main

import "fmt"

// This function takes two pointers to integers as arguments
// and returns the sum of the values pointed to by the pointers
func sum(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {
	myVar1 := 10
	myVar2 := 20
	myVarSum := sum(&myVar1, &myVar2)

	fmt.Println(myVarSum)
	// The value of myVar1 is changed to 50
	fmt.Println(myVar1)
}
