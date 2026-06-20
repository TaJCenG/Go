package main

import "fmt"

func main() {
	fmt.Println("Welcome to Calculator")
	fmt.Println("Addition of 10 and 5 is:", addition(10, 5))
	fmt.Println("Subtraction of 10 and 5 is:", subtraction(10, 5))
	fmt.Println("Multiplication of 10 and 5 is:", multiplication(10, 5))
	result, err := division(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division of 10 and 5 is:", result)
	}
}

func addition(a int, b int) int {
	return a + b
}

func subtraction(a int, b int) int {
	return a - b
}

func multiplication(a int, b int) int {
	return a * b
}

func division(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("error: division by zero")
	}
	return a / b, nil
}
