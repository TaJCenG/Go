package main

import "fmt"

// Interface
type Shape interface {
	Area() float64
}

// Struct
type Circle struct {
	Radius float64
}

// Method
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Function that accepts any Shape
func printArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	c := Circle{Radius: 5}
	printArea(c) // Circle automatically satisfies Shape
}

// Shape is the interface.

// Circle has a method Area() float64.

// Because Circle matches the interface’s method signature, Go says: “Circle is a Shape.”

// No implements keyword needed.

//Go → Implicit: if the methods match, the type satisfies the interface automatically.
