package main

import (
	"fmt"
	"strconv"
)

func addi(a int, b int) int {
	return a + b
}

func showPackageConstant() {
	fmt.Println("Accessing package-level FullName again:")
}

func packagelevel() {
	fmt.Println("Package-level FullName:")
}

func printStatus(status interface{}) {
	fmt.Println("Current status:", status)
}

func getNameandAge() (string, int) {
	return "Taj", 20
}

func getCoordinates() (int, int) {
	return 10, 20
}

func stringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}
