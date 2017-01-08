package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	c <- 1
	fmt.Println(<-c)
}

// This results in a deadlock.
// Can you determine why? because nobody read
// And what would you do to fix it? you need a goroutine

