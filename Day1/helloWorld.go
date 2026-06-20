package main

import "fmt"

func main() {
	fmt.Println("hello world") //recommended approach
	println("hello world")
	fmt.Println(add(1, 2))
}

func add(a int, b int) int {
	return a + b
}
