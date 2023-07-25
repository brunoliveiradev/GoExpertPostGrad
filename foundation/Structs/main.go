package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	twitter := User{
		Name:   "Elon Musk",
		Age:    100,
		Active: true,
	}

	twitter.Active = false

	fmt.Printf("User name: %s, Age: %d, Active: %t\n", twitter.Name, twitter.Age, twitter.Active)
}
