package main

import "fmt"

type Client struct {
	name string
}

func (c Client) setName(name string) {
	c.name = name
	fmt.Println("The name of the client inside setName is", c.name)
}

type Account struct {
	balance int
}

// The simulation method is not changing the value of the balance because it is not a pointer (it is a copy of the value)
func (acc Account) simulation(value int) int {
	acc.balance += value
	fmt.Println("The balance of the account inside simulation is", acc.balance)
	return acc.balance
}

// The * is used to indicate that the method is a pointer, and it will change the original value of the balance
func (acc *Account) deposit(value int) int {
	acc.balance += value
	fmt.Println("The balance of the account inside deposit is", acc.balance)
	return acc.balance
}

func main() {

	newClient := Client{
		name: "Bruno",
	}

	newClient.setName("Bruno Henrique")
	fmt.Println("The name of the client is", newClient.name)

	Account := Account{
		balance: 100,
	}

	// The simulation method is not changing the value of the balance because it is not a pointer
	Account.simulation(200)
	fmt.Println("The balance of the account after simulation is", Account.balance)

	// The deposit method is changing the value of the balance because it is a pointer
	Account.deposit(200)
	fmt.Println("The balance of the account after deposit is", Account.balance)

}
