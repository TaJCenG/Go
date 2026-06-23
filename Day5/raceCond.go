package main

import (
	"fmt"
	"sync"
)

var counter = 0

func increment(wg *sync.WaitGroup) {
	defer wg.Done() //defer tells Go: “run this statement at the very end of the current function, just before it returns.”
	for i := 0; i < 10; i++ {
		counter++ // unsafe: multiple goroutines writing
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go increment(&wg)
	go increment(&wg)
	wg.Wait()
	fmt.Println("Counter:", counter) // unpredictable result
}
defer = schedule something to run at the end of the function.

// wg.Done() = tell the WaitGroup “this goroutine is finished.”

// Together, they make concurrency safe and predictable.