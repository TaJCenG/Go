package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // signal completion
	fmt.Println("Worker", id, "starting")
	fmt.Println("Worker", id, "done")
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)         // add one goroutine
		go worker(i, &wg) // start worker
	}

	wg.Wait() // wait for all workers
	fmt.Println("All workers finished")
}

//& is the address-of operator in Go.

//It gives you the memory address of a variable, not the value itself.
