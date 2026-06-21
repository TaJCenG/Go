package main

import (
	"fmt"
)

func main() {
	arr := [3]int{1, 2, 3} //Fixed size, cannot grow or shrink.
	fmt.Println(arr[1])    // 1

	nums := []int{1, 2, 3} //More flexible than arrays — they are views over arrays.
	nums = append(nums, 4) //Can grow/shrink dynamically.
	fmt.Println(nums)      // [1 2 3 4]
	//Internally: a slice has pointer to array + length + capacity.
	//If you append beyond capacity, Go allocates a new underlying array.

	m := make(map[string]int)
	m["apple"] = 5
	m["banana"] = 10
	fmt.Println(m["apple"])  // 5
	fmt.Println(m["banana"]) // 10

	a := [3]int{1, 2, 3}
	b := a
	b[0] = 99
	fmt.Println(a[0]) // still 1

	s := []int{1, 2, 3}
	t := s
	t[0] = 99
	fmt.Println(s[0]) // 99 (same underlying array)

}
