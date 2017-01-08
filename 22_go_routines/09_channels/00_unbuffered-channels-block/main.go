package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	 make(chan int, 100)
	 The capacity, in number of elements, sets the size of the buffer in the channel.
	 If the capacity is zero or absent, the channel is unbuffered and communication succeeds only
	 when both a sender and receiver are ready
	 Otherwise, the channel is buffered and communication succeeds without blocking if the buffer
	 is not full (sends) or not empty (receives). A nil channel is never ready for communication.

	 ch <- v   Send v to channel ch.
	 v := <-ch Receive from ch, and assign value to v.
	 */
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Hola")
			c <- i
		}
	}()


	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	time.Sleep(time.Second)
}
