package main

import "fmt"

func main() {
	ch := make(chan int, 2) // buffer size 2

	ch <- 1
	ch <- 2
	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2
}

//Channels can be buffered (hold multiple values) instead of blocking immediately.
