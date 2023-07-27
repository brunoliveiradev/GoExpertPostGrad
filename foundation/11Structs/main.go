package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Active bool
	Adress
}

type Adress struct {
	Stret   string
	Number  int
	City    string
	State   string
	ZipCode int
}

type Account interface {
	Disable()
}

func (u User) Disable() {
	u.Active = false
	fmt.Printf("User %s is disabled \n", u.Name)
}

func DisableAccount(acc Account) {
	acc.Disable()
}

func main() {
	twitter := User{
		Name:   "Elon Musk",
		Age:    100,
		Active: true,
	}

	DisableAccount(twitter)

	fmt.Printf("User name: %s, Age: %d, Active: %t\n", twitter.Name, twitter.Age, twitter.Active)
}
