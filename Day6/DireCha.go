package main

import "fmt"

func send(ch chan<- int) { // send-only
	ch <- 42
}

func receive(ch <-chan int) { // receive-only
	fmt.Println(<-ch)
}

func main() {
	ch := make(chan int)
	go send(ch)
	receive(ch)
}

//You can restrict channels to send-only or receive-only in function parameters.
