package main

import "fmt"

func main() {
	ch := make(chan string) // create channel

	go func() {
		ch <- "Hello from goroutine" // send
	}()

	msg := <-ch // receive
	fmt.Println(msg)
}
