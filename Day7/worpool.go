package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
	}
}

func main() {
	jobs := make(chan int, 5)
	var wg sync.WaitGroup

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
	fmt.Println("All jobs done")

}

// ctx, cancel := context.WithCancel(context.Background())
// go func() {
//     // simulate shutdown after 5s
//     time.Sleep(5 * time.Second)
//     cancel()
// }()
