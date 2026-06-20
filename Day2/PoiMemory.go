package main

import "fmt"

func createPointer() *int {
	x := 10
	fmt.Println("Value of x:", x) // prints 10
	return &x                     // escapes to heap
}

func main() {
	p := createPointer()
	fmt.Println("Value from pointer:", *p) // prints 10
}
