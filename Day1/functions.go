package main

import (
	"fmt"
	"strconv"
)

func addi(a int, b int) int {
	return a + b
}

var one int = 1
var two int = 2
var three int = 3
var four int = 4
var five int = 5

func getNameandAge() (string, int) {
	return "Taj", 20
}

func getCoordinates() (int, int) {
	return 10, 20
}

func stringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func main() {
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
}
