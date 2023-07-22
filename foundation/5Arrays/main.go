package main

import "fmt"

var (
	a = "Class about array"
)

func main() {
	fmt.Println(a)

	var array [3]int
	array[0] = 10
	array[1] = 20
	array[2] = 30

	fmt.Println("Last element of the array has the value:", array[len(array)-1])

	for indice, value := range array {
		fmt.Printf("O valor do indice %d é %d \n", indice, value)
	}
}
