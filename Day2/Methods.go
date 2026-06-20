package main

import "fmt"

//Methods are functions with a receiver (like attaching behavior to a struct).
type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
}

func main() {
	taj := Person{Name: "Taj", Age: 20}
	taj.Greet()
}
