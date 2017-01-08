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
		// el channel ya no puede recibir mas mensajes
		close(c)
	}()

	// usamos range para recorrer los mensajes encolados
	for n := range c {
		fmt.Println(n)
	}
}
