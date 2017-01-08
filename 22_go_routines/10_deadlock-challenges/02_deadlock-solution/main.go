package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	fmt.Println(<-c)
}

// A "go" statement starts the execution of a function call as an independent concurrent thread of control,
// or goroutine, within the same address space