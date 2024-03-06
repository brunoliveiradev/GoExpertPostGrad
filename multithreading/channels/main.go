package main

import (
	"fmt"
	"sync"
)

func main() {
	// Demonstrating basic channel operation: sending and receiving.
	basicChannelDemo()

	// Demonstrating a channel to signal and block the main goroutine.
	signalAndBlockDemo()

	// Demonstrating channel operations with publisher and subscriber pattern.
	publishSubscribeDemo()

	// Demonstrating channel and WaitGroup to synchronize multiple goroutines.
	waitGroupDemo(10)
}

func basicChannelDemo() {
	channel := make(chan string)
	go reverseNameAndSendToChannel("banana", channel)
	msg := <-channel
	fmt.Printf("The reversed name is: %s\n", msg)
}

func reverseNameAndSendToChannel(name string, channel chan string) {
	var reversedName string
	for i := len(name) - 1; i >= 0; i-- {
		reversedName += string(name[i])
	}
	channel <- reversedName // Send reversed name to the channel.
}

func signalAndBlockDemo() {
	forever := make(chan bool)
	//<-forever  // using forever just after will block the main goroutine forever and cause deadlock kept here for example

	go func() {
		fmt.Println("Signal and block demo running.")
		forever <- true // Send a signal to unblock the main goroutine.
	}()

	<-forever // Wait for signal.
}

func publishSubscribeDemo() {
	ch := make(chan int)
	go publishIntoChannel(ch, 10)
	consumeFromChannel(ch) // By not using goroutine to consume, we want to wait for the consumption to complete before moving on in the main goroutine
}

func publishIntoChannel(channel chan int, size int) {
	for i := range size {
		channel <- i
	}
	fmt.Println("Closing the channel")
	close(channel) // It's important to close the channel to signal that no more values will be sent.
}

func consumeFromChannel(channel chan int) {
	for x := range channel {
		fmt.Printf("Received: %d\n", x)
	}
}

func waitGroupDemo(size int) {
	channel := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(size)

	go func() {
		for i := range size {
			channel <- i
		}
		fmt.Println("WaitGroup Channel closed.")
		close(channel) // Ensure to close the channel after publishing to prevent deadlock.
	}()

	// By using a goroutine to consume the channel, you ensure that the consuming process is non-blocking and to perform operations concurrently
	go func() {
		for x := range channel {
			fmt.Printf("WaitGroup demo received: %d\n", x)
			wg.Done()
		}
	}()

	wg.Wait() // Wait for all goroutines to finish.
}
