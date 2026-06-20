//A struct is a collection of fields (like a lightweight class).

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Car struct {
	Brand string
	Year  int
}

func main() {
	p := Person{Name: "Taj", Age: 20}
	fmt.Println("Name:", p.Name, "Age:", p.Age)
	fmt.Printf("Person struct: %+v\n", p)
	fmt.Println("Person:", p.Name, p.Age)

	c := Car{Brand: "Tesla", Year: 2025}
	fmt.Println("Car:", c.Brand, c.Year)

	// Nesting
	owner := struct {
		Person
		Car
	}{
		Person: Person{Name: "Taj", Age: 20},
		Car:    Car{Brand: "Tesla", Year: 2025},
	}
	fmt.Println("Owner:", owner.Name, "Car:", owner.Brand)
}
