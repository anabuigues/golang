package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	// imprime el primer valor y despues esta esperando los siguientes pero la
	// rutina anterior ya a terminado por lo que se queda en deadlock!
	for {
		fmt.Println(<-c)
	}
}

// This prints the number, but we have received this error:
// fatal error: all goroutines are asleep - deadlock!
// Where is the deadlock?
// Why are all goroutines asleep?
// How can we fix this?
