package main

import "fmt"

func main() {
	var myVar interface{} = "Random String!"

	// Go doesn't know the type of myVar because it's an interface and it can be anything
	println(myVar)

	// Type assertion is a way to tell Go what type the variable is
	println(myVar.(string))

	// If the type assertion is wrong, Go will return a zero value with false.
	// If you don't check the status, Go will throw an error called panic in case of false results at runtime
	res, ok := myVar.(int)
	fmt.Printf("The result is %v and the status is %v \n", res, ok)
}
