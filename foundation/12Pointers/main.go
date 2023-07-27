package main

import "fmt"

func main() {

	// Memory address -> &variable -> 0xc0000b4008 -> value
	// Pointer -> *variable -> 0xc0000b4008 -> value
	// Dereferencing -> *pointer -> value -> &variable -> 0xc0000b4008
	// Assigning -> *pointer = value -> value -> &variable -> 0xc0000b4008
	// Assigning -> pointer = &variable -> 0xc0000b4008 -> value

	a := 10
	fmt.Printf("Address of a: %p\n", &a)

	// pointerExample := &a is same as var pointerExample *int = &a
	// the & operator generates a pointer to its operand and is called the address operator
	var pointerExample *int = &a
	fmt.Printf("Address of pointerExample: %p, and is (pointerExample == &a) equal? : %t\n", pointerExample, pointerExample == &a)

	// *pointerExample = value is same as a = value
	*pointerExample = 20
	fmt.Printf("Value of a: %d and value of *pointerExample: %d\n", a, *pointerExample)

	// b:= &a is same as var b *int = &a
	// the difference is that b is a pointer to an int and a is an int
	// b is a pointer to a, so b stores the address of a
	b := &a
	fmt.Printf("Address of b: %p, and is (b == &a) equal? : %t\n", b, b == &a)

	// *b = value is same as a = value
	*b = 30
	fmt.Printf("Value of a: %d and value of *b: %d\n", a, *b)
}
