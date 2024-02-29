package main

import "fmt"

// Thread 1
func main() {
	// create a channel to communicate between the main goroutine and the other goroutines
	// the channel will be used to wait for the other goroutines to finish
	// the channel will be used to send a value to the main goroutine
	// the main goroutine will wait for the value to be sent
	// the other goroutines will send the value to the channel
	// when they finish
	channel := make(chan string)

	// Thread 2
	go func() {
		// print the reversed name
		var reversedName string
		for i := len("banana") - 1; i >= 0; i-- {
			reversedName += string("banana"[i])
		}
		fmt.Println(reversedName)
		// send a value to the channel
		channel <- reversedName
		channel <- "Cant print this" // this will not be printed because the main goroutine will wait for the value to be sent
	}()

	// Thread 1 will wait for the value to be sent
	msg := <-channel
	fmt.Printf("The reversed name is: %s\n", msg)
}
