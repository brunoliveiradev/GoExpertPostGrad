package main

import (
	"fmt"
	"time"
)

func task(name string, order string) {
	// print the reversed name
	var reversedName string
	for i := len(name) - 1; i >= 0; i-- {
		reversedName += string(name[i])
		fmt.Printf("i: %d, Task %s running\n", i, order)
	}
	fmt.Println(reversedName)
}

// Thread 1
func main() {
	// create 4 goroutines to run the task function
	go task("banana", "A") // Thread 2
	go task("apple", "B")  // Thread 3
	go task("orange", "C") // Thread 4
	go task("socorram me subi no onibus em marrocos", "D")
	// wait for the goroutines to finish
	// nothing will be printed
	// because the main goroutine will exit before the other goroutines finish
	// so the other goroutines will be terminated
	// to fix that we can use a channel to wait for the goroutines to finish
	// or we can use a wait group
	// or we can use a time.Sleep
	// or we can use a sync.Mutex, etc...
	// Doing nothing will cause the other goroutines to be terminated
	time.Sleep(6 * time.Second)
	fmt.Println("Main goroutine finished")
}
