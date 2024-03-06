package main

import (
	"fmt"
	"sync"
	"time"
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

	// Demonstrating channel direction.
	channelDirectionDemo()

	// Demonstrating channel select statement.
	channelSelectDemo()

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

// channelDirectionDemo demonstrates the use of send-ONLY and receive-ONLY channel directions.
// By default, channels in Go are bidirectional, allowing both sending and receiving of values.
// A send-only channel (chan<- Type) can only be used to send data.
// A receive-only channel (<-chan Type) can only be used to receive data.
func channelDirectionDemo() {
	ch := make(chan int) // Create a "normal" bidirectional channel.

	wg := sync.WaitGroup{}
	wg.Add(2) // Wait for two goroutines.

	// Start a goroutine to send data into the send-only channel.
	go func() {
		defer wg.Done()
		sendData(ch)
	}()

	// Start a goroutine to receive data from the receive-only channel.
	go func() {
		defer wg.Done()
		receiveData(ch)
	}()

	wg.Wait() // Wait for both goroutines to finish.
	fmt.Println("Channel direction demonstration completed.")
}

// sendData simulates sending data to a send-only channel (chan<- Type).
// It accepts a send-only channel, denoted by (chan<- int), which means the channel can only be used to send integers.
func sendData(ch chan<- int) {
	// Simulate sending data.
	for i := 0; i < 5; i++ {
		ch <- i // Send data into the channel.
	}
	close(ch)
	fmt.Println("All data sent to the channel.")
}

// receiveData simulates receiving data from a receive-only channel (<-chan Type).
// It accepts a receive-only channel, denoted by (<-chan int), indicating the channel can only be used to receive integers.
func receiveData(ch <-chan int) {
	// Receive data until the channel is closed.
	for value := range ch {
		fmt.Printf("Received value: %d\n", value)
	}
	fmt.Println("All data received from the channel.")
}

// channelSelectDemo demonstrates the use of the select statement to receive from the first channel that is ready.
func channelSelectDemo() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(300 * time.Millisecond)
		ch1 <- "Hello"
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		ch2 <- "World"
	}()

	// Use the select statement to receive from the first channel that is ready.
	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	}
}
