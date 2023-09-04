package main

import "fmt"

func main() {

	var x interface{} = 10
	var y interface{} = "Hello, World!"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("The type of the variable is %T and the value is %v \n", t, t)
}
