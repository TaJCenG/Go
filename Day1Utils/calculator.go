package main

import "fmt"

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
