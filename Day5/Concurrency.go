package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(name, "step", i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go task("Task A")
	go task("Task B")

	time.Sleep(time.Second * 3) // wait long enough
}
