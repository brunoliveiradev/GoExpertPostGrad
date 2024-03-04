package main

import "fmt"

func main() {
	// create a channel to communicate between the main goroutine and the other goroutines
	channel := make(chan string)

	// Thread 2
	go reverseNameAndSendToChannel("banana", channel)

	// Thread 1 will wait for the value to be sent
	msg := <-channel
	fmt.Printf("The reversed name is: %s\n", msg)

	// Thread 3
	usingForever(10)
}

// The function first reverses the input string 'name' and then sends the reversed string to the provided channel.
func reverseNameAndSendToChannel(name string, channel chan string) {
	// print the reversed name
	var reversedName string
	for i := len(name) - 1; i >= 0; i-- {
		reversedName += string(name[i])
	}
	fmt.Println(reversedName)

	// send a value to the channel which can be used to communicate between different goroutines
	channel <- reversedName
}

func usingForever(times int) {
	forever := make(chan bool)

	//<-forever // this will block the main goroutine forever and cause deadlock
	// to fix that we can use another goroutine to receive the value from the channel

	go func() {
		for i := range times {
			fmt.Print(i, " ")
		}
		fmt.Println()
		forever <- true // send a value to the channel and unblock the main goroutine
	}()

	<-forever // this will unload the value from the channel and unblock the main goroutine
}
