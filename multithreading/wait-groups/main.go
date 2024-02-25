package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, order string, waitGroup *sync.WaitGroup) {
	// print the reversed name
	var reversedName string
	for i := len(name) - 1; i >= 0; i-- {
		reversedName += string(name[i])
		fmt.Printf("i: %d, Task %s running\n", i, order)
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Println(reversedName)
	waitGroup.Done()
}

func main() {
	// create 3 goroutines to run the task function
	// and pass the waitGroup to the task function
	// so the task function can call waitGroup.Done() to notify the waitGroup that the task is done
	// and the waitGroup.Wait() will wait for all the tasks to finish
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(3)

	go task("banana", "A", &waitGroup)
	go task("apple", "B", &waitGroup)
	go task("orange", "C", &waitGroup)

	waitGroup.Wait()
}
