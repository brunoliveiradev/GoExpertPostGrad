package main

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

var count uint64

func main() {
	//concurrencySimpleProblem()

	//concurrencySimpleProblemWithSolution()
	concurrencySimpleProblemWithMoreViableSolution()
}

// This function will create a simple web server that will count the number of visits
// but when called by multiple clients, concurrently, it will not work as expected
func concurrencySimpleProblem() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		count++
		w.Write([]byte(fmt.Sprintf("Você recebeu mais uma visita! Total: %d!", count)))
	})

	http.ListenAndServe(":3000", nil)
}

// This solution uses a mutex to lock the count variable
func concurrencySimpleProblemWithSolution() {
	m := sync.Mutex{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m.Lock() // lock the count variable, which will prevent other goroutines from accessing it
		count++
		m.Unlock()
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você recebeu mais uma visita! Total: %d!", count)))
	})

	http.ListenAndServe(":3003", nil)
}

// This solution uses a atomic operation to increment the count variable
func concurrencySimpleProblemWithMoreViableSolution() {
	//m := sync.Mutex{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//m.Lock()
		atomic.AddUint64(&count, 1) // internally it uses a mutex to lock the count variable
		//m.Unlock()
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você recebeu mais uma visita! Total: %d!", count)))
	})

	http.ListenAndServe(":3001", nil)
}
