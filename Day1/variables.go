package main

import "fmt"

var name string = "Taj"
var age int = 20
var height = "5"

//firstName := "Taju" only allow in function scope, not in package scope

const FullName string = "Tajuddin"

const (
	Pending = iota
	Running
	Completed
)

func main() {
	fmt.Println("name:" + name)
	fmt.Println("age: ", fmt.Sprint(age))
	fmt.Println("height: " + height)
	firstname := "taju"
	fmt.Println("firstname:", firstname)
	// Prints the package-level constant
	fmt.Println("Package-level FullName:", FullName)

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

}

func showPackageConstant() {
	fmt.Println("Accessing package-level FullName again:", FullName)
}

func packagelevel() {
	fmt.Println("Package-level FullName:", FullName)
}

func printStatus(status int) {
	switch status {
	case Pending:
		fmt.Println("Status: Pending")
	case Running:
		fmt.Println("Status: Running")
	case Completed:
		fmt.Println("Status: Completed")
	default:
		fmt.Println("Status: Unknown")
	}
}
