package main

import (
	"fmt"
	"time"
)

func main() {
	// Demonstrating load balancer using channels.
	loadBalancerDemo()
}

func loadBalancerDemo() {
	dataCh := make(chan int)
	qtWorkers := 10

	// Create workers to process the data
	for i := 0; i < qtWorkers; i++ {
		go worker(i, dataCh)
	}

	// Send data to the workers
	for i := 0; i < 100; i++ {
		dataCh <- i
	}
}

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received: %d\n", workerId, x)
		time.Sleep(500 * time.Millisecond)
	}
}
