package main

import "fmt"

var (
	a = "Class about Slices"
)

func main() {
	fmt.Println(a)

	mySlice := []int{0, 1, 2, 3, 5, 8, 13, 21, 42}
	fmt.Printf("len = %d, cap = %d, values = %v \n", len(mySlice), cap(mySlice), mySlice)

	fmt.Printf("len = %d, cap = %d, values = %v \n", len(mySlice[:0]), cap(mySlice[:0]), mySlice[:0])

	fmt.Printf("len = %d, cap = %d, values = %v \n", len(mySlice[:4]), cap(mySlice[:4]), mySlice[:4])

	fmt.Printf("len = %d, cap = %d, values = %v \n", len(mySlice[2:]), cap(mySlice[2:]), mySlice[2:])

	fmt.Println("Len was 9 and will be 18 at the next step. \n" +
		"When you add a new element at the end of slice they double the capacity. \n" +
		"Its a good pratice to initialize a slice with the proximity capacity that you need.")

	mySlice = append(mySlice, 66)
	fmt.Printf("len = %d, cap = %d, values = %v \n", len(mySlice[2:]), cap(mySlice[:2]), mySlice[:2])

	fmt.Printf("len = %d, cap = %d, values = %v \n", len(mySlice), cap(mySlice), mySlice)

}
