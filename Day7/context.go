package main

import (
	"context"
	"fmt"
	"time"
)

// context is a package in Go that lets you control goroutines (cancel them, set deadlines, pass values).
func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped:", ctx.Err())
			return
		default:
			fmt.Println("Working...")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)

	time.Sleep(3 * time.Second)
	cancel() // stop worker
	time.Sleep(time.Second)
}

// ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
// defer cancel()

// select {
// case <-time.After(3 * time.Second):
//     fmt.Println("Finished work")
// case <-ctx.Done():
//     fmt.Println("Timeout:", ctx.Err())
//}
