package main

import "fmt"

//Value receiver: gets a copy of the struct. Changes don’t affect the original.
// func (p Person) IncrementAge() {
// 	p.Age++ // only changes the copy
// }

//Pointer receiver: gets a pointer to the struct. Changes affect the original.
// func (p *Person) IncrementAgee() {
// 	p.Age++ // modifies the actual struct
// }

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Println("Hello, I am", p.Name)
}

func (p *Person) Birthday() {
	p.Age++
}

func main() {
	x := 10
	px := &x         // px points to x
	fmt.Println(*px) // dereference → prints 10
	*px = 20
	fmt.Println(x) // now x is 20
	taj := Person{Name: "Taj", Age: 20}
	taj.Greet()
	taj.Birthday()
	fmt.Println(taj.Age) // prints 21

}
