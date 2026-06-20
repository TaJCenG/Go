package main

import "fmt"

func main() {
	// Using new
	p := new(int)
	*p = 42
	fmt.Println("Pointer address:", p)
	fmt.Println("Pointer value:", *p)

	// Using make
	nums := make([]int, 3)
	nums[0], nums[1], nums[2] = 10, 20, 30
	fmt.Println("Slice:", nums)
}
