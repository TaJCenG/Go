package main

import (
	"fmt"
)

func main() {

	firstname := "taju"
	fmt.Println("firstname:", firstname)
	// Prints the package-level constant
	// Declare a new constant with the same name inside main
	const FullName = "Taj"

	// Prints the function-level constant (shadows the package one)
	fmt.Println("Function-level FullName:", FullName)

	// To show both clearly, we can use a different name for the local one
	const LocalFullName = "Taj"
	fmt.Println("LocalFullName:", LocalFullName)

	// The package-level constant is still intact outside of main
	showPackageConstant()
	packagelevel()
	printStatus(Pending)
	printStatus(Running)
	printStatus(Completed)
	printStatus(4)

	result := addi(2, 3)
	fmt.Println(result)
	fmt.Println(addi(2, 2))
	fmt.Println(addi(one, two))
	fmt.Println(addi(three, four))
	fmt.Println(addi(five, five))

	name, age := getNameandAge()
	fmt.Println("Name: " + name)
	fmt.Println("Age: ", age+1)

	x, _ := getCoordinates() // ignore y
	fmt.Println("X coordinate:", x)

	num, err := stringToInt("123")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Converted number:", num)
	}

	fmt.Println("Welcome to Calculator")
	fmt.Println("Addition of 10 and 5 is:", addition(10, 5))
	fmt.Println("Subtraction of 10 and 5 is:", subtraction(10, 5))
	fmt.Println("Multiplication of 10 and 5 is:", multiplication(10, 5))
	result, err = division(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division of 10 and 5 is:", result)
	}
}
