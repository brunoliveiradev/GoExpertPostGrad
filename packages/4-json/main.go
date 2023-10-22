package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Account struct {
	Number  int     `json:"number"` // json tag is used to specify the name of the field in the JSON object
	Balance float64 `json:"balance"`
}

func main() {
	account := Account{Number: 0001, Balance: 1000.50}

	// Convert struct to JSON - Marshal returns a byte array
	accountJson, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}
	// Convert byte array to string
	println(string(accountJson))

	// json.NewEncoder with os.Stdout as the destination is a convenient way to write JSON in a stream to the terminal
	// without having to save it to a file first
	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		panic(err)
	}

	// Convert JSON to struct - Unmarshal takes a byte array and a pointer to a struct
	fakeJson := []byte(`{"number":444,"balance":1660.50}`)

	var jsonAccount Account
	err = json.Unmarshal(fakeJson, &jsonAccount)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%.2f\n", jsonAccount.Balance)
}
